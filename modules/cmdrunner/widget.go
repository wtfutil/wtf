package cmdrunner

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget contains the data for this widget
type Widget struct {
	view.TextWidget

	settings *Settings

	m       sync.Mutex
	buffer  *bytes.Buffer
	running bool
}

// NewWidget creates a new instance of the widget
func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

		settings: settings,
		buffer:   &bytes.Buffer{},
	}

	widget.View.SetWrap(true)
	widget.View.SetScrollable(true)

	return &widget
}

func (widget *Widget) content() (string, string, bool) {
	result := widget.buffer.String()

	ansiTitle := tview.TranslateANSI(widget.CommonSettings().Title)
	if ansiTitle == defaultTitle {
		ansiTitle = tview.TranslateANSI(widget.String())
	}
	ansiResult := tview.TranslateANSI(result)

	return ansiTitle, ansiResult, false
}

// Refresh executes the command and updates the view with the results
func (widget *Widget) Refresh() {
	widget.m.Lock()
	defer widget.m.Unlock()

	widget.execute()
	widget.Redraw(widget.content)
	if widget.settings.tail {
		widget.View.ScrollToEnd()
	}
}

// String returns the string representation of the widget
func (widget *Widget) String() string {
	args := strings.Join(widget.settings.args, " ")

	if args != "" {
		return fmt.Sprintf(" %s %s ", widget.settings.cmd, args)
	}

	return fmt.Sprintf(" %s ", widget.settings.cmd)
}

func (widget *Widget) Write(p []byte) (n int, err error) {
	widget.m.Lock()
	defer widget.m.Unlock()

	// Write the new data into the buffer
	n, err = widget.buffer.Write(p)

	// Remove lines that exceed maxLines
	lines := widget.countLines()
	if lines > widget.settings.maxLines {
		widget.drainLines(lines - widget.settings.maxLines)
	}

	// Redraw the widget
	widget.Redraw(widget.content)

	return
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) execute() {
	// Make sure the command is not already running
	if widget.running {
		return
	}

	// Reset the buffer
	widget.buffer.Reset()

	// Indicate that the command is running
	widget.running = true

	// Setup the command to run
	cmd := exec.Command(widget.settings.cmd, widget.settings.args...)
	cmd.Stdout = widget
	cmd.Env = widget.environment()

	// Run the command and wait for it to exit in another Go-routine
	go func() {
		err := cmd.Run()

		// The command has exited, print any error messages
		widget.m.Lock()
		if err != nil {
			widget.buffer.WriteString(err.Error())
		}
		widget.running = false
		widget.m.Unlock()
	}()
}

// countLines counts the lines of data in the buffer
func (widget *Widget) countLines() int {
	return bytes.Count(widget.buffer.Bytes(), []byte{'\n'})
}

// drainLines removed the first n lines from the buffer
func (widget *Widget) drainLines(n int) {
	for i := 0; i < n; i++ {
		widget.buffer.ReadBytes('\n')
	}
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
