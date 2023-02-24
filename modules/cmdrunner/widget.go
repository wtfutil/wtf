package cmdrunner

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/creack/pty"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/view"
)

// Widget contains the data for this widget
type Widget struct {
	view.TextWidget

	settings *Settings

	m          sync.Mutex
	buffer     *bytes.Buffer
	runChan    chan bool
	redrawChan chan bool
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
		buffer:   &bytes.Buffer{},
	}

	widget.View.SetWrap(true)
	widget.View.SetScrollable(true)

	widget.runChan = make(chan bool)
	widget.redrawChan = make(chan bool)
	go runCommandLoop(&widget)
	go redrawLoop(&widget)
	widget.runChan <- true

	return &widget
}

// Refresh signals the runCommandLoop to continue, or triggers a re-draw if the
// command is still running.
func (widget *Widget) Refresh() {
	// Try to run the command. If the command is still running, let it keep
	// running and do a refresh instead. Otherwise, the widget will redraw when
	// the command completes.
	select {
	case widget.runChan <- true:
	default:
		widget.redrawChan <- true
	}
}

// String returns the string representation of the widget
func (widget *Widget) String() string {
	args := strings.Join(widget.settings.args, " ")

	if args != "" {
		return fmt.Sprintf("%s %s", widget.settings.cmd, args)
	}

	return widget.settings.cmd
}

func (widget *Widget) Write(p []byte) (n int, err error) {
	widget.m.Lock()
	defer widget.m.Unlock()

	// Write the new data into the buffer
	n, err = widget.buffer.Write(p)

	// Remove lines that exceed maxLines
	lines := widget.countLines()
	if lines > widget.settings.maxLines {
		err = widget.drainLines(lines - widget.settings.maxLines)
	}

	return n, err
}

/* -------------------- Unexported Functions -------------------- */

// countLines counts the lines of data in the buffer
func (widget *Widget) countLines() int {
	return bytes.Count(widget.buffer.Bytes(), []byte{'\n'})
}

// drainLines removed the first n lines from the buffer
func (widget *Widget) drainLines(n int) error {
	for i := 0; i < n; i++ {
		_, err := widget.buffer.ReadBytes('\n')
		if err != nil {
			return err
		}
	}

	return nil
}

func (widget *Widget) environment() []string {
	envs := os.Environ()
	envs = append(
		envs,
		fmt.Sprintf("WTF_WIDGET_WIDTH=%d", widget.settings.width),
		fmt.Sprintf("WTF_WIDGET_HEIGHT=%d", widget.settings.height),
	)
	return envs
}

func runCommandLoop(widget *Widget) {
	// Run the command forever in a loop. Refresh() will put a value into the
	// channel to signal the loop to continue.
	for {
		<-widget.runChan
		widget.resetBuffer()
		cmd := exec.Command(widget.settings.cmd, widget.settings.args...)
		cmd.Env = widget.environment()
		cmd.Dir = widget.settings.workingDir
		var err error
		if widget.settings.pty {
			err = runCommandPty(widget, cmd)
		} else {
			err = runCommand(widget, cmd)
		}
		if err != nil {
			widget.handleError(err)
		}
		widget.redrawChan <- true
	}
}

func runCommand(widget *Widget, cmd *exec.Cmd) error {
	cmd.Stdout = widget
	return cmd.Run()
}

func runCommandPty(widget *Widget, cmd *exec.Cmd) error {
	f, err := pty.Start(cmd)
	// The command has exited, print any error messages
	if err != nil {
		return err
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
		return err
	}
	return cmd.Wait()
}

func (widget *Widget) handleError(err error) {
	widget.m.Lock()
	defer widget.m.Unlock()
	_, writeErr := widget.buffer.WriteString(err.Error())
	if writeErr != nil {
		return
	}
}

func redrawLoop(widget *Widget) {
	for {
		widget.Redraw(widget.content)
		if widget.settings.tail {
			widget.View.ScrollToEnd()
		}
		<-widget.redrawChan
	}
}

func (widget *Widget) content() (string, string, bool) {
	widget.m.Lock()
	result := widget.buffer.String()
	widget.m.Unlock()

	ansiTitle := tview.TranslateANSI(tview.Escape(widget.CommonSettings().Title))
	if ansiTitle == defaultTitle {
		ansiTitle = tview.TranslateANSI(tview.Escape(widget.String()))
	}
	ansiResult := tview.TranslateANSI(tview.Escape(result))

	return ansiTitle, ansiResult, false
}

func (widget *Widget) resetBuffer() {
	widget.m.Lock()
	defer widget.m.Unlock()

	widget.buffer.Reset()
}
