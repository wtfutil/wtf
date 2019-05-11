package flags

import (
	"fmt"
	"os"
	"path/filepath"

	goFlags "github.com/jessevdk/go-flags"
	"github.com/wtfutil/wtf/utils"
)

type Flags struct {
	Config  string `short:"c" long:"config" optional:"yes" description:"Path to config file"`
	Profile bool   `short:"p" long:"profile" optional:"yes" description:"Profile application memory usage"`
	Version bool   `short:"v" long:"version" description:"Show version info"`
}

func NewFlags() *Flags {
	flags := Flags{}
	return &flags
}

/* -------------------- Exported Functions -------------------- */

func (flags *Flags) ConfigFilePath() string {
	return flags.Config
}

func (flags *Flags) Display(version string) {
	if flags.HasVersion() {
		fmt.Println(version)
		os.Exit(0)
	}
}

func (flags *Flags) HasConfig() bool {
	return len(flags.Config) > 0
}

func (flags *Flags) HasVersion() bool {
	return flags.Version == true
}

func (flags *Flags) Parse() {
	parser := goFlags.NewParser(flags, goFlags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*goFlags.Error); ok && flagsErr.Type == goFlags.ErrHelp {
			os.Exit(0)
		}
	}

	// If no config file is explicitly passed in as a param,
	// set the flag to the default config file
	if !flags.HasConfig() {
		homeDir, err := utils.Home()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		flags.Config = filepath.Join(homeDir, ".config", "wtf", "config.yml")
	}
}
