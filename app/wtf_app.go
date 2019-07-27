package app

import (
	"log"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/maker"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

// WtfApp is the container for a collection of widgets that are all constructed from a single
// configuration file and displayed together
type WtfApp struct {
	App            *tview.Application
	Config         *config.Config
	ConfigFilePath string
	Display        *Display
	FocusTracker   wtf.FocusTracker
	IsCustomConfig bool
	Pages          *tview.Pages
	Widgets        []wtf.Wtfable
}

// NewWtfApp creates and returns an instance of WtfApp
func NewWtfApp(app *tview.Application, config *config.Config, configFilePath string, isCustom bool) *WtfApp {
	wtfApp := WtfApp{
		App:            app,
		Config:         config,
		ConfigFilePath: configFilePath,
		IsCustomConfig: isCustom,
		Pages:          tview.NewPages(),
	}

	wtfApp.Widgets = maker.MakeWidgets(wtfApp.App, wtfApp.Pages, wtfApp.Config)
	wtfApp.Display = NewDisplay(wtfApp.Widgets, wtfApp.Config)
	wtfApp.FocusTracker = wtf.NewFocusTracker(wtfApp.App, wtfApp.Widgets, wtfApp.Config)

	wtfApp.App.SetInputCapture(wtfApp.keyboardIntercept)
	wtfApp.Pages.AddPage("grid", wtfApp.Display.Grid, true, true)
	wtfApp.App.SetRoot(wtfApp.Pages, true)

	wtf.ValidateWidgets(wtfApp.Widgets)

	return &wtfApp
}

/* -------------------- Exported Functions -------------------- */

// Start initializes the app
func (wtfApp *WtfApp) Start() {
	wtfApp.scheduleWidgets()
	go wtfApp.watchForConfigChanges()
}

// Stop kills all the currently-running widgets in this app
func (wtfApp *WtfApp) Stop() {
	wtfApp.disableAllWidgets()
}

/* -------------------- Unexported Functions -------------------- */

func (wtfApp *WtfApp) disableAllWidgets() {
	for _, widget := range wtfApp.Widgets {
		widget.Disable()
	}
}

func (wtfApp *WtfApp) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	// These keys are global keys used by the app. Widgets should not implement these keys
	switch event.Key() {
	case tcell.KeyCtrlR:
		wtfApp.refreshAllWidgets()
		return nil
	case tcell.KeyTab:
		wtfApp.FocusTracker.Next()
		return nil
	case tcell.KeyBacktab:
		wtfApp.FocusTracker.Prev()
		return nil
	case tcell.KeyEsc:
		wtfApp.FocusTracker.None()
		return nil
	}

	// Checks to see if any widget has been assigned the pressed key as its focus key
	if wtfApp.FocusTracker.FocusOn(string(event.Rune())) {
		return nil
	}

	// If no specific widget has focus, then allow the key presses to fall through to the app
	if !wtfApp.FocusTracker.IsFocused {
		switch string(event.Rune()) {
		case "/":
			return nil
		}
	}

	return event
}

func (wtfApp *WtfApp) refreshAllWidgets() {
	for _, widget := range wtfApp.Widgets {
		go widget.Refresh()
	}
}

func (wtfApp *WtfApp) scheduleWidgets() {
	for _, widget := range wtfApp.Widgets {
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

				config := cfg.LoadWtfConfigFile(wtfApp.ConfigFilePath, wtfApp.IsCustomConfig)

				newApp := NewWtfApp(wtfApp.App, config, wtfApp.ConfigFilePath, wtfApp.IsCustomConfig)
				newApp.Start()
			case err := <-watch.Error:
				log.Fatalln(err)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch config file for changes.
	absPath, _ := utils.ExpandHomeDir(wtfApp.ConfigFilePath)
	if err := watch.Add(absPath); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := watch.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
