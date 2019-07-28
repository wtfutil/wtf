package main

// Generators
// To generate the skeleton for a new TextWidget use 'WTF_WIDGET_NAME=MySuperAwesomeWidget go generate -run=text
//go:generate -command text go run generator/textwidget.go
//go:generate text

import (
	"fmt"
	"log"
	"os"

	"github.com/logrusorgru/aurora"
	"github.com/olebedev/config"
	"github.com/pkg/profile"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/app"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/wtf"
)

var tviewApp *tview.Application

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

/* -------------------- Functions -------------------- */

func setTerm(config *config.Config) {
	term := config.UString("wtf.term", os.Getenv("TERM"))
	err := os.Setenv("TERM", term)
	if err != nil {
		fmt.Printf("\n%s Failed to set $TERM to %s.\n", aurora.Red("ERROR"), aurora.Yellow(term))
		os.Exit(1)
	}
}

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Manage the configuration directories and config file
	cfg.Initialize()

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	// Load the configuration file
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath(), flags.HasCustomConfig())
	flags.RenderIf(version, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	wtf.Init(config.UString("wtf.openFileUtil", "open"))

	setTerm(config)

	// Build the application
	tviewApp = tview.NewApplication()
	wtfApp := app.NewWtfApp(tviewApp, config, flags.Config, flags.HasCustomConfig())
	wtfApp.Start()

	if err := tviewApp.Run(); err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}
