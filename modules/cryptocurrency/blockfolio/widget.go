package blockfolio

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

type Widget struct {
	view.TextWidget

	device_token string
	settings     *Settings
}

func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		device_token: settings.deviceToken,
		settings:     settings,
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {

	widget.Redraw(widget.content)
}

/* -------------------- Unexported Functions -------------------- */
func (widget *Widget) content() (string, string, bool) {
	positions, err := Fetch(widget.device_token)
	title := widget.CommonSettings().Title
	if err != nil {
		return title, err.Error(), true
	}

	res := ""
	totalFiat := float32(0.0)

	for i := 0; i < len(positions.PositionList); i++ {
		colorForChange := widget.settings.colors.grows

		if positions.PositionList[i].TwentyFourHourPercentChangeFiat <= 0 {
			colorForChange = widget.settings.colors.drop
		}

		totalFiat += positions.PositionList[i].HoldingValueFiat

		if widget.settings.displayHoldings {
			res += fmt.Sprintf(
				"[%s]%-6s - %5.2f ([%s]%.3fk [%s]%.2f%s)\n",
				widget.settings.colors.name,
				positions.PositionList[i].Coin,
				positions.PositionList[i].Quantity,
				colorForChange,
				positions.PositionList[i].HoldingValueFiat/1000,
				colorForChange,
				positions.PositionList[i].TwentyFourHourPercentChangeFiat,
				"%",
			)
		} else {
			res += fmt.Sprintf(
				"[%s]%-6s - %5.2f ([%s]%.2f%s)\n",
				widget.settings.colors.name,
				positions.PositionList[i].Coin,
				positions.PositionList[i].Quantity,
				colorForChange,
				positions.PositionList[i].TwentyFourHourPercentChangeFiat,
				"%",
			)
		}
	}

	if widget.settings.displayHoldings {
		res += fmt.Sprintf("\n[%s]Total value: $%.3fk", "green", totalFiat/1000)
	}

	return title, res, true
}

// always the same
const magic = "edtopjhgn2345piuty89whqejfiobh89-2q453"

type Position struct {
	Coin                            string  `json:"coin"`
	LastPriceFiat                   float32 `json:"lastPriceFiat"`
	TwentyFourHourPercentChangeFiat float32 `json:"twentyFourHourPercentChangeFiat"`
	Quantity                        float32 `json:"quantity"`
	HoldingValueFiat                float32 `json:"holdingValueFiat"`
}

type AllPositionsResponse struct {
	PositionList []Position `json:"positionList"`
}

func MakeApiRequest(token string, method string) ([]byte, error) {
	client := &http.Client{}
	url := "https://api-v0.blockfolio.com/rest/" + method + "/" + token + "?use_alias=true&fiat_currency=USD"
	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("magic", magic)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func GetAllPositions(token string) (*AllPositionsResponse, error) {
	jsn, _ := MakeApiRequest(token, "get_all_positions")
	var parsed AllPositionsResponse

	err := json.Unmarshal(jsn, &parsed)
	if err != nil {
		log.Fatalf("Failed to parse json %v", err)
		return nil, err
	}
	return &parsed, err
}

func Fetch(token string) (*AllPositionsResponse, error) {
	return GetAllPositions(token)
}
