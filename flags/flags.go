package flags

import (
	"fmt"
	"os"
	"path/filepath"

	goFlags "github.com/jessevdk/go-flags"
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/help"
	"github.com/wtfutil/wtf/utils"
)

// Flags is the container for command line flag data
type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Module  string `short:"m" long:"module" optional:"yes" description:"Display info about a specific module, i.e.: 'wtf -m=todo'"`
	Profile bool   `short:"p" long:"profile" optional:"yes" description:"Profile application memory usage"`
	Version bool   `short:"v" long:"version" description:"Show version info"`

	hasCustom bool
}

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
func (flags *Flags) RenderIf(version string, config *config.Config) {
	if flags.HasModule() {
		help.Display(flags.Module, config)
		os.Exit(0)
	}

	if flags.HasVersion() {
		fmt.Println(version)
		os.Exit(0)
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
	return flags.Version == true
}

// Parse parses the incoming flags
func (flags *Flags) Parse() {
	parser := goFlags.NewParser(flags, goFlags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*goFlags.Error); ok && flagsErr.Type == goFlags.ErrHelp {
			os.Exit(0)
		}
	}

	// If we have a custom config, then we're done parsing parameters, we don't need to
	// generate the default value
	flags.hasCustom = (len(flags.Config) > 0)
	if flags.hasCustom == true {
		return
	}

	// If no config file is explicitly passed in as a param then set the flag to the default config file
	homeDir, err := utils.Home()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	flags.Config = filepath.Join(homeDir, ".config", "wtf", "config.yml")
}
