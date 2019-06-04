package newrelic

import (
	"sort"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"github.com/wtfutil/wtf/wtf"
)

type Widget struct {
	wtf.KeyboardWidget
	wtf.MultiSourceWidget
	wtf.TextWidget

	Clients []*Client

	settings *Settings
}

func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:    wtf.NewKeyboardWidget(app, pages, settings.common),
		MultiSourceWidget: wtf.NewMultiSourceWidget(settings.common, "applicationID", "applicationIDs"),
		TextWidget:        wtf.NewTextWidget(app, settings.common, true),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	for _, id := range wtf.ToInts(widget.settings.applicationIDs) {
		widget.Clients = append(widget.Clients, NewClient(widget.settings.apiKey, id))
	}

	sort.Slice(widget.Clients, func(i, j int) bool {
		return widget.Clients[i].applicationId < widget.Clients[j].applicationId
	})

	widget.SetDisplayFunction(widget.display)

	widget.KeyboardWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() (string, string, bool) {
	app, appErr := widget.client.Application()
	deploys, depErr := widget.client.Deployments()

	appName := "error"
	if appErr == nil {
		appName = app.Name
	}

	var content string
	title := fmt.Sprintf("%s - [green]%s[white]", widget.CommonSettings().Title, appName)
	wrap := false
	if depErr != nil {
		wrap = true
		content = depErr.Error()
	} else {
		content += fmt.Sprintf(
			" %s\n",
			"[red]Latest Deploys[white]",
		)

		revisions := []string{}

		for _, deploy := range deploys {
			if (deploy.Revision != "") && utils.DoesNotInclude(revisions, deploy.Revision) {
				lineColor := "white"
				if wtf.IsToday(deploy.Timestamp) {
					lineColor = "lightblue"
				}

				revLen := 8
				if revLen > len(deploy.Revision) {
					revLen = len(deploy.Revision)
				}

				content += fmt.Sprintf(
					" [green]%s[%s] %s %-.16s[white]\n",
					deploy.Revision[0:revLen],
					lineColor,
					deploy.Timestamp.Format("Jan 02 15:04 MST"),
					utils.NameFromEmail(deploy.User),
				)

				revisions = append(revisions, deploy.Revision)

				if len(revisions) == widget.settings.deployCount {
					break
				}
			}
		}
	}

	return title, content, wrap
}

func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) currentData() *Client {
	if len(widget.Clients) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Clients) {
		return nil
	}

	return widget.Clients[widget.Idx]
}
