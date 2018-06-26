package toplist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/olebedev/config"
)

// Config is a pointer to the global config object
var Config *config.Config
var baseURL = "https://min-api.cryptocompare.com/data/top/exchanges"

type textColors struct {
	from struct {
		name        string
		displayName string
	}
	to struct {
		name  string
		field string
		value string
	}
}

// Widget Toplist Widget
type Widget struct {
	Result string

	RefreshInterval int

	list *cList

	colors textColors
}

// NewWidget Make new toplist widget
func NewWidget() *Widget {
	widget := Widget{}

	widget.list = &cList{}
	widget.setList()
	widget.config()

	return &widget
}

func (widget *Widget) setList() {
	currenciesMap, _ := Config.Map("wtf.mods.cryptolive.top")

	for fromCurrency := range currenciesMap {
		displayName := Config.UString("wtf.mods.cryptolive.top."+fromCurrency+".displayName", "")
		limit := Config.UInt("wtf.mods.cryptolive.top."+fromCurrency+".limit", 1)
		widget.list.addItem(fromCurrency, displayName, limit, makeToList(fromCurrency, limit))
	}
}

func makeToList(fCurrencyName string, limit int) (list []*tCurrency) {
	toList, _ := Config.List("wtf.mods.cryptolive.top." + fCurrencyName + ".to")

	for _, toCurrency := range toList {
		list = append(list, &tCurrency{
			name: toCurrency.(string),
			info: make([]tInfo, limit),
		})
	}

	return
}

func (widget *Widget) config() {
	// set colors
	widget.colors.from.name = Config.UString("wtf.mods.cryptolive.colors.top.from.name", "coral")
	widget.colors.from.displayName = Config.UString("wtf.mods.cryptolive.colors.top.from.displayName", "grey")
	widget.colors.to.name = Config.UString("wtf.mods.cryptolive.colors.top.to.name", "red")
	widget.colors.to.field = Config.UString("wtf.mods.cryptolive.colors.top.to.field", "white")
	widget.colors.to.value = Config.UString("wtf.mods.cryptolive.colors.top.to.value", "value")
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh(wg *sync.WaitGroup) {
	if len(widget.list.items) == 0 {
		return
	}

	widget.updateData()

	widget.display()
	wg.Done()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) updateData() {
	defer func() {
		recover()
	}()

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	for _, fromCurrency := range widget.list.items {
		for _, toCurrency := range fromCurrency.to {

			request := makeRequest(fromCurrency.name, toCurrency.name, fromCurrency.limit)
			response, _ := client.Do(request)

			var jsonResponse responseInterface

			err := json.NewDecoder(response.Body).Decode(&jsonResponse)

			if err != nil {
				os.Exit(1)
			}

			for idx, info := range jsonResponse.Data {
				toCurrency.info[idx] = tInfo{
					exchange:    info.Exchange,
					volume24h:   info.Volume24h,
					volume24hTo: info.Volume24hTo,
				}
			}

		}
	}
}

func makeRequest(fsym, tsym string, limit int) *http.Request {
	url := fmt.Sprintf("%s?fsym=%s&tsym=%s&limit=%d", baseURL, fsym, tsym, limit)
	request, _ := http.NewRequest("GET", url, nil)
	return request
}
