package pepinohttpservice

import (
	"encoding/json"
	"fmt"
	"os"
)

// DatabaseHTTPServiceConfig contains configuration values for the
// DatabaseHTTPService object
type DatabaseHTTPServiceConfig struct {
	Host                   string
	Port                   int
	Password               string
	TLSEnable              bool
	TLSCertFile            string
	TLSKeyFile             string
	DataPath               string
	TmpPath                string
	EnableStoredProcedures bool
}

// LoadFromJSONFile initializes a DatabaseHTTPServiceConfig from a JSON File
func (cfg *DatabaseHTTPServiceConfig) LoadFromJSONFile(filePath string) error {
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
