package pepinoservice

import (
	"encoding/json"
	"fmt"
	"os"
)

// ServiceConfig is ...
type ServiceConfig struct {
	Port int
}

// LoadFromJSONFile is ...
func (cfg *ServiceConfig) LoadFromJSONFile(filePath string) error {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error loading config from JSON file: \n\t%s", err.Error())
	}
	err = json.Unmarshal(fileBytes, cfg)
	if err != nil {
		return fmt.Errorf("error loading config from JSON file: \n\t%s", err.Error())
	}
	return nil
}
