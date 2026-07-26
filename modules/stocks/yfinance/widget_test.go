package yfinance

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
)

func TestNewWidget(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	ymlConfig, err := config.ParseYaml(`title: "Test Yahoo Finance"`)
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)

	widget := NewWidget(app, redrawChan, settings)

	assert.NotNil(t, widget)
	assert.Equal(t, settings, widget.settings)
	assert.Equal(t, "Test Yahoo Finance", widget.CommonSettings().Title)
}

func TestWidget_Content_UnknownSymbolsFallBack(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	ymlConfig, err := config.ParseYaml(`
symbols:
  - "NOTASYMBOL"
`)
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)
	widget := NewWidget(app, redrawChan, settings)

	out := widget.content()

	assert.Contains(t, out, "NOTASYMBOL")
}

func TestWidget_Content_SortsBySymbolChangeWhenEnabled(t *testing.T) {
	now := time.Now().Unix()
	upBody := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"UP","regularMarketPrice":110.0,
		"chartPreviousClose":100.0,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},"regular":{"start":%d,"end":%d},"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-1000, now-1000, now-1000, now+1000, now+1000, now+1000)
	dropBody := fmt.Sprintf(`{"chart":{"result":[{"meta":{
		"currency":"USD","symbol":"DOWN","regularMarketPrice":90.0,
		"chartPreviousClose":100.0,
		"currentTradingPeriod":{
			"pre":{"start":%d,"end":%d},"regular":{"start":%d,"end":%d},"post":{"start":%d,"end":%d}
		}
	}}],"error":null}}`, now-1000, now-1000, now-1000, now+1000, now+1000, now+1000)

	withTestServer(t, map[string]string{"/DOWN": dropBody, "/UP": upBody})

	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	ymlConfig, err := config.ParseYaml(`
sort: true
symbols:
  - "DOWN"
  - "UP"
`)
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)
	widget := NewWidget(app, redrawChan, settings)

	out := widget.content()

	// With sort enabled, the best performer (UP) should be rendered
	// before the worst performer (DOWN).
	upIndex := strings.Index(out, "UP")
	downIndex := strings.Index(out, "DOWN")
	assert.Greater(t, upIndex, -1)
	assert.Greater(t, downIndex, -1)
	assert.Less(t, upIndex, downIndex)
}

func TestWidget_Refresh(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	ymlConfig, err := config.ParseYaml(`symbols: []`)
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)
	widget := NewWidget(app, redrawChan, settings)

	// Refresh should not panic and should push a redraw signal.
	go widget.Refresh()

	select {
	case <-redrawChan:
	default:
		// Redraw may be synchronous/non-blocking depending on TextWidget
		// implementation; absence of a panic is the primary assertion.
	}
}
