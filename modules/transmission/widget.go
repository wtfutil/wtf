package transmission

import (
	"fmt"
	"strings"

	"github.com/hekmon/transmissionrpc"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

// Widget is the container for transmission data
type Widget struct {
	wtf.KeyboardWidget
	wtf.ScrollableWidget

	settings *Settings
	torrents []*transmissionrpc.Torrent
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   wtf.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: wtf.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves torrent data from the Transmission daemon
func (widget *Widget) Fetch() ([]*transmissionrpc.Torrent, error) {
	transmissionbt, err := transmissionrpc.New(widget.settings.host, widget.settings.username, widget.settings.password, nil)
	if err != nil {
		return nil, err
	}

	torrents, err := transmissionbt.TorrentGetAll()
	if err != nil {
		return nil, err
	}

	return torrents, nil
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	torrents, err := widget.Fetch()
	if err != nil {
		widget.SetItemCount(0)
		widget.ScrollableWidget.Redraw(widget.CommonSettings.Title, err.Error(), false)
		return
	}

	widget.torrents = torrents
	widget.SetItemCount(len(torrents))

	widget.display()
}

// HelpText returns the help text for this widget
func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

// Next selects the next item in the list
func (widget *Widget) Next() {
	widget.ScrollableWidget.Next()
}

// Prev selects the previous item in the list
func (widget *Widget) Prev() {
	widget.ScrollableWidget.Prev()
}

// Unselect clears the selection of list items
func (widget *Widget) Unselect() {
	widget.ScrollableWidget.Unselect()
	widget.RenderFunction()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) display() {
	if len(widget.torrents) == 0 {
		widget.ScrollableWidget.Redraw(widget.CommonSettings.Title, "no torrents", false)
		return
	}

	content := widget.contentFrom(widget.torrents)
	widget.ScrollableWidget.Redraw(widget.CommonSettings.Title, content, false)
}

func (widget *Widget) contentFrom(data []*transmissionrpc.Torrent) string {
	str := ""

	for idx, torrent := range data {
		torrName := *torrent.Name

		row := fmt.Sprintf(
			"[%s] %s %s%s[white]",
			widget.RowColor(idx),
			widget.torrentPercentDone(torrent),
			widget.torrentState(torrent),
			tview.Escape(widget.prettyTorrentName(torrName)),
		)

		str += wtf.HighlightableHelper(widget.View, row, idx, len(torrName))
	}

	return str
}

func (widget *Widget) prettyTorrentName(name string) string {
	str := strings.Replace(name, "[", "(", -1)
	str = strings.Replace(str, "]", ")", -1)

	return str
}

func (widget *Widget) torrentPercentDone(torrent *transmissionrpc.Torrent) string {
	pctDone := *torrent.PercentDone
	str := fmt.Sprintf("%3d", int(pctDone*100))

	if pctDone == 0.0 {
		str = "[gray::b]" + str
	} else if pctDone == 1.0 {
		str = "[green::b]" + str
	} else {
		str = "[lightblue::b]" + str
	}

	return str + "[white]"
}

func (widget *Widget) torrentState(torrent *transmissionrpc.Torrent) string {
	str := ""

	switch *torrent.Status {
	case transmissionrpc.TorrentStatusStopped:
		str += "[gray]"
	case transmissionrpc.TorrentStatusDownload:
		str += "[lightblue]"
	case transmissionrpc.TorrentStatusSeed:
		str += "[green]"
	}

	return str
}
