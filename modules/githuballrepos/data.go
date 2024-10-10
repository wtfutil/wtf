package githuballrepos

import (
	"fmt"
	"strings"
)

// WidgetData holds all the data for the widget
type WidgetData struct {
	PRsOpenedByMe         int
	PRReviewRequestsCount int
	OpenIssuesCount       int

	MyPRs            []PR
	PRReviewRequests []PR
	WatchedPRs       []PR
}

// PR represents a single pull request
type PR struct {
	Title      string
	URL        string
	Author     string
	Repository string
}

// FormatCounters returns a formatted string of counters
func (d *WidgetData) FormatCounters() string {
	return fmt.Sprintf(
		"PRs opened by me: %d\nPR review requests: %d\nOpen issues: %d\n",
		d.PRsOpenedByMe,
		d.PRReviewRequestsCount,
		d.OpenIssuesCount,
	)
}

// FormatPRs returns a formatted string of PRs
func (d *WidgetData) FormatPRs() string {
	var sb strings.Builder

	sb.WriteString("[green]My PRs:[white]\n")
	for _, pr := range d.MyPRs {
		sb.WriteString(fmt.Sprintf("- %s (%s)\n", pr.Title, pr.Repository))
	}

	sb.WriteString("\n[yellow]PR Review Requests:[white]\n")
	for _, pr := range d.PRReviewRequests {
		sb.WriteString(fmt.Sprintf("- %s (%s)\n", pr.Title, pr.Repository))
	}

	sb.WriteString("\n[blue]Watched PRs:[white]\n")
	for _, pr := range d.WatchedPRs {
		sb.WriteString(fmt.Sprintf("- %s (%s by %s)\n", pr.Title, pr.Repository, pr.Author))
	}

	return sb.String()
}
