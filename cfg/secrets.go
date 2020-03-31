package cfg

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/docker/docker-credential-helpers/client"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/logger"
)

type Secret struct {
	Service  string
	Secret   string
	Username string
	Store    string
}

// Configure the secret for a service.
//
// Does not overwrite explicitly configured values, so is safe to call
// if username and secret were explicitly set in module config.
//
// Input:
// * service: URL or identifier for service if configured by user. Not all
//   modules support or need this. Optional, defaults to serviceDefault.
// * serviceDefault: Default URL or identifier for service. Must be unique,
//   using the API URL is customary, but using the module name is reasonable.
//   Required, secrets cannot be stored unless associated with a service.
//
// Output:
// * username: If a user/subdomain/identifier specific to the service is
//   configurable, it can be saved as a "username". Optional.
// * secret: The secret for service. Optional.
func ConfigureSecret(
	globalConfig *config.Config,
	service string,
	serviceDefault string,
	username *string,
	secret *string, // unfortunate order dependency...
) {
	notWanted := func(out *string) bool {
		return out == nil && *out != ""
	}

	// Don't try to fetch from cred store if nothing is wanted.
	if notWanted(secret) && notWanted(username) {
		return
	}

	if service == "" {
		service = serviceDefault
	}

	if service == "" {
		return
	}

	cred, err := FetchSecret(globalConfig, service)

	if err != nil {
		logger.Log(fmt.Sprintf("Loading secret failed: %s", err.Error()))
		return
	}

	if cred == nil {
		// No secret store configued.
		return
	}

	if username != nil && *username == "" {
		*username = cred.Username
	}

	if secret != nil && *secret == "" {
		*secret = cred.Secret
	}
}

// Fetch secret for `service`. Service is customarily a URL, but can be any
// identifier uniquely used by wtf to identify the service, such as the name
// of the module.  nil is returned if the secretStore global property is not
// present or the secret is not found in that store.
func FetchSecret(globalConfig *config.Config, service string) (*Secret, error) {
	prog := newProgram(globalConfig)

	if prog == nil {
		// No secret store configured.
		return nil, nil
	}

	cred, err := client.Get(prog.runner, service)

	if err != nil {
		return nil, fmt.Errorf("get %v from %v: %w", service, prog.store, err)
	}

	return &Secret{
		Service:  cred.ServerURL,
		Secret:   cred.Secret,
		Username: cred.Username,
		Store:    prog.store,
	}, nil
}

func StoreSecret(globalConfig *config.Config, secret *Secret) error {
	prog := newProgram(globalConfig)

	if prog == nil {
		return errors.New("Cannot store secrets: wtf.secretStore is not configured")
	}

	cred := &credentials.Credentials{
		ServerURL: secret.Service,
		Username:  secret.Username,
		Secret:    secret.Secret,
	}

	// docker-credential requires a username, but it isn't necessary for
	// all services. Use a default if a username was not set.
	if cred.Username == "" {
		cred.Username = "default"
	}

	err := client.Store(prog.runner, cred)

	if err != nil {
		return fmt.Errorf("store %v: %w", prog.store, err)
	}

	return nil
}

type program struct {
	store  string
	runner client.ProgramFunc
}

func newProgram(globalConfig *config.Config) *program {
	secretStore := globalConfig.UString("wtf.secretStore", "(none)")

	if secretStore == "(none)" {
		return nil
	}

	if secretStore == "" {
		switch runtime.GOOS {
		case "windows":
			secretStore = "winrt"
		case "darwin":
			secretStore = "osxkeychain"
		default:
			secretStore = "secretservice"
		}

	}

	return &program{
		secretStore,
		client.NewShellProgramFunc("docker-credential-" + secretStore),
	}
}
