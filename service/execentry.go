package pepinoservice

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
)

// ExecEntry loads the database to memory
// and executes the given entryName as a Python script.
// StandardOutput is returned
func (svc *DatabaseService) ExecEntry(dbName string, entryName string, input []byte) ([]byte, error) {
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
	cmd.Stdin = bytes.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("cannot execute entry: \n\t%s", err.Error())
	}
	return out.Bytes(), nil
}
