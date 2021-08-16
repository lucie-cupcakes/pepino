package pepinohttpservice

import (
	"fmt"
	"net/http"
	"os"
	"strings"
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

	execParam := uriValues.Get("exec")
	if execParam == "" || strings.ToLower(execParam) == "false" {
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
	} else if strings.ToLower(execParam) == "true" {
		bodyReader, err := r.httpRequest.GetBody()
		if err != nil {
			r.writeError(http.StatusInternalServerError, err.Error())
			return
		}
		execResult, err := r.dbHTTPService.dbService.ExecEntry(dbName, entryName, bodyReader)
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
		_, err = rw.Write(execResult)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
	}
}
