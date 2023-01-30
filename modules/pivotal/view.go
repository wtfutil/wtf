package pivotal

import (
	"fmt"
	"github.com/wtfutil/wtf/utils"
)

type PivotalSource struct {
	client    *PivotalClient
	name      string
	filter    string
	widget    *Widget
	Err       error
	stories   []Story
	max_items int
}

// NewPivotalSource returns a new Pivotal Filter source with a name
func NewPivotalSource(name string, filter string, client *PivotalClient, widget *Widget) *PivotalSource {
	source := PivotalSource{
		name:   name,
		filter: filter,
		client: client,
		widget: widget,
	}
	source.loadStories()
	return &source
}

func (source *PivotalSource) loadStories() {
	search, err := source.client.searchStories(source.filter)
	if err != nil {
		source.stories = nil
		source.Err = err
		source.setItemCount(0)
	} else {
		source.stories = search.Stories.Stories
		source.Err = err
		source.setItemCount(len(source.stories))
	}
}

// Open: Will open Pivotal search url with filter applied using the utils helper
func (source *PivotalSource) Open() {
	sel := source.widget.GetSelected()
	projectID := source.client.projectId
	if sel >= 0 && sel < source.getItemCount() {
		story := &source.stories[sel]
		baseURL := "https://www.pivotaltracker.com/n/projects/"
		ticketURL := fmt.Sprintf("%s%s/stories/%d", baseURL, projectID, story.ID)
		utils.OpenFile(ticketURL)
	}
}

// OpenPulls will open the GitHub Pull Requests URL using the utils helper
func (source *PivotalSource) OpenPulls() {
	sel := source.widget.GetSelected()
	if sel >= 0 && sel < source.getItemCount() {
		story := &source.stories[sel]
		if len(story.PullRequests) > 0 {
			pr := story.PullRequests[0]
			ticketURL := fmt.Sprintf("%s%s/%s/pull/%d", pr.HostURL, pr.Owner, pr.Repo, pr.Number)
			utils.OpenFile(ticketURL)
		}
	}
}

/* -------------------- Counts -------------------- */

func (source *PivotalSource) getItemCount() int {
	if source.stories == nil {
		return 0
	}
	return len(source.stories)
}
func (source *PivotalSource) setItemCount(count int) {
	source.max_items = count
}

/* -------------------- Unexported Functions -------------------- */
