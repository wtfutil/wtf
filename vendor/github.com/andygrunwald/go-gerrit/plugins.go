package gerrit

import (
	"fmt"
)

// PluginsService contains Plugin related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html
type PluginsService struct {
	client *Client
}

// PluginInfo entity describes a plugin.
type PluginInfo struct {
	ID       string `json:"id"`
	Version  string `json:"version"`
	IndexURL string `json:"index_url,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}

// PluginInput entity describes a plugin that should be installed.
type PluginInput struct {
	URL string `json:"url"`
}

// PluginOptions specifies the different options for the ListPlugins call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#list-plugins
type PluginOptions struct {
	// All enabled that all plugins are returned (enabled and disabled).
	All bool `url:"all,omitempty"`
}

// ListPlugins lists the plugins installed on the Gerrit server.
// Only the enabled plugins are returned unless the all option is specified.
//
// To be allowed to see the installed plugins, a user must be a member of a group that is granted the 'View Plugins' capability or the 'Administrate Server' capability.
// The entries in the map are sorted by plugin ID.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#list-plugins
func (s *PluginsService) ListPlugins(opt *PluginOptions) (*map[string]PluginInfo, *Response, error) {
	u := "plugins/"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]PluginInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetPluginStatus retrieves the status of a plugin on the Gerrit server.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#get-plugin-status
func (s *PluginsService) GetPluginStatus(pluginID string) (*PluginInfo, *Response, error) {
	u := fmt.Sprintf("plugins/%s/gerrit~status", pluginID)
	return s.requestWithPluginInfoResponse("GET", u, nil)
}

// InstallPlugin installs a new plugin on the Gerrit server.
// If a plugin with the specified name already exists it is overwritten.
//
// Note: if the plugin provides its own name in the MANIFEST file, then the plugin name from the MANIFEST file has precedence over the {plugin-id} above.
//
// The plugin jar can either be sent as binary data in the request body or a URL to the plugin jar must be provided in the request body inside a PluginInput entity.
//
// As response a PluginInfo entity is returned that describes the plugin.
// If an existing plugin was overwritten the response is “200 OK”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-projects.html#set-dashboard
func (s *PluginsService) InstallPlugin(pluginID string, input *PluginInput) (*PluginInfo, *Response, error) {
	u := fmt.Sprintf("plugins/%s", pluginID)
	return s.requestWithPluginInfoResponse("PUT", u, input)
}

// EnablePlugin enables a plugin on the Gerrit server.
//
// As response a PluginInfo entity is returned that describes the plugin.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#enable-plugin
func (s *PluginsService) EnablePlugin(pluginID string) (*PluginInfo, *Response, error) {
	u := fmt.Sprintf("plugins/%s/gerrit~enable", pluginID)
	return s.requestWithPluginInfoResponse("POST", u, nil)
}

// DisablePlugin disables a plugin on the Gerrit server.
//
// As response a PluginInfo entity is returned that describes the plugin.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#disable-plugin
func (s *PluginsService) DisablePlugin(pluginID string) (*PluginInfo, *Response, error) {
	u := fmt.Sprintf("plugins/%s/gerrit~disable", pluginID)
	return s.requestWithPluginInfoResponse("POST", u, nil)
}

// ReloadPlugin reloads a plugin on the Gerrit server.
//
// As response a PluginInfo entity is returned that describes the plugin.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-plugins.html#disable-plugin
func (s *PluginsService) ReloadPlugin(pluginID string) (*PluginInfo, *Response, error) {
	u := fmt.Sprintf("plugins/%s/gerrit~reload", pluginID)
	return s.requestWithPluginInfoResponse("POST", u, nil)
}

func (s *PluginsService) requestWithPluginInfoResponse(method, u string, input interface{}) (*PluginInfo, *Response, error) {
	req, err := s.client.NewRequest(method, u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(PluginInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}
