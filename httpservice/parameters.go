package pepinohttpservice

import "errors"

// RequestParameters is the struct to deserialize all the incoming HTTPRequests
type RequestParameters struct {
	Password  string
	Arguments map[string]string
}

// RequireArgument is ...
func (p *RequestParameters) RequireArgument(argumentName string) error {
	if p.Arguments == nil {
		return errors.New("missing arguments dictionary")
	}
	if _, ok := p.Arguments[argumentName]; !ok {
		return errors.New("argument not found: " + argumentName)
	}
	return nil
}
