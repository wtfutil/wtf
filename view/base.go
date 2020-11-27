package view

import (
	"fmt"
	"sync"

	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/utils"
)

type Base struct {
	bordered        bool
	commonSettings  *cfg.Common
	enabled         bool
	enabledMutex    *sync.Mutex
	focusChar       string
	focusable       bool
	name            string
	quitChan        chan bool
	refreshInterval int
	refreshing      bool
}

// NewBase creates and returns an instance of the Base module, the lowest-level
// primitive module from which all others are derived
func NewBase(commonSettings *cfg.Common) *Base {
	base := &Base{
		commonSettings: commonSettings,

		bordered:        commonSettings.Bordered,
		enabled:         commonSettings.Enabled,
		enabledMutex:    &sync.Mutex{},
		focusChar:       commonSettings.FocusChar(),
		focusable:       commonSettings.Focusable,
		name:            commonSettings.Name,
		quitChan:        make(chan bool),
		refreshInterval: commonSettings.RefreshInterval,
		refreshing:      false,
	}

	return base
}

/* -------------------- Exported Functions -------------------- */

// Bordered returns whether or not this widget should be drawn with a border
func (base *Base) Bordered() bool {
	return base.bordered
}

// BorderColor returns the color that the border of this widget should be drawn in
func (base *Base) BorderColor() string {
	if base.Focusable() {
		return base.commonSettings.Colors.BorderTheme.Focusable
	}

	return base.commonSettings.Colors.BorderTheme.Unfocusable
}

func (base *Base) CommonSettings() *cfg.Common {
	return base.commonSettings
}

func (base *Base) ConfigText() string {
	return utils.HelpFromInterface(cfg.Common{})
}

func (base *Base) ContextualTitle(defaultStr string) string {
	switch {
	case defaultStr == "" && base.FocusChar() == "":
		return ""
	case defaultStr != "" && base.FocusChar() == "":
		return fmt.Sprintf(" %s ", defaultStr)
	case defaultStr == "" && base.FocusChar() != "":
		return fmt.Sprintf(" [darkgray::u]%s[::-][white] ", base.FocusChar())
	}

	return fmt.Sprintf(" %s [darkgray::u]%s[::-][white] ", defaultStr, base.FocusChar())
}

func (base *Base) Disable() {
	base.enabledMutex.Lock()
	base.enabled = false
	base.enabledMutex.Unlock()
}

func (base *Base) Disabled() bool {
	base.enabledMutex.Lock()
	result := !base.enabled
	base.enabledMutex.Unlock()
	return result
}

func (base *Base) Enabled() bool {
	base.enabledMutex.Lock()
	result := base.enabled
	base.enabledMutex.Unlock()
	return result
}

func (base *Base) Focusable() bool {
	base.enabledMutex.Lock()
	result := base.enabled && base.focusable
	base.enabledMutex.Unlock()
	return result
}

func (base *Base) FocusChar() string {
	return base.focusChar
}

func (base *Base) Name() string {
	return base.name
}

func (base *Base) QuitChan() chan bool {
	return base.quitChan
}

// Refreshing returns TRUE if the base is currently refreshing its data, FALSE if it is not
func (base *Base) Refreshing() bool {
	return base.refreshing
}

// RefreshInterval returns how often, in seconds, the base will return its data
func (base *Base) RefreshInterval() int {
	return base.refreshInterval
}

func (base *Base) SetFocusChar(char string) {
	base.focusChar = char
}

func (base *Base) Stop() {
	base.enabledMutex.Lock()
	base.enabled = false
	base.enabledMutex.Unlock()
	base.quitChan <- true
}

func (base *Base) String() string {
	return base.name
}
