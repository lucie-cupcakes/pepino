package pepinoservice

import (
	"errors"
	"fmt"
	"os"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

func (h *ServiceHTTPHandler) handleGETMethod() {
	fmt.Println("pepino service: GET request")
	if h.handleError(h.loadServiceRequest()) {
		return
	}

	if h.handleError(h.checkPassword(h.serviceRequest.Password)) {
		return
	}

	svcReq := *h.serviceRequest
	if h.handleError(svcReq.RequireArgument("DatabaseName")) {
		return
	}

	if h.handleError(svcReq.RequireArgument("EntryName")) {
		return
	}
	dbName := svcReq.Arguments["DatabaseName"]
	entryName := svcReq.Arguments["EntryName"]

	dbPtr, dbPtrFound := h.service.Databases[dbName]
	if !dbPtrFound {
		var dbLocal engine.Database
		dbLocal.New(dbName)
		if !dbLocal.HasSavedData() {
			h.handleError(errors.New("the database " + dbName + " is empty"))
			return
		}
		if h.handleError(dbLocal.Load()) {
			return
		}
		dbPtr = &dbLocal
		h.service.Databases[dbName] = dbPtr
	}

	rw := *h.responseWriter
	entryValue, entryFound := dbPtr.Entries[entryName]
	if !entryFound {
		rw.WriteHeader(404)
		rw.Header().Add("Content-Type", "text/plain")
		rw.Write([]byte("the entry " + entryName + " is not found"))
		return
	}

	rw.WriteHeader(200)
	rw.Header().Add("Content-Type", "application/octet-stream")
	_, err := rw.Write(entryValue)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
