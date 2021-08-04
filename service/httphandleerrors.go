package pepinoservice

import (
	"encoding/json"
	"fmt"
	"os"
)

// ServiceHTTPRequestError is ...
type ServiceHTTPRequestError struct {
	Code        int
	Description string
}

func (h *ServiceHTTPHandler) handleError(err error) bool {
	//@TODO: Respond with actual error codes for know errors.
	if err == nil {
		return false
	}
	fmt.Fprintln(os.Stderr, err)
	var httpErr ServiceHTTPRequestError
	httpErr.Code = 999
	httpErr.Description = err.Error()
	errBytes, err := json.Marshal(httpErr)
	rw := *h.responseWriter
	rw.WriteHeader(500)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		rw.Write(errBytes)
	}
	return true
}
