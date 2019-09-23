package devto

import "errors"

//Confugration errors
var (
	ErrMissingRequiredParameter = errors.New("a required parameter is missing")
)

//Config contains the elements required to initialize a
// devto client.
type Config struct {
	APIKey       string
	InsecureOnly bool
}

//NewConfig build a devto configuration instance with the
//required parameters.
//
//It takes a boolean (p) as first parameter to indicate if
//you need access to endpoints which require authentication,
//and a API key as second paramenter, if p is set to true and
//you don't provide an API key, it will return an error.
func NewConfig(p bool, k string) (c *Config, err error) {
	if p == true && k == "" {
		return nil, ErrMissingRequiredParameter
	}

	return &Config{
		InsecureOnly: !p,
		APIKey:       k,
	}, nil
}
