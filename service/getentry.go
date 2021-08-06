package pepinoservice

import (
	"errors"
)

// GetEntry automatically loads the database to memory and returns
// given entry value if it exists.
func (svc *DatabaseService) GetEntry(dbName string, entryName string) ([]byte, error) {
	if !svc.initialized {
		return nil, errors.New("object DatabaseService is not initialized")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return nil, err
	}
	entryValue, entryFound := dbPtr.Entries[entryName]
	if !entryFound {
		return nil, errors.New("entry not found")
	}
	return entryValue, nil
}
