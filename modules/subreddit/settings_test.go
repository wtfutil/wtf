package subreddit

import (
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name              string
		yaml              string
		wantSubreddit     string
		wantNumberOfPosts int
		wantSortOrder     string
		wantTopTimePeriod string
	}{
		{
			name: "all fields specified",
			yaml: `
subreddit: "golang"
numberOfPosts: 5
sortOrder: "top"
topTimePeriod: "week"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`,
			wantSubreddit:     "golang",
			wantNumberOfPosts: 5,
			wantSortOrder:     "top",
			wantTopTimePeriod: "week",
		},
		{
			name: "defaults applied",
			yaml: `
subreddit: "programming"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`,
			wantSubreddit:     "programming",
			wantNumberOfPosts: 10,
			wantSortOrder:     "hot",
			wantTopTimePeriod: "all",
		},
		{
			name: "rising sort",
			yaml: `
subreddit: "news"
sortOrder: "rising"
numberOfPosts: 20
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`,
			wantSubreddit:     "news",
			wantNumberOfPosts: 20,
			wantSortOrder:     "rising",
			wantTopTimePeriod: "all",
		},
	}

	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yamlConfig, err := config.ParseYaml(tt.yaml)
			if err != nil {
				t.Fatalf("failed to parse yaml: %v", err)
			}

			settings := NewSettingsFromYAML("subreddit", yamlConfig, globalConfig)

			if settings.subreddit != tt.wantSubreddit {
				t.Errorf("subreddit: got %q, want %q", settings.subreddit, tt.wantSubreddit)
			}
			if settings.numberOfPosts != tt.wantNumberOfPosts {
				t.Errorf("numberOfPosts: got %d, want %d", settings.numberOfPosts, tt.wantNumberOfPosts)
			}
			if settings.sortOrder != tt.wantSortOrder {
				t.Errorf("sortOrder: got %q, want %q", settings.sortOrder, tt.wantSortOrder)
			}
			if settings.topTimePeriod != tt.wantTopTimePeriod {
				t.Errorf("topTimePeriod: got %q, want %q", settings.topTimePeriod, tt.wantTopTimePeriod)
			}
			if settings.Common == nil {
				t.Error("expected non-nil Common settings")
			}
		})
	}
}

func TestSettings_ConfigText(t *testing.T) {
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlStr := `
subreddit: "test"
enabled: true
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalConfig, _ := config.ParseYaml(globalStr)
	yamlConfig, _ := config.ParseYaml(yamlStr)

	settings := NewSettingsFromYAML("subreddit", yamlConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)
	widget := NewWidget(app, redrawChan, nil, settings)

	text := widget.ConfigText()
	if text == "" {
		t.Error("expected non-empty config text")
	}
	// Should contain help text from struct tags
	if !strings.Contains(text, "subreddit") || !strings.Contains(text, "Subreddit") {
		t.Errorf("expected config text to reference subreddit field, got %q", text)
	}
}
