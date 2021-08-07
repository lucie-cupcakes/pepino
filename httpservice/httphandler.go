package pepinohttpservice

import (
	"net/http"
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
		r.writeError(http.StatusBadRequest, "Invalid HTTP Method")
	}
}
