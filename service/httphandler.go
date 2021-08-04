package pepinoservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ServiceHTTPHandler is the state of an ongoing HTTPRequest
type ServiceHTTPHandler struct {
	service        *Service
	request        *http.Request
	responseWriter *http.ResponseWriter
	serviceRequest *ServiceHTTPRequest
	initialized    bool
}

// New initializes a ServiceHTTPHandler object
func (h *ServiceHTTPHandler) New(svc *Service, r *http.Request,
	rw *http.ResponseWriter) {
	h.request = r
	h.responseWriter = rw
	h.service = svc
	h.serviceRequest = nil
	h.initialized = true
}

func (h *ServiceHTTPHandler) loadServiceRequest() error {
	req := h.request
	if !strings.HasSuffix(strings.ToLower(req.Header.Get("Content-Type")), "/json") {
		return errors.New("invalid Content-Type")
	}

	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	var svcReq ServiceHTTPRequest
	json.Unmarshal(bodyBytes, &svcReq)
	h.serviceRequest = &svcReq

	if svcReq.Arguments == nil {
		return errors.New("missing arguments dictionary")
	}

	// DEBUG:
	jBytes, err := json.Marshal(&svcReq)
	if err == nil {
		fmt.Println("Request: " + string(jBytes))
	}
	return nil
}

// Handle is the main function that should be called after initializing
// the ServiceHTTPHandler object.
// It executes the corresponding HTTP method corresponding to the HTTP.Request
func (h *ServiceHTTPHandler) Handle() {
	switch h.request.Method {
	case "GET":
		h.handleGETMethod()
	case "POST":
		h.handlePOSTMethod()
	case "DELETE":
		h.handleDELETEMethod()
	default:
		h.handleError(errors.New("warning: HTTP method \"" +
			h.request.Method + "\" has no request handler."))
	}
}

func (h *ServiceHTTPHandler) checkPassword(pwd string) error {
	//@TODO: Use hash insted of Plain password.
	if pwd == h.service.Config.Password {
		return nil
	}
	return errors.New("invalid password")
}

func (svc *Service) getHTTPHandle() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var h ServiceHTTPHandler
		h.New(svc, r, &rw)
		h.Handle()
	})
}
