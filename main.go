package main

import (
	pepinohttpservice "github.com/lucie-cupcakes/pepino/httpservice"
)

var (
	dbHTTPService pepinohttpservice.DatabaseHTTPService
)

func init() {
	cfg := pepinohttpservice.DatabaseHTTPServiceConfig{}
	err := cfg.LoadFromJSONFile("./config.json")
	if err != nil {
		panic(err)
	}
	dbHTTPService.Initialize(&cfg)
}

func main() {
	err := dbHTTPService.Listen()
	if err != nil {
		panic(err)
	}
}
