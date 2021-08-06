package pepinoservice

import (
	"errors"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

func (svc *DatabaseService) loadDatabaseToMemory(dbName string,
	createIt bool) (*engine.Database, error) {
	dbPtr, dbPtrFound := svc.databases[dbName]
	if !dbPtrFound {
		var dbLocal engine.Database
		dbLocal.Initialize(dbName, svc.dataPath)
		if !createIt {
			if !dbLocal.HasSavedData() {
				return nil, errors.New("database is empty")
			}
			err := dbLocal.Load()
			if err != nil {
				return nil, err
			}
		}
		dbPtr = &dbLocal
		svc.databases[dbName] = dbPtr
	}
	return dbPtr, nil
}
