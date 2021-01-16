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
	"github.com/pkg/profile"
	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/app"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

var (
	date    = "dev"
	version = "dev"
)

var appMan app.WtfAppManager
var tviewApp *tview.Application

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()

	// Load the configuration file
	cfg.Initialize(flags.HasCustomConfig())
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath())

	wtf.SetTerminal(config)

	flags.RenderIf(version, date, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openFileUtil := config.UString("wtf.openFileUtil", "open")
	openURLUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
	utils.Init(openFileUtil, openURLUtil)

	/* Initialize the App Manager */
	tviewApp = tview.NewApplication()
	appMan = app.NewAppManager(config, tviewApp)
	appMan.MakeNewWtfApp(flags.Config)
	tviewApp.SetInputCapture(appMan.KeyboardIntercept)

	currentApp, err := appMan.CurrentWtfApp()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}

	currentApp.Run()
}
