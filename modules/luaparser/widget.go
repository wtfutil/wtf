package luaparser

import (
	"errors"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	lua "github.com/yuin/gopher-lua"
)

const (
	errUnconvertableLuaString = "could not convert output to Lua string"
	errUndefinedLuaFile       = "no lua file defined in configuration"
)

// Widget is the container for the functionality of this module
type Widget struct {
	L *lua.LState

	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of widget
func NewWidget(tviewApp *tview.Application, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, nil, settings.Common),

		settings: settings,
	}

	widget.View.SetWordWrap(false)

	widget.L = lua.NewState()
	filePath, err := utils.ExpandHomeDir(widget.settings.filePath)
	if err != nil {
		return nil
	}

	if err = widget.L.DoFile(filePath); err != nil {
		return nil
	}

	return widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh redraws the widget content with new data
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	err := widget.validate()
	if err != nil {
		return widget.CommonSettings().Title, err.Error(), true
	}

	content, err := widget.parseLua()
	if err != nil {
		return widget.CommonSettings().Title, err.Error(), true
	}

	return widget.CommonSettings().Title, content, true
}

func (widget *Widget) parseLua() (string, error) {
	if err := widget.L.CallByParam(lua.P{
		Fn:      widget.L.GetGlobal("main"), // execute the Lua function called "main"
		NRet:    1,                          // expect one return value
		Protect: true,                       // return an error or panic?
	}); err != nil {
		return "", err
	}

	// Pop the last value off the stack (presumably that's the return value
	// from the function executed above)
	if output, ok := widget.L.Get(-1).(*lua.LTable); ok {
		// wid := output.RawGetString("widget").(*lua.LTable)
		// title := wid.RawGetString("widget.title").String()
		// return title, nil

		return output.RawGetString("time").String(), nil
	}

	return "", errors.New(errUnconvertableLuaString)
}

func (widget *Widget) validate() error {
	if widget.settings.filePath == "" {
		return errors.New(errUndefinedLuaFile)
	}

	return nil
}
