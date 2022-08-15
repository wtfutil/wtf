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

type SecretLoadParams struct {
	name         string
	globalConfig *config.Config
	service      string

	secret *string
}

// Load module secrets.
//
// The credential helpers impose this structure:
//
//	SERVICE is mapped to a SECRET and USERNAME
//
// Only SECRET is secret, SERVICE and USERNAME are not, so this
// API doesn't expose USERNAME.
//
// SERVICE was intended to be the URL of an API server, but
// for hosted services that do not have or need a configurable
// API server, its easier to just use the module name as the
// SERVICE:
//
//	   cfg.ModuleSecret(name, globalConfig, &settings.apiKey).Load()
//
//	The user will use the module name as the service, and the API key as
//	the secret, for example:
//
//	   % wtfutil save-secret circleci
//	   Secret: ...
//
// If a module (such as pihole, jenkins, or github) might have multiple
// instantiations each using a different API service (with its own unique
// API key), then the module should use the API URL to lookup the secret.
// For example, for github:
//
//	cfg.ModuleSecret(name, globalConfig, &settings.apiKey).
//	    Service(settings.baseURL).
//	    Load()
//
// The user will use the API URL as the service, and the API key as the
// secret, for example, with github configured as:
//
//	   -- config.yml
//	   mods:
//	     github:
//	       baseURL: "https://github.mycompany.com/api/v3"
//	       ...
//
//	the secret must be saved as:
//
//	   % wtfutil save-secret https://github.mycompany.com/api/v3
//	   Secret: ...
//
//	If baseURL is not set in the configuration it will be the modules
//	default, and the SERVICE will default to the module name, "github",
//	and the user must save the secret as:
//
//	   % wtfutil save-secret github
//	   Secret: ...
//
//	Ideally, the individual module documentation would describe the
//	SERVICE name to use to save the secret.
func ModuleSecret(name string, globalConfig *config.Config, secret *string) *SecretLoadParams {
	return &SecretLoadParams{
		name:         name,
		globalConfig: globalConfig,
		secret:       secret,
		service:      name, // Default the service to the module name
	}
}

func (slp *SecretLoadParams) Service(service string) *SecretLoadParams {
	if service != "" {
		slp.service = service
	}
	return slp
}

func (slp *SecretLoadParams) Load() {
	configureSecret(
		slp.globalConfig,
		slp.service,
		slp.secret,
	)
}

type Secret struct {
	Service  string
	Secret   string
	Username string
	Store    string
}

func configureSecret(
	globalConfig *config.Config,
	service string,
	secret *string,
) {
	if service == "" {
		return
	}

	if secret == nil {
		return
	}

	// Don't overwrite the secret if it was configured with yaml
	if *secret != "" {
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
		return errors.New("cannot store secrets: wtf.secretStore is not configured")
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
