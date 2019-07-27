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
	"github.com/logrusorgru/aurora"
	"github.com/olebedev/config"
	"github.com/pkg/profile"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/flags"
	"github.com/wtfutil/wtf/maker"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

var focusTracker wtf.FocusTracker
var runningWidgets []wtf.Wtfable

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
	// These keys are global keys used by the app. Widgets should not implement these keys
	switch event.Key() {
	case tcell.KeyCtrlR:
		refreshAllWidgets(runningWidgets)
		return nil
	case tcell.KeyTab:
		focusTracker.Next()
		return nil
	case tcell.KeyBacktab:
		focusTracker.Prev()
		return nil
	case tcell.KeyEsc:
		focusTracker.None()
		return nil
	}

	// This function checks to see if any widget has been assigned the pressed key as its
	// focus key
	if focusTracker.FocusOn(string(event.Rune())) {
		return nil
	}

	// If no specific widget has focus, then allow the key presses to fall through to the app
	if !focusTracker.IsFocused {
		switch string(event.Rune()) {
		case "/":
			return nil
		}
	}

	return event
}

func refreshAllWidgets(widgets []wtf.Wtfable) {
	for _, widget := range widgets {
		go widget.Refresh()
	}
}

func setTerm(config *config.Config) {
	term := config.UString("wtf.term", os.Getenv("TERM"))
	err := os.Setenv("TERM", term)
	if err != nil {
		fmt.Printf("\n%s Failed to set $TERM to %s.\n", aurora.Red("ERROR"), aurora.Yellow(term))
		os.Exit(1)
	}
}

func watchForConfigChanges(app *tview.Application, configFilePath string, isCustomConfig bool, grid *tview.Grid, pages *tview.Pages) {
	watch := watcher.New()
	absPath, _ := utils.ExpandHomeDir(configFilePath)

	// Notify write events
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				// Disable all widgets to stop scheduler goroutines and remove widgets from memory
				disableAllWidgets(runningWidgets)

				config := cfg.LoadWtfConfigFile(absPath, false)

				widgets := maker.MakeWidgets(app, pages, config)
				runningWidgets = widgets

				wtf.ValidateWidgets(widgets)

				focusTracker = wtf.NewFocusTracker(app, widgets, config)

				display := wtf.NewDisplay(widgets, config)
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

/* -------------------- Main -------------------- */

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Manage the configuration directories and file
	cfg.Initialize()

	// Parse and handle flags
	flags := flags.NewFlags()
	flags.Parse()
	config := cfg.LoadWtfConfigFile(flags.ConfigFilePath(), flags.HasCustomConfig())
	flags.RenderIf(version, config)

	if flags.Profile {
		defer profile.Start(profile.MemProfile).Stop()
	}

	setTerm(config)

	app := tview.NewApplication()
	pages := tview.NewPages()

	widgets := maker.MakeWidgets(app, pages, config)
	runningWidgets = widgets

	wtf.ValidateWidgets(widgets)

	focusTracker = wtf.NewFocusTracker(app, widgets, config)

	display := wtf.NewDisplay(widgets, config)
	pages.AddPage("grid", display.Grid, true, true)

	app.SetInputCapture(keyboardIntercept)

	go watchForConfigChanges(app, flags.Config, flags.HasCustomConfig(), display.Grid, pages)

	wtf.Init(config.UString("wtf.openFileUtil", "open"))

	if err := app.SetRoot(pages, true).Run(); err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}
