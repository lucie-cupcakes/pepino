package pepinohttpservice

import (
	"fmt"
	"os"
	"strconv"
)

func (r *request) writeError(statusCode int, description string) {
	fmt.Fprintln(os.Stderr, "HTTP Error "+
		strconv.Itoa(statusCode)+": "+description)
	rw := *r.httpResponseWriter
	rw.WriteHeader(statusCode)
	rw.Header().Add("Content-Type", "text/plain")
	_, err := rw.Write([]byte(description))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error writing HTTP response: "+err.Error())
	}
}
