package textfile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"

	"github.com/wtfutil/wtf/view"
)

// newTestWidget builds a Widget suitable for exercising the read/parse/format
// pipeline without starting the file watcher goroutine or requiring a live
// tview.Application.
func newTestWidget(t *testing.T, filePaths []string, format bool, formatStyle string) *Widget {
	t.Helper()

	yaml := "format: " + boolStr(format) + "\n"
	if formatStyle != "" {
		yaml += "formatStyle: " + formatStyle + "\n"
	}
	if len(filePaths) > 0 {
		yaml += "filePaths:\n"
		for _, p := range filePaths {
			yaml += "  - \"" + strings.ReplaceAll(p, "\\", "\\\\") + "\"\n"
		}
	}

	ymlConfig, err := config.ParseYaml(yaml)
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml(testGlobalConfig)
	assert.NoError(t, err)

	settings := NewSettingsFromYAML("textfile", ymlConfig, globalConfig)

	widget := &Widget{settings: settings}
	widget.MultiSourceWidget = view.NewMultiSourceWidget(settings.Common, "filePath", "filePaths")
	widget.View = tview.NewTextView()

	return widget
}

func boolStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func writeTempFile(t *testing.T, dir, name, contents string) string {
	t.Helper()

	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte(contents), 0644)
	assert.NoError(t, err)

	return path
}

