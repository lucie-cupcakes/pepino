package pepinoservice

import (
	engine "github.com/lucie-cupcakes/pepino/engine"
)

// DatabaseService provides an interface to access pepino databases.
type DatabaseService struct {
	databases   map[string]*engine.Database
	dataPath    string
	initialized bool
}

// Initialize DatabaseService object
func (svc *DatabaseService) Initialize(dataPath string) {
	svc.databases = make(map[string]*engine.Database)
	// @TODO: Validate dataPath exists.
	svc.dataPath = dataPath
	svc.initialized = true
}
