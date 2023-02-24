package transmission

import (
	"context"
	"errors"
	"sync"

	"github.com/hekmon/transmissionrpc/v2"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for transmission data
type Widget struct {
	view.ScrollableWidget

	client   *transmissionrpc.Client
	settings *Settings
	mu       sync.Mutex
	torrents []transmissionrpc.Torrent
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()

	go buildClient(&widget)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves torrent data from the Transmission daemon
func (widget *Widget) Fetch() ([]transmissionrpc.Torrent, error) {
	if widget.client == nil {
		return nil, errors.New("client was not initialized")
	}

	torrents, err := widget.client.TorrentGetAll(context.Background())
	if err != nil {
		return nil, err
	}

	out := make([]transmissionrpc.Torrent, 0)
	for _, torrent := range torrents {
		if widget.settings.hideComplete {
			if *torrent.PercentDone == 1.0 {
				continue
			}
		}

		out = append(out, torrent)
	}

	return out, nil
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	torrents, err := widget.Fetch()
	count := 0

	if err == nil {
		count = len(torrents)
	}

	widget.mu.Lock()
	widget.err = err
	widget.torrents = torrents
	widget.SetItemCount(count)
	widget.mu.Unlock()

	widget.display()
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

// buildClient creates a persisten transmission client
func buildClient(widget *Widget) {
	widget.mu.Lock()
	defer widget.mu.Unlock()

	client, err := transmissionrpc.New(widget.settings.host, widget.settings.username, widget.settings.password,
		&transmissionrpc.AdvancedConfig{
			Port:   widget.settings.port,
			RPCURI: widget.settings.url,
			HTTPS:  widget.settings.https,
		})
	if err != nil {
		client = nil
	}

	widget.client = client
}

func (widget *Widget) currentTorrent() *transmissionrpc.Torrent {
	if len(widget.torrents) == 0 {
		return nil
	}

	if len(widget.torrents) <= widget.Selected {
		return nil
	}

	return &widget.torrents[widget.Selected]
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

	removePayload := transmissionrpc.TorrentRemovePayload{
		IDs:             ids,
		DeleteLocalData: false,
	}

	err := widget.client.TorrentRemove(context.Background(), removePayload)
	if err != nil {
		return
	}

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

	var err error
	if *currTorrent.Status == transmissionrpc.TorrentStatusStopped {
		err = widget.client.TorrentStartIDs(context.Background(), ids)
	} else {
		err = widget.client.TorrentStopIDs(context.Background(), ids)
	}

	if err != nil {
		return
	}

	widget.display()
}
