package steam

import (
	"context"
	"fmt"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"golang.org/x/sync/errgroup"
)

type Widget struct {
	view.ScrollableWidget

	settings *Settings
	err      error
	steam    *Steam
	players  []*Player
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),
		settings:         settings,
		steam:            NewClient(&ClientOpts{key: settings.key}),
	}

	widget.SetRenderFunction(widget.Render)
	widget.initializeKeyboardControls()

	return widget
}

func (widget *Widget) Refresh() {
	errg, _ := errgroup.WithContext(context.Background())
	players := make([]*Player, len(widget.settings.userIds))

	for i, id := range widget.settings.userIds {
		func(idx int, id string) {
			errg.Go(func() error {
				status, err := widget.steam.Status(id)
				if err != nil {
					return err
				}
				players[idx] = status
				return nil
			})
		}(i, id)
	}

	if err := errg.Wait(); err != nil {
		widget.err = err
		widget.players = nil
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		if len(players) <= widget.settings.numberOfResults {
			widget.players = players
		} else {
			widget.players = players[:widget.settings.numberOfResults]
		}
		widget.SetItemCount(len(widget.players))
	}

	widget.Render()
}

func (widget *Widget) Render() {
	widget.Redraw(widget.content)
}

func friendlyStatus(personastate int) string {
	switch personastate {
	case 0:
		return "Offline"
	case 1:
		return "Online"
	case 2:
		return "Busy"
	case 3:
		return "Away"
	case 4:
		return "Snooze"
	case 5:
		return "Looking to Trade"
	case 6:
		return "Looking to Play"
	}
	return ""
}

func (widget *Widget) content() (string, string, bool) {
	var title = "Steam Statuses"

	if widget.CommonSettings().Title != "" {
		title = widget.CommonSettings().Title
	}

	if widget.err != nil {
		return title, widget.err.Error(), true
	}

	if len(widget.players) == 0 {
		return title, "No data", false
	}

	var str string

	for idx, player := range widget.players {
		status := friendlyStatus(player.Personastate)

		row := fmt.Sprintf(
			"[white]%s: [yellow]%s",
			player.Personaname,
			status,
		)

		if len(player.Gameextrainfo) > 0 {
			row += " [red](" + player.Gameextrainfo + ")"
		}

		str += utils.HighlightableHelper(widget.View, row, idx, len(player.Personaname))
	}

	return title, str, false
}
