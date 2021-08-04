package pepinoservice

import "fmt"

func (h *ServiceHTTPHandler) handleGETMethod() {
	fmt.Println("pepino service: GET request")
	if h.handleError(h.loadServiceRequest()) {
		return
	}

	if h.handleError(h.checkPassword(h.serviceRequest.Password)) {
		return
	}

	rw := *h.responseWriter
	rw.Header().Add("Content-Type", "text/plain")
	rw.Write([]byte("Good :)\n"))
}
