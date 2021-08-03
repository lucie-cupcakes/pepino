package pepinoengine

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"os"
)

// PepinoDatabase contains information an specific db
type PepinoDatabase struct {
	Name    string
	Entries map[string][]byte
}

// New creates a new PepinoDatabase object
func (d *PepinoDatabase) New(name string) {
	d.Name = name
	d.Entries = make(map[string][]byte)
}

// Save stores object state
func (d *PepinoDatabase) Save() error {
	//@TODO: the Database folder path should be absolute and stored
	// into a config object or something.
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(d.Entries)
	if err != nil {
		return fmt.Errorf("error saving PepinoDatabase object:"+
			"\n\tSerialization error: %s", err.Error())
	}
	err = os.WriteFile("./data/"+d.Name, buff.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error saving PepinoDatabase object:"+
			"\n\tos.WriteFile() error: %s", err.Error())
	}
	return nil
}

// HasSavedData tells whatever a PepinoDatabase object has saved data
// available to restore
func (d *PepinoDatabase) HasSavedData() bool {
	if _, err := os.Stat("./data/" + d.Name); err == nil {
		return true
	} else {
		return false
	}
}

// Load restores object state
func (d *PepinoDatabase) Load() error {
	if !d.HasSavedData() {
		return errors.New("error loading PepinoDatabase object:" +
			"\n\tthe object does not have saved data")
	}

	var buff bytes.Buffer
	fBytes, err := os.ReadFile("./data/" + d.Name)
	if err != nil {
		return fmt.Errorf("error loading PepinoDatabase object:"+
			"\n\tos.ReadFile() error: %s", err.Error())
	}
	_, err = buff.Write(fBytes)
	if err != nil {
		return fmt.Errorf("error loading PepinoDatabase object:"+
			"\n\tbuff.Write() error: %s", err.Error())
	}
	dec := gob.NewDecoder(&buff)
	err = dec.Decode(&d.Entries)
	if err != nil {
		return fmt.Errorf("error loading PepinoDatabase object:"+
			"\n\tgobDecoder.Decode() error: %s", err.Error())
	}
	return nil
}
