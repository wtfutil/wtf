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
	"github.com/wtfutil/wtf/utils"
)

var tviewApp *tview.Application

var (
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

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	hasCustom := flags.HasCustomConfig()
	cfg.Initialize(hasCustom)

	// Load the configuration file
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath())
	flags.RenderIf(version, date, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openUrlUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
	utils.Init(config.UString("wtf.openFileUtil", "open"), openUrlUtil)

	setTerm(config)

	// Build the application
	tviewApp = tview.NewApplication()
	wtfApp := app.NewWtfApp(tviewApp, config, flags.Config)
	wtfApp.Start()

	if err := tviewApp.Run(); err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}
