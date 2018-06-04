package bittrex

import (
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

type TextColors struct {
	base struct {
		name        string
		displayName string
	}
	market struct {
		name  string
		field string
		value string
	}
}

var ok = true
var errorText = ""
var started = false
var baseURL = "https://bittrex.com/api/v1.1/public/getmarketsummary"

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.TextWidget
	summaryList
	updateInterval int
	TextColors
}

// NewWidget Make new instance of widget
func NewWidget() *Widget {

	widget := Widget{
		TextWidget:     wtf.NewTextWidget(" $ Bittrex ", "bittrex", false),
		summaryList:    summaryList{},
		updateInterval: Config.UInt("wtf.mods.bittrex.updateInterval", 10),
	}

	started = false
	ok = true
	errorText = ""

	widget.config()
	widget.setSummaryList()

	return &widget
}

func (widget *Widget) config() {
	widget.TextColors.base.name = Config.UString("wtf.mods.bittrex.colors.base.name", "red")
	widget.TextColors.base.displayName = Config.UString("wtf.mods.bittrex.colors.base.displayName", "grey")
	widget.TextColors.market.name = Config.UString("wtf.mods.bittrex.colors.market.name", "red")
	widget.TextColors.market.field = Config.UString("wtf.mods.bittrex.colors.market.field", "coral")
	widget.TextColors.market.value = Config.UString("wtf.mods.bittrex.colors.market.value", "white")
}

func (widget *Widget) setSummaryList() {
	sCurrencies, _ := Config.Map("wtf.mods.bittrex.summary")
	for fromCurrencyName := range sCurrencies {
		displayName, _ := Config.String("wtf.mods.bittrex.summary." + fromCurrencyName + ".displayName")
		toCurrencyList := makeSummaryToList(fromCurrencyName)
		widget.summaryList.addSummaryItem(fromCurrencyName, displayName, toCurrencyList)
	}
}

func makeSummaryToList(currencyName string) []*tCurrency {
	tCurrencyList := []*tCurrency{}

	configToList, _ := Config.List("wtf.mods.bittrex.summary." + currencyName + ".market")
	for _, toCurrencyName := range configToList {
		tCurrencyList = append(tCurrencyList, makeToCurrency(toCurrencyName.(string)))
	}

	return tCurrencyList
}

func makeToCurrency(name string) *tCurrency {
	return &tCurrency{
		name: name,
		summaryInfo: summaryInfo{
			High:           "-1",
			Low:            "-1",
			Volume:         "-1",
			Last:           "-1",
			OpenBuyOrders:  "-1",
			OpenSellOrders: "-1",
		},
	}
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	if started == false {
		go func() {
			for {
				widget.updateSummary()
				time.Sleep(time.Second * time.Duration(widget.updateInterval))
			}
		}()
		started = true
	}

	widget.UpdateRefreshedAt()

	widget.display()

}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) updateSummary() {
	// In case if anything bad happened!
	defer func() {
		recover()
	}()

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	for _, fromCurrency := range widget.summaryList.items {
		for _, toCurrency := range fromCurrency.to {
			request := makeRequest(fromCurrency.name, toCurrency.name)
			response, err := client.Do(request)

			if err != nil {
				ok = false
				errorText = "Please Check Your Internet Connection!"
				break
			} else {
				ok = true
				errorText = ""
			}

			if response.StatusCode != http.StatusOK {
				errorText = response.Status
				ok = false
				break
			} else {
				ok = true
				errorText = ""
			}

			defer response.Body.Close()
			jsonResponse := summaryResponse{}
			decoder := json.NewDecoder(response.Body)
			decoder.Decode(&jsonResponse)

			if !jsonResponse.Success {
				ok = false
				errorText = fmt.Sprintf("%s-%s: %s", fromCurrency.name, toCurrency.name, jsonResponse.Message)
				break
			}
			ok = true
			errorText = ""

			toCurrency.Last = fmt.Sprintf("%f", jsonResponse.Result[0].Last)
			toCurrency.High = fmt.Sprintf("%f", jsonResponse.Result[0].High)
			toCurrency.Low = fmt.Sprintf("%f", jsonResponse.Result[0].Low)
			toCurrency.Volume = fmt.Sprintf("%f", jsonResponse.Result[0].Volume)
			toCurrency.OpenBuyOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenBuyOrders)
			toCurrency.OpenSellOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenSellOrders)
		}
	}

	widget.display()
}

func makeRequest(fName, tName string) *http.Request {
	url := fmt.Sprintf("%s?market=%s-%s", baseURL, fName, tName)
	request, _ := http.NewRequest("GET", url, nil)

	return request
}
