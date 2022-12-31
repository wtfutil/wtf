package progress

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

var errShellUndefined = errors.New("command shell not defined in $SHELL environment variable")

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget

	settings *Settings

	minimum float64
	maximum float64
	current float64
	percent float64

	padding string

	shell string

	err error
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.common),

		settings: settings,

		minimum: settings.minimum,
		maximum: settings.maximum,
		current: settings.current,

		shell: os.Getenv("SHELL"),

		padding: strings.Repeat(" ", settings.padding),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	var err error

	if cmd := widget.settings.minimumCmd; cmd != "" {
		widget.minimum, err = widget.execValueCmd(cmd)
		if err != nil {
			widget.err = fmt.Errorf("minimumCmd execution failed: %w", err)
			widget.display()
			return
		}
	}

	if cmd := widget.settings.maximumCmd; cmd != "" {
		widget.maximum, err = widget.execValueCmd(cmd)
		if err != nil {
			widget.err = fmt.Errorf("maximumCmd execution failed: %w", err)
			widget.display()
			return
		}
	}

	if cmd := widget.settings.currentCmd; cmd != "" {
		widget.current, err = widget.execValueCmd(cmd)
		if err != nil {
			widget.err = fmt.Errorf("currentCmd execution failed: %w", err)
			widget.display()
			return
		}
	}

	widget.calcPercent()

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	if widget.err != nil {
		return "[red]Error: " + widget.err.Error()
	}

	pb := widget.buildProgress()

	return widget.padding + tview.TranslateANSI(pb.ViewAs(widget.percent)) + widget.padding
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (widget *Widget) execValueCmd(cmd string) (float64, error) {
	if widget.shell == "" {
		return -1, errShellUndefined
	}

	out, err := exec.Command(widget.shell, "-c", cmd).Output()
	if err != nil {
		return -1, err
	}

	outStr := strings.TrimSpace(string(out))

	val, err := strconv.ParseFloat(outStr, 64)
	if err != nil {
		return -1, fmt.Errorf("failed to parse command output '%s' as float64: %w", outStr, err)
	}

	return val, nil
}

func (widget *Widget) buildProgress() *progress.Model {
	_, _, width, _ := widget.View.GetInnerRect()

	pOpts := []progress.Option{
		progress.WithWidth(width - (widget.settings.padding * 2)),
	}

	if !widget.settings.showPercentage {
		pOpts = append(pOpts, progress.WithoutPercentage())
	}

	if widget.settings.colors.solid != "" {
		pOpts = append(pOpts, progress.WithSolidFill(widget.settings.colors.solid))
	} else {
		pOpts = append(pOpts, progress.WithGradient(
			widget.settings.colors.gradientA,
			widget.settings.colors.gradientB,
		))
	}

	pb := progress.New(pOpts...)
	return &pb
}

func (widget *Widget) calcPercent() {
	if widget.maximum == 0 {
		if widget.current > 100 {
			widget.percent = 1
		}

		if widget.current < 0 {
			widget.percent = 0
		}

		widget.percent = widget.current / 100
		return
	}

	if widget.current > widget.maximum {
		widget.percent = 1
		return
	}

	if widget.current < widget.minimum {
		widget.percent = 0
		return
	}

	widget.percent = (widget.current - widget.minimum) / (widget.maximum - widget.minimum)
}
