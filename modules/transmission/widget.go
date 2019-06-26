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
	wtf.TextWidget

	settings *Settings
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		KeyboardWidget: wtf.NewKeyboardWidget(app, pages, settings.common),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)
	widget.KeyboardWidget.SetView(widget.View)

	return widget
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

	var content string
	if err != nil {
		content = err.Error()
	} else {
		content = widget.contentFrom(torrents)
	}

	widget.Redraw(widget.CommonSettings.Title, content, false)
}

// HelpText returns the help text for this widget
func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(data []*transmissionrpc.Torrent) string {
	str := ""

	for _, torrent := range data {
		str += fmt.Sprintf(
			" %s %s%s[white]\n",
			widget.torrentPercentDone(torrent),
			widget.torrentState(torrent),
			widget.prettyTorrentName(*torrent.Name),
		)
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
