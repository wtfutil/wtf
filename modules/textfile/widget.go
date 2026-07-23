package textfile

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"github.com/radovskyb/watcher"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

const (
	pollingIntervalms = 100
)

type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	settings    *Settings
	fileWatcher *watcher.Watcher
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "filePath", "filePaths"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	// Don't use a timer for this widget, watch for filesystem changes instead
	widget.settings.RefreshInterval = 0

	widget.initializeKeyboardControls()

	widget.SetDisplayFunction(widget.Refresh)
	widget.View.SetWordWrap(true)
	widget.View.SetWrap(settings.wrapText)

	go widget.watchForFileChanges()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh is only called once on start-up. Its job is to display the
// text files that first time. After that, the watcher takes over
func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	title := fmt.Sprintf(
		"[%s]%s[white]",
		widget.settings.Colors.Title,
		widget.CurrentSource(),
	)

	_, _, width, _ := widget.View.GetRect()
	text := widget.settings.PaginationMarker(len(widget.Sources), widget.Idx, width) + "\n"

	if widget.settings.format {
		text += widget.formattedText()
	} else {
		text += widget.plainText()
	}

	return title, text, widget.settings.wrapText
}

func (widget *Widget) formattedText() string {
	filePath, _ := utils.ExpandHomeDir(widget.CurrentSource())
	return formatFile(filePath, widget.settings.formatStyle)
}

func (widget *Widget) plainText() string {
	filePath, _ := utils.ExpandHomeDir(filepath.Clean(widget.CurrentSource()))
	return readPlainFile(filePath)
}

// formatFile reads the file at filePath, applies syntax highlighting using the
// given chroma style name, and returns terminal256-formatted output.
func formatFile(filePath, styleName string) string {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err.Error()
	}
	defer func() { _ = file.Close() }()

	lexer := lexers.Match(filePath)
	if lexer == nil {
		lexer = lexers.Fallback
	}

	style := styles.Get(styleName)
	if style == nil {
		style = styles.Fallback
	}
	formatter := formatters.Get("terminal256")
	if formatter == nil {
		formatter = formatters.Fallback
	}

	contents, _ := io.ReadAll(file)
	str := string(contents)
	str = tview.Escape(str)
	iterator, _ := lexer.Tokenise(nil, str)

	var buf bytes.Buffer
	err = formatter.Format(&buf, style, iterator)
	if err != nil {
		return err.Error()
	}

	return tview.TranslateANSI(buf.String())
}

// readPlainFile reads the file at filePath and returns its content with tview tags escaped.
func readPlainFile(filePath string) string {
	text, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return err.Error()
	}
	return tview.Escape(string(text))
}

func (widget *Widget) watchForFileChanges() {
	widget.fileWatcher = watcher.New()
	watch := widget.fileWatcher
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
			case quit := <-widget.QuitChan():
				if quit {
					watch.Close()
					return
				}
			}
		}
	}()

	// Watch each textfile for changes
	for _, source := range widget.Sources {
		fullPath, err := utils.ExpandHomeDir(source)
		if err == nil {
			e := watch.Add(fullPath)
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
		}
	}

	// Start the watching process - it'll check for changes every pollingIntervalms.
	if err := watch.Start(time.Millisecond * pollingIntervalms); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
