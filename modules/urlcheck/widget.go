package urlcheck

import (
	"net/http"
	"text/template"
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	settings         *Settings          // settings from the configuration file
	urlList          []*urlResult       // list of a collection of useful properies of the url
	client           *http.Client       // the http client shared with all the requestes across all the refreshes
	timeout          time.Duration      // the timeout for a single request
	PreparedTemplate *template.Template // the test template shared across the refreshes
	templateString   string             // the string needed to parse the template and shared across all the widget refreshes
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	maxUrl := len(settings.urls)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
		urlList:  make([]*urlResult, maxUrl),
		client:   &http.Client{},
		timeout:  time.Duration(settings.requestTimeout) + time.Second,
	}

	widget.init()
	widget.View.SetWrap(false)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.check()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

// The string passed from the settings are checked and prepared for processing
func (widget *Widget) init() {

	// Prepare the template for the results
	widget.PrepareTemplate()

	for i, urlString := range widget.settings.urls {
		widget.urlList[i] = newUrlResult(urlString)
	}
}

// Do the actual requests and check the responses at every widget refresh
func (widget *Widget) check() {
	for _, urlRes := range widget.urlList {
		if urlRes.IsValid {
			urlRes.ResultCode, urlRes.ResultMessage = DoRequest(urlRes.Url, widget.timeout, widget.client)
		}
	}
}

// Format and displays the results at every refresh
func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.FormatResult(), false
	})
}
