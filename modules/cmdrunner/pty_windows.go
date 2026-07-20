//go:build windows

package cmdrunner

import (
	"errors"
	"os/exec"
)

// runCommandPty is not supported on Windows. PTY mode requires Unix-specific
// syscalls (SIGWINCH) and the creack/pty library which does not support Windows.
func runCommandPty(widget *Widget, cmd *exec.Cmd) error {
	return errors.New("pty mode is not supported on Windows")
}
