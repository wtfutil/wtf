package bargraph

import (
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/wtfutil/wtf/view"
)

func testSettings(t *testing.T) *Settings {
	t.Helper()
	ymlConfig, err := config.ParseYaml("enabled: true")
	require.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	require.NoError(t, err)
	return NewSettingsFromYAML("bargraph", ymlConfig, globalConfig)
}

func testWidget(t *testing.T) *Widget {
	t.Helper()
	app := tview.NewApplication()
	redrawChan := make(chan bool, 10)
	settings := testSettings(t)
	return NewWidget(app, redrawChan, settings)
}

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		expectedTitle string
	}{
		{
			name:          "default values",
			yaml:          "{}",
			expectedTitle: defaultTitle,
		},
		{
			name:          "custom title",
			yaml:          "title: \"My Graph\"",
			expectedTitle: "My Graph",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			require.NoError(t, err)
			globalConfig, err := config.ParseYaml("wtf: {}")
			require.NoError(t, err)

			settings := NewSettingsFromYAML("bargraph", ymlConfig, globalConfig)

			assert.NotNil(t, settings)
			assert.NotNil(t, settings.Common)
			assert.Equal(t, tt.expectedTitle, settings.Title)
		})
	}
}

func TestDefaultConstants(t *testing.T) {
	assert.Equal(t, false, defaultFocusable)
	assert.Equal(t, "Bargraph", defaultTitle)
}

func TestNewWidget(t *testing.T) {
	widget := testWidget(t)

	assert.NotNil(t, widget)
	assert.NotNil(t, widget.View)
	assert.NotNil(t, widget.tviewApp)
}

func TestMakeGraph(t *testing.T) {
	widget := testWidget(t)

	// MakeGraph should not panic and should write content to the view
	MakeGraph(widget)

	content := widget.View.GetText(false)
	assert.NotEmpty(t, content)

	// Should have 8 lines (one per bar)
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	assert.Equal(t, 8, len(lines))

	// Each line should contain a time label in HH:MM format
	for _, line := range lines {
		assert.Contains(t, line, ":")
	}
}

func TestMakeGraph_BarsHaveValidFormat(t *testing.T) {
	widget := testWidget(t)
	MakeGraph(widget)

	content := widget.View.GetText(false)
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")

	for _, line := range lines {
		// Each line should contain the bar character "|" (default starChar)
		assert.Contains(t, line, "|")
		// Each line should contain the color marker
		assert.Contains(t, line, "red")
	}
}

func TestMakeGraph_MultipleCallsOverwrite(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 10) // larger buffer to avoid blocking
	settings := testSettings(t)
	widget := NewWidget(app, redrawChan, settings)

	MakeGraph(widget)
	first := widget.View.GetText(false)
	assert.NotEmpty(t, first)

	// Clear and regenerate
	widget.View.Clear()
	MakeGraph(widget)
	second := widget.View.GetText(false)
	assert.NotEmpty(t, second)

	// Both should have 8 lines
	lines1 := strings.Split(strings.TrimRight(first, "\n"), "\n")
	lines2 := strings.Split(strings.TrimRight(second, "\n"), "\n")
	assert.Equal(t, 8, len(lines1))
	assert.Equal(t, 8, len(lines2))
}

func TestBuildStars_EmptyData(t *testing.T) {
	result := view.BuildStars([]view.Bar{}, 20, "|")
	assert.Equal(t, "", result)
}

func TestBuildStars_SingleBar(t *testing.T) {
	bars := []view.Bar{
		{Label: "10:00", Percent: 50, LabelColor: "green"},
	}
	result := view.BuildStars(bars, 20, "*")

	assert.Contains(t, result, "10:00")
	assert.Contains(t, result, "green")
	// 50% of 20 = 10 stars
	assert.Contains(t, result, strings.Repeat("*", 10))
	// Value label defaults to percent string
	assert.Contains(t, result, "50")
}

func TestBuildStars_CustomValueLabel(t *testing.T) {
	bars := []view.Bar{
		{Label: "A", Percent: 80, ValueLabel: "80%", LabelColor: "blue"},
	}
	result := view.BuildStars(bars, 10, "#")

	assert.Contains(t, result, "80%")
	// 80% of 10 = 8 chars
	assert.Contains(t, result, strings.Repeat("#", 8))
}

