package finnhub

import (
    "fmt"
	"testing"
)
func TestFinnhubClient(t *testing.T){
	testClient := &Client {
		symbols: []string{
			"AAPL",
			"MSFT",
		},
		apiKey: "",
	}

	r, err := testClient.Getquote()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r[0].Stock, r[0].C)
	fmt.Println(r[1].Stock, r[1].C)

}