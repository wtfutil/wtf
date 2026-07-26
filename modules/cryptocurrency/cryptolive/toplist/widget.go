package toplist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var baseURL = "https://min-api.cryptocompare.com/data/top/exchanges"

// Widget Toplist Widget
type Widget struct {
	Result string

	RefreshInterval time.Duration

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
	for symbol, currency := range widget.settings.top {
		toList := widget.makeToList(symbol, currency.limit)
		widget.list.addItem(symbol, currency.displayName, currency.limit, toList)
	}
}

func (widget *Widget) makeToList(symbol string, limit int) (list []*tCurrency) {
	for _, to := range widget.settings.top[symbol].to {
		list = append(list, &tCurrency{
			name: to.(string),
			info: make([]tInfo, limit),
		})
	}

	return
}

/* -------------------- Exported Functions -------------------- */

// Refresh & update after interval time
func (widget *Widget) Refresh(wg *sync.WaitGroup) {
	if len(widget.list.items) != 0 {

		widget.updateData()

		widget.display()
	}
	wg.Done()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) updateData() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in updateSummary()", r)
		}
	}()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for _, fromCurrency := range widget.list.items {
		for _, toCurrency := range fromCurrency.to {

			request := makeRequest(fromCurrency.name, toCurrency.name, fromCurrency.limit)
			response, err := client.Do(request)

			if err != nil {
				widget.Result = fmt.Sprintf("API error: %v", err)
				return
			}

			if response.StatusCode != http.StatusOK {
				widget.Result = fmt.Sprintf("API error: %s (check API key)", response.Status)
				_ = response.Body.Close()
				return
			}

			var jsonResponse responseInterface

			err = json.NewDecoder(response.Body).Decode(&jsonResponse)
			_ = response.Body.Close()

			if err != nil {
				widget.Result = fmt.Sprintf("JSON decode error: %v", err)
				return
			}

			for idx, info := range jsonResponse.Data {
				if idx >= len(toCurrency.info) {
					break
				}
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
	request, _ := http.NewRequest("GET", url, http.NoBody)
	return request
}
