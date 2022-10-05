package flags

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/chzyer/readline"
	goFlags "github.com/jessevdk/go-flags"
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/help"
)

// Flags is the container for command line flag data
type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Module  string `short:"m" long:"module" optional:"yes" description:"Display info about a specific module, i.e.: 'wtfutil -m=todo'"`
	Profile bool   `short:"p" long:"profile" optional:"yes" description:"Profile application memory usage"`
	Version bool   `short:"v" long:"version" description:"Show version info"`
	// Work-around go-flags misfeatures. If any sub-command is defined
	// then `wtf` (no sub-commands, the common usage), is warned about.
	Opt struct {
		Cmd  string   `positional-arg-name:"command"`
		Args []string `positional-arg-name:"args"`
	} `positional-args:"yes"`

	hasCustom bool
}

var EXTRA = `
Commands:
  save-secret <service>
    service      Service URL or module name of secret.
  Save a secret into the secret store. The secret will be prompted for.
  Requires wtf.secretStore to be configured.  See individual modules for
  information on what service and secret means for their configuration,
  not all modules use secrets.
`

// NewFlags creates an instance of Flags
func NewFlags() *Flags {
	flags := Flags{}
	return &flags
}

/* -------------------- Exported Functions -------------------- */

// ConfigFilePath returns the path to the currently-loaded config file
func (flags *Flags) ConfigFilePath() string {
	return flags.Config
}

// RenderIf displays special-case information based on the flags passed
// in, if any flags were passed in
func (flags *Flags) RenderIf(config *config.Config) {
	if flags.HasModule() {
		help.Display(flags.Module, config)
		os.Exit(0)
	}

	if flags.HasVersion() {
		info, _ := debug.ReadBuildInfo()
		version := "dev"
		date := "now"
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				version = setting.Value
			} else if setting.Key == "vcs.time" {
				date = setting.Value
			}
		}
		fmt.Printf("%s (%s)\n", version, date)
		os.Exit(0)
	}

	if flags.Opt.Cmd == "" {
		return
	}

	switch cmd := flags.Opt.Cmd; cmd {
	case "save-secret":
		var service, secret string
		args := flags.Opt.Args

		if len(args) < 1 || args[0] == "" {
			fmt.Fprintf(os.Stderr, "save-secret: service required, see `%s --help`\n", os.Args[0])
			os.Exit(1)
		}

		service = args[0]

		if len(args) > 1 {
			fmt.Fprintf(os.Stderr, "save-secret: too many arguments, see `%s --help`\n", os.Args[0])
			os.Exit(1)
		}

		b, err := readline.Password("Secret: ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		secret = string(b)
		secret = strings.TrimSpace(secret)

		if secret == "" {
			fmt.Fprintf(os.Stderr, "save-secret: secret required, see `%s --help`\n", os.Args[0])
			os.Exit(1)
		}

		err = cfg.StoreSecret(config, &cfg.Secret{
			Service:  service,
			Secret:   secret,
			Username: "default",
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Saving secret for service %q: %s\n", service, err.Error())
			os.Exit(1)
		}

		fmt.Printf("Saved secret for service %q\n", service)
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "Command `%s` is not supported, try `%s --help`\n", cmd, os.Args[0])
		os.Exit(1)
	}
}

// HasCustomConfig returns TRUE if a config path was passed in, FALSE if one was not
func (flags *Flags) HasCustomConfig() bool {
	return flags.hasCustom
}

// HasModule returns TRUE if a module name was passed in, FALSE if one was not
func (flags *Flags) HasModule() bool {
	return len(flags.Module) > 0
}

// HasVersion returns TRUE if the version flag was passed in, FALSE if it was not
func (flags *Flags) HasVersion() bool {
	return flags.Version
}

// Parse parses the incoming flags
func (flags *Flags) Parse() {
	parser := goFlags.NewParser(flags, goFlags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*goFlags.Error); ok && flagsErr.Type == goFlags.ErrHelp {
			fmt.Println(EXTRA)
			os.Exit(0)
		}
	}

	// If we have a custom config, then we're done parsing parameters, we don't need to
	// generate the default value
	flags.hasCustom = (len(flags.Config) > 0)
	if flags.hasCustom {
		return
	}

	// If no config file is explicitly passed in as a param then set the flag to the default config file
	configDir, err := cfg.WtfConfigDir()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	flags.Config = filepath.Join(configDir, "config.yml")
}
