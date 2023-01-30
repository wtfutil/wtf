package pivotal

import (
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

// A Widget represents a Todoist widget
type Widget struct {
	view.MultiSourceWidget
	view.ScrollableWidget
	settings *Settings

	client        *PivotalClient
	projectClient map[string]*PivotalClient
	sources       []*PivotalSource
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {

	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "customQuery", "customQueries"),
		ScrollableWidget:  view.NewScrollableWidget(tviewApp, redrawChan, pages, settings.Common),

		settings:      settings,
		client:        NewPivotalClient(settings.apiToken, settings.projectId),
		projectClient: make(map[string]*PivotalClient),
	}

	widget.loadSources()

	// Add the client to projectClient list
	widget.projectClient[widget.settings.projectId] = widget.client

	//Build the Souce lists
	widget.sources = widget.buildPivotalSources()

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.SetDisplayFunction(widget.display)

	return &widget
}

func (widget *Widget) loadSources() {
	var queries []string
	for _, query := range widget.settings.customQueries {
		queries = append(queries, query.title)
	}
	widget.Sources = queries
}

func (widget *Widget) buildPivotalSources() []*PivotalSource {
	var sources []*PivotalSource

	for _, query := range widget.settings.customQueries {
		client := widget.client
		// Make sure that we have a viable Pivotal Client
		if query.project != "" && query.project != widget.client.projectId {
			nclient, ok := widget.projectClient[query.project]
			if !ok {
				nclient = NewPivotalClient(widget.settings.apiToken, query.project)
			}
			client = nclient
		}

		sources = append(sources,
			NewPivotalSource(query.title, query.filter, client, widget))
	}

	return sources
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) CurrentSource() *PivotalSource {
	if len(widget.sources) == 0 {
		return nil
	}

	return widget.sources[widget.Idx]

}

func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}
	widget.SetItemCount(widget.CurrentSource().getItemCount())
	widget.display()
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Open() {
	widget.CurrentSource().Open()
}
func (widget *Widget) OpenPulls() {
	widget.CurrentSource().OpenPulls()
}
