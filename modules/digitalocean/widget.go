package digitalocean

import (
	"context"
	"errors"

	"github.com/digitalocean/godo"
	"github.com/rivo/tview"
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

// Widget is the container for transmission data
type Widget struct {
	view.KeyboardWidget
	view.ScrollableWidget

	client   *godo.Client
	droplets []godo.Droplet
	settings *Settings
	err      error
}

// NewWidget creates a new instance of a widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		KeyboardWidget:   view.NewKeyboardWidget(app, pages, settings.common),
		ScrollableWidget: view.NewScrollableWidget(app, settings.common),

		settings: settings,
	}

	widget.SetRenderFunction(widget.display)
	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.KeyboardWidget.SetView(widget.View)

	widget.createClient()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves torrent data from the Transmission daemon
func (widget *Widget) Fetch() error {
	if widget.client == nil {
		return errors.New("client could not be initialized")
	}

	var err error
	widget.droplets, err = widget.fetchDroplets()
	if err != nil {
		return err
	}

	return nil
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

// HelpText returns the help text for this widget
func (widget *Widget) HelpText() string {
	return widget.KeyboardWidget.HelpText()
}

// Next selects the next item in the list
func (widget *Widget) Next() {
	widget.ScrollableWidget.Next()
}

// Prev selects the previous item in the list
func (widget *Widget) Prev() {
	widget.ScrollableWidget.Prev()
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

func (widget *Widget) fetchDroplets() ([]godo.Droplet, error) {
	dropletList := []godo.Droplet{}
	opts := &godo.ListOptions{}

	for {
		droplets, resp, err := widget.client.Droplets.List(context.Background(), opts)
		if err != nil {
			return dropletList, err
		}

		for _, d := range droplets {
			dropletList = append(dropletList, d)
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
