package main

import (
	pepinosvc "github.com/lucie-cupcakes/pepino/service"
)

func main() {
	var svc pepinosvc.Service

	err := svc.Config.LoadFromJSONFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = svc.ListenAndHandleRequests()
	if err != nil {
		panic(err)
	}
}
