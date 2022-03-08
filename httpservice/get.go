package pepinohttpservice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func (r *request) handleGETMethodGetEntry(dbName string, entryName string) {
	rw := *r.httpResponseWriter
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

func (r *request) handleGETMethodExecEntry(dbName string, entryName string) {
	rw := *r.httpResponseWriter
	var bodyReader io.ReadCloser
	var err error
	if r.httpRequest.ContentLength > 0 {
		bodyReader, err = r.httpRequest.GetBody()
		if err != nil {
			r.writeError(http.StatusInternalServerError, err.Error())
			return
		}
	}
	cmdEnv := make(map[string]string)
	cmdEnv["PEPINODB_LURI"] = (func() string {
		res := strings.Builder{}
		if r.dbHTTPService.config.TLSEnable {
			res.WriteString("https://localhost:")
		} else {
			res.WriteString("http://localhost:")
		}
		res.WriteString(strconv.Itoa(r.dbHTTPService.config.Port))
		res.WriteString("/?password=")
		res.WriteString(url.QueryEscape(r.dbHTTPService.config.Password))
		return res.String()
	})()
	cmdEnv["PEPINODB_HOST"] = r.dbHTTPService.config.Host
	cmdEnv["PEPINODB_PORT"] = strconv.Itoa(r.dbHTTPService.config.Port)
	cmdEnv["PEPINODB_TLS"] = (func() string {
		if r.dbHTTPService.config.TLSEnable {
			return "True"
		}
		return "False"
	})()
	cmdEnv["PEPINODB_PWD"] = r.dbHTTPService.config.Password
	cmdEnv["PEPINODB_DB"] = dbName
	cmdEnv["PEPINODB_SCRIPT"] = entryName
	cmdEnv["PEPINODB_URI_PARAMS"] = r.httpRequest.URL.String()
	execResult, err := r.dbHTTPService.dbService.ExecStoredProcedure(dbName, entryName, bodyReader, cmdEnv)
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

func (r *request) handleGETMethodEntryExists(dbName string, entryName string) {
	rw := *r.httpResponseWriter
	entryExists, err := r.dbHTTPService.dbService.EntryExists(dbName, entryName)
	if err != nil {
		if err.Error() == "database is empty" {
			rw.WriteHeader(http.StatusNotFound)
			return
		}
		r.writeError(http.StatusInternalServerError, err.Error())
		return
	}
	if entryExists {
		rw.WriteHeader(http.StatusOK)
		return
	}
	rw.WriteHeader(http.StatusNotFound)
}

func (r *request) handleGETMethodListEntries(dbName string) {
	rw := *r.httpResponseWriter
	entryList, err := r.dbHTTPService.dbService.ListEntries(dbName)
	if err != nil {
		r.writeError(http.StatusInternalServerError, err.Error())
	}
	entryListJSON, err := json.Marshal(&entryList)
	if err != nil {
		r.writeError(http.StatusInternalServerError, err.Error())
	}
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "text/json")
	_, err = rw.Write(entryListJSON)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func (r *request) handleGETMethod() {
	uri := r.httpRequest.URL.String()
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

	listParam := uriValues.Get("list")
	if strings.ToLower(listParam) == "true" {
		r.handleGETMethodListEntries(dbName)
		return
	}

	entryName := uriValues.Get("entry")
	if entryName == "" {
		r.writeError(http.StatusBadRequest, "missing URI parameter: entry")
		return
	}

	existsParam := uriValues.Get("exists")
	if strings.ToLower(existsParam) == "true" {
		r.handleGETMethodEntryExists(dbName, entryName)
		return
	}

	execParam := uriValues.Get("exec")
	if strings.ToLower(execParam) == "true" {
		r.handleGETMethodExecEntry(dbName, entryName)
		return
	}
	r.handleGETMethodGetEntry(dbName, entryName)
}
