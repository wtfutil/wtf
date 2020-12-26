package digitalocean

import (
	"context"
	"errors"
	"fmt"

	"github.com/digitalocean/godo"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
	"golang.org/x/oauth2"
)

/* -------------------- Oauth2 Token -------------------- */

type tokenSource struct {
	AccessToken string
}

// Token creates and returns an Oauth2 token
func (t *tokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

/* -------------------- Widget -------------------- */

// Widget is the container for droplet data
type Widget struct {
	view.ScrollableWidget

	app      *tview.Application
	client   *godo.Client
	droplets []*Droplet
	pages    *tview.Pages
	settings *Settings

	err error
}

// NewWidget creates a new instance of a widget
func NewWidget(tviewApp *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		ScrollableWidget: view.NewScrollableWidget(tviewApp, pages, settings.Common),

		app:      tviewApp,
		pages:    pages,
		settings: settings,
	}

	widget.initializeKeyboardControls()

	widget.View.SetScrollable(true)

	widget.SetRenderFunction(widget.display)

	widget.createClient()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves droplet data
func (widget *Widget) Fetch() error {
	if widget.client == nil {
		return errors.New("client could not be initialized")
	}

	var err error
	widget.droplets, err = widget.dropletsFetch()
	return err
}

// Next selects the next item in the list
func (widget *Widget) Next() {
	widget.ScrollableWidget.Next()
}

// Prev selects the previous item in the list
func (widget *Widget) Prev() {
	widget.ScrollableWidget.Prev()
}

// Refresh updates the data for this widget and displays it onscreen
func (widget *Widget) Refresh() {
	err := widget.Fetch()
	if err != nil {
		widget.err = err
		widget.SetItemCount(0)
	} else {
		widget.err = nil
		widget.SetItemCount(len(widget.droplets))
	}

	widget.display()
}

// Unselect clears the selection of list items
func (widget *Widget) Unselect() {
	widget.ScrollableWidget.Unselect()
	widget.RenderFunction()
}

/* -------------------- Unexported Functions -------------------- */

// createClient create a persisten DigitalOcean client for use in the calls below
func (widget *Widget) createClient() {
	tokenSource := &tokenSource{
		AccessToken: widget.settings.apiKey,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	widget.client = godo.NewClient(oauthClient)
}

// currentDroplet returns the currently-selected droplet, if there is one
// Returns nil if no droplet is selected
func (widget *Widget) currentDroplet() *Droplet {
	if len(widget.droplets) == 0 {
		return nil
	}

	if len(widget.droplets) <= widget.Selected {
		return nil
	}

	return widget.droplets[widget.Selected]
}

// dropletsFetch uses the DigitalOcean API to fetch information about all the available droplets
func (widget *Widget) dropletsFetch() ([]*Droplet, error) {
	dropletList := []*Droplet{}
	opts := &godo.ListOptions{}

	for {
		doDroplets, resp, err := widget.client.Droplets.List(context.Background(), opts)
		if err != nil {
			return dropletList, err
		}

		for _, doDroplet := range doDroplets {
			droplet := NewDroplet(doDroplet)
			dropletList = append(dropletList, droplet)
		}

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return dropletList, err
		}

		// Set the page we want for the next request
		opts.Page = page + 1
	}

	return dropletList, nil
}

/* -------------------- Droplet Actions -------------------- */

// dropletDestroy destroys the selected droplet
func (widget *Widget) dropletDestroy() {
	currDroplet := widget.currentDroplet()
	if currDroplet == nil {
		return
	}

	_, err := widget.client.Droplets.Delete(context.Background(), currDroplet.ID)
	if err != nil {
		return
	}

	widget.dropletRemoveSelected()
	widget.Refresh()
}

// dropletEnabledPrivateNetworking enabled private networking on the selected droplet
func (widget *Widget) dropletEnabledPrivateNetworking() {
	currDroplet := widget.currentDroplet()
	if currDroplet == nil {
		return
	}

	_, _, err := widget.client.DropletActions.EnablePrivateNetworking(context.Background(), currDroplet.ID)
	if err != nil {
		return
	}

	widget.Refresh()
}

// dropletRemoveSelected removes the currently-selected droplet from the internal list of droplets
func (widget *Widget) dropletRemoveSelected() {
	currDroplet := widget.currentDroplet()
	if currDroplet != nil {
		widget.droplets[len(widget.droplets)-1], widget.droplets[widget.Selected] = widget.droplets[widget.Selected], widget.droplets[len(widget.droplets)-1]
		widget.droplets = widget.droplets[:len(widget.droplets)-1]
	}
}

// dropletRestart restarts the selected droplet
func (widget *Widget) dropletRestart() {
	currDroplet := widget.currentDroplet()
	if currDroplet == nil {
		return
	}

	_, _, err := widget.client.DropletActions.Reboot(context.Background(), currDroplet.ID)
	if err != nil {
		return
	}
	widget.Refresh()
}

// dropletShutDown powers down the selected droplet
func (widget *Widget) dropletShutDown() {
	currDroplet := widget.currentDroplet()
	if currDroplet == nil {
		return
	}

	_, _, err := widget.client.DropletActions.Shutdown(context.Background(), currDroplet.ID)
	if err != nil {
		return
	}
	widget.Refresh()
}

/* -------------------- Common Actions -------------------- */

// showInfo shows a modal window with information about the selected droplet
func (widget *Widget) showInfo() {
	droplet := widget.currentDroplet()
	if droplet == nil {
		return
	}

	closeFunc := func() {
		widget.pages.RemovePage("info")
		widget.app.SetFocus(widget.View)
	}

	propTable := newDropletPropertiesTable(droplet).render()
	propTable += utils.CenterText("Esc to close", 80)

	modal := view.NewBillboardModal(propTable, closeFunc)
	modal.SetTitle(fmt.Sprintf("  %s  ", droplet.Name))

	widget.pages.AddPage("info", modal, false, true)
	widget.app.SetFocus(modal)

	widget.app.QueueUpdateDraw(func() {
		widget.app.Draw()
	})
}