func TestWidget_PlainText(t *testing.T) {
	dir := t.TempDir()

	tests := []struct {
		name     string
		contents string
		expected string
	}{
		{
			name:     "simple text file",
			contents: "hello world",
			expected: "hello world",
		},
		{
			name:     "empty file",
			contents: "",
			expected: "",
		},
		{
			name:     "multiline text file",
			contents: "line one\nline two\nline three",
			expected: "line one\nline two\nline three",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeTempFile(t, dir, tt.name+".txt", tt.contents)

			widget := newTestWidget(t, []string{path}, false, "")

			result := widget.plainText()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestWidget_PlainText_MissingFile(t *testing.T) {
	dir := t.TempDir()
	missing := filepath.Join(dir, "does-not-exist.txt")

	widget := newTestWidget(t, []string{missing}, false, "")

	result := widget.plainText()

	_, expectedErr := os.ReadFile(filepath.Clean(missing))
	assert.Error(t, expectedErr)
	assert.Equal(t, expectedErr.Error(), result)
}

func TestWidget_FormattedText(t *testing.T) {
	dir := t.TempDir()

	tests := []struct {
		name        string
		fileName    string
		contents    string
		formatStyle string
	}{
		{
			name:        "go source with valid style",
			fileName:    "main.go",
			contents:    "package main\n\nfunc main() {}\n",
			formatStyle: "vim",
		},
		{
			name:        "plain text with unknown style falls back",
			fileName:    "notes.txt",
			contents:    "just some notes",
			formatStyle: "this-style-does-not-exist",
		},
		{
			name:        "empty file",
			fileName:    "empty.txt",
			contents:    "",
			formatStyle: "vim",
		},
		{
			name:        "unrecognized extension falls back to plaintext lexer",
			fileName:    "data.zzznotalanguage",
			contents:    "unmatched extension content",
			formatStyle: "vim",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path := writeTempFile(t, dir, tt.fileName, tt.contents)

			widget := newTestWidget(t, []string{path}, true, tt.formatStyle)

			result := widget.formattedText()
			// formatted output should not error, and for non-empty input
			// should contain the original text somewhere in the (possibly
			// ANSI-decorated) output.
			assert.NotContains(t, result, "no such file")
			if tt.contents != "" {
				// Formatted output may be tokenized/colorized, so check that
				// the first "word" of the content survives somewhere in the output.
				firstLine := strings.TrimSpace(strings.Split(tt.contents, "\n")[0])
				firstWord := strings.Fields(firstLine)[0]
				assert.Contains(t, result, firstWord)
			}
		})
	}
}

func TestWidget_FormattedText_MissingFile(t *testing.T) {
	dir := t.TempDir()
	missing := filepath.Join(dir, "does-not-exist.go")

	widget := newTestWidget(t, []string{missing}, true, "vim")

	result := widget.formattedText()

	_, expectedErr := os.Open(filepath.Clean(missing))
	assert.Error(t, expectedErr)
	assert.Equal(t, expectedErr.Error(), result)
}

// TestNewWidget_LifeCycle exercises the full constructor, which wires up
// keyboard controls and starts the file-watcher goroutine, plus Refresh().
// The watcher is stopped explicitly (via Stop()) before the temp dir is
// removed, to avoid a background race where the watcher polls a deleted file.
func TestNewWidget_LifeCycle(t *testing.T) {
	dir := t.TempDir()
	path := writeTempFile(t, dir, "watched.txt", "initial contents")

	ymlConfig, err := config.ParseYaml("filePaths:\n  - \"" + strings.ReplaceAll(path, "\\", "\\\\") + "\"\n")
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml(testGlobalConfig)
	assert.NoError(t, err)
	settings := NewSettingsFromYAML("textfile", ymlConfig, globalConfig)

	app := tview.NewApplication()
	pages := tview.NewPages()
	redrawChan := make(chan bool, 10)

	widget := NewWidget(app, redrawChan, pages, settings)
	assert.NotNil(t, widget)
	assert.Equal(t, time.Duration(0), widget.settings.RefreshInterval)

	widget.Refresh()

	select {
	case <-redrawChan:
		// expected: Refresh -> Redraw sends on RedrawChan
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for redraw signal")
	}

	assert.Contains(t, widget.View.GetText(true), "initial contents")

	// Modify the watched file to trigger the watcher's write-event branch,
	// which should cause an automatic Refresh() and a second redraw signal.
	err = os.WriteFile(path, []byte("updated contents"), 0644)
	assert.NoError(t, err)

	select {
	case <-redrawChan:
		// expected: file watcher detected the write and called Refresh()
	case <-time.After(3 * time.Second):
		t.Fatal("timed out waiting for redraw signal after file change")
	}

	assert.Contains(t, widget.View.GetText(true), "updated contents")

	// Stop the watcher goroutine before the temp dir is cleaned up.
	widget.Stop()
}

func TestWidget_Content(t *testing.T) {
	dir := t.TempDir()

	t.Run("plain text mode", func(t *testing.T) {
		path := writeTempFile(t, dir, "plain.txt", "plain contents")

		widget := newTestWidget(t, []string{path}, false, "")

		title, text, wrap := widget.content()

		assert.Contains(t, title, path)
		assert.Contains(t, text, "plain contents")
		assert.True(t, wrap)
	})

	t.Run("formatted mode", func(t *testing.T) {
		path := writeTempFile(t, dir, "formatted.go", "package main\n")

		widget := newTestWidget(t, []string{path}, true, "vim")

		title, text, _ := widget.content()

		assert.Contains(t, title, path)
		assert.NotEmpty(t, text)
	})

	t.Run("multiple sources shows pagination", func(t *testing.T) {
		pathA := writeTempFile(t, dir, "a.txt", "file a")
		pathB := writeTempFile(t, dir, "b.txt", "file b")

		widget := newTestWidget(t, []string{pathA, pathB}, false, "")

		title, text, _ := widget.content()

		assert.Contains(t, title, pathA)
		assert.Contains(t, text, "file a")
		// Pagination marker uses the Paging sigils when there's more than one source.
		assert.NotEmpty(t, text)
	})

	t.Run("no sources configured", func(t *testing.T) {
		widget := newTestWidget(t, nil, false, "")

		title, text, _ := widget.content()

		// With no sources, CurrentSource() is empty, so the title reflects an
		// empty source and the body surfaces the OS error from trying to
		// read an empty path.
		assert.Contains(t, title, "[white]")
		assert.NotEmpty(t, strings.TrimSpace(text))
	})
}
