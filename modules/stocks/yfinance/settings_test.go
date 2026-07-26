package yfinance

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	ymlConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)

	assert.NotNil(t, settings)
	assert.Equal(t, defaultTitle, settings.common.Title)
	assert.Equal(t, "greenyellow", settings.colors.bigup)
	assert.Equal(t, "green", settings.colors.up)
	assert.Equal(t, "firebrick", settings.colors.drop)
	assert.Equal(t, "red", settings.colors.bigdrop)
	assert.False(t, settings.sort)
	assert.Empty(t, settings.symbols)
}

func TestNewSettingsFromYAML_CustomValues(t *testing.T) {
	ymlConfig, err := config.ParseYaml(`
title: "My Watchlist"
refreshInterval: 30
sort: true
colors:
  bigup: "blue"
  up: "cyan"
  drop: "orange"
  bigdrop: "magenta"
symbols:
  - "AMD"
  - "^NSEI"
  - "EURUSD=X"
`)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("yfinance", ymlConfig, globalConfig)

	assert.NotNil(t, settings)
	assert.Equal(t, "My Watchlist", settings.common.Title)
	assert.True(t, settings.sort)
	assert.Equal(t, "blue", settings.colors.bigup)
	assert.Equal(t, "cyan", settings.colors.up)
	assert.Equal(t, "orange", settings.colors.drop)
	assert.Equal(t, "magenta", settings.colors.bigdrop)
	assert.Equal(t, []string{"AMD", "^NSEI", "EURUSD=X"}, settings.symbols)
}
