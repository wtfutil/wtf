package progress

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/muesli/reflow/ansi"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
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

	percent := widget.formatPercent(widget.percent)
	bar := widget.buildProgressBar(percent)
	barView := tview.TranslateANSI(bar.ViewAs(widget.percent))

	var sb strings.Builder

	switch widget.settings.showPercentage {
	case "left":
		sb.WriteString(widget.padding + percent + barView + widget.padding)
	case "right":
		sb.WriteString(widget.padding + barView + percent + widget.padding)
	case "above":
		centered := utils.CenterText(percent, bar.Width+widget.settings.padding*2)
		sb.WriteString(centered + "\n" + widget.padding + barView + widget.padding)
	case "below":
		centered := utils.CenterText(percent, bar.Width+widget.settings.padding*2)
		sb.WriteString(widget.padding + barView + widget.padding + "\n" + centered)
	default:
		sb.WriteString(widget.padding + barView + widget.padding)
	}

	return sb.String()
}

func (widget *Widget) display() {
	title := widget.CommonSettings().Title

	if widget.settings.showPercentage == "titleLeft" {
		title = widget.formatPercent(widget.percent) + " " + title
	} else if widget.settings.showPercentage == "titleRight" {
		title = title + " " + widget.formatPercent(widget.percent)
	}

	widget.Redraw(func() (string, string, bool) {
		return title, widget.content(), false
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

func (widget *Widget) buildProgressBar(percent string) *progress.Model {
	pOpts := []progress.Option{
		progress.WithWidth(widget.calcBarWidth(percent)),
		progress.WithoutPercentage(),
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

func (widget *Widget) formatPercent(p float64) string {
	switch widget.settings.showPercentage {
	case "left":
		return fmt.Sprintf("%.0f%% ", p*100)
	case "right":
		return fmt.Sprintf(" %.0f%%", p*100)
	case "none":
		return ""
	default:
		return fmt.Sprintf("%.0f%%", p*100)
	}
}

func (widget *Widget) calcBarWidth(percent string) int {
	_, _, width, _ := widget.View.GetInnerRect()
	width -= widget.settings.padding * 2

	if widget.settings.showPercentage == "left" || widget.settings.showPercentage == "right" {
		width -= ansi.PrintableRuneWidth(percent)
	}

	return width
}
