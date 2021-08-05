package pepinoservice

import (
	"errors"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

// GetEntry automatically loads the database to memory and returns
// given entry value if it exists.
func (svc *DatabaseService) GetEntry(dbName string, entryName string) ([]byte, error) {
	if !svc.initialized {
		return nil, errors.New("object DatabaseService is not initialized")
	}
	dbPtr, dbPtrFound := svc.databases[dbName]
	if !dbPtrFound {
		var dbLocal engine.Database
		dbLocal.Initialize(dbName, svc.dataPath)
		if !dbLocal.HasSavedData() {
			return nil, errors.New("database is empty")
		}
		err := dbLocal.Load()
		if err != nil {
			return nil, err
		}
		dbPtr = &dbLocal
		svc.databases[dbName] = dbPtr
	}
	entryValue, entryFound := dbPtr.Entries[entryName]
	if !entryFound {
		return nil, errors.New("entry not found")
	}
	return entryValue, nil
}
