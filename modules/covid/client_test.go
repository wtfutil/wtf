package covid

import (
	"testing"
)

func TestLatestCases(t *testing.T) {
	latestCasesToAssert, err := LatestCases()
	if err != nil {
		t.Error("LatestCases() returned an error")
	}
	if latestCasesToAssert.Latest.Confirmed == 0 {
		t.Error("LatestCases() should return a non 0 integer")
	}
}

func (widget *Widget) TestCountryCases(t *testing.T) {
	countryList := []string{"US", "FR"}
	c := make([]interface{}, len(countryList))
	for i, v := range countryList {
		c[i] = v
	}
	latestCountryCasesToAssert, err := widget.LatestCountryCases(c)
	if err != nil {
		t.Error("LatestCountryCases() returned an error")
	}
	if len(latestCountryCasesToAssert) == 0 {
		t.Error("LatestCountryCases() should not be empty")
	}

}
