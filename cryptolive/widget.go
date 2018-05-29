package cryptolive

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cizixs/gohttp"
	"github.com/olebedev/config"
	"github.com/senorprogrammer/wtf/wtf"
)

// Config is a pointer to the global config object
var Config *config.Config

var started = false

// Widget define wtf widget to register widget later
type Widget struct {
	wtf.TextWidget

	// time interval for send http request
	updateInterval int

	*list
}

// NewWidget Make new instance of widget
func NewWidget() *Widget {
	started = false
	widget := Widget{
		TextWidget:     wtf.NewTextWidget(" $ CryptoLive ", "cryptolive", false),
		updateInterval: Config.UInt("wtf.mods.cryptolive.updateInterval", 10),
	}

	currenciesMap, _ := Config.Map("wtf.mods.cryptolive.currencies")

	var currencies []*fromCurrency

	for currency := range currenciesMap {
		displayName, _ := Config.String("wtf.mods.cryptolive.currencies." + currency + ".displayName")
		toCList, _ := Config.List("wtf.mods.cryptolive.currencies." + currency + ".to")
		var toList []*toCurrency
		for _, v := range toCList {
			toList = append(toList, &toCurrency{
				name:  v.(string),
				price: -1,
			})
		}
		currencies = append(currencies, &fromCurrency{
			name:        currency,
			displayName: displayName,
			to:          toList,
		})
	}

	widget.list = &list{
		items: currencies,
	}

	return &widget
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
				widget.updateCurrencies()
				time.Sleep(time.Duration(widget.updateInterval) * time.Second)
			}
		}()

	}

	started = true

	widget.UpdateRefreshedAt()
	widget.View.Clear()
	display(widget)
}

/* -------------------- Unexported Functions -------------------- */

func display(widget *Widget) {
	str := ""
	for _, item := range widget.list.items {
		str += fmt.Sprintf("[coral]%s[gray](%s):\n", item.displayName, item.name)
		for _, toItem := range item.to {
			str += fmt.Sprintf("\t%s[%s]: %f\n", toItem.name, "green", toItem.price)
		}
		str += "\n"
	}

	fmt.Fprintf(
		widget.View,
		"\n%s",
		str,
	)
}

func (widget *Widget) updateCurrencies() {
	defer func() {
		recover()
	}()
	for _, fromCurrency := range widget.list.items {
		request := gohttp.New().Path("data", "price").Query("fsym", fromCurrency.name)
		tsyms := ""
		for _, toCurrency := range fromCurrency.to {
			tsyms += fmt.Sprintf("%s,", toCurrency.name)
		}

		response, err := request.Query("tsyms", tsyms).Get("https://min-api.cryptocompare.com")
		if err != nil {
		}

		jsonResponse := &cResponse{}
		response.AsJSON(jsonResponse)

		responseRef := reflect.Indirect(reflect.ValueOf(jsonResponse))
		for idx, toCurrency := range fromCurrency.to {
			fromCurrency.to[idx].price = responseRef.FieldByName(toCurrency.name).Interface().(float32)
		}

	}
}
