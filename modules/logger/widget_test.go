package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

const testGlobalConfig = "wtf:\n  colors:\n    theme: default\n"

// newTestSettings creates a Settings suitable for testing.
func newTestSettings(t *testing.T) *Settings {
	t.Helper()

	ymlConfig, err := config.ParseYaml("enabled: true\n")
	if err != nil {
		t.Fatal(err)
	}
	globalConfig, err := config.ParseYaml(testGlobalConfig)
	if err != nil {
		t.Fatal(err)
	}

	return &Settings{
		Common: cfg.NewCommonSettingsFromModule("logger", defaultTitle, defaultFocusable, ymlConfig, globalConfig),
	}
}

// newTestWidget creates a Widget with the TextWidget properly initialized.
func newTestWidget(t *testing.T, filePath string) *Widget {
	t.Helper()

	settings := newTestSettings(t)
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	widget := &Widget{
		TextWidget: view.NewTextWidget(app, redrawChan, nil, settings.Common),
		filePath:   filePath,
		settings:   settings,
	}

	return widget
}

func TestContent_FileDoesNotExist(t *testing.T) {
	// When the widget's filePath points to a nonexistent file,
	// tailFile returns empty and content returns empty body.
	dir := t.TempDir()
	t.Setenv("WTF_CONFIG_DIR", dir)

	// Create the log.txt so LogFileMissing() returns false
	logFile := filepath.Join(dir, "log.txt")
	if err := os.WriteFile(logFile, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	// But point the widget at a different nonexistent file
	widget := newTestWidget(t, filepath.Join(dir, "nonexistent.txt"))

	title, body, wrap := widget.content()

	if title != defaultTitle {
		t.Errorf("expected title %q, got %q", defaultTitle, title)
	}
	if body != "" {
		t.Errorf("expected empty body, got %q", body)
	}
	if wrap != false {
		t.Errorf("expected wrap=false, got %v", wrap)
	}
}

func TestContent_FormatsLogLines(t *testing.T) {
	tests := []struct {
		name     string
		logData  string
		wantSubs []string // substrings expected in output
	}{
		{
			name:    "standard log line with 4+ chunks",
			logData: "2023/01/15 10:30:45 INFO something happened here\n",
			wantSubs: []string{
				"[green]2023/01/15[white]",
				"[yellow]10:30:45[white]",
				"something happened here",
			},
		},
		{
			name:    "multiple log lines reversed",
			logData: "2023/01/15 10:00:00 INFO first message\n2023/01/15 11:00:00 INFO second message\n",
			wantSubs: []string{
				"second message",
				"first message",
			},
		},
		{
			name:     "line with fewer than 4 chunks is skipped",
			logData:  "short line\n",
			wantSubs: []string{},
		},
		{
			name:    "exactly 4 chunks",
			logData: "2023/01/15 10:30:45 INFO message\n",
			wantSubs: []string{
				"[green]2023/01/15[white]",
				"[yellow]10:30:45[white]",
				"message",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := t.TempDir()
			t.Setenv("WTF_CONFIG_DIR", dir)

			logFile := filepath.Join(dir, "log.txt")
			if err := os.WriteFile(logFile, []byte(tt.logData), 0644); err != nil {
				t.Fatal(err)
			}

			widget := newTestWidget(t, logFile)
			_, body, wrap := widget.content()

			for _, sub := range tt.wantSubs {
				if !strings.Contains(body, sub) {
					t.Errorf("expected body to contain %q, got:\n%s", sub, body)
				}
			}
			if wrap != false {
				t.Errorf("expected wrap=false")
			}
		})
	}
}

func TestContent_LineOrderIsReversed(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("WTF_CONFIG_DIR", dir)

	logData := "2023/01/01 01:00:00 INFO alpha\n2023/01/01 02:00:00 INFO beta\n2023/01/01 03:00:00 INFO gamma\n"
	logFile := filepath.Join(dir, "log.txt")
	if err := os.WriteFile(logFile, []byte(logData), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(t, logFile)
	_, body, _ := widget.content()

	gammaIdx := strings.Index(body, "gamma")
	betaIdx := strings.Index(body, "beta")
	alphaIdx := strings.Index(body, "alpha")

	if gammaIdx == -1 || betaIdx == -1 || alphaIdx == -1 {
		t.Fatalf("expected all lines in body, got:\n%s", body)
	}
	if gammaIdx > betaIdx || betaIdx > alphaIdx {
		t.Errorf("expected reversed order (gamma, beta, alpha), got gamma@%d beta@%d alpha@%d", gammaIdx, betaIdx, alphaIdx)
	}
}

func TestTailFile_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "log.txt")
	if err := os.WriteFile(logFile, []byte(""), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(t, logFile)
	lines := widget.tailFile()

	// An empty file has size 0, so bufferSize=0, ReadAt reads 0 bytes,
	// strings.Split("", "\n") returns [""], and reversal keeps it.
	// The key thing is it doesn't panic.
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			t.Errorf("expected only empty lines for empty file, got %q", line)
		}
	}
}

func TestTailFile_NonexistentFile(t *testing.T) {
	widget := newTestWidget(t, "/nonexistent/path/log.txt")
	lines := widget.tailFile()

	if len(lines) != 0 {
		t.Errorf("expected empty slice for missing file, got %v", lines)
	}
}

