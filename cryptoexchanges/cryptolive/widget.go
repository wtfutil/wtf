package cryptolive

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

var started = false
var baseURL = "https://min-api.cryptocompare.com/data/price"
var ok = true

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.TextWidget

	*list
}

// NewWidget Make new instance of widget
func NewWidget() *Widget {
	started = false
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" CryptoLive ", "cryptolive", false),
	}

	widget.setList()

	return &widget
}

func (widget *Widget) setList() {
	currenciesMap, _ := Config.Map("wtf.mods.cryptolive.currencies")

	widget.list = &list{}

	for currency := range currenciesMap {
		displayName, _ := Config.String("wtf.mods.cryptolive.currencies." + currency + ".displayName")
		toList := getToList(currency)
		widget.list.addItem(currency, displayName, toList)
	}

}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh() {
	if widget.Disabled() {
		return
	}

	if started == false {
		// this code should run once
		go func() {
			for {
				widget.updateCurrencies()
				time.Sleep(time.Duration(widget.RefreshInterval()) * time.Second)
			}
		}()

	}

	started = true

	widget.UpdateRefreshedAt()
	widget.View.Clear()

	if !ok {
		widget.View.SetText(
			fmt.Sprint("Please check your internet connection!"),
		)
		return
	}
	display(widget)
}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	str := ""
	var (
		fromNameColor        = Config.UString("wtf.mods.cryptolive.colors.from.name", "coral")
		fromDisplayNameColor = Config.UString("wtf.mods.cryptolive.colors.from.displayName", "grey")
		toNameColor          = Config.UString("wtf.mods.cryptolive.colors.to.name", "white")
		toPriceColor         = Config.UString("wtf.mods.cryptolive.colors.to.price", "green")
	)
	for _, item := range widget.list.items {
		str += fmt.Sprintf(" [%s]%s[%s] (%s)\n", fromNameColor, item.displayName, fromDisplayNameColor, item.name)
		for _, toItem := range item.to {
			str += fmt.Sprintf("\t[%s]%s: [%s]%f\n", toNameColor, toItem.name, toPriceColor, toItem.price)
		}
		str += "\n"
	}

	widget.View.SetText(fmt.Sprintf("\n%s", str))
}

func getToList(fromName string) []*toCurrency {
	toNames, _ := Config.List("wtf.mods.cryptolive.currencies." + fromName + ".to")

	var toList []*toCurrency

	for _, to := range toNames {
		toList = append(toList, &toCurrency{
			name:  to.(string),
			price: 0,
		})
	}

	return toList
}

func (widget *Widget) updateCurrencies() {
	defer func() {
		recover()
	}()
	for _, fromCurrency := range widget.list.items {

		var (
			client       http.Client
			jsonResponse cResponse
		)

		client = http.Client{
			Timeout: time.Duration(5 * time.Second),
		}

		request := makeRequest(fromCurrency)
		response, err := client.Do(request)

		if err != nil {
			ok = false
		} else {
			ok = true
		}

		defer response.Body.Close()

		_ = json.NewDecoder(response.Body).Decode(&jsonResponse)

		setPrices(&jsonResponse, fromCurrency)
	}

	display(widget)
}

func makeRequest(currency *fromCurrency) *http.Request {
	fsym := currency.name
	tsyms := ""
	for _, to := range currency.to {
		tsyms += fmt.Sprintf("%s,", to.name)
	}

	url := fmt.Sprintf("%s?fsym=%s&tsyms=%s", baseURL, fsym, tsyms)
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
	}

	return request
}

func setPrices(response *cResponse, currencry *fromCurrency) {
	for idx, toCurrency := range currencry.to {
		currencry.to[idx].price = (*response)[toCurrency.name]
	}
}
