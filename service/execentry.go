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

// ExecEntry loads the database to memory
// and executes the given entryName as a Python script.
// StandardOutput is returned
func (svc *DatabaseService) ExecEntry(dbName string, entryName string,
	inputReader io.ReadCloser, env map[string]string) ([]byte, error) {
	if !svc.initialized {
		return nil, errors.New("object DatabaseService is not initialized")
	}
	if !strings.HasSuffix(entryName, ".py") {
		return nil, errors.New("only entries ending with .py can be executed")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return nil, err
	}
	entryValue, entryFound := dbPtr.Entries[entryName]
	if !entryFound {
		return nil, errors.New("entry not found")
	}
	entryID, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	entryFPath := "/tmp/" + entryID.String() + "-" + entryName
	err = os.WriteFile(entryFPath, entryValue, 0664)
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	defer os.Remove(entryFPath)
	cmd := exec.Command("python3", entryFPath)
	if inputReader != nil {
		cmd.Stdin = inputReader
	}
	cmd.Env = os.Environ()
	if env != nil && len(env) > 0 {
		for k, v := range env {
			cmd.Env = append(cmd.Env, k+"="+v)
		}
	}
	var cmdStdOut bytes.Buffer
	var cmdStdErr bytes.Buffer
	cmd.Stdout = &cmdStdOut
	cmd.Stderr = &cmdStdErr
	err = cmd.Run()
	if err != nil {
		if strings.HasPrefix(err.Error(), "exit status") && cmdStdErr.Len() > 0 {
			return nil, errors.New(cmdStdErr.String())
		}
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	return cmdStdOut.Bytes(), nil
}
