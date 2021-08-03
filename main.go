package main

import (
	pepinoengine "github.com/lucie-cupcakes/pepino/engine"
)

func main() {
	var db pepinoengine.PepinoDatabase
	db.New("test-db")
	db.Entries["hello"] = []byte("Hello world!")
	db.Save()

}
