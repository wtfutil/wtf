// Package pluggable is an attempt at dynamically loading
// user-compiled plugins.
// Might be a hacky clusterfuck, who knows? ¯\_(ツ)_/¯
package pluggable

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"

	"github.com/olebedev/config"
	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/wtf"
)

const (
	ExportedSymbol string = "WTFModule"
)

// ExternalModule represents a module which is dynamically loaded
// by wtfutil. It does its own settings initialization via Initialize(),
// most likely in a similar way to how the MakeWidget function in app/
// loads individual widgets.
type ExternalModule interface {
	// Initialize takes the following parameters:
	// - name of the module
	// - module's local config
	// - wtf's global config
	// - a *tview.Application (widget can do whatever it wants with it, including nothing)
	// - a *tview.Pages (can also be ignored if not needed by the widget)
	Initialize(
		string,
		*config.Config,
		*config.Config,
		*tview.Application,
		*tview.Pages,
	) wtf.Wtfable
}

// LoadExternalModule attempts to load a "plugin" module by way of loading a pre-compiled .so file,
// compiled by way of `go build -buildmode=plugin [...things and stuff...]`.
func LoadExternalModule(
	tviewApp *tview.Application,
	pages *tview.Pages,
	moduleName string,
	pluginConfig map[string]interface{},
	ymlConfig *config.Config,
	globalConfig *config.Config,
) wtf.Wtfable {
	pluginSoPath, _ := pluginConfig["path"].(string)

	validPath, isValid := validatePath(pluginSoPath)
	if !isValid {
		return nil
	}

	// Attempt to load the .so file at the path
	p, err := plugin.Open(validPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Attempt to find an exported variable in the loaded .so by name
	// (specifically, "WTFModule")
	pl, err := p.Lookup(ExportedSymbol)
	if err != nil {
		fmt.Printf("Symbol %s not found in file %s; did you remember to export it?\n", ExportedSymbol, validPath)
		return nil
	}

	// Check that the exported symbol is, in fact, an ExternalModule.
	plug, ok := pl.(ExternalModule)
	if !ok {
		fmt.Printf("Module at path %s is not a valid ExternalModule; consider writing a valid one.\n", validPath)
		return nil
	}

	// finally, initialize that bad boi
	return plug.Initialize(moduleName, ymlConfig, globalConfig, tviewApp, pages)
}

func validatePath(p string) (validFilepath string, isValid bool) {
	// let's program defensively
	if p == "" {
		return "", false
	}
	// make sure default return values are set
	validFilepath = ""
	isValid = false

	cleaned, err := filepath.Abs(p)
	if err != nil {
		return
	}

	info, err := os.Stat(cleaned)
	if err != nil {
		return
	}

	validFilepath = cleaned
	isValid = !info.IsDir()

	return
}
