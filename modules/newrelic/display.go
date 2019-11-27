package newrelic

import (
	"fmt"

	nr "github.com/wtfutil/wtf/modules/newrelic/client"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) content() (string, string, bool) {
	client := widget.currentData()
	if client == nil {
		return widget.CommonSettings().Title, " NewRelic data unavailable ", false
	}
	app, appErr := client.Application()
	deploys, depErr := client.Deployments()

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
		content = widget.contentFrom(deploys)
	}

	return title, content, wrap
}

func (widget *Widget) contentFrom(deploys []nr.ApplicationDeployment) string {
	str := fmt.Sprintf(
		" %s\n",
		fmt.Sprintf(
			"[%s]Latest Deploys[white]",
			widget.settings.common.Colors.Subheading,
		),
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

			str += fmt.Sprintf(
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

	return str
}
