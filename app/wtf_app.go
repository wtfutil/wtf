package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"

	// Bring in the extended set of terminfo definitions
	_ "github.com/gdamore/tcell/terminfo/extended"

	"github.com/logrusorgru/aurora"
	"github.com/olebedev/config"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

// WtfApp is the container for a collection of widgets that are all constructed from a single
// configuration file and displayed together
type WtfApp struct {
	TViewApp *tview.Application

	config         *config.Config
	configFilePath string
	display        *Display
	focusTracker   FocusTracker
	pages          *tview.Pages
	validator      *ModuleValidator
	widgets        []wtf.Wtfable
}

// NewWtfApp creates and returns an instance of WtfApp
func NewWtfApp(config *config.Config, tviewApp *tview.Application, configFilePath string) WtfApp {
	wtfApp := WtfApp{
		TViewApp: tviewApp,

		config:         config,
		configFilePath: configFilePath,
		pages:          tview.NewPages(),
	}

	wtfApp.widgets = MakeWidgets(wtfApp.TViewApp, wtfApp.pages, wtfApp.config)
	wtfApp.display = NewDisplay(wtfApp.widgets, wtfApp.config)
	wtfApp.focusTracker = NewFocusTracker(wtfApp.TViewApp, wtfApp.widgets, wtfApp.config)
	wtfApp.validator = NewModuleValidator()

	wtfApp.pages.AddPage("grid", wtfApp.display.Grid, true, true)

	wtfApp.validator.Validate(wtfApp.widgets)

	wtfApp.setBackgroundColor()

	wtfApp.TViewApp.SetRoot(wtfApp.pages, true)

	return wtfApp
}

/* -------------------- Exported Functions -------------------- */

// FirstWidget returns the first wiget in the set of widgets, or
// an error if there are none
func (wtfApp *WtfApp) FirstWidget() (wtf.Wtfable, error) {
	if len(wtfApp.widgets) < 1 {
		return nil, errors.New("cannot get first widget. no widgets defined")
	}

	return wtfApp.widgets[0], nil
}

// Run starts the underlying tview app
func (wtfApp *WtfApp) Run() {
	if err := wtfApp.TViewApp.Run(); err != nil {
		fmt.Printf("\n%s %v\n", aurora.Red("ERROR"), err)
		os.Exit(1)
	}
}

// Start initializes the app
func (wtfApp *WtfApp) Start() {
	go wtfApp.scheduleWidgets()
	go wtfApp.watchForConfigChanges()
}

// Stop kills all the currently-running widgets in this app
func (wtfApp *WtfApp) Stop() {
	wtfApp.stopAllWidgets()
	wtfApp.TViewApp.Stop()
}

/* -------------------- Unexported Functions -------------------- */

func (wtfApp *WtfApp) setBackgroundColor() {
	var bgColor tcell.Color

	firstWidget, err := wtfApp.FirstWidget()
	if err != nil {
		bgColor = wtf.ColorFor("black")
	} else {
		bgColor = wtf.ColorFor(firstWidget.CommonSettings().Colors.WidgetTheme.Background)
	}

	wtfApp.pages.Box.SetBackgroundColor(bgColor)
}

func (wtfApp *WtfApp) stopAllWidgets() {
	for _, widget := range wtfApp.widgets {
		widget.Stop()
	}
}

func (wtfApp *WtfApp) refreshAllWidgets() {
	for _, widget := range wtfApp.widgets {
		go widget.Refresh()
	}
}

func (wtfApp *WtfApp) scheduleWidgets() {
	for _, widget := range wtfApp.widgets {
		go Schedule(widget)
	}
}

func (wtfApp *WtfApp) watchForConfigChanges() {
	watch := watcher.New()

	// Notify write events
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				wtfApp.Stop()

				config := cfg.LoadWtfConfigFile(wtfApp.configFilePath)
				newApp := NewWtfApp(config, wtfApp.TViewApp, wtfApp.configFilePath)
				openURLUtil := utils.ToStrs(config.UList("wtf.openUrlUtil", []interface{}{}))
				utils.Init(config.UString("wtf.openFileUtil", "open"), openURLUtil)

				newApp.Start()
			case err := <-watch.Error:
				if err == watcher.ErrWatchedFileDeleted {
					// Usually happens because the watcher looks for the file as the OS is updating it
					continue
				}
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	absPath, _ := utils.ExpandHomeDir(wtfApp.configFilePath)
	if err := watch.Add(absPath); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
