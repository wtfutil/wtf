package covid

import (
	"encoding/json"
	"testing"
)

func Test_CasesInclude(t *testing.T) {
	// The api does not seem to return the correct recovered numbers
	responseBody := `{"latest":{"confirmed":3093619,"deaths":73018,"recovered":0},"locations":[]}`
	latestData := Cases{}
	_ = json.Unmarshal([]byte(responseBody), &latestData)
	expectedConfirmed := 3093619
	expectedDeaths := 73018
	actualConfirmed := latestData.Latest.Confirmed
	actualDeaths := latestData.Latest.Deaths

	if expectedConfirmed != actualConfirmed {
		t.Errorf("\nexpected: %v\n     got: %v", expectedConfirmed, actualConfirmed)
	}
	if expectedDeaths != actualDeaths {
		t.Errorf("\nexpected: %v\n     got: %v", expectedDeaths, actualDeaths)
	}
}
