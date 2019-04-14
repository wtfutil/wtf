package toplist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/wtfutil/wtf/wtf"
)

var baseURL = "https://min-api.cryptocompare.com/data/top/exchanges"

// Widget Toplist Widget
type Widget struct {
	Result string

	RefreshInterval int

	list     *cList
	settings *Settings
}

// NewWidget Make new toplist widget
func NewWidget(settings *Settings) *Widget {
	widget := Widget{
		settings: settings,
	}

	widget.list = &cList{}
	widget.setList()

	return &widget
}

func (widget *Widget) setList() {
	for fromCurrency := range widget.settings.top {
		displayName := wtf.Config.UString("wtf.mods.cryptolive.top."+fromCurrency+".displayName", "")
		limit := wtf.Config.UInt("wtf.mods.cryptolive.top."+fromCurrency+".limit", 1)
		widget.list.addItem(fromCurrency, displayName, limit, makeToList(fromCurrency, limit))
	}
}

func makeToList(fCurrencyName string, limit int) (list []*tCurrency) {
	toList, _ := wtf.Config.List("wtf.mods.cryptolive.top." + fCurrencyName + ".to")

	for _, toCurrency := range toList {
		list = append(list, &tCurrency{
			name: toCurrency.(string),
			info: make([]tInfo, limit),
		})
	}

	return
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
