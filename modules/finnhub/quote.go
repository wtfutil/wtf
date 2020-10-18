package finnhub

type Quote struct {
	C  float64 `json:"c"`
	H  float64 `json:"h"`
	L  float64 `json:"l"`
	O  float64 `json:"o"`
	Pc float64 `json:"pc"`
	T  int     `json:"t"`

	Stock string
}