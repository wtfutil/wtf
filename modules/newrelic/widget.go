package newrelic

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
	nr "github.com/yfronto/newrelic"
)

type Widget struct {
	wtf.TextWidget
	client *Client
}

func NewWidget(app *tview.Application) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(app, "New Relic", "newrelic", false),
		client:     NewClient(apiKey(), wtf.Config.UInt("wtf.mods.newrelic.applicationId")),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	app, appErr := widget.client.Application()
	deploys, depErr := widget.client.Deployments()

	appName := "error"
	if appErr == nil {
		appName = app.Name
	}

	widget.View.SetTitle(widget.ContextualTitle(fmt.Sprintf("%s - [green]%s[white]", widget.Name, appName)))
	widget.View.Clear()

	var content string
	if depErr != nil {
		widget.View.SetWrap(true)
		content = depErr.Error()
	} else {
		widget.View.SetWrap(false)
		content = widget.contentFrom(deploys)
	}

	widget.View.SetText(content)
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) contentFrom(deploys []nr.ApplicationDeployment) string {
	str := fmt.Sprintf(
		" %s\n",
		"[red]Latest Deploys[white]",
	)

	revisions := []string{}

	for _, deploy := range deploys {
		if (deploy.Revision != "") && wtf.Exclude(revisions, deploy.Revision) {
			lineColor := "white"
			if wtf.IsToday(deploy.Timestamp) {
				lineColor = "lightblue"
			}

			revLen := 8
			if revLen > len(deploy.Revision) {
				revLen = len(deploy.Revision)
			}

			str = str + fmt.Sprintf(
				" [green]%s[%s] %s %-.16s[white]\n",
				deploy.Revision[0:revLen],
				lineColor,
				deploy.Timestamp.Format("Jan 02 15:04 MST"),
				wtf.NameFromEmail(deploy.User),
			)

			revisions = append(revisions, deploy.Revision)

			if len(revisions) == wtf.Config.UInt("wtf.mods.newrelic.deployCount", 5) {
				break
			}
		}
	}

	return str
}

func apiKey() string {
	return wtf.Config.UString(
		"wtf.mods.newrelic.apiKey",
		os.Getenv("WTF_NEW_RELIC_API_KEY"),
	)
}
