package textfile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wtfutil/wtf/view"
)

// newTestWidget creates a minimal Widget suitable for testing formattedText/plainText.
func newTestWidget(filePath string, format bool, formatStyle string) *Widget {
	msw := view.MultiSourceWidget{
		Idx:     0,
		Sources: []string{filePath},
	}

	settings := &Settings{
		format:      format,
		formatStyle: formatStyle,
	}

	return &Widget{
		MultiSourceWidget: msw,
		settings:          settings,
	}
}

// --- Tests for pure functions (no widget needed) ---

func TestFormatFile_KnownFileType(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "hello.go")
	content := "package main\n\nfunc main() {}\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result := formatFile(filePath, "monokai")

	if !strings.Contains(result, "\x1b[") && !strings.Contains(result, "[") {
		t.Error("expected formatted output to contain color/escape sequences")
	}
	if !strings.Contains(result, "package") {
		t.Errorf("expected output to contain 'package', got: %s", result)
	}
}

func TestFormatFile_NonExistentFile(t *testing.T) {
	result := formatFile("/nonexistent/path/file.go", "monokai")

	if !strings.Contains(result, "no such file") && !strings.Contains(result, "cannot find") &&
		!strings.Contains(result, "The system cannot find") {
		t.Errorf("expected error message for non-existent file, got: %s", result)
	}
}

func TestFormatFile_UnknownExtension(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "data.xyz123unknown")
	content := "some random content\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result := formatFile(filePath, "monokai")

	if strings.Contains(result, "no such file") || strings.Contains(result, "cannot find") {
		t.Errorf("unexpected error for unknown extension: %s", result)
	}
	if !strings.Contains(result, "some random content") {
		t.Errorf("expected output to contain original content, got: %s", result)
	}
}

func TestFormatFile_JSONFile(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "config.json")
	content := `{"key": "value", "num": 42}`
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result := formatFile(filePath, "vim")

	if !strings.Contains(result, "key") {
		t.Errorf("expected output to contain 'key', got: %s", result)
	}
}

func TestFormatFile_FallbackStyle(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "test.py")
	content := "def hello():\n    print('hi')\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	// Invalid style name should fall back gracefully
	result := formatFile(filePath, "nonexistent-style-name")

	if !strings.Contains(result, "hello") {
		t.Errorf("expected output to contain 'hello', got: %s", result)
	}
}

func TestReadPlainFile_ReturnsContent(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "sample.txt")
	content := "Hello, World!\nLine two.\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result := readPlainFile(filePath)

	if result != content {
		t.Errorf("expected %q, got %q", content, result)
	}
}

func TestReadPlainFile_EscapesTviewTags(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "tags.txt")
	content := "[red]text[white]\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	result := readPlainFile(filePath)

	// tview.Escape doubles the opening bracket: [red] -> [[red]
	if strings.Contains(result, "[red]") && !strings.Contains(result, "[[red]") {
		t.Errorf("expected tview tags to be escaped, got: %s", result)
	}
}

func TestReadPlainFile_NonExistentFile(t *testing.T) {
	result := readPlainFile("/nonexistent/path/file.txt")

	if !strings.Contains(result, "no such file") && !strings.Contains(result, "cannot find") &&
		!strings.Contains(result, "The system cannot find") {
		t.Errorf("expected error message for non-existent file, got: %s", result)
	}
}

// --- Tests for widget methods (thin wrappers) ---

func TestWidget_FormattedText(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "main.go")
	if err := os.WriteFile(filePath, []byte("package main\n"), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(filePath, true, "monokai")
	result := widget.formattedText()

	if !strings.Contains(result, "package") {
		t.Errorf("expected widget.formattedText() to contain 'package', got: %s", result)
	}
}

func TestWidget_PlainText(t *testing.T) {
	dir := t.TempDir()
	filePath := filepath.Join(dir, "note.txt")
	content := "just plain text\n"
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	widget := newTestWidget(filePath, false, "")
	result := widget.plainText()

	if result != content {
		t.Errorf("expected %q, got %q", content, result)
	}
}
