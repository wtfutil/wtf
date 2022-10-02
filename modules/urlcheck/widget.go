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

	settings         *Settings
	urlList          []*urlResult
	client           *http.Client
	timeout          time.Duration
	PreparedTemplate *template.Template
	templateString   string
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	maxUrl := len(settings.urls)

	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings: settings,
		urlList:  make([]*urlResult, maxUrl),
		client:   GetClient(),
		timeout:  time.Duration(settings.requestTimeout) + time.Second,
	}

	widget.PrepareTemplate()

	widget.View.SetWrap(false)
	widget.init()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	widget.check()
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) init() {
	for i, urlString := range widget.settings.urls {
		widget.urlList[i] = newUrlResult(urlString)
	}
}

func (widget *Widget) check() {
	for _, urlRes := range widget.urlList {
		if urlRes.IsValid {
			urlRes.ResultCode, urlRes.ResultMessage = DoRequest(urlRes.Url, widget.timeout, widget.client)
		}
	}
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.FormatResult(), false
	})
}
