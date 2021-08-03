package main

import (
	"fmt"
	"strconv"

	pepinoengine "github.com/lucie-cupcakes/pepino/engine"
)

func main() {
	var db pepinoengine.PepinoDatabase
	db.New("test-db")
	if db.HasSavedData() {
		err := db.Load()
		if err != nil {
			panic(err)
		}
		fmt.Println("The database has been loaded.")
		fmt.Println("there are " + strconv.Itoa(len(db.Entries)) + " entries.")
		for k, v := range db.Entries {
			fmt.Println("key{" + k + "} = \"" + string(v) + "\"")
		}
	} else {
		fmt.Println("The database has been created.")
		db.Entries["hello"] = []byte("Hello world!")
		err := db.Save()
		if err != nil {
			panic(err)
		}
		fmt.Println("The database has been saved.")
	}

}
