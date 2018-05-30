package wtf

import (
	"fmt"
	"os"
	"path/filepath"

	flags "github.com/jessevdk/go-flags"
)

type CommandFlags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Version bool   `short:"v" long:"version" description:"Show Version Info"`
}

func NewCommandFlags() *CommandFlags {
	cmdFlags := CommandFlags{}
	return &cmdFlags
}

/* -------------------- Exported Functions -------------------- */

func (cmdFlags *CommandFlags) Parse(version string) {
	parser := flags.NewParser(cmdFlags, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		}
	}

	if len(cmdFlags.Config) == 0 {
		homeDir, err := Home()
		if err != nil {
			os.Exit(1)
		}

		cmdFlags.Config = filepath.Join(homeDir, ".wtf", "config.yml")
		fmt.Printf(">> A: %s\n", cmdFlags.Config)
	}

	if cmdFlags.Version {
		fmt.Printf("Version: %s\n", version)
		os.Exit(0)
	}
}
