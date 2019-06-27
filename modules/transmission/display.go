package transmission

import (
	"fmt"
	"strings"

	"github.com/hekmon/transmissionrpc"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

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

func (widget *Widget) display() {
	if len(widget.torrents) == 0 {
		widget.ScrollableWidget.Redraw(widget.CommonSettings.Title, "no torrents", false)
		return
	}

	content := widget.contentFrom(widget.torrents)
	widget.ScrollableWidget.Redraw(widget.CommonSettings.Title, content, false)
}

func (widget *Widget) prettyTorrentName(name string) string {
	str := strings.Replace(name, "[", "(", -1)
	str = strings.Replace(str, "]", ")", -1)

	return str
}

func (widget *Widget) torrentPercentDone(torrent *transmissionrpc.Torrent) string {
	pctDone := *torrent.PercentDone
	str := fmt.Sprintf("%3d%%", int(pctDone*100))

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
