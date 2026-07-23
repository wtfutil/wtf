package cmdrunner

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"

	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

// newTestCommon builds a real *cfg.Common (with colors populated) so that
// widgets relying on it (e.g. via view.NewTextWidget) don't hit nil pointers.
func newTestCommon(t *testing.T) *cfg.Common {
	t.Helper()

	moduleConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)

	globalConfig, err := config.ParseYaml("wtf: {}")
	assert.NoError(t, err)

	return cfg.NewCommonSettingsFromModule("cmdrunner", defaultTitle, defaultFocusable, moduleConfig, globalConfig)
}

func Test_expandTilde(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("could not determine home dir: %v", err)
	}

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			name: "tilde with subpath is expanded",
			path: "~/bin/script.sh",
			want: filepath.Join(home, "bin/script.sh"),
		},
		{
			name: "bare tilde is expanded",
			path: "~",
			want: home,
		},
		{
			name: "path without tilde is unchanged",
			path: "/usr/bin/ping",
			want: "/usr/bin/ping",
		},
		{
			name: "relative path is unchanged",
			path: ".",
			want: ".",
		},
		{
			name: "empty string is unchanged",
			path: "",
			want: "",
		},
		{
			name: "user-specific tilde falls back to original value",
			path: "~someuser/foo",
			want: "~someuser/foo",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := expandTilde(tt.path)
			if actual != tt.want {
				t.Errorf("expandTilde(%q) = %q, want %q", tt.path, actual, tt.want)
			}
		})
	}
}

func Test_expandTildes(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("could not determine home dir: %v", err)
	}

	input := []string{"~/a", "-I", "cisco.com", "~/b/c"}
	want := []string{
		filepath.Join(home, "a"),
		"-I",
		"cisco.com",
		filepath.Join(home, "b/c"),
	}

	actual := expandTildes(input)
	if len(actual) != len(want) {
		t.Fatalf("expandTildes(%v) = %v, want %v", input, actual, want)
	}
	for i := range want {
		if actual[i] != want[i] {
			t.Errorf("expandTildes(%v)[%d] = %q, want %q", input, i, actual[i], want[i])
		}
	}
}

// Test_runCommandLoop_tildeExpansion is an end-to-end check that mirrors the
// exec.Command construction in runCommandLoop: it points HOME/USERPROFILE at
// a temp dir, drops a marker file there, and then runs a real command whose
// cmd/args/workingDir are all `~`-prefixed - exactly what a user's cmdrunner
// config would look like. This confirms the fix actually resolves `~` at
// execution time, not just that the helper functions return the right string.
func Test_runCommandLoop_tildeExpansion(t *testing.T) {
	tempHome := t.TempDir()
	if runtime.GOOS == "windows" {
		t.Setenv("USERPROFILE", tempHome)
	} else {
		t.Setenv("HOME", tempHome)
	}

	const markerContents = "hello from tilde expansion"
	markerPath := filepath.Join(tempHome, "wtf_cmdrunner_test_marker.txt")
	if err := os.WriteFile(markerPath, []byte(markerContents), 0o644); err != nil {
		t.Fatalf("could not write marker file: %v", err)
	}

	// cmd/args as a user would write them in a wtf config: cmd reads the
	// marker file via a `~`-prefixed path, workingDir is also `~`.
	var cmdName string
	var args []string
	if runtime.GOOS == "windows" {
		cmdName = "cmd"
		args = []string{"/c", "type", "~\\wtf_cmdrunner_test_marker.txt"}
	} else {
		cmdName = "cat"
		args = []string{"~/wtf_cmdrunner_test_marker.txt"}
	}

	cmd := exec.Command(expandTilde(cmdName), expandTildes(args)...)
	cmd.Dir = expandTilde("~")

	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("command failed: %v", err)
	}

	if !strings.Contains(string(out), markerContents) {
		t.Errorf("expected output to contain %q, got %q", markerContents, string(out))
	}
}

func TestWidget_String(t *testing.T) {
	tests := []struct {
		name     string
		cmd      string
		args     []string
		expected string
	}{
		{
			name:     "command with no args",
			cmd:      "whoami",
			args:     []string{},
			expected: "whoami",
		},
		{
			name:     "command with a single arg",
			cmd:      "echo",
			args:     []string{"hello"},
			expected: "echo hello",
		},
		{
			name:     "command with multiple args",
			cmd:      "curl",
			args:     []string{"-I", "cisco.com"},
			expected: "curl -I cisco.com",
		},
		{
			name:     "empty command",
			cmd:      "",
			args:     []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: &Settings{cmd: tt.cmd, args: tt.args},
			}

			assert.Equal(t, tt.expected, widget.String())
		})
	}
}

