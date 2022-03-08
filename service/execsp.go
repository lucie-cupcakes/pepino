package pepinoservice

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

// ExecStoredProcedure loads the meta data of an sp and executes it.
// StandardOutput is returned
func (svc *DatabaseService) ExecStoredProcedure(dbName string, entryName string,
	inputReader io.ReadCloser, env map[string]string) ([]byte, error) {
	if !svc.initialized {
		return nil, errors.New("object DatabaseService is not initialized")
	}
	if !svc.enableSp {
		return nil, errors.New("stored procedures disabled")
	}
	if !strings.HasPrefix(entryName, "sp_") {
		return nil, errors.New("only entries starting with 'sp_' can be executed")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return nil, err
	}
	spValue, spFound := dbPtr.Entries[entryName]
	if !spFound {
		return nil, errors.New("entry not found")
	}

	spMeta := StoredProcedureMetadata{}

	spMetaBytes, spMetaBytesFound := dbPtr.Entries[entryName+"_meta"]
	if spMetaBytesFound {
		err = spMeta.LoadFromJSONByteArray(spMetaBytes)
		if err != nil {
			return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("cannot execute entry: \n\tmeta entry does not exists")
	}

	entryID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}

	spDirPath := svc.tmpPath + entryName + "_" + entryID.String()
	err = os.Mkdir(spDirPath, 0755)
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	defer os.RemoveAll(spDirPath)
	err = os.WriteFile(spDirPath+"/entry", spValue, 0664)
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}

	if spMeta.IsTar {
		var cmdTarStdErr bytes.Buffer
		cmdTar := exec.Command("tar", "-xf", "entry")
		cmdTar.Dir = spDirPath
		cmdTar.Stderr = &cmdTarStdErr
		err = cmdTar.Run()
		if err != nil {
			if strings.HasPrefix(err.Error(), "exit status") && cmdTarStdErr.Len() > 0 {
				return nil, errors.New(cmdTarStdErr.String())
			}
			return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
		}
	}
	entryFileName := "entry"
	if spMeta.IsTar {
		entryFileName = spMeta.TarEntryPoint
	}
	cmd := exec.Command(spMeta.Interpreter, entryFileName)
	var cmdStdOut bytes.Buffer
	var cmdStdErr bytes.Buffer
	cmd.Stdout = &cmdStdOut
	cmd.Stderr = &cmdStdErr
	cmd.Dir = spDirPath
	if inputReader != nil {
		cmd.Stdin = inputReader
	}
	cmd.Env = os.Environ()
	if env != nil && len(env) > 0 {
		for k, v := range env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
	err = cmd.Run()
	if err != nil {
		if strings.HasPrefix(err.Error(), "exit status") && cmdStdErr.Len() > 0 {
			return nil, errors.New(cmdStdErr.String())
		}
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	return cmdStdOut.Bytes(), nil
}
