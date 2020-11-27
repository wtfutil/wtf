package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	_ "github.com/gdamore/tcell/terminfo/extended"
	"github.com/logrusorgru/aurora"
	"github.com/olebedev/config"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/support"
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
	ghUser         *support.GitHubUser
	pages          *tview.Pages
	validator      *ModuleValidator
	widgets        []wtf.Wtfable
}

// NewWtfApp creates and returns an instance of WtfApp
func NewWtfApp(tviewApp *tview.Application, config *config.Config, configFilePath string) WtfApp {
	wtfApp := WtfApp{
		TViewApp: tviewApp,

		config:         config,
		configFilePath: configFilePath,
		pages:          tview.NewPages(),
	}

	wtfApp.TViewApp.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	wtfApp.widgets = MakeWidgets(wtfApp.TViewApp, wtfApp.pages, wtfApp.config)
	wtfApp.display = NewDisplay(wtfApp.widgets, wtfApp.config)
	wtfApp.focusTracker = NewFocusTracker(wtfApp.TViewApp, wtfApp.widgets, wtfApp.config)

	githubAPIKey := readGitHubAPIKey(wtfApp.config)
	wtfApp.ghUser = support.NewGitHubUser(githubAPIKey)

	wtfApp.validator = NewModuleValidator()

	wtfApp.pages.AddPage("grid", wtfApp.display.Grid, true, true)

	wtfApp.validator.Validate(wtfApp.widgets)

	firstWidget := wtfApp.widgets[0]
	wtfApp.pages.Box.SetBackgroundColor(
		wtf.ColorFor(
			firstWidget.CommonSettings().Colors.WidgetTheme.Background,
		),
	)

	wtfApp.TViewApp.SetInputCapture(wtfApp.keyboardIntercept)
	wtfApp.TViewApp.SetRoot(wtfApp.pages, true)

	return wtfApp
}

/* -------------------- Exported Functions -------------------- */

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

	go func() { _ = wtfApp.ghUser.Load() }()
}

// Stop kills all the currently-running widgets in this app
func (wtfApp *WtfApp) Stop() {
	wtfApp.stopAllWidgets()
}

/* -------------------- Unexported Functions -------------------- */

func (wtfApp *WtfApp) stopAllWidgets() {
	for _, widget := range wtfApp.widgets {
		widget.Stop()
	}
}

func (wtfApp *WtfApp) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	// These keys are global keys used by the app. Widgets should not implement these keys
	switch event.Key() {
	case tcell.KeyCtrlC:
		wtfApp.Stop()
		wtfApp.TViewApp.Stop()
		wtfApp.DisplayExitMessage()
	case tcell.KeyCtrlR:
		wtfApp.refreshAllWidgets()
		return nil
	case tcell.KeyTab:
		wtfApp.focusTracker.Next()
	case tcell.KeyBacktab:
		wtfApp.focusTracker.Prev()
		return nil
	case tcell.KeyEsc:
		wtfApp.focusTracker.None()
	}

	// Checks to see if any widget has been assigned the pressed key as its focus key
	if wtfApp.focusTracker.FocusOn(string(event.Rune())) {
		return nil
	}

	// If no specific widget has focus, then allow the key presses to fall through to the app
	if !wtfApp.focusTracker.IsFocused {
		switch string(event.Rune()) {
		case "/":
			return nil
		default:
		}
	}

	return event
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
				newApp := NewWtfApp(wtfApp.TViewApp, config, wtfApp.configFilePath)
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
