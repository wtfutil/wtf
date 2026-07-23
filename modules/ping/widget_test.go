package ping

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	probing "github.com/prometheus-community/pro-bing"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

// expectedStatusLine mirrors the padding logic in Widget.content so tests
// stay correct if the minimum column width or padding character changes.
func expectedStatusLine(nameWidth int, label, status string) string {
	return fmt.Sprintf("[white]%-*s: %s", nameWidth, label, status)
}

func Test_Widget_content(t *testing.T) {
	tests := []struct {
		name     string
		hosts    []Host
		expected string
	}{
		{
			name:     "no hosts produces empty content",
			hosts:    []Host{},
			expected: "",
		},
		{
			name: "single host up",
			hosts: []Host{
				{Label: "example.com", Hostname: "example.com", Up: true},
			},
			expected: expectedStatusLine(12, "example.com", "[green]Up"),
		},
		{
			name: "single host down",
			hosts: []Host{
				{Label: "example.com", Hostname: "example.com", Up: false},
			},
			expected: expectedStatusLine(12, "example.com", "[red]DOWN"),
		},
		{
			name: "multiple hosts mixed status",
			hosts: []Host{
				{Label: "up-host", Hostname: "up.example.com", Up: true},
				{Label: "down-host", Hostname: "down.example.com", Up: false},
			},
			// nameWidth stays at the 12-char minimum since both labels are shorter.
			expected: strings.Join([]string{
				expectedStatusLine(12, "up-host", "[green]Up"),
				expectedStatusLine(12, "down-host", "[red]DOWN"),
			}, "\n"),
		},
		{
			name: "long label widens name column",
			hosts: []Host{
				{Label: "a-very-long-hostname-label", Hostname: "long.example.com", Up: true},
			},
			// nameWidth grows to len(label)+2 since it exceeds the 12-char minimum.
			expected: expectedStatusLine(28, "a-very-long-hostname-label", "[green]Up"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{hosts: tt.hosts}

			assert.Equal(t, tt.expected, widget.content())
		})
	}
}

func Test_Widget_content_shortLabelsPadToMinimumWidth(t *testing.T) {
	widget := &Widget{
		hosts: []Host{
			{Label: "a", Hostname: "a.example.com", Up: true},
		},
	}

	content := widget.content()

	// nameWidth defaults to 12 when all labels are shorter than that.
	assert.Equal(t, expectedStatusLine(12, "a", "[green]Up"), content)
}

func testSettings(title string, hosts []Host) *Settings {
	return &Settings{
		common: &cfg.Common{Title: title},
		hosts:  hosts,
	}
}

func Test_NewWidget(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := testSettings("Test Pings", []Host{
		{Label: "example.com", Hostname: "example.com"},
	})

	widget := NewWidget(app, redrawChan, settings)

	assert.NotNil(t, widget)
	assert.Equal(t, settings, widget.settings)
	assert.Equal(t, settings.hosts, widget.hosts)
	assert.Equal(t, "Test Pings", widget.CommonSettings().Title)
}

func Test_NewWidget_NoHosts(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := testSettings("Test Pings", []Host{})

	widget := NewWidget(app, redrawChan, settings)

	assert.NotNil(t, widget)
	assert.Empty(t, widget.hosts)
}

func Test_Widget_display(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := testSettings("Test Pings", []Host{
		{Label: "up-host", Hostname: "up.example.com", Up: true},
		{Label: "down-host", Hostname: "down.example.com", Up: false},
	})

	widget := NewWidget(app, redrawChan, settings)

	widget.display()

	// display() pushes a value onto RedrawChan once the view is updated.
	select {
	case <-redrawChan:
		// expected
	case <-time.After(time.Second):
		t.Fatal("display() did not signal RedrawChan")
	}

	expectedContent := strings.Join([]string{
		expectedStatusLine(12, "up-host", "[green]Up"),
		expectedStatusLine(12, "down-host", "[red]DOWN"),
	}, "\n")

	assert.Equal(t, expectedContent, widget.TextView().GetText(false))
	assert.Equal(t, "Test Pings", widget.CommonSettings().Title)
}

// fakePinger is a test double for pinger that avoids sending real ICMP
// packets. runErr, when set, is returned from Run(); otherwise stats is
// returned from Statistics().
type fakePinger struct {
	runErr error
	stats  *probing.Statistics
}

func (f *fakePinger) Run() error {
	return f.runErr
}

func (f *fakePinger) Statistics() *probing.Statistics {
	return f.stats
}

