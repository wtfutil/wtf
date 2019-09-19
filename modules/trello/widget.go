package trello

import (
	"fmt"

	"github.com/adlio/trello"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings *Settings
}

func NewWidget(app *tview.Application, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(app, settings.common),

		settings: settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {

	client := trello.NewClient(
		widget.settings.apiKey,
		widget.settings.accessToken,
	)

	// Get the cards
	searchResult, err := GetCards(
		client,
		widget.settings.username,
		widget.settings.board,
		widget.settings.list,
	)

	var title string
	content := ""

	wrap := false
	if err != nil {
		wrap = true
		title = widget.CommonSettings().Title
		content = err.Error()
	} else {
		title = fmt.Sprintf(
			"[white]%s: [green]%s ",
			widget.CommonSettings().Title,
			widget.settings.board,
		)
		for list, cardArray := range searchResult.TrelloCards {
			content += fmt.Sprintf(" [red]%s[white]\n", list)

			for _, card := range cardArray {
				content += fmt.Sprintf(" %s[white]\n", card.Name)
			}
			content = fmt.Sprintf("%s\n", content)
		}
	}

	return title, content, wrap
}
