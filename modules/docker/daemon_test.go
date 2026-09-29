package docker

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_resolvePidFilePath(t *testing.T) {
	tests := []struct {
		name       string
		configured string
		daemonHost string
		expected   string
	}{
		{"unset", "", "unix:///var/run/docker.sock", ""},
		{"explicit path", "/tmp/docker.pid", "unix:///var/run/docker.sock", "/tmp/docker.pid"},
		{"auto, rootful", "auto", "unix:///var/run/docker.sock", "/var/run/docker.pid"},
		{"auto, rootless", "auto", "unix:///run/user/1000/docker.sock", "/run/user/1000/docker.pid"},
		{"auto, remote", "auto", "tcp://192.168.1.10:2376", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, resolvePidFilePath(tt.configured, tt.daemonHost))
		})
	}
}

func Test_daemonIsRunning(t *testing.T) {
	dir := t.TempDir()

	pidFile := func(name, contents string) string {
		path := filepath.Join(dir, name)
		assert.NoError(t, os.WriteFile(path, []byte(contents), 0o644))
		return path
	}

	tests := []struct {
		name    string
		path    string
		running bool
		known   bool
	}{
		{"no path", "", false, false},
		{"absent file", filepath.Join(dir, "absent.pid"), false, true},
		{"live pid", pidFile("live.pid", strconv.Itoa(os.Getpid())), true, true},
		{"stale pid", pidFile("stale.pid", "0"), false, true},
		{"malformed", pidFile("garbage.pid", "nope"), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			running, known := daemonIsRunning(tt.path)
			assert.Equal(t, tt.known, known)
			assert.Equal(t, tt.running, running)
		})
	}
}
