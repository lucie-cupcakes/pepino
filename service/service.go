package pepinoservice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	engine "github.com/lucie-cupcakes/pepino/engine"
)

// Service contains the internal of the HTTP Service for the database
type Service struct {
	Databases   map[string]*engine.Database
	Config      *ServiceConfig
	initialized bool
}

// New initializes the Service object
func (s *Service) New(cfg *ServiceConfig) {
	s.Config = cfg
	s.initialized = true
}

// ListenAndHandleRequests starts HTTP service to handle database requests
func (s *Service) ListenAndHandleRequests() error {
	if !s.initialized {
		return errors.New("the object Service is not initialized")
	}
	hostStr := s.Config.Host + ":" + strconv.Itoa(s.Config.Port)
	if s.Config.TLSEnable {
		fmt.Println("pepino service: Listening on HTTPS@" + hostStr)
		http.ListenAndServeTLS(hostStr, s.Config.TLSCertFile, s.Config.TLSKeyFile, s.getHTTPHandle())
	} else {
		fmt.Println("pepino service: Listening on HTTP@" + hostStr)
		http.ListenAndServe(hostStr, s.getHTTPHandle())
	}
	return nil
}
