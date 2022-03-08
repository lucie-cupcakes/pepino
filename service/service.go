package pepinoservice

import (
	"fmt"
	"os"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

// DatabaseService provides an interface to access pepino databases.
type DatabaseService struct {
	databases   map[string]*engine.Database
	dataPath    string
	tmpPath     string
	enableSp    bool
	initialized bool
}

// Initialize DatabaseService object
func (svc *DatabaseService) Initialize(dataPath string, tmpPath string, enableSp bool) error {
	svc.databases = make(map[string]*engine.Database)
	// @TODO: Validate dataPath exists.
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return fmt.Errorf("cannot initialize DatabaseService object:\n\tdataPath \"%s\"does not exists", dataPath)
	}
	if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
		return fmt.Errorf("cannot initialize DatabaseService object:\n\ttmpPath: \"%s\" does not exists", dataPath)
	}
	svc.dataPath = dataPath
	svc.tmpPath = tmpPath
	svc.enableSp = enableSp
	svc.initialized = true
	return nil
}
