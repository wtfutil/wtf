package newrelic

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/senorprogrammer/wtf/wtf"
	nr "github.com/yfronto/newrelic"
)

var Config *config.Config

type Widget struct {
	wtf.BaseWidget

	View *tview.TextView
}

func NewWidget() *Widget {
	widget := Widget{
		BaseWidget: wtf.BaseWidget{
			Name:        "New Relic",
			RefreshedAt: time.Now(),
			RefreshInt:  Config.UInt("wtf.newrelic.refreshInterval", 900),
		},
	}

	widget.addView()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	app, _ := Application()
	deploys, _ := Deployments()

	widget.View.SetTitle(fmt.Sprintf(" New Relic: [green]%s[white] ", app.Name))
	widget.RefreshedAt = time.Now()

	widget.View.Clear()
	fmt.Fprintf(widget.View, "%s", widget.contentFrom(deploys))
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) addView() {
	view := tview.NewTextView()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorGray)
	view.SetDynamicColors(true)
	view.SetTitle(widget.Name)
	view.SetWrap(false)

	widget.View = view
}

func (widget *Widget) contentFrom(deploys []nr.ApplicationDeployment) string {
	str := "\n"
	str = str + " [red]Latest Deploys[white]\n"

	revisions := []string{}

	for _, deploy := range deploys {
		if (deploy.Revision != "") && wtf.Exclude(revisions, deploy.Revision) {
			str = str + fmt.Sprintf(
				" [green]%4s[white] %s %s\n",
				deploy.Revision[len(deploy.Revision)-4:],
				deploy.Timestamp.Format("Jan 2 15:04"),
				wtf.NameFromEmail(deploy.User),
			)

			revisions = append(revisions, deploy.Revision)

			if len(revisions) == Config.UInt("wtf.newrelic.deployCount", 5) {
				break
			}
		}
	}

	return str
}
