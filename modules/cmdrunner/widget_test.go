package cmdrunner

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

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
