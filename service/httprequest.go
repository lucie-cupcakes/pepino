package pepinoservice

import "errors"

// ServiceHTTPRequest is ...
type ServiceHTTPRequest struct {
	Password  string
	Arguments map[string]string
}

// RequireArgument is ...
func (svcReq *ServiceHTTPRequest) RequireArgument(argumentName string) error {
	if svcReq.Arguments == nil {
		return errors.New("missing arguments dictionary")
	}
	if _, ok := svcReq.Arguments[argumentName]; !ok {
		return errors.New("argument not found: " + argumentName)
	}
	return nil
}
