package pepinohttpservice

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	service "github.com/lucie-cupcakes/pepino/service"
)

// DatabaseHTTPService contains the internals of the HTTP Service
// for the DatabaseService
type DatabaseHTTPService struct {
	dbService   *service.DatabaseService
	config      *DatabaseHTTPServiceConfig
	initialized bool
}

// Initialize : setup DatabaseHTTPService object
func (s *DatabaseHTTPService) Initialize(cfg *DatabaseHTTPServiceConfig) error {
	s.config = cfg
	s.initialized = true
	var dbSvc service.DatabaseService
	err := dbSvc.Initialize(cfg.DataPath, cfg.TmpPath, cfg.EnableStoredProcedures)
	if err != nil {
		return err
	}
	s.dbService = &dbSvc
	return nil
}

func (s *DatabaseHTTPService) getHTTPHandle() http.Handler {
	return http.HandlerFunc(func(httpRw http.ResponseWriter, httpReq *http.Request) {
		var r request
		r.initialize(s, httpReq, &httpRw)
		r.handle()
	})
}

// Listen starts HTTP service to handle requests
func (s *DatabaseHTTPService) Listen() error {
	if !s.initialized {
		return errors.New("the object Service is not initialized")
	}
	hostStr := s.config.Host + ":" + strconv.Itoa(s.config.Port)
	if s.config.TLSEnable {
		fmt.Println("pepino service: Listening on HTTPS@" + hostStr)
		http.ListenAndServeTLS(hostStr, s.config.TLSCertFile, s.config.TLSKeyFile, s.getHTTPHandle())
	} else {
		fmt.Println("pepino service: Listening on HTTP@" + hostStr)
		http.ListenAndServe(hostStr, s.getHTTPHandle())
	}
	return nil
}
