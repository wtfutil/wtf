//go:build !windows

package cmdrunner

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/creack/pty"
	"github.com/wtfutil/wtf/logger"
)

func runCommandPty(widget *Widget, cmd *exec.Cmd) error {
	f, err := pty.Start(cmd)
	// The command has exited, print any error messages
	if err != nil {
		if widget.settings.ptySuppressErrors {
			return cmd.Wait()
		} else {
			return err
		}
	}

	// Make sure to close the pty at the end.
	defer func() { _ = f.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, f); err != nil {
				logger.Log(fmt.Sprintf("error resizing pty: %s", err))
			}
		}
	}()
	ch <- syscall.SIGWINCH                        // Initial resize.
	defer func() { signal.Stop(ch); close(ch) }() // Cleanup signals when done.

	// Extract output
	_, err = io.Copy(widget.buffer, f)
	if err != nil {
		if widget.settings.ptySuppressErrors && errors.Is(err, syscall.EIO) {
			return cmd.Wait()
		}
		return err
	}
	return cmd.Wait()
}
