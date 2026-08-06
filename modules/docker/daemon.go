package docker

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/docker/pkg/pidfile"
)

// resolvePidFilePath expands "auto" into a guess at the pid file's location, or
// "" when there is no local pid file to read. Set the path explicitly if the
// daemon runs with a --pidfile that isn't next to its socket.
func resolvePidFilePath(configured string, daemonHost string) string {
	if configured != "auto" {
		return configured
	}

	socket, ok := strings.CutPrefix(daemonHost, "unix://")
	if !ok {
		return ""
	}

	return filepath.Join(filepath.Dir(socket), "docker.pid")
}

// daemonIsRunning reports whether dockerd is up, judged only from its pid file:
// under socket activation, asking the API is what starts it.
func daemonIsRunning(path string) (running bool, known bool) {
	if path == "" {
		return false, false
	}

	pid, err := pidfile.Read(path)
	if os.IsNotExist(err) {
		return false, true
	}
	if err != nil {
		return false, false
	}

	// pidfile.Read reports 0 for a stale file.
	return pid != 0, true
}
