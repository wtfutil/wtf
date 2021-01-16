package app

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/support"
)

// WtfAppManager handles the instances of WtfApp, ensuring that they're displayed as requested
type WtfAppManager struct {
	WtfApps []WtfApp

	config      *config.Config
	ghUser      *support.GitHubUser
	selectedIdx int
	tviewApp    *tview.Application
}

// NewAppManager creates and returns an instance of AppManager
func NewAppManager(config *config.Config, tviewApp *tview.Application) WtfAppManager {
	appMan := WtfAppManager{
		WtfApps: []WtfApp{},

		config: config,

		tviewApp: tviewApp,
	}

	appMan.tviewApp.SetBeforeDrawFunc(func(s tcell.Screen) bool {
		s.Clear()
		return false
	})

	githubAPIKey := readGitHubAPIKey(config)
	appMan.ghUser = support.NewGitHubUser(githubAPIKey)

	go func() { _ = appMan.ghUser.Load() }()

	return appMan
}

// MakeNewWtfApp creates and starts a new instance of WtfApp from a set of configuration params
func (appMan *WtfAppManager) MakeNewWtfApp(configFilePath string) {
	wtfApp := NewWtfApp(appMan.config, appMan.tviewApp, configFilePath)
	appMan.Add(wtfApp)

	wtfApp.Start()
}

/* -------------------- Exported Functions -------------------- */

// Add adds a WtfApp to the collection of apps that the AppManager manages.
// This app is then available for display onscreen.
func (appMan *WtfAppManager) Add(wtfApp WtfApp) {
	appMan.WtfApps = append(appMan.WtfApps, wtfApp)
}

// CurrentWtfApp returns the currently-displaying instance of WtfApp
func (appMan *WtfAppManager) CurrentWtfApp() (WtfApp, error) {
	appCount := len(appMan.WtfApps)

	if appCount < 1 {
		return WtfApp{}, fmt.Errorf("no wtf apps defined, cannot select current app: %d", appCount)
	}

	if appMan.selectedIdx < 0 || appMan.selectedIdx >= appCount {
		return WtfApp{}, fmt.Errorf("invalid app index selected: %d", appMan.selectedIdx)
	}

	return appMan.WtfApps[appMan.selectedIdx], nil
}

// NextWtfApp cycles the WtfApps forward by one, making the next one in the list
// the current one. If there are none after the current one, it wraps around.
func (appMan *WtfAppManager) NextWtfApp() {
	appMan.selectedIdx++

	if appMan.selectedIdx >= len(appMan.WtfApps) {
		appMan.selectedIdx = 0
	}
}

// PrevWtfApp cycles the WtfApps backwards by one, making the previous one in the
// list the current one. If there are none before the current one, it wraps around.
func (appMan *WtfAppManager) PrevWtfApp() {
	appMan.selectedIdx--

	if appMan.selectedIdx < 0 {
		appMan.selectedIdx = len(appMan.WtfApps) - 1
	}
}

// KeyboardIntercept controls all the top-level keyboard input handling.
func (appMan *WtfAppManager) KeyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	currentWtfApp, err := appMan.CurrentWtfApp()
	if err != nil {
		return nil
	}

	// These keys are global keys used by the app. Widgets should not implement these keys
	switch event.Key() {
	case tcell.KeyCtrlC:
		currentWtfApp.Stop()
		appMan.DisplayExitMessage()
	case tcell.KeyCtrlR:
		currentWtfApp.refreshAllWidgets()
		return nil
	case tcell.KeyCtrlSpace:
		fmt.Println(">> Next app")
		appMan.NextWtfApp()
		return nil
	case tcell.KeyTab:
		currentWtfApp.focusTracker.Next()
	case tcell.KeyBacktab:
		currentWtfApp.focusTracker.Prev()
		return nil
	case tcell.KeyEsc:
		currentWtfApp.focusTracker.None()
	}

	// Checks to see if any widget has been assigned the pressed key as its focus key
	if currentWtfApp.focusTracker.FocusOn(string(event.Rune())) {
		return nil
	}

	// If no specific widget has focus, then allow the key presses to fall through to the app
	if !currentWtfApp.focusTracker.IsFocused {
		switch string(event.Rune()) {
		case "/":
			return nil
		default:
		}
	}

	return event
}
