package pepinohttpservice

import (
	"log"
	"net/http"
	"time"
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
	t0 := time.Now()
	defer func() {
		td := time.Since(t0)
		log.Printf("request completed in %v", td)
	}()
	httpMethod := r.httpRequest.Method
	switch httpMethod {
	case "GET":
		r.handleGETMethod()
	case "POST":
		r.handlePOSTMethod()
	case "DELETE":
		r.handleDELETEMethod()
	default:
		r.writeError(http.StatusBadRequest, "Invalid HTTP Method")
	}
}
