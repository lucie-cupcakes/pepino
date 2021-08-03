package pepinoengine

import (
	"bytes"
	"encoding/gob"
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
			"\n\tos.WriteFile error: %s", err.Error())
	}
	return nil
}

// Load restores object state
func (d *PepinoDatabase) Load() {
	fmt.Println("pepinoengine: TODO: Load()")
}
