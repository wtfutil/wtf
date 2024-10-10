package githuballrepos

import (
	"time"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/view"
)

// Widget is the main struct for the GitHubAllRepos module
type Widget struct {
	view.TextWidget

	settings *Settings
	data     *WidgetData
	cache    Cacher
	client   GitHubFetcher

	currentTab      int
	selectedPRIndex int
	isItemSelected  bool
	lastUpdateTime  time.Time
	isSelected      bool
	carouselChan    chan bool

	testMode bool
}

// NewWidget creates and returns an instance of Widget
func NewWidget(app *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := &Widget{
		TextWidget:      view.NewTextWidget(app, redrawChan, pages, settings.common),
		settings:        settings,
		data:            &WidgetData{},
		cache:           NewCache(settings.CachePath),
		client:          NewGitHubClient(settings.APIKey),
		currentTab:      0,
		selectedPRIndex: -1,
		isItemSelected:  false,
		isSelected:      false,
		carouselChan:    make(chan bool),
	}

	widget.View.SetWrap(false)
	widget.View.SetScrollable(true)

	widget.View.SetTitle(settings.common.Title)

	widget.initializeKeyboardControls()
	widget.startCarousel()

	logger.Log("GitHubAllRepos widget created")

	return widget
}

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {

	logger.Log("Refresh started")
	if widget.cache.IsValid() {

		logger.Log("Cache is valid")
		widget.data = widget.cache.Get()
	} else {
		logger.Log("Fetching new data")
		widget.data = widget.client.FetchData(widget.settings.Organizations, widget.settings.Username)
		widget.cache.Set(widget.data)
	}
	widget.lastUpdateTime = time.Now()
	logger.Log("Refresh completed")

	if !widget.testMode {
		widget.display()
	}
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}

func (widget *Widget) content() string {
	var str string

	// Display counters
	str += "[yellow]Counters:[white]\n"
	str += widget.data.FormatCounters()

	// Display tabs
	str += "\n[blue]Tabs:[white] "
	tabs := []string{"My PRs", "PR Review Requests", "Watched PRs"}
	for i, tab := range tabs {
		if i == widget.currentTab {
			str += "[green]" + tab + "[white] "
		} else {
			str += tab + " "
		}
	}
	str += "\n\n"

	// Display PRs for the current tab
	var prs []PR
	switch widget.currentTab {
	case 0:
		prs = widget.data.MyPRs
	case 1:
		prs = widget.data.PRReviewRequests
	case 2:
		prs = widget.data.WatchedPRs
	}

	for i, pr := range prs {
		if i == widget.selectedPRIndex {
			str += "[green]> " + pr.Title + " (" + pr.Repository + ")[white]\n"
		} else {
			str += "  " + pr.Title + " (" + pr.Repository + ")\n"
		}
	}

	return str
}

func (widget *Widget) startCarousel() {
	// skipcq: GOOO-E1007
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if !widget.isSelected && !widget.isItemSelected {
					widget.nextTabWithoutRefresh()
					widget.Refresh()
				}
			case isSelected := <-widget.carouselChan:
				if isSelected {
					ticker.Stop()
				} else {
					ticker = time.NewTicker(10 * time.Second)
				}
			}
		}
	}()
}

func (widget *Widget) nextTabWithoutRefresh() {
	widget.currentTab = (widget.currentTab + 1) % 3
	widget.selectedPRIndex = -1
}

// SetSelected sets the selection state of the widget
func (widget *Widget) SetSelected(selected bool) {
	widget.isSelected = selected
	if !selected {
		widget.isItemSelected = false
		widget.selectedPRIndex = -1
	}
	widget.carouselChan <- selected
	widget.Refresh()
}

// Implements the Focusable interface
func (widget *Widget) Focused() {
	widget.SetSelected(true)
	widget.Refresh()
}

// Implements the Focusable interface
func (widget *Widget) Unfocused() {
	widget.SetSelected(false)
	widget.selectedPRIndex = -1
	widget.Refresh()
}