func TestWidget_Write_And_LineDraining(t *testing.T) {
	tests := []struct {
		name          string
		maxLines      int
		input         string
		expectedLines int
	}{
		{
			name:          "under the limit keeps everything",
			maxLines:      5,
			input:         "one\ntwo\nthree\n",
			expectedLines: 3,
		},
		{
			name:          "at the limit keeps everything",
			maxLines:      3,
			input:         "one\ntwo\nthree\n",
			expectedLines: 3,
		},
		{
			name:          "over the limit drains the oldest lines",
			maxLines:      2,
			input:         "one\ntwo\nthree\nfour\n",
			expectedLines: 2,
		},
		{
			name:          "zero max lines drains everything with a trailing newline",
			maxLines:      0,
			input:         "one\ntwo\n",
			expectedLines: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: &Settings{maxLines: tt.maxLines},
				buffer:   &bytes.Buffer{},
			}

			n, err := widget.Write([]byte(tt.input))

			assert.NoError(t, err)
			assert.Equal(t, len(tt.input), n)
			assert.Equal(t, tt.expectedLines, widget.countLines())
		})
	}
}

func TestWidget_drainLines_ErrorWhenNotEnoughLines(t *testing.T) {
	widget := &Widget{
		settings: &Settings{maxLines: 100},
		buffer:   bytes.NewBufferString("only one line\n"),
	}

	err := widget.drainLines(5)

	assert.Error(t, err)
}

func TestWidget_environment(t *testing.T) {
	widget := &Widget{
		settings: &Settings{width: 42, height: 24},
	}

	envs := widget.environment()

	assert.Contains(t, envs, "WTF_WIDGET_WIDTH=42")
	assert.Contains(t, envs, "WTF_WIDGET_HEIGHT=24")
	// Also inherits the process environment.
	assert.Greater(t, len(envs), 2)
}

func TestWidget_handleError(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		wantContain string
	}{
		{
			name:        "generic error",
			err:         assert.AnError,
			wantContain: assert.AnError.Error(),
		},
		{
			name:        "exec error",
			err:         exec.ErrNotFound,
			wantContain: exec.ErrNotFound.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: &Settings{maxLines: 100},
				buffer:   &bytes.Buffer{},
			}

			widget.handleError(tt.err)

			assert.Contains(t, widget.buffer.String(), tt.wantContain)
		})
	}
}

func TestWidget_resetBuffer(t *testing.T) {
	widget := &Widget{
		settings: &Settings{maxLines: 100},
		buffer:   bytes.NewBufferString("some previous output"),
	}

	widget.resetBuffer()

	assert.Equal(t, "", widget.buffer.String())
}

// echoCommand returns an *exec.Cmd that safely prints text to stdout on the
// current platform without depending on a shell being on PATH in an unusual
// location.
func echoCommand(text string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/c", "echo", text)
	}
	return exec.Command("echo", text)
}

func TestRunCommand(t *testing.T) {
	tests := []struct {
		name        string
		cmd         func() *exec.Cmd
		wantErr     bool
		wantContain string
	}{
		{
			name: "successful echo command",
			cmd: func() *exec.Cmd {
				return echoCommand("hello wtf")
			},
			wantErr:     false,
			wantContain: "hello wtf",
		},
		{
			name: "command that does not exist",
			cmd: func() *exec.Cmd {
				return exec.Command("this-binary-should-not-exist-xyz123")
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := &Widget{
				settings: &Settings{maxLines: 100},
				buffer:   &bytes.Buffer{},
			}

			err := runCommand(widget, tt.cmd())

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Contains(t, widget.buffer.String(), tt.wantContain)
			}
		})
	}
}

func TestWidget_content(t *testing.T) {
	common := newTestCommon(t)
	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, common),
		settings:   &Settings{cmd: "whoami", args: []string{}},
		buffer:     bytes.NewBufferString("some output\n"),
	}

	title, content, wrap := widget.content()

	assert.False(t, wrap)
	// The default title falls back to the rendered command string.
	assert.Contains(t, title, "whoami")
	assert.Contains(t, content, "some output")
}

func TestWidget_content_UsesCommandAsTitleWhenDefault(t *testing.T) {
	common := newTestCommon(t)
	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 1)

	widget := &Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, common),
		settings:   &Settings{cmd: "ping", args: []string{"-c", "1"}},
		buffer:     &bytes.Buffer{},
	}

	title, _, _ := widget.content()

	assert.Contains(t, title, "ping -c 1")
}

func TestWidget_Refresh_Integration(t *testing.T) {
	common := newTestCommon(t)
	settings := &Settings{
		Common:   common,
		cmd:      echoCommand("integration-test").Path,
		args:     echoCommand("integration-test").Args[1:],
		maxLines: 100,
	}

	tviewApp := tview.NewApplication()
	redrawChan := make(chan bool, 10)

	widget := NewWidget(tviewApp, redrawChan, settings)

	// The redraw loop fires once immediately (before the command has run) and
	// again once the command completes, so poll until the expected output
	// shows up in the buffer or we time out.
	deadline := time.After(5 * time.Second)
	for {
		select {
		case <-redrawChan:
			_, content, _ := widget.content()
			if bytes.Contains([]byte(content), []byte("integration-test")) {
				return
			}
		case <-deadline:
			_, content, _ := widget.content()
			t.Fatalf("timed out waiting for command output, last content: %q", content)
		}
	}
}
