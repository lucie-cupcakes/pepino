package pepinohttpservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type request struct {
	dbHTTPService      *DatabaseHTTPService
	httpRequest        *http.Request
	httpResponseWriter *http.ResponseWriter
	initialized        bool
}

func (r *request) initialize(dbHTTPService *DatabaseHTTPService,
	httpRequest *http.Request, httpRw *http.ResponseWriter) {
	r.httpRequest = httpRequest
	r.httpResponseWriter = httpRw
	r.dbHTTPService = dbHTTPService
	r.initialized = true
}

func (r *request) handle() {
	httpMethod := r.httpRequest.Method
	switch httpMethod {
	case "GET":
		r.handleGETMethod()
	case "POST":
		r.handlePOSTMethod()
	case "DELETE":
		r.handleDELETEMethod()
	default:
		r.handleError(errors.New("warning: HTTP method \"" +
			httpMethod + "\" has no request handler."))
	}
}

func (r *request) checkPassword(pwd string) error {
	//@TODO: Use hash insted of Plain password.
	if pwd == r.dbHTTPService.config.Password {
		return nil
	}
	return errors.New("invalid password")
}

func (r *request) handleGETMethod() {
	rw := *r.httpResponseWriter
	uri := r.httpRequest.URL.Path

	fmt.Println("pepino service: GET request " + uri)

	uriValues := r.httpRequest.URL.Query()

	password := uriValues.Get("password")
	if r.checkPassword(password) != nil {
		fmt.Fprintln(os.Stderr, "forbidden: invalid password")
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	dbName := uriValues.Get("db")
	if dbName == "" {
		fmt.Fprintln(os.Stderr, "bad request: missing URI parameter: db")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	entryName := uriValues.Get("entry")
	if entryName == "" {
		fmt.Fprintln(os.Stderr, "bad request: missing URI parameter: entry")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	entryValue, err := r.dbHTTPService.dbService.GetEntry(dbName, entryName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/octet-stream")
	_, err = rw.Write(entryValue)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

func (r *request) handlePOSTMethod() {
	rw := *r.httpResponseWriter
	uri := r.httpRequest.URL.Path

	fmt.Println("pepino service: POST request " + uri)

	uriValues := r.httpRequest.URL.Query()

	password := uriValues.Get("password")
	if r.checkPassword(password) != nil {
		fmt.Fprintln(os.Stderr, "forbidden: invalid password")
		rw.WriteHeader(http.StatusForbidden)
		return
	}

	dbName := uriValues.Get("db")
	if dbName == "" {
		fmt.Fprintln(os.Stderr, "bad request: missing URI parameter: db")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	entryName := uriValues.Get("entry")
	if entryName == "" {
		fmt.Fprintln(os.Stderr, "bad request: missing URI parameter: entry")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	entryValue, err := ioutil.ReadAll(r.httpRequest.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = r.dbHTTPService.dbService.PutEntry(dbName, entryName, entryValue)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (r *request) handleDELETEMethod() {
	fmt.Println("pepino service: DELETE request")
}

// RequestError is ...
type RequestError struct {
	Code        int
	Description string
}

func (r *request) handleError(err error) bool {
	//@TODO: Respond with actual error codes for know errors.
	if err == nil {
		return false
	}
	fmt.Fprintln(os.Stderr, err)
	var reqErr RequestError
	reqErr.Code = 999
	reqErr.Description = err.Error()
	errBytes, err := json.Marshal(&reqErr)
	rw := *r.httpResponseWriter
	rw.WriteHeader(500)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		if _, err := rw.Write(errBytes); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	return true
}
