package pepinoservice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

// Service contains the internal of the HTTP Service for the database
type Service struct {
	Databases   map[string]*engine.Database
	Config      *ServiceConfig
	httpHandler func(rw http.ResponseWriter, r *http.Request)
	initialized bool
}

// New initializes the Service object
func (s *Service) New(cfg *ServiceConfig) {
	s.Config = cfg
	s.httpHandler = func(rw http.ResponseWriter, r *http.Request) {
		switch strings.ToUpper(r.Method) {
		case "GET":
			fmt.Println("pepino service: GET request")
		case "POST":
			fmt.Println("pepino service: POST request")
		case "DELETE":
			fmt.Println("pepino service: DELETE request")
		}
	}
	s.initialized = true
}

// ListenAndHandleRequests starts HTTP service to handle database requests
func (s *Service) ListenAndHandleRequests() error {
	if !s.initialized {
		return errors.New("the object Service is not initialized")
	}
	hostStr := "localhost:" + strconv.Itoa(s.Config.Port)
	fmt.Println("pepino service: Listening on HTTP " + hostStr)
	http.ListenAndServe(hostStr, http.HandlerFunc(s.httpHandler))
	return nil
}
