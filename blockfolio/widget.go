package blockfolio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andrewzolotukhin/wtf/wtf"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

// Config is a pointer to the global config object
var Config *config.Config

type Widget struct {
	wtf.TextWidget

	app          *tview.Application
	device_token string
}

func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget:   wtf.NewTextWidget(" Blockfolio ", "blockfolio", false),
		device_token: Config.UString("wtf.mods.blockfolio.device_token"),
	}

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	widget.UpdateRefreshedAt()
	widget.View.SetTitle(" Blockfolio ")

	positions, err := Fetch(widget.device_token)
	if err != nil {
		return
	}
	widget.View.SetText(fmt.Sprintf("%s", contentFrom(positions)))
}

/* -------------------- Unexported Functions -------------------- */
func contentFrom(positions *AllPositionsResponse) string {
	res := ""
	colorName := Config.UString("wtf.mods.blockfolio.colors.name")
	colorGrows := Config.UString("wtf.mods.blockfolio.colors.grows")
	colorDrop := Config.UString("wtf.mods.blockfolio.colors.drop")
	var totalFiat float32
	totalFiat = 0.0
	for i := 0; i < len(positions.PositionList); i++ {
		colorForChange := colorGrows
		if positions.PositionList[i].TwentyFourHourPercentChangeFiat <= 0 {
			colorForChange = colorDrop
		}
		totalFiat += positions.PositionList[i].HoldingValueFiat
		res = res + fmt.Sprintf("[%s]%-6s - %5.2f ([%s]%.2fk [%s]%.2f%s)\n", colorName, positions.PositionList[i].Coin, positions.PositionList[i].Quantity, colorForChange, positions.PositionList[i].HoldingValueFiat/1000, colorForChange, positions.PositionList[i].TwentyFourHourPercentChangeFiat, "%")
	}
	res = res + fmt.Sprintf("\n[%s]Total value: $%.2fk", "green", totalFiat/1000)

	return res
}

//always the same
const magic = "edtopjhgn2345piuty89whqejfiobh89-2q453"

type Position struct {
	Coin                            string  `json:coin`
	LastPriceFiat                   float32 `json:lastPriceFiat`
	TwentyFourHourPercentChangeFiat float32 `json:twentyFourHourPercentChangeFiat`
	Quantity                        float32 `json:quantity`
	HoldingValueFiat                float32 `json:holdingValueFiat`
}

type AllPositionsResponse struct {
	PositionList []Position `json:positionList`
}

func MakeApiRequest(token string, method string) ([]byte, error) {
	client := &http.Client{}
	url := "https://api-v0.blockfolio.com/rest/" + method + "/" + token + "?use_alias=true&fiat_currency=USD"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("magic", magic)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func GetAllPositions(token string) (*AllPositionsResponse, error) {
	jsn, err := MakeApiRequest(token, "get_all_positions")
	var parsed AllPositionsResponse

	err = json.Unmarshal(jsn, &parsed)
	if err != nil {
		log.Fatalf("Failed to parse json %v", err)
		return nil, err
	}
	return &parsed, err
}

func Fetch(token string) (*AllPositionsResponse, error) {
	return GetAllPositions(token)
}
