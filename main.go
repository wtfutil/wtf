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

	"github.com/logrusorgru/aurora/v4"
	"github.com/pkg/profile"

	"github.com/wtfutil/wtf/app"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

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

	flags.RenderIf(config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	openFileUtil := config.UString("wtf.openFileUtil", "open")
	openURLUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
	utils.Init(openFileUtil, openURLUtil)

	/* Initialize the App Manager */
	appMan := app.NewAppManager()
	appMan.MakeNewWtfApp(config, flags.Config)

	currentApp, err := appMan.Current()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}

	err = currentApp.Execute()
	if err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}
