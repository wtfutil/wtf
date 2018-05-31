package wtf

import (
	"fmt"
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
)

type CommandFlags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Module  string `short:"m" long:"module" optional:"yes" description:"Display info about a specific module, i.e.: 'wtf -m=todo'"`
	Version bool   `short:"v" long:"version" description:"Show Version Info"`
}

func NewCommandFlags() *CommandFlags {
	cmdFlags := CommandFlags{}
	return &cmdFlags
}

/* -------------------- Exported Functions -------------------- */

func (cmdFlags *CommandFlags) HasConfig() bool {
	return len(cmdFlags.Config) > 0
}

func (cmdFlags *CommandFlags) HasModule() bool {
	return len(cmdFlags.Module) > 0
}

func (cmdFlags *CommandFlags) Parse(version string) {
	parser := flags.NewParser(cmdFlags, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
	}

	if !cmdFlags.HasConfig() {
		homeDir, err := Home()
		if err != nil {
			os.Exit(1)
		}

		cmdFlags.Config = filepath.Join(homeDir, ".wtf", "config.yml")
	}

	if cmdFlags.Version {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}
}
