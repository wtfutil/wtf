package exchangerates

import (
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

type Response struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

func FetchExchangeRates(settings *Settings) (map[string]map[string]float64, error) {
	out := map[string]map[string]float64{}

	for base, rates := range settings.rates {
		resp, err := http.Get(fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s", base))
		if err != nil {
			return nil, err
		}
		defer func() { _ = resp.Body.Close() }()

		var data Response
		err = utils.ParseJSON(&data, resp.Body)
		if err != nil {
			return nil, err
		}

		out[base] = map[string]float64{}

		for _, currency := range rates {
			rate, ok := data.Rates[currency]
			if ok {
				out[base][currency] = rate
			}
		}
	}

	return out, nil
}
