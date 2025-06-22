package azuredevops

import (
	"fmt"
	"strings"

	azrBuild "github.com/microsoft/azure-devops-go-api/azuredevops/build"
	"github.com/pkg/errors"
)

func (widget *Widget) getBuildStats() string {
	projName := widget.settings.projectName
	statusFilter := azrBuild.BuildStatusValues.All
	top := widget.settings.maxRows
	builds, err := widget.cli.GetBuilds(widget.ctx, azrBuild.GetBuildsArgs{Project: &projName, StatusFilter: &statusFilter, Top: &top})
	if err != nil {
		return errors.Wrap(err, "could not get builds").Error()
	}

	result := ""
	for _, build := range builds.Value {
		num := *build.BuildNumber
		branch := *build.SourceBranch
		reason := *build.Reason
		triggers := *build.TriggerInfo
		if reason == azrBuild.BuildReasonValues.PullRequest {
			branch = triggers["pr.sourceBranch"]
		}
		branch = strings.TrimPrefix(branch, "refs/heads/")
		status := *build.Status
		statusDisplay := "[white:grey]unknown"

		switch status {
		case azrBuild.BuildStatusValues.InProgress:
			statusDisplay = "[white:blue]in progress"
		case azrBuild.BuildStatusValues.Cancelling:
			statusDisplay = "[white:orange]in cancelling"
		case azrBuild.BuildStatusValues.Postponed, azrBuild.BuildStatusValues.NotStarted:
			statusDisplay = "[white:blue]waiting"
		case azrBuild.BuildStatusValues.Completed:

			buildResult := *build.Result

			switch buildResult {
			case azrBuild.BuildResultValues.Succeeded:
				statusDisplay = "[white:green]succeeded"
			case azrBuild.BuildResultValues.Failed:
				statusDisplay = "[white:red]failed"
			case azrBuild.BuildResultValues.Canceled:
				statusDisplay = "[white:darkgrey]cancelled"
			case azrBuild.BuildResultValues.PartiallySucceeded:
				statusDisplay = "[white:magenta]partially"
			}
		}

		result += fmt.Sprintf("%s[-:-:-] #%s %s (%s) \n", statusDisplay, num, branch, reason)
	}

	if result == "" {
		result = "no builds found"
	}

	return result
}