func TestBuildStars_DefaultLabelColor(t *testing.T) {
	bars := []view.Bar{
		{Label: "X", Percent: 25, LabelColor: ""},
	}
	result := view.BuildStars(bars, 20, "|")

	// Empty LabelColor should fall back to "default"
	assert.Contains(t, result, "default")
}

func TestBuildStars_LabelAlignment(t *testing.T) {
	bars := []view.Bar{
		{Label: "A", Percent: 10, LabelColor: "red"},
		{Label: "LongLabel", Percent: 90, LabelColor: "red"},
	}
	result := view.BuildStars(bars, 10, "|")
	lines := strings.Split(strings.TrimRight(result, "\n"), "\n")

	require.Equal(t, 2, len(lines))
	// Shorter label should be padded to match longest label length
	// "A" is 1 char, "LongLabel" is 9 chars, so "A" gets 8 spaces padding
	assert.True(t, strings.HasPrefix(lines[0], "A"+strings.Repeat(" ", 8)))
	assert.True(t, strings.HasPrefix(lines[1], "LongLabel"))
}

func TestBuildStars_ZeroPercent(t *testing.T) {
	bars := []view.Bar{
		{Label: "empty", Percent: 0, LabelColor: "white"},
	}
	result := view.BuildStars(bars, 20, "|")

	assert.Contains(t, result, "empty")
	// 0% should produce no bar characters between color tags
	assert.NotContains(t, result, "||")
}

func TestBuildStars_HundredPercent(t *testing.T) {
	bars := []view.Bar{
		{Label: "full", Percent: 100, LabelColor: "blue"},
	}
	result := view.BuildStars(bars, 20, "|")

	// 100% of 20 = 20 chars
	assert.Contains(t, result, strings.Repeat("|", 20))
}

func TestBuildStars_TableDriven(t *testing.T) {
	tests := []struct {
		name        string
		bars        []view.Bar
		maxStars    int
		starChar    string
		expectStars int
		expectLabel string
	}{
		{
			name:        "25% of 20",
			bars:        []view.Bar{{Label: "T", Percent: 25, LabelColor: "red"}},
			maxStars:    20,
			starChar:    "*",
			expectStars: 5,
			expectLabel: "25",
		},
		{
			name:        "75% of 40",
			bars:        []view.Bar{{Label: "T", Percent: 75, LabelColor: "red"}},
			maxStars:    40,
			starChar:    "#",
			expectStars: 30,
			expectLabel: "75",
		},
		{
			name:        "10% of 10",
			bars:        []view.Bar{{Label: "T", Percent: 10, LabelColor: "red"}},
			maxStars:    10,
			starChar:    "=",
			expectStars: 1,
			expectLabel: "10",
		},
		{
			name:        "custom value label",
			bars:        []view.Bar{{Label: "T", Percent: 50, ValueLabel: "half", LabelColor: "red"}},
			maxStars:    20,
			starChar:    "|",
			expectStars: 10,
			expectLabel: "half",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := view.BuildStars(tt.bars, tt.maxStars, tt.starChar)

			assert.Contains(t, result, strings.Repeat(tt.starChar, tt.expectStars))
			assert.Contains(t, result, tt.expectLabel)
		})
	}
}

func TestRefresh(t *testing.T) {
	widget := testWidget(t)

	// Refresh should populate the view without panicking
	widget.Refresh()

	content := widget.View.GetText(false)
	assert.NotEmpty(t, content)
}

func TestRefresh_WhenDisabled(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 10)

	ymlConfig, err := config.ParseYaml("enabled: false")
	require.NoError(t, err)
	globalConfig, err := config.ParseYaml("wtf: {}")
	require.NoError(t, err)

	settings := NewSettingsFromYAML("bargraph", ymlConfig, globalConfig)
	widget := NewWidget(app, redrawChan, settings)

	widget.Refresh()

	// When disabled, the view should remain empty
	content := widget.View.GetText(false)
	assert.Empty(t, content)
}
