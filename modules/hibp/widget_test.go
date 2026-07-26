package hibp

import (
	"errors"
	"testing"

	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

func newTestWidgetWithColors(since, ok, pwned string) *Widget {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{Title: "HIBP"},
		since:  since,
	}
	settings.ok = ok
	settings.pwned = pwned

	return NewWidget(app, redrawChan, settings)
}

func TestSinceDateForTitle(t *testing.T) {
	tests := []struct {
		name     string
		since    string
		expected string
	}{
		{"no since", "", ""},
		{"valid since", "2019-06-22", " since Jun 22, 2019"},
		{"malformed since falls back to raw string", "not-a-date", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := newTestWidgetWithColors(tt.since, "white", "red")
			assert.Equal(t, tt.expected, widget.sinceDateForTitle())
		})
	}
}

func TestWidget_Content(t *testing.T) {
	t.Run("error state shows error message", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")
		widget.err = errors.New("boom")

		title, content, isErr := widget.content()

		assert.Equal(t, "HIBP", title)
		assert.Equal(t, "boom", content)
		assert.True(t, isErr)
	})

	t.Run("no breaches uses ok color", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")
		widget.statuses = []*Status{
			NewStatus("safe@example.com", []Breach{}),
		}

		title, content, isErr := widget.content()

		assert.Equal(t, "HIBP", title)
		assert.False(t, isErr)
		assert.Contains(t, content, "[white]safe@example.com[white]")
	})

	t.Run("breaches use pwned color", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")
		widget.statuses = []*Status{
			NewStatus("pwned@example.com", []Breach{{Name: "Adobe", Date: "2013-10-04"}}),
		}

		title, content, isErr := widget.content()

		assert.Equal(t, "HIBP", title)
		assert.False(t, isErr)
		assert.Contains(t, content, "[red]pwned@example.com[white]")
	})

	t.Run("title includes since date", func(t *testing.T) {
		widget := newTestWidgetWithColors("2019-06-22", "white", "red")
		widget.statuses = []*Status{}

		title, content, isErr := widget.content()

		assert.Equal(t, "HIBP since Jun 22, 2019", title)
		assert.Equal(t, "", content)
		assert.False(t, isErr)
	})

	t.Run("mixed statuses render each account", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")
		widget.statuses = []*Status{
			NewStatus("safe@example.com", nil),
			NewStatus("pwned@example.com", []Breach{{Name: "Adobe"}}),
		}

		_, content, isErr := widget.content()

		assert.False(t, isErr)
		assert.Contains(t, content, "[white]safe@example.com[white]")
		assert.Contains(t, content, "[red]pwned@example.com[white]")
	})
}

func TestNewWidget(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	settings := &Settings{
		Common: &cfg.Common{Title: "HIBP"},
	}

	widget := NewWidget(app, redrawChan, settings)

	assert.NotNil(t, widget)
	assert.Equal(t, settings, widget.settings)
}

func TestWidget_Fetch(t *testing.T) {
	t.Run("empty accounts list returns empty statuses", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")

		statuses, err := widget.Fetch([]string{})

		assert.NoError(t, err)
		assert.Empty(t, statuses)
	})

	t.Run("accounts containing an empty string yield a nil status entry", func(t *testing.T) {
		widget := newTestWidgetWithColors("", "white", "red")

		statuses, err := widget.Fetch([]string{""})

		assert.NoError(t, err)
		assert.Len(t, statuses, 1)
		assert.Nil(t, statuses[0])
	})
}

func TestWidget_Refresh(t *testing.T) {
	widget := newTestWidgetWithColors("", "white", "red")

	// With no accounts configured, Refresh should complete without error and
	// leave the widget in a clean, non-error state.
	widget.Refresh()

	assert.Nil(t, widget.err)
	assert.Empty(t, widget.statuses)
}
