package app

import "errors"

// WtfAppManager handles the instances of WtfApp, ensuring that they're displayed as requested
type WtfAppManager struct {
	WtfApps []*WtfApp

	selected int
}

// NewAppManager creates and returns an instance of AppManager
func NewAppManager() WtfAppManager {
	appMan := WtfAppManager{
		selected: 0,
	}

	return appMan
}

// AddApp adds a WtfApp to the collection of apps that the AppManager manages.
// This app is then available for display onscreen.
func (appMan *WtfAppManager) AddApp(wtfApp *WtfApp) error {
	appMan.WtfApps = append(appMan.WtfApps, wtfApp)
	return nil
}

// Current returns the currently-displaying instance of WtfApp
func (appMan *WtfAppManager) Current() (*WtfApp, error) {
	if appMan.selected < 0 || appMan.selected > (len(appMan.WtfApps)-1) {
		return nil, errors.New("invalid app index selected")
	}

	return appMan.WtfApps[appMan.selected], nil
}

// Next cycles the WtfApps forward by one, making the next one in the list
// the current one. If there are none after the current one, it wraps around.
func (appMan *WtfAppManager) Next() (*WtfApp, error) {
	appMan.selected++
	if appMan.selected >= len(appMan.WtfApps) {
		appMan.selected = 0
	}

	return appMan.Current()
}

// Prev cycles the WtfApps backwards by one, making the previous one in the
// list the current one. If there are none before the current one, it wraps around.
func (appMan *WtfAppManager) Prev() (*WtfApp, error) {
	appMan.selected--
	if appMan.selected < 0 {
		appMan.selected = len(appMan.WtfApps) - 1
	}

	return appMan.Current()
}
