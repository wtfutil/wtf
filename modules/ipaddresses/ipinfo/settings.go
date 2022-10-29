package ipinfo

import (
	"fmt"
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
	log "github.com/wtfutil/wtf/logger"
)

const (
	defaultFocusable                 = false
	defaultTitle                     = "IPInfo"
	ipV4             protocolVersion = "v4"
	ipV6             protocolVersion = "v6"
	auto             protocolVersion = "auto"
)

type protocolVersion string

func (pv protocolVersion) String() string {
	switch pv {
	case ipV4:
		return "v4"
	case ipV6:
		return "v6"
	default:
		return "auto"
	}
}

func newProtocolVersion(str string) (protocolVersion, error) {
	switch str {
	case "v4":
		return ipV4, nil
	case "v6":
		return ipV6, nil
	case "auto":
		return auto, nil
	default:
		return "", fmt.Errorf("%s module: Unsupported protocol version: '%s'", defaultTitle, str)
	}
}

type Settings struct {
	*cfg.Common

	apiToken        string          `help:"An api token" optional:"true"`
	protocolVersion protocolVersion `help:"IP protocol version to display. Possible options are: 'v4' to show only IpV4 address, 'v6' to show only IpV6 address and 'auto' (default) to show the address preferred by OS." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {
	settings := Settings{
		Common: cfg.NewCommonSettingsFromModule(name, defaultTitle, defaultFocusable, ymlConfig, globalConfig),

		apiToken:        ymlConfig.UString("apiToken", ""),
		protocolVersion: auto,
	}

	pv, err := newProtocolVersion(ymlConfig.UString("protocolVersion", auto.String()))
	if err != nil {
		log.Log(err.Error())
		log.Log(fmt.Sprintf("%s module: Use '%s' protocol version as a default", defaultTitle, auto))
	} else {
		settings.protocolVersion = pv
	}

	settings.SetDocumentationPath("ipaddress/ipinfo")

	return &settings
}