func Test_hostUpFromStatistics(t *testing.T) {
	tests := []struct {
		name     string
		stats    *probing.Statistics
		expected bool
	}{
		{
			name:     "packets received means host is up",
			stats:    &probing.Statistics{PacketsRecv: 1},
			expected: true,
		},
		{
			name:     "multiple packets received means host is up",
			stats:    &probing.Statistics{PacketsRecv: 5},
			expected: true,
		},
		{
			name:     "no packets received means host is down",
			stats:    &probing.Statistics{PacketsRecv: 0},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, hostUpFromStatistics(tt.stats))
		})
	}
}

func Test_Widget_doPings(t *testing.T) {
	tests := []struct {
		name          string
		hosts         []Host
		pingerFactory func(hostname string) (pinger, error)
		expectedUp    []bool
	}{
		{
			name: "all hosts respond",
			hosts: []Host{
				{Label: "a", Hostname: "a.example.com"},
				{Label: "b", Hostname: "b.example.com"},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				return &fakePinger{stats: &probing.Statistics{PacketsRecv: 1}}, nil
			},
			expectedUp: []bool{true, true},
		},
		{
			name: "no hosts respond",
			hosts: []Host{
				{Label: "a", Hostname: "a.example.com"},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				return &fakePinger{stats: &probing.Statistics{PacketsRecv: 0}}, nil
			},
			expectedUp: []bool{false},
		},
		{
			name: "pinger creation error leaves host down",
			hosts: []Host{
				{Label: "a", Hostname: "bad host"},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				return nil, errors.New("invalid hostname")
			},
			expectedUp: []bool{false},
		},
		{
			name: "pinger run error leaves host down",
			hosts: []Host{
				{Label: "a", Hostname: "unreachable.example.com"},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				return &fakePinger{runErr: errors.New("network unreachable")}, nil
			},
			expectedUp: []bool{false},
		},
		{
			name: "previously up host resets to down before re-pinging",
			hosts: []Host{
				{Label: "a", Hostname: "a.example.com", Up: true},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				return &fakePinger{runErr: errors.New("timeout")}, nil
			},
			expectedUp: []bool{false},
		},
		{
			name:          "no hosts is a no-op",
			hosts:         []Host{},
			pingerFactory: func(hostname string) (pinger, error) { return nil, nil },
			expectedUp:    []bool{},
		},
		{
			name: "mixed results across hosts",
			hosts: []Host{
				{Label: "up", Hostname: "up.example.com"},
				{Label: "down", Hostname: "down.example.com"},
			},
			pingerFactory: func(hostname string) (pinger, error) {
				if hostname == "up.example.com" {
					return &fakePinger{stats: &probing.Statistics{PacketsRecv: 1}}, nil
				}
				return &fakePinger{stats: &probing.Statistics{PacketsRecv: 0}}, nil
			},
			expectedUp: []bool{true, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{hosts: tt.hosts, pingerFactory: tt.pingerFactory}

			widget.doPings()

			gotUp := make([]bool, len(widget.hosts))
			for i, h := range widget.hosts {
				gotUp[i] = h.Up
			}
			assert.Equal(t, tt.expectedUp, gotUp)
		})
	}
}

func Test_newRealPinger(t *testing.T) {
	t.Run("invalid hostname returns an error", func(t *testing.T) {
		p, err := newRealPinger("this is not a valid hostname")

		assert.Error(t, err)
		assert.Nil(t, p)
	})

	t.Run("valid hostname builds a pinger without sending packets", func(t *testing.T) {
		p, err := newRealPinger("localhost")

		assert.NoError(t, err)
		assert.NotNil(t, p)
		// p satisfies the pinger interface (Run/Statistics); we don't call
		// Run() here since that requires real network/ICMP privileges.
	})
}

func Test_Widget_Refresh(t *testing.T) {
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	settings := testSettings("Test Pings", []Host{
		{Label: "example.com", Hostname: "example.com"},
	})

	widget := NewWidget(app, redrawChan, settings)
	widget.pingerFactory = func(hostname string) (pinger, error) {
		return &fakePinger{stats: &probing.Statistics{PacketsRecv: 1}}, nil
	}

	widget.Refresh()

	select {
	case <-redrawChan:
		// expected
	case <-time.After(time.Second):
		t.Fatal("Refresh() did not signal RedrawChan via display()")
	}

	assert.True(t, widget.hosts[0].Up)
	assert.Equal(t, expectedStatusLine(12, "example.com", "[green]Up"), widget.TextView().GetText(false))
}
