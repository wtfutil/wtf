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
		res, err := http.Get(fmt.Sprintf("https://api.exchangeratesapi.io/latest?base=%s", base))
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var resp Response
		err = utils.ParseJSON(&resp, res.Body)
		if err != nil {
			return nil, err
		}

		out[base] = map[string]float64{}

		for _, currency := range rates {
			rate, ok := resp.Rates[currency]
			if ok {
				out[base][currency] = rate
			}
		}
	}

	return out, nil
}
