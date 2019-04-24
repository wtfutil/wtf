package main

// Generators
// To generate the skeleton for a new TextWidget use 'WTF_WIDGET_NAME=MySuperAwesomeWidget go generate -run=text
//go:generate -command text go run generator/textwidget.go
//go:generate text

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/pkg/profile"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/maker"
	"github.com/wtfutil/wtf/wtf"
)

var focusTracker wtf.FocusTracker
var runningWidgets []wtf.Wtfable

// Config parses the config.yml file and makes available the settings within
var Config *config.Config

var (
	commit  = "dev"
	date    = "dev"
	version = "dev"
)

/* -------------------- Functions -------------------- */

func disableAllWidgets(widgets []wtf.Wtfable) {
	for _, widget := range widgets {
		widget.Disable()
	}
}

func keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyCtrlR:
		refreshAllWidgets(runningWidgets)
	case tcell.KeyTab:
		focusTracker.Next()
	case tcell.KeyBacktab:
		focusTracker.Prev()
	case tcell.KeyEsc:
		focusTracker.None()
	}

	if focusTracker.FocusOn(string(event.Rune())) {
		return nil
	}

	return event
}

func loadConfigFile(filePath string) {
	Config = cfg.LoadConfigFile(filePath)
	wtf.Config = Config
}

func refreshAllWidgets(widgets []wtf.Wtfable) {
	for _, widget := range widgets {
		go widget.Refresh()
	}
}

func setTerm() {
	err := os.Setenv("TERM", Config.UString("wtf.term", os.Getenv("TERM")))
	if err != nil {
		return
	}
}

func watchForConfigChanges(app *tview.Application, configFilePath string, grid *tview.Grid, pages *tview.Pages) {
	watch := watcher.New()
	absPath, _ := wtf.ExpandHomeDir(configFilePath)

	// Notify write events
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				// Disable all widgets to stop scheduler goroutines and remove widgets from memory
				disableAllWidgets(runningWidgets)

				loadConfigFile(absPath)

				widgets := maker.MakeWidgets(app, pages, Config)
				wtf.ValidateWidgets(widgets)
				runningWidgets = widgets

				focusTracker = wtf.NewFocusTracker(app, widgets)

				display := wtf.NewDisplay(widgets)
				pages.AddPage("grid", display.Grid, true, true)
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	if err := watch.Add(absPath); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}

func enableAppRefresh(app *tview.Application) {
	defaultInterval := 100

	refreshInterval := Config.UInt("wtf.refreshInterval", defaultInterval)
	if refreshInterval < defaultInterval {
		refreshInterval = defaultInterval
	}

	interval := time.Duration(refreshInterval) * time.Millisecond

	tick := time.NewTicker(interval)
	quit := make(chan struct{})

	for {
		select {
		case <-tick.C:
			app.Draw()
		case <-quit:
			tick.Stop()
			return
		}
	}
}

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flags := flags.NewFlags()
	flags.Parse()
	flags.Display(version)

	cfg.MigrateOldConfig()
	cfg.CreateConfigDir()
	cfg.CreateConfigFile()
	loadConfigFile(flags.ConfigFilePath())

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	setTerm()

	app := tview.NewApplication()
	pages := tview.NewPages()

	widgets := maker.MakeWidgets(app, pages, Config)
	wtf.ValidateWidgets(widgets)
	runningWidgets = widgets

	focusTracker = wtf.NewFocusTracker(app, widgets)

	display := wtf.NewDisplay(widgets)
	pages.AddPage("grid", display.Grid, true, true)

	app.SetInputCapture(keyboardIntercept)

	go watchForConfigChanges(app, flags.Config, display.Grid, pages)
	go enableAppRefresh(app)

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
