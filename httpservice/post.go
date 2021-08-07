package pepinohttpservice

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *request) handlePOSTMethod() {
	rw := *r.httpResponseWriter
	uri := r.httpRequest.URL.Path

	fmt.Println("pepino service: POST request " + uri)

	uriValues := r.httpRequest.URL.Query()

	password := uriValues.Get("password")
	if !r.checkPassword(password) {
		r.writeError(http.StatusForbidden, "invalid password")
		return
	}

	dbName := uriValues.Get("db")
	if dbName == "" {
		r.writeError(http.StatusBadRequest, "missing URI parameter: db")
		return
	}

	entryName := uriValues.Get("entry")
	if entryName == "" {
		r.writeError(http.StatusBadRequest, "missing URI parameter: entry")
		return
	}

	entryValue, err := ioutil.ReadAll(r.httpRequest.Body)
	if err != nil {
		r.writeError(http.StatusInternalServerError, err.Error())
		return
	}

	err = r.dbHTTPService.dbService.PutEntry(dbName, entryName, entryValue)
	if err != nil {
		r.writeError(http.StatusInternalServerError, err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
}
