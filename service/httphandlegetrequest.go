package pepinoservice

import (
	"errors"
	"fmt"
)

func (h *ServiceHTTPHandler) handleGETMethod() {
	fmt.Println("pepino service: GET request")
	if h.handleError(h.loadServiceRequest()) {
		return
	}

	if h.handleError(h.checkPassword(h.serviceRequest.Password)) {
		return
	}

	//rw := *h.responseWriter
	//rw.Header().Add("Content-Type", "text/plain")
	//rw.Write([]byte("Good :)\n"))

	svcReq := *h.serviceRequest

	if _, ok := svcReq.Arguments["DatabaseName"]; !ok {
		h.handleError(errors.New("argument not found: DatabaseName"))
		return
	}

	if _, ok := svcReq.Arguments["DatabaseName"]; !ok {
		h.handleError(errors.New("argument not found: DatabaseName"))
		return
	}
}
