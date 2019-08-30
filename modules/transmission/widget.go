package transmission

import (
	"errors"

	"github.com/hekmon/transmissionrpc"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for transmission data
type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	client   *transmissionrpc.Client
	settings *Settings
	torrents []*transmissionrpc.Torrent
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common, true),

		settings: settings,
	}

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	// Create a persisten transmission client for use in the calls below
	client, err := transmissionrpc.New(widget.settings.host, widget.settings.username, widget.settings.password,
		&transmissionrpc.AdvancedConfig{
			Port: widget.settings.port,
		})
	if err != nil {
		client = nil
	}
	widget.client = client

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves torrent data from the Transmission daemon
func (widget *Widget) Fetch() ([]*transmissionrpc.Torrent, error) {
	if widget.client == nil {
		return nil, errors.New("client could not be initialized")
	}

	torrents, err := widget.client.TorrentGetAll()
	if err != nil {
		return nil, err
	}

	return torrents, nil
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	torrents, err := widget.Fetch()
	if err != nil {
		widget.err = err
		widget.torrents = nil
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.torrents = torrents
		widget.SetItemCount(len(torrents))
	}

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

func (widget *Widget) currentTorrent() *transmissionrpc.Torrent {
	if len(widget.torrents) == 0 {
		return nil
	}

	if len(widget.torrents) <= widget.Selected {
		return nil
	}

	return widget.torrents[widget.Selected]
}

// deleteSelected removes the selected torrent from transmission
// This action is non-destructive, it does not delete the files on the host
func (widget *Widget) deleteSelectedTorrent() {
	if widget.client == nil {
		return
	}

	currTorrent := widget.currentTorrent()
	if currTorrent == nil {
		return
	}

	ids := []int64{*currTorrent.ID}

	removePayload := &transmissionrpc.TorrentRemovePayload{
		IDs:             ids,
		DeleteLocalData: false,
	}

	widget.client.TorrentRemove(removePayload)

	widget.display()
}

// pauseUnpauseTorrent either pauses or unpauses the downloading and seeding of the selected torrent
func (widget *Widget) pauseUnpauseTorrent() {
	if widget.client == nil {
		return
	}

	currTorrent := widget.currentTorrent()
	if currTorrent == nil {
		return
	}

	ids := []int64{*currTorrent.ID}

	if *currTorrent.Status == transmissionrpc.TorrentStatusStopped {
		widget.client.TorrentStartIDs(ids)
	} else {
		widget.client.TorrentStopIDs(ids)
	}

	widget.display()
}
