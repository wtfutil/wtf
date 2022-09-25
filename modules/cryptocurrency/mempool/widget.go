package mempool

import (
	"fmt"
	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

// Widget is the container for your module's data
type Widget struct {
	view.TextWidget
	settings *Settings
}

// NewWidget creates and returns an instance of Widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, pages, settings.common),

		settings: settings,
	}
	return &widget
}

type feeStruct struct {
	FastFee     int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
	EcoFee      int `json:"economyFee"`
}

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {
	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) content() string {
	return getBTCTxFees()
}

func getBTCTxFees() string {
	url := "https://mempool.space/api/v1/fees/recommended"
	resp, err := http.Get(url)
	if err != nil {
		logger.Log(fmt.Sprintf("[mempool] Error: Failed to make request to mempool. Reason: %s", err))
		return "[mempool] error callng mempool API"
	}
	defer resp.Body.Close()

	parsed := feeStruct{}
	err = utils.ParseJSON(&parsed, resp.Body)
	if err != nil {
		logger.Log(fmt.Sprintf("[mempool] Error: Failed to decode JSON data from mempool. Reason: %s", err))
		return "[mempool] error parsing JSON from mempool API"
	}

	finalStr := ""
	finalStr += fmt.Sprintf("%-7s %2d sat/vB\n", "Fast", parsed.FastFee)
	finalStr += fmt.Sprintf("%-7s %2d sat/vB\n", "30 min", parsed.HalfHourFee)
	finalStr += fmt.Sprintf("%-7s %2d sat/vB\n", "60 min", parsed.HourFee)
	finalStr += fmt.Sprintf("%-7s %2d sat/vB\n", "Eco", parsed.EcoFee)

	return finalStr
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
