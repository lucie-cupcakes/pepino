package pepinoservice

import (
	"encoding/json"
	"fmt"
)

// StoredProcedureMetadata contains information about an stored procedure
type StoredProcedureMetadata struct {
	Interpreter   string `json:"intepreter"`
	IsTar         bool   `json:"isTar"`
	TarEntryPoint string `json:"tarEntryPoint"`
}

// LoadFromJSONByteArray initializes ExecMeta object from JSON byte array.
func (spMeta *StoredProcedureMetadata) LoadFromJSONByteArray(fileBytes []byte) error {
	err := json.Unmarshal(fileBytes, spMeta)
	if err != nil {
		return fmt.Errorf("error loading config from JSON file: \n\t%s", err.Error())
	}
	return nil
}
