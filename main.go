package main

import (
	"fmt"
	"log"
	"os"

	// Blank import of tzdata embeds the timezone database to allow Windows hosts to find timezone
	// information even if the timezone database is not available on the local system. See release
	// notes at https://golang.org/doc/go1.15#time/tzdata for details. This prevents "no timezone
	// data available" errors in clocks module.
	_ "time/tzdata"

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

func makeWtfApp(config *config.Config, flagConfig string) app.WtfApp {
	tviewApp = tview.NewApplication()
	wtfApp := app.NewWtfApp(tviewApp, config, flagConfig)
	wtfApp.Start()

	return wtfApp
}

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	// Load the configuration file
	cfg.Initialize(flags.HasCustomConfig())
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath())
	setTerm(config)

	flags.RenderIf(version, date, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openFileUtil := config.UString("wtf.openFileUtil", "open")
	openURLUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
	utils.Init(openFileUtil, openURLUtil)

	/* Initialize the App Manager */

	wtfApp := makeWtfApp(config, flags.Config)

	appMan := app.NewAppManager()
	appMan.Add(&wtfApp)

	currentApp, err := appMan.Current()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}

	currentApp.Run()
}
