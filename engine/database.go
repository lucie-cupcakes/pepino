package pepinoengine

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

// Database contains information an specific db
type Database struct {
	Name    string
	Entries map[string][]byte
}

// New creates a new Database object
func (d *Database) New(name string) {
	d.Name = name
	d.Entries = make(map[string][]byte)
}

// Save stores object state
func (d *Database) Save() error {
	//@TODO: the Database folder path should be absolute and stored
	// into a config object or something.
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(d.Entries)
	if err != nil {
		return fmt.Errorf("error saving Database object:"+
			"\n\tSerialization error: %s", err.Error())
	}
	err = os.WriteFile("./data/"+d.Name, buff.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error saving Database object:"+
			"\n\tos.WriteFile() error: %s", err.Error())
	}
	return nil
}

// HasSavedData indicates if a Database object has saved data
// available to restore
func (d *Database) HasSavedData() bool {
	_, err := os.Stat("./data/" + d.Name)
	return err == nil
}

// Load restores object state
func (d *Database) Load() error {
	if !d.HasSavedData() {
		return errors.New("error loading Database object:" +
			"\n\tthe object does not have saved data")
	}

	var buff bytes.Buffer
	fBytes, err := os.ReadFile("./data/" + d.Name)
	if err != nil {
		return fmt.Errorf("error loading Database object:"+
			"\n\tos.ReadFile() error: %s", err.Error())
	}
	_, err = buff.Write(fBytes)
	if err != nil {
		return fmt.Errorf("error loading Database object:"+
			"\n\tbuff.Write() error: %s", err.Error())
	}
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&d.Entries)
	if err != nil {
		return fmt.Errorf("error loading Database object:"+
			"\n\tgobDecoder.Decode() error: %s", err.Error())
	}
	return nil
}
