package system

import (
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
	"gotest.tools/assert"
)

func newTestSettings(t *testing.T) *Settings {
	t.Helper()

	ymlConfig, err := config.ParseYaml("{}")
	assert.NilError(t, err)

	globalConfig, err := config.ParseYaml("{}")
	assert.NilError(t, err)

	return NewSettingsFromYAML("system", ymlConfig, globalConfig)
}

func Test_prettyDate(t *testing.T) {
	tests := []struct {
		name     string
		date     string
		expected string
		wantErr  bool
	}{
		{
			name:     "valid timestamp",
			date:     "2021-03-05T14:30:00-0700",
			expected: "Mar  5, 14:30",
		},
		{
			name:     "valid timestamp, single digit day",
			date:     "2020-01-02T09:05:00+0000",
			expected: "Jan  2, 09:05",
		},
		{
			name:     "valid timestamp, double digit day",
			date:     "2020-12-25T23:59:00+0000",
			expected: "Dec 25, 23:59",
		},
		{
			name:    "empty date",
			date:    "",
			wantErr: true,
		},
		{
			name:    "malformed date",
			date:    "not-a-date",
			wantErr: true,
		},
		{
			name:    "wrong format",
			date:    "2021-03-05",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				Date: tt.date,
			}

			actual := widget.prettyDate()

			if tt.wantErr {
				// On error, prettyDate returns the error message itself,
				// which will not match the "Jan _2, 15:04" pretty format.
				assert.Assert(t, actual != "")
				assert.Assert(t, !strings.Contains(actual, ","))
			} else {
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func Test_NewWidget(t *testing.T) {
	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := newTestSettings(t)

	widget := NewWidget(tviewApp, redrawChan, "2021-01-01T00:00:00+0000", "v1.2.3", settings)

	assert.Assert(t, widget != nil)
	assert.Equal(t, "2021-01-01T00:00:00+0000", widget.Date)
	assert.Equal(t, "v1.2.3", widget.Version)
	assert.Assert(t, widget.settings != nil)
	assert.Assert(t, widget.systemInfo != nil)
}

func Test_display(t *testing.T) {
	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := newTestSettings(t)

	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),
		Date:       "2021-03-05T14:30:00-0700",
		Version:    "v1.2.3",
		settings:   settings,
		systemInfo: &SystemInfo{
			ProductName:    "TestOS",
			ProductVersion: "TestOS 1.0",
			BuildVersion:   "12345",
		},
	}

	title, content, wrap := widget.display()

	assert.Equal(t, settings.Title, title)
	assert.Assert(t, strings.Contains(content, "Mar  5, 14:30"))
	assert.Assert(t, strings.Contains(content, "v1.2.3"))
	assert.Assert(t, strings.Contains(content, "TestOS 1.0"))
	assert.Assert(t, strings.Contains(content, "12345"))
	assert.Equal(t, false, wrap)
}

func Test_Refresh(t *testing.T) {
	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := newTestSettings(t)

	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),
		Date:       "2021-03-05T14:30:00-0700",
		Version:    "v1.2.3",
		settings:   settings,
		systemInfo: &SystemInfo{
			ProductName:    "TestOS",
			ProductVersion: "TestOS 1.0",
			BuildVersion:   "12345",
		},
	}

	widget.Refresh()

	select {
	case redrawn := <-redrawChan:
		assert.Equal(t, true, redrawn)
	case <-time.After(time.Second):
		t.Fatal("expected Refresh to signal the redraw channel")
	}

	assert.Assert(t, strings.Contains(widget.TextView().GetText(true), "v1.2.3"))
}
