package bittrex

type summaryList struct {
	items []*bCurrency
}

// Base Currency
type bCurrency struct {
	name        string
	displayName string
	markets     []*mCurrency
}

// Market Currency
type mCurrency struct {
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

func (list *summaryList) addSummaryItem(name, displayName string, marketList []*mCurrency) {
	list.items = append(list.items, &bCurrency{
		name:        name,
		displayName: displayName,
		markets:     marketList,
	})
}
