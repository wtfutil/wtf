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
		if status == azrBuild.BuildStatusValues.InProgress {
			statusDisplay = "[white:blue]in progress"
		} else if status == azrBuild.BuildStatusValues.Cancelling {
			statusDisplay = "[white:orange]in cancelling"
		} else if (status == azrBuild.BuildStatusValues.Postponed) || (status == azrBuild.BuildStatusValues.NotStarted) {
			statusDisplay = "[white:blue]waiting"
		} else if status == azrBuild.BuildStatusValues.Completed {
			buildResult := *build.Result
			if buildResult == azrBuild.BuildResultValues.Succeeded {
				statusDisplay = "[white:green]succeeded"
			} else if buildResult == azrBuild.BuildResultValues.Failed {
				statusDisplay = "[white:red]failed"
			} else if buildResult == azrBuild.BuildResultValues.Canceled {
				statusDisplay = "[white:darkgrey]cancelled"
			} else if buildResult == azrBuild.BuildResultValues.PartiallySucceeded {
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
