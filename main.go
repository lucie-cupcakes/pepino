package main

import (
	pepinosvc "github.com/lucie-cupcakes/pepino/service"
)

func main() {
	var svc pepinosvc.Service
	var cfg pepinosvc.ServiceConfig

	err := cfg.LoadFromJSONFile("./config.json")
	if err != nil {
		panic(err)
	}
	svc.New(&cfg)

	err = svc.ListenAndHandleRequests()
	if err != nil {
		panic(err)
	}
}
