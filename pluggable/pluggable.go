// Package pluggable is an attempt at dynamically loading
// user-compiled plugins.
// Might be a hacky clusterfuck, who knows? ¯\_(ツ)_/¯
package pluggable

import (
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/olebedev/config"
	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/wtf"
)

const (
	RPCPluginName string = "wtf_module"
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

// LoadExternalModule attempts to load a "plugin" module by way of hacky net/rpc bullcrap.
func LoadExternalModule(
	tviewApp *tview.Application,
	pages *tview.Pages,
	moduleName string,
	pluginConfig map[string]interface{},
	ymlConfig *config.Config,
	globalConfig *config.Config,
) wtf.Wtfable {
	cmdPath := pluginConfig["path"].(string)

	client := plugin.NewClient(&plugin.ClientConfig{
		Plugins: map[string]plugin.Plugin{
			RPCPluginName: &RPCModuleLoader{},
		},
		Cmd: exec.Command("sh", "-c", cmdPath),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}

	rawMod, err := rpcClient.Dispense(RPCPluginName)
	if err != nil {
		panic(err)
	}

	modLoader := rawMod.(ModuleLoader)
	externalMod, err := modLoader.Module()
	if err != nil {
		panic(err)
	}

	return externalMod.Initialize(moduleName, ymlConfig, globalConfig, tviewApp, pages)
}
