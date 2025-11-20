package cdsfavorites

import (
	"fmt"
	"strconv"

	"github.com/ovh/cds/sdk"
	"github.com/ovh/cds/sdk/cdsclient"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.MultiSourceWidget
	view.TextWidget

	workflows []sdk.Workflow

	client cdsclient.Interface

	settings *Settings
	Selected int
	maxItems int
	Items    []int64
}

// NewWidget creates a new instance of the widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		MultiSourceWidget: view.NewMultiSourceWidget(settings.Common, "workflow", "workflows"),
		TextWidget:        view.NewTextWidget(tviewApp, redrawChan, pages, settings.Common),

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetRegions(true)
	widget.SetDisplayFunction(widget.display)

	widget.Unselect()

	widget.client = cdsclient.New(cdsclient.Config{
		Host:                               settings.apiURL,
		BuiltinConsumerAuthenticationToken: settings.token,
	})

	config, _ := widget.client.ConfigUser()

	if config.URLUI != "" {
		widget.settings.uiURL = config.URLUI
	}

	widget.workflows = widget.buildWorkflowsCollection()
	return &widget
}

/* -------------------- Exported Functions -------------------- */

// SetItemCount sets the amount of workflows throughout the widgets display creation
func (widget *Widget) SetItemCount(items int) {
	widget.maxItems = items
}

// GetItemCount returns the amount of workflows calculated so far as an int
func (widget *Widget) GetItemCount() int {
	return widget.maxItems
}

// GetSelected returns the index of the currently highlighted item as an int
func (widget *Widget) GetSelected() int {
	if widget.Selected < 0 {
		return 0
	}
	return widget.Selected
}

// Next cycles the currently highlighted text down
func (widget *Widget) Next() {
	widget.Selected++
	if widget.Selected >= widget.maxItems {
		widget.Selected = 0
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected))
	widget.View.ScrollToHighlight()
}

// Prev cycles the currently highlighted text up
func (widget *Widget) Prev() {
	widget.Selected--
	if widget.Selected < 0 {
		widget.Selected = widget.maxItems - 1
	}
	widget.View.Highlight(strconv.Itoa(widget.Selected))
	widget.View.ScrollToHighlight()
}

// Unselect stops highlighting the text and jumps the scroll position to the top
func (widget *Widget) Unselect() {
	widget.Selected = -1
	widget.View.Highlight()
	widget.View.ScrollToBeginning()
}

// Refresh reloads the data
func (widget *Widget) Refresh() {
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

/* buildWorkflowsCollection returns a slice of Workflow that have been
 * bookmarked by the user. Previously this method would returns workflows that
 * have been favorited but that concept was replaced with bookmarks in CDS.
 */
func (widget *Widget) buildWorkflowsCollection() []sdk.Workflow {

	/*
	   package main

	   import (
	       "context"
	       "fmt"
	       "log"
	       "net/http"

	       "github.com/ovh/cds/sdk"
	       "github.com/ovh/cds/sdk/cdsclient"
	   )

	   func getBookmarkedWorkflows(client cdsclient.Interface) ([]sdk.Workflow, error) {
	       // 1. Fetch the bookmarks directly using the REST endpoint
	       // The SDK doesn't always expose a helper for this, so we use .Request()
	       var bookmarks []sdk.Bookmark
	       _, err := client.Request(context.Background(), http.MethodGet, "/v2/user/me/bookmarks", nil, &bookmarks)
	       if err != nil {
	           return nil, fmt.Errorf("failed to fetch bookmarks: %v", err)
	       }

	       var bookmarkedWorkflows []sdk.Workflow

	       // 2. Iterate through bookmarks
	       for _, b := range bookmarks {
	           // Filter for Workflow bookmarks (Type is often "workflow" or defined by a constant)
	           if b.Type != sdk.BookmarkTypeWorkflow {
	               continue
	           }

	           // 3. Fetch the actual Workflow details
	           // Bookmarks only contain metadata (ProjectKey, Name), not the full struct
	           wf, err := client.WorkflowGet(b.ProjectKey, b.WorkflowName)
	           if err != nil {
	               log.Printf("Warning: Could not fetch workflow %s/%s: %v", b.ProjectKey, b.WorkflowName, err)
	               continue
	           }

	           bookmarkedWorkflows = append(bookmarkedWorkflows, *wf)
	       }

	       return bookmarkedWorkflows, nil
	   }
	*/

	workflows := []sdk.Workflow{}

	user, err := widget.client.UserGetMe()
	if err != nil {
		workflows
	}

	for _, bookmark := range user.Bookmarks {
		if bookmark.Type == "workflow" {
		}
	}

	data, _ := widget.client.Navbar()
	for _, v := range data {
		if v.Favorite && v.WorkflowName != "" {
			workflows = append(workflows, sdk.Workflow{ProjectKey: v.Key, Name: v.WorkflowName})
		}
	}
	return workflows
}

func (widget *Widget) currentCDSWorkflow() *sdk.Workflow {
	if len(widget.workflows) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.workflows) {
		widget.Idx = 0
	}

	p := widget.workflows[widget.Idx]
	return &p
}

func (widget *Widget) openWorkflow() {
	currentSelection := widget.View.GetHighlights()
	if widget.Selected >= 0 && currentSelection[0] != "" {
		wf := widget.currentCDSWorkflow()
		url := fmt.Sprintf("%s/project/%s/workflow/%s/run/%d",
			widget.settings.uiURL, wf.ProjectKey, wf.Name, widget.Items[widget.Selected])
		utils.OpenFile(url)
	}
}
