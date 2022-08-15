package bittrex

import (
	"encoding/json"
	"fmt"
	"time"

	"net/http"

	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/view"
)

const (
	baseURL = "https://bittrex.com/api/v1.1/public/getmarketsummary"
)

var (
	errorText = ""
	ok        = true
)

// Widget define wtf widget to register widget later
type Widget struct {
	view.TextWidget

	settings *Settings
	summaryList
}

// NewWidget Make new instance of widget
func NewWidget(tviewApp *tview.Application, redrawChan chan bool, settings *Settings) *Widget {
	widget := Widget{
		TextWidget: view.NewTextWidget(tviewApp, redrawChan, nil, settings.Common),

		settings:    settings,
		summaryList: summaryList{},
	}

	ok = true
	errorText = ""

	widget.setSummaryList()

	return &widget
}

func (widget *Widget) setSummaryList() {
	for symbol, currency := range widget.settings.summary.currencies {
		mCurrencyList := widget.makeSummaryMarketList(currency.market)
		widget.summaryList.addSummaryItem(symbol, currency.displayName, mCurrencyList)
	}
}

func (widget *Widget) makeSummaryMarketList(market []interface{}) []*mCurrency {
	mCurrencyList := []*mCurrency{}

	for _, marketSymbol := range market {
		mCurrencyList = append(mCurrencyList, makeMarketCurrency(marketSymbol.(string)))
	}

	return mCurrencyList
}

func makeMarketCurrency(name string) *mCurrency {
	return &mCurrency{
		name: name,
		summaryInfo: summaryInfo{
			High:           "",
			Low:            "",
			Volume:         "",
			Last:           "",
			OpenBuyOrders:  "",
			OpenSellOrders: "",
		},
	}
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	widget.updateSummary()

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) updateSummary() {
	// In case if anything bad happened!
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in updateSummary()", r)
		}
	}()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, baseCurrency := range widget.summaryList.items {
		for _, mCurrency := range baseCurrency.markets {
			request := makeRequest(baseCurrency.name, mCurrency.name)
			response, err := client.Do(request)

			ok = true
			errorText = ""

			if err != nil {
				ok = false
				errorText = "Please Check Your Internet Connection!"
				break
			}

			if response.StatusCode != http.StatusOK {
				errorText = response.Status
				ok = false
				break
			}

			defer func() { _ = response.Body.Close() }()
			jsonResponse := summaryResponse{}
			decoder := json.NewDecoder(response.Body)
			err = decoder.Decode(&jsonResponse)
			if err != nil {
				errorText = "Could not parse JSON!"
				break
			}

			if !jsonResponse.Success {
				ok = false
				errorText = fmt.Sprintf("%s-%s: %s", baseCurrency.name, mCurrency.name, jsonResponse.Message)
				break
			}
			ok = true
			errorText = ""

			mCurrency.Last = fmt.Sprintf("%f", jsonResponse.Result[0].Last)
			mCurrency.High = fmt.Sprintf("%f", jsonResponse.Result[0].High)
			mCurrency.Low = fmt.Sprintf("%f", jsonResponse.Result[0].Low)
			mCurrency.Volume = fmt.Sprintf("%f", jsonResponse.Result[0].Volume)
			mCurrency.OpenBuyOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenBuyOrders)
			mCurrency.OpenSellOrders = fmt.Sprintf("%d", jsonResponse.Result[0].OpenSellOrders)
		}
	}

	widget.display()
}

func makeRequest(baseName, marketName string) *http.Request {
	url := fmt.Sprintf("%s?market=%s-%s", baseURL, baseName, marketName)
	request, _ := http.NewRequest("GET", url, http.NoBody)

	return request
}
