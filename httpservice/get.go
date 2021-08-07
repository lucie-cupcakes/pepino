package pepinohttpservice

import (
	"fmt"
	"net/http"
	"os"
)

func (r *request) handleGETMethod() {
	rw := *r.httpResponseWriter
	uri := r.httpRequest.URL.Path

	fmt.Println("pepino service: GET request " + uri)

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

	entryValue, err := r.dbHTTPService.dbService.GetEntry(dbName, entryName)
	if err != nil {
		if err.Error() == "entry not found" {
			r.writeError(http.StatusNotFound, err.Error())
		} else {
			r.writeError(http.StatusInternalServerError, err.Error())
		}
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/octet-stream")
	_, err = rw.Write(entryValue)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