func TestTailFile_LargerThanBuffer(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "log.txt")

	// Create content larger than maxBufferSize (1024 bytes)
	var sb strings.Builder
	for i := 0; i < 100; i++ {
		sb.WriteString("2023/01/15 10:30:45 INFO this is line number that should fill buffer\n")
	}
	data := sb.String()
	if int64(len(data)) <= maxBufferSize {
		t.Fatal("test data must exceed maxBufferSize")
	}

	if err := os.WriteFile(logFile, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(t, logFile)
	lines := widget.tailFile()

	// Should return lines from the tail, not the full file
	if len(lines) == 0 {
		t.Fatal("expected non-empty result")
	}

	// The tail should contain content from the end of the file
	joined := strings.Join(lines, "\n")
	if !strings.Contains(joined, "fill buffer") {
		t.Errorf("expected tail content from end of file, got:\n%s", joined)
	}
}

func TestTailFile_ExactlyBufferSize(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "log.txt")

	// Create content exactly maxBufferSize bytes
	line := "2023/01/15 10:30:45 INFO msg\n"
	repeatCount := int(maxBufferSize) / len(line)
	remainder := int(maxBufferSize) - (repeatCount * len(line))
	data := strings.Repeat(line, repeatCount) + strings.Repeat("x", remainder)

	if int64(len(data)) != maxBufferSize {
		t.Fatalf("expected %d bytes, got %d", maxBufferSize, len(data))
	}

	if err := os.WriteFile(logFile, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(t, logFile)
	lines := widget.tailFile()

	if len(lines) == 0 {
		t.Error("expected non-empty result for file at buffer size")
	}
}

func TestTailFile_SmallerThanBuffer(t *testing.T) {
	dir := t.TempDir()
	logFile := filepath.Join(dir, "log.txt")

	data := "2023/01/15 10:30:45 INFO short\n"
	if err := os.WriteFile(logFile, []byte(data), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(t, logFile)
	lines := widget.tailFile()

	if len(lines) == 0 {
		t.Fatal("expected non-empty result")
	}

	joined := strings.Join(lines, " ")
	if !strings.Contains(joined, "short") {
		t.Errorf("expected content from file, got: %v", lines)
	}
}

func TestNewSettingsFromYAML(t *testing.T) {
	tests := []struct {
		name          string
		yaml          string
		expectedTitle string
	}{
		{
			name:          "default title",
			yaml:          "enabled: true\n",
			expectedTitle: defaultTitle,
		},
		{
			name:          "custom title",
			yaml:          "enabled: true\ntitle: MyLog\n",
			expectedTitle: "MyLog",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymlConfig, err := config.ParseYaml(tt.yaml)
			if err != nil {
				t.Fatal(err)
			}
			globalConfig, err := config.ParseYaml(testGlobalConfig)
			if err != nil {
				t.Fatal(err)
			}

			settings := NewSettingsFromYAML("logger", ymlConfig, globalConfig)

			if settings.Common == nil {
				t.Fatal("expected Common to be set")
			}
			if settings.Title != tt.expectedTitle {
				t.Errorf("expected title %q, got %q", tt.expectedTitle, settings.Title)
			}
		})
	}
}

func TestSettings_DefaultFocusable(t *testing.T) {
	if defaultFocusable != true {
		t.Errorf("expected defaultFocusable=true, got %v", defaultFocusable)
	}
}

func TestSettings_DefaultTitle(t *testing.T) {
	if defaultTitle != "Logger" {
		t.Errorf("expected defaultTitle=%q, got %q", "Logger", defaultTitle)
	}
}

func TestConfigText(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("WTF_CONFIG_DIR", dir)

	widget := newTestWidget(t, filepath.Join(dir, "log.txt"))
	text := widget.ConfigText()

	// ConfigText should return something (from utils.HelpFromInterface)
	if text == "" {
		t.Error("expected non-empty ConfigText")
	}
}

func TestMaxBufferSize(t *testing.T) {
	if maxBufferSize != 1024 {
		t.Errorf("expected maxBufferSize=1024, got %d", maxBufferSize)
	}
}

func TestNewWidget(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("WTF_CONFIG_DIR", dir)

	settings := newTestSettings(t)
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	widget := NewWidget(app, redrawChan, settings)

	if widget == nil {
		t.Fatal("expected non-nil widget")
	}
	if widget.settings != settings {
		t.Error("expected settings to be assigned")
	}
	if widget.filePath == "" {
		t.Error("expected filePath to be set from LogFilePath()")
	}
}

func TestRefresh(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("WTF_CONFIG_DIR", dir)

	logFile := filepath.Join(dir, "log.txt")
	if err := os.WriteFile(logFile, []byte("2023/01/15 10:30:45 INFO test message\n"), 0644); err != nil {
		t.Fatal(err)
	}

	settings := newTestSettings(t)
	app := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	widget := NewWidget(app, redrawChan, settings)

	// Refresh should not panic
	widget.Refresh()

	// Drain the redraw channel to confirm it was triggered
	select {
	case <-redrawChan:
		// success
	default:
		// It's fine if the channel was already drained
	}
}
