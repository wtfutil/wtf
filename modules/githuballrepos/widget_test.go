package githuballrepos

import (
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/logger"
)

type mockCache struct{}

func (m *mockCache) Get() *WidgetData {
	logger.Log("Mock cache: Get called")
	return &WidgetData{}
}

func (m *mockCache) Set(data *WidgetData) {
	logger.Log("Mock cache: Set called")
}

func (m *mockCache) IsValid() bool {
	logger.Log("Mock cache: IsValid called")
	return false
}

type mockGitHubClient struct{}

func (m *mockGitHubClient) FetchData(orgs []string, username string) *WidgetData {
	logger.Log("Mock GitHub client: FetchData called")
	return &WidgetData{}
}

func Test_Refresh(t *testing.T) {
	widget := createTestWidget()

	done := make(chan bool)
	go func() {
		logger.Log("Starting Refresh")
		widget.Refresh()
		logger.Log("Refresh completed")
		done <- true
	}()

	select {
	case <-done:
		// Test passed
	case <-time.After(5 * time.Second):
		t.Fatal("Test timed out")
	}

	assert.NotNil(t, widget.data)
	assert.True(t, time.Since(widget.lastUpdateTime) < time.Second)
}

func Test_NewWidget(t *testing.T) {
	app := tview.NewApplication()
	pages := tview.NewPages()

	mockYmlConfig, _ := config.ParseYaml(`
common:
  refreshInterval: 1
  title: "Mock Title"
`)
	mockGlobalConfig, _ := config.ParseYaml(`
wtf:
  colors:
    background: "black"
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`)

	settings := &Settings{
		common: cfg.NewCommonSettingsFromModule(
			"githuballrepos",
			"githuballrepos",
			true,
			mockYmlConfig,
			mockGlobalConfig,
		),
		APIKey:        "mock_api_key",
		Organizations: []string{"org1", "org2"},
		Username:      "mock_username",
		CachePath:     "/tmp/mock_cache",
	}

	widget := NewWidget(app, make(chan bool), pages, settings)

	assert.NotNil(t, widget)
	assert.Equal(t, "githuballrepos", widget.CommonSettings().Title)
	assert.False(t, widget.isItemSelected)
	assert.Equal(t, -1, widget.selectedPRIndex)
	assert.Equal(t, 0, widget.currentTab)
}

func Test_Next(t *testing.T) {
	widget := createTestWidget()
	widget.data = &WidgetData{
		MyPRs: []PR{{Title: "PR1"}, {Title: "PR2"}, {Title: "PR3"}},
	}

	widget.Next()

	assert.Equal(t, 0, widget.selectedPRIndex)
	assert.True(t, widget.isItemSelected)

	widget.Next()
	assert.Equal(t, 1, widget.selectedPRIndex)

	widget.Next()
	assert.Equal(t, 2, widget.selectedPRIndex)

	widget.Next()
	assert.Equal(t, 0, widget.selectedPRIndex)
}

func Test_Prev(t *testing.T) {
	widget := createTestWidget()
	widget.data = &WidgetData{
		MyPRs: []PR{{Title: "PR1"}, {Title: "PR2"}, {Title: "PR3"}},
	}

	widget.Prev()
	assert.Equal(t, 2, widget.selectedPRIndex)
	assert.True(t, widget.isItemSelected)

	widget.Prev()
	assert.Equal(t, 1, widget.selectedPRIndex)

	widget.Prev()
	assert.Equal(t, 0, widget.selectedPRIndex)

	widget.Prev()
	assert.Equal(t, 2, widget.selectedPRIndex)
}

func Test_NextTab(t *testing.T) {
	widget := createTestWidget()
	widget.isItemSelected = true
	widget.selectedPRIndex = 1

	widget.nextTab()
	assert.Equal(t, 1, widget.currentTab)
	assert.Equal(t, -1, widget.selectedPRIndex)
	assert.False(t, widget.isItemSelected)

	widget.nextTab()
	assert.Equal(t, 2, widget.currentTab)

	widget.nextTab()
	assert.Equal(t, 0, widget.currentTab)
}

func Test_PrevTab(t *testing.T) {
	widget := createTestWidget()
	widget.isItemSelected = true
	widget.selectedPRIndex = 1

	widget.prevTab()
	assert.Equal(t, 2, widget.currentTab)
	assert.Equal(t, -1, widget.selectedPRIndex)
	assert.False(t, widget.isItemSelected)

	widget.prevTab()
	assert.Equal(t, 1, widget.currentTab)

	widget.prevTab()
	assert.Equal(t, 0, widget.currentTab)
}

func Test_SetSelected(t *testing.T) {
	widget := createTestWidget()
	widget.isItemSelected = true
	widget.selectedPRIndex = 1

	widget.SetSelected(false)
	assert.False(t, widget.isSelected)
	assert.False(t, widget.isItemSelected)
	assert.Equal(t, -1, widget.selectedPRIndex)

	widget.SetSelected(true)
	assert.True(t, widget.isSelected)
	assert.False(t, widget.isItemSelected)
	assert.Equal(t, -1, widget.selectedPRIndex)
}

func createTestWidget() *Widget {
	logger.Log("Creating widget")

	app := tview.NewApplication()
	pages := tview.NewPages()

	mockYmlConfig, _ := config.ParseYaml(`
common:
  refreshInterval: 1
  title: "Mock Title"
`)
	mockGlobalConfig, _ := config.ParseYaml(`
wtf:
  colors:
    background: "black"
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`)

	settings := &Settings{
		common: cfg.NewCommonSettingsFromModule(
			"githuballrepos",
			"Mock Title", // Use "Mock Title" here instead of "githuballrepos"
			true,
			mockYmlConfig,
			mockGlobalConfig,
		),
		APIKey:        "mock_api_key",
		Organizations: []string{"org1", "org2"},
		Username:      "mock_username",
		CachePath:     "/tmp/mock_cache",
	}

	widget := NewWidget(app, make(chan bool), pages, settings)

	// Replace cache and client with mocks
	widget.cache = &mockCache{}
	widget.client = &mockGitHubClient{}
	widget.testMode = true

	logger.Log("Widget created")
	return widget
}
