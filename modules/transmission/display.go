package transmission

import (
	"fmt"
	"strings"

	"github.com/hekmon/transmissionrpc/v2"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
)

func (widget *Widget) display() {
	widget.ScrollableWidget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	widget.mu.Lock()
	defer widget.mu.Unlock()

	title := widget.CommonSettings().Title
	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(widget.torrents) == 0 {
		return title, "No data", false
	}

	str := ""

	for idx, torrent := range widget.torrents {
		torrName := *torrent.Name

		row := fmt.Sprintf(
			"[%s] %s %s %s%s[white]",
			widget.RowColor(idx),
			widget.torrentPercentDone(torrent),
			widget.torrentSeedRatio(torrent),
			widget.torrentState(torrent),
			tview.Escape(widget.prettyTorrentName(torrName)),
		)

		str += utils.HighlightableHelper(widget.View, row, idx, len(torrName))
	}

	return title, str, false
}

func (widget *Widget) prettyTorrentName(name string) string {
	str := strings.ReplaceAll(name, "[", "(")
	str = strings.ReplaceAll(str, "]", ")")

	return str
}

func (widget *Widget) torrentPercentDone(torrent transmissionrpc.Torrent) string {
	pctDone := *torrent.PercentDone
	str := fmt.Sprintf("%3d%%↓", int(pctDone*100))

	switch pctDone {
	case 0.0:
		str = "[gray::b]" + str
	case 1.0:
		str = "[green::b]" + str
	default:
		str = "[lightblue::b]" + str
	}

	return str + "[white]"
}

func (widget *Widget) torrentSeedRatio(torrent transmissionrpc.Torrent) string {
	seedRatio := *torrent.UploadRatio

	if seedRatio < 0 {
		seedRatio = 0
	}

	return fmt.Sprintf("[green]%3d%%↑", int(seedRatio*100))
}

func (widget *Widget) torrentState(torrent transmissionrpc.Torrent) string {
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
