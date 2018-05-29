package cryptolive

type list struct {
	items []*fromCurrency
}

type fromCurrency struct {
	name        string
	displayName string
	to          []*toCurrency
}

type toCurrency struct {
	name  string
	price float32
}

type cResponse struct {
	BTC  float32 `json:"BTC"`
	HBZ  float32 `json:"HBZ"`
	ETH  float32 `json:"ETH"`
	EOS  float32 `json:"EOS"`
	BCH  float32 `json:"BCH"`
	TRX  float32 `json:"TRX"`
	XRP  float32 `json:"XRP"`
	LTC  float32 `json:"LTC"`
	ETC  float32 `json:"ETC"`
	ADA  float32 `json:"ADA"`
	CMT  float32 `json:"CMT"`
	DASH float32 `json:"DASH"`
	ZEC  float32 `json:"ZEC"`
	IOT  float32 `json:"IOT"`
	ONT  float32 `json:"ONT"`
	NEO  float32 `json:"NEO"`
	BTG  float32 `json:"BTG"`
	LSK  float32 `json:"LSK"`
	ELA  float32 `json:"ELA"`
	DTA  float32 `json:"DTA"`
	NANO float32 `json:"NANO"`
	WTC  float32 `json:"WTC"`
	DOGE float32 `json:"DOGE"`
	USD  float32 `json:"USD"`
	EUR  float32 `json:"EUR"`
}
