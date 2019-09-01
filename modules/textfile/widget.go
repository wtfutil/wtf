package textfile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/chroma/formatters"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	pollingIntervalms = 100
)

type Widget struct {
	view.KeyboardWidget
	view.MultiSourceWidget
	view.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    view.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: view.NewMultiSourceWidget(settings.common, "filePath", "filePaths"),
		TextWidget:        view.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	// Don't use a timer for this widget, watch for filesystem changes instead
	widget.settings.common.RefreshInterval = 0

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.SetDisplayFunction(widget.Refresh)
	widget.View.SetWordWrap(true)
	widget.View.SetWrap(settings.wrapText)

	widget.KeyboardWidget.SetView(widget.View)

	go widget.watchForFileChanges()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh is only called once on start-up. Its job is to display the
// text files that first time. After that, the watcher takes over
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf("[green]%s[white]", widget.CurrentSource())

	_, _, width, _ := widget.View.GetRect()
	text := widget.settings.common.SigilStr(len(widget.Sources), widget.Idx, width) + "\n"

	if widget.settings.format {
		text += widget.formattedText()
	} else {
		text += widget.plainText()
	}

	return title, text, widget.settings.wrapText
}

func (widget *Widget) fileName() string {
	return filepath.Base(widget.CurrentSource())
}

func (widget *Widget) formattedText() string {
	filePath, _ := utils.ExpandHomeDir(widget.CurrentSource())

	file, err := os.Open(filePath)
	if err != nil {
		return err.Error()
	}

	lexer := lexers.Match(filePath)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(widget.settings.formatStyle)
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	contents, _ := ioutil.ReadAll(file)
	iterator, _ := lexer.Tokenise(nil, string(contents))

	var buf bytes.Buffer
	formatter.Format(&buf, style, iterator)

	return tview.TranslateANSI(buf.String())
}

func (widget *Widget) plainText() string {
	filePath, _ := utils.ExpandHomeDir(widget.CurrentSource())

	text, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err.Error()
	}
	return string(text)
}

func (widget *Widget) watchForFileChanges() {
	watch := watcher.New()
	watch.FilterOps(watcher.Write)

	go func() {
		for {
			select {
			case <-watch.Event:
				widget.Refresh()
			case err := <-watch.Error:
				fmt.Println(err)
				os.Exit(1)
			case <-watch.Closed:
				return
			}
		}
	}()

	// Watch each textfile for changes
	for _, source := range widget.Sources {
		fullPath, err := utils.ExpandHomeDir(source)
		if err == nil {
			if err := watch.Add(fullPath); err != nil {
				// Ignore it, don't care about a file that doesn't exist
			}
		}
	}

	// Start the watching process - it'll check for changes every pollingIntervalms.
	if err := watch.Start(time.Millisecond * pollingIntervalms); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
