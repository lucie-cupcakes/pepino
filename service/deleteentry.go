package pepinoservice

import (
	"errors"
)

// DeleteEntry automatically loads the database to memory and deletes an entry
// if it exists.
func (svc *DatabaseService) DeleteEntry(dbName string, entryName string) error {
	if !svc.initialized {
		return errors.New("object DatabaseService is not initialized")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return err
	}
	_, entryFound := dbPtr.Entries[entryName]
	if !entryFound {
		return errors.New("entry not found")
	}
	delete(dbPtr.Entries, entryName)
	err = dbPtr.Save()
	if err != nil {
		return err
	}
	return nil
}
