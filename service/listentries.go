package pepinoservice

import (
	"errors"
)

// ListEntries loads the database to memory
// and returns the entries name in a database
func (svc *DatabaseService) ListEntries(dbName string) ([]string, error) {
	if !svc.initialized {
		return nil, errors.New("object DatabaseService is not initialized")
	}
	dbPtr, err := svc.loadDatabaseToMemory(dbName, false)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(dbPtr.Entries))
	for k := range dbPtr.Entries {
		keys = append(keys, k)
	}
	return keys, nil
}
