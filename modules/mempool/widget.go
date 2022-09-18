package mempool

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/logger"
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

/* -------------------- Exported Functions -------------------- */

// Refresh updates the onscreen contents of the widget
func (widget *Widget) Refresh() {

	// The last call should always be to the display function
	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

type feeStruct struct {
	FastFee     int `json:"fastestFee"`
	HalfHourFee int `json:"halfHourFee"`
	HourFee     int `json:"hourFee"`
	EcoFee      int `json:"economyFee"`
	MinFee      int `json:"minimumFee"`
}

// {
// 	fastestFee: 14,
// 	halfHourFee: 1,
// 	hourFee: 1,
// 	economyFee: 1,
// 	minimumFee: 1
// }

func (widget *Widget) content() string {
	// var fees map[string]interface{}
	return callAPI()
}

func callAPI() string {
	url := "https://mempool.space/api/v1/fees/recommendeddd"
	resp, err := http.Get(url)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to make request to mempool. ERROR: %s", err))
		return "ERROR"
	}
	defer resp.Body.Close()

	// var parsed feeStruct
	// body, err := io.ReadAll(resp.Body)
	// newErr := json.Unmarshal(body, &parsed)

	parsed := feeStruct{}
	err = json.NewDecoder(resp.Body).Decode(&parsed)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to make decode JSON data from mempool. ERROR: %s", err))
		return "ERROR"
	}
	// decoder := json.NewDecoder(resp.Body)
	// err = decoder.Decode(&jsonResponse)
	// log.Fatal(parsed.EcoFee)
	return fmt.Sprint(parsed.EcoFee)
}

func (widget *Widget) display() {
	widget.Redraw(func() (string, string, bool) {
		return widget.CommonSettings().Title, widget.content(), false
	})
}
