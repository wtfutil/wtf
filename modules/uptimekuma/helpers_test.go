package uptimekuma

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

// makeTestCommon creates a minimal Common settings for unit tests.
func makeTestCommon() *cfg.Common {
	return &cfg.Common{
		Colors: cfg.ColorTheme{
			TextTheme: cfg.TextTheme{
				Text: "white",
			},
		},
		Title:   "Uptime Kuma",
		Enabled: true,
	}
}

// makeTestWidget creates a Widget with a real TextWidget for integration tests.
func makeTestWidget(url string) *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	common := makeTestCommon()
	settings := &Settings{
		common: common,
		url:    url,
	}
	return NewWidget(app, redrawChan, nil, settings)
}
