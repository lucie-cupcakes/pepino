package pepinohttpservice

import (
	"fmt"
	"net/http"
)

func (r *request) handleDELETEMethod() {
	rw := *r.httpResponseWriter
	uri := r.httpRequest.URL.String()

	fmt.Println("pepino service: DELETE request " + uri)

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

	err := r.dbHTTPService.dbService.DeleteEntry(dbName, entryName)
	if err != nil {
		r.writeError(http.StatusInternalServerError, err.Error())
		return
	}

	rw.WriteHeader(http.StatusOK)
}
