package pepinohttpservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type request struct {
	dbHTTPService      *DatabaseHTTPService
	parameters         *RequestParameters
	httpRequest        *http.Request
	httpResponseWriter *http.ResponseWriter
	initialized        bool
}

func (r *request) initialize(dbHTTPService *DatabaseHTTPService,
	httpRequest *http.Request, httpRw *http.ResponseWriter) {
	r.httpRequest = httpRequest
	r.httpResponseWriter = httpRw
	r.dbHTTPService = dbHTTPService
	r.parameters = nil
	r.initialized = true
}

func (r *request) loadParameters() error {
	if !r.initialized {
		return errors.New("object outgoingRequest is not initialized")
	}
	if r.parameters != nil {
		return errors.New("parameters already loaded")
	}
	httpRequest := r.httpRequest
	if !strings.HasSuffix(strings.ToLower(httpRequest.Header.Get("Content-Type")), "/json") {
		return errors.New("invalid Content-Type")
	}

	bodyBytes, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		return err
	}
	var p RequestParameters
	err = json.Unmarshal(bodyBytes, &p)
	if err != nil {
		//@TODO: FormatError
		return err
	}
	r.parameters = &p

	if p.Arguments == nil {
		return errors.New("missing arguments dictionary")
	}

	// DEBUG:
	jBytes, err := json.Marshal(&p)
	if err == nil {
		fmt.Println("Request: " + string(jBytes))
	}
	return nil
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
	fmt.Println("pepino service: GET request")
	if r.handleError(r.loadParameters()) {
		return
	}
	if r.handleError(r.checkPassword(r.parameters.Password)) {
		return
	}

	if r.handleError(r.parameters.RequireArgument("DatabaseName")) {
		return
	}

	if r.handleError(r.parameters.RequireArgument("EntryName")) {
		return
	}

	entryValue, err := r.dbHTTPService.dbService.GetEntry(r.parameters.Arguments["DatabaseName"], r.parameters.Arguments["EntryName"])

	if r.handleError(err) {
		return
	}
	rw := *r.httpResponseWriter
	rw.WriteHeader(200)
	rw.Header().Add("Content-Type", "application/octet-stream")
	_, err = rw.Write(entryValue)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}

	/*rw.WriteHeader(404)
	rw.Header().Add("Content-Type", "text/plain")
	rw.Write([]byte("the entry " + entryName + " is not found"))
	*/
}

func (r *request) handlePOSTMethod() {
	fmt.Println("pepino service: POST request")
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
