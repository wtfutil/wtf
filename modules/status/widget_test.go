package status

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
)

func testWidget(title string) *Widget {
	ymlConfig, _ := config.ParseYaml(`
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
`)
	globalConfig, _ := config.ParseYaml(`
wtf:
  colors:
    background: black
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
`)

	common := cfg.NewCommonSettingsFromModule("status", title, defaultFocusable, ymlConfig, globalConfig)
	settings := &Settings{Common: common}

	app := tview.NewApplication()
	redrawChan := make(chan bool)

	return NewWidget(app, redrawChan, settings)
}

func TestAnimation_CyclesThroughIcons(t *testing.T) {
	expectedIcons := []string{"|", "/", "-", "\\", "|"}

	tests := []struct {
		name         string
		startIcon    int
		wantContent  string
		wantNextIcon int
	}{
		{"first icon (pipe)", 0, "|", 1},
		{"second icon (slash)", 1, "/", 2},
		{"third icon (dash)", 2, "-", 3},
		{"fourth icon (backslash)", 3, "\\", 4},
		{"fifth icon (pipe again)", 4, "|", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := testWidget("Status")
			widget.CurrentIcon = tt.startIcon

			_, content, wrap := widget.animation()

			if content != expectedIcons[tt.startIcon] {
				t.Errorf("animation() content = %q, want %q", content, tt.wantContent)
			}
			if widget.CurrentIcon != tt.wantNextIcon {
				t.Errorf("animation() next CurrentIcon = %d, want %d", widget.CurrentIcon, tt.wantNextIcon)
			}
			if wrap != false {
				t.Errorf("animation() wrap = %v, want false", wrap)
			}
		})
	}
}

func TestAnimation_WrapsAround(t *testing.T) {
	widget := testWidget("Status")
	widget.CurrentIcon = 4

	widget.animation()

	if widget.CurrentIcon != 0 {
		t.Errorf("after last icon, CurrentIcon = %d, want 0", widget.CurrentIcon)
	}
}

func TestAnimation_ReturnsTitle(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantTitle string
	}{
		{"default title", "Status", "Status"},
		{"custom title", "My Custom Status", "My Custom Status"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := testWidget(tt.title)

			title, _, _ := widget.animation()

			if title != tt.wantTitle {
				t.Errorf("animation() title = %q, want %q", title, tt.wantTitle)
			}
		})
	}
}

func TestAnimation_FullCycle(t *testing.T) {
	widget := testWidget("Status")
	expectedSequence := []string{"|", "/", "-", "\\", "|"}

	for i, expected := range expectedSequence {
		_, content, _ := widget.animation()
		if content != expected {
			t.Errorf("cycle iteration %d: got %q, want %q", i, content, expected)
		}
	}

	// After a full cycle, should wrap and start over
	_, content, _ := widget.animation()
	if content != "|" {
		t.Errorf("after full cycle, got %q, want %q", content, "|")
	}
}

func TestNewWidget_InitialState(t *testing.T) {
	widget := testWidget("Status")

	if widget.CurrentIcon != 0 {
		t.Errorf("NewWidget() CurrentIcon = %d, want 0", widget.CurrentIcon)
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	ymlConfig, _ := config.ParseYaml(`
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
`)
	globalConfig, _ := config.ParseYaml(`
wtf:
  colors:
    background: black
    border:
      focusable: darkslateblue
      focused: orange
      normal: gray
`)

	settings := NewSettingsFromYAML("teststatus", ymlConfig, globalConfig)

	if settings.Common == nil {
		t.Fatal("NewSettingsFromYAML() Common is nil")
	}
	if settings.Title != defaultTitle {
		t.Errorf("NewSettingsFromYAML() Title = %q, want %q", settings.Title, defaultTitle)
	}
	if settings.Focusable != defaultFocusable {
		t.Errorf("NewSettingsFromYAML() Focusable = %v, want %v", settings.Focusable, defaultFocusable)
	}
}
