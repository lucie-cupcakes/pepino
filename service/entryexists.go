package pepinoservice

import (
	"errors"
)

// EntryExists loads the database to memory and indicates if an entry exists
// on Error: false will be returned.
func (svc *DatabaseService) EntryExists(dbName string, entryName string) (bool, error) {
	if !svc.initialized {
		return false, errors.New("object DatabaseService is not initialized")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return false, err
	}
	_, entryFound := dbPtr.Entries[entryName]
	return entryFound, nil
}
