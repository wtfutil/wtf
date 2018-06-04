package bittrex

type summaryList struct {
	items []*fCurrency
}

// fCurrency From Currency
type fCurrency struct {
	name        string
	displayName string
	to          []*tCurrency
}

// tCurrency To Currency
type tCurrency struct {
	name string
	summaryInfo
}

type summaryInfo struct {
	Low            string
	High           string
	Volume         string
	Last           string
	OpenSellOrders string
	OpenBuyOrders  string
}

type summaryResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		MarketName     string  `json:"MarketName"`
		High           float64 `json:"High"`
		Low            float64 `json:"Low"`
		Last           float64 `json:"Last"`
		Volume         float64 `json:"Volume"`
		OpenSellOrders int     `json:"OpenSellOrders"`
		OpenBuyOrders  int     `json:"OpenBuyOrders"`
	} `json:"result"`
}

func (list *summaryList) addSummaryItem(name, displayName string, toList []*tCurrency) {
	list.items = append(list.items, &fCurrency{
		name:        name,
		displayName: displayName,
		to:          toList,
	})
}
