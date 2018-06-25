package toplist

type cList struct {
	items []*fCurrency
}

type fCurrency struct {
	name, displayName string
	limit             int
	to                []*tCurrency
}

type tCurrency struct {
	name string
	info []tInfo
}

type tInfo struct {
	exchange               string
	volume24h, volume24hTo float32
}

type responseInterface struct {
	Response string `json:"Response"`
	Data     []struct {
		Exchange    string  `json:"exchange"`
		FromSymbol  string  `json:"fromSymbol"`
		ToSymbol    string  `json:"toSymbol"`
		Volume24h   float32 `json:"volume24h"`
		Volume24hTo float32 `json:"volume24hTo"`
	} `json:"Data"`
}

func (list *cList) addItem(name, displayName string, limit int, to []*tCurrency) {
	list.items = append(list.items, &fCurrency{
		name:        name,
		displayName: displayName,
		limit:       limit,
		to:          to,
	})
}
