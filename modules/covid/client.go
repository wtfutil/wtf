package covid

import (
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

const covidTrackerAPIURL = "https://coronavirus-tracker-api.herokuapp.com/v2/"

// LatestCases queries the /latest endpoint, does not take any query parameters
func LatestCases() (*Cases, error) {
	latestURL := covidTrackerAPIURL + "latest"
	resp, err := http.Get(latestURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var latestGlobalCases Cases
	err = utils.ParseJSON(&latestGlobalCases, resp.Body)
	if err != nil {
		return nil, err
	}

	return &latestGlobalCases, nil
}

// LatestCountryCases queries the /locations endpoint, takes a query parameter: the country code
func (widget *Widget) LatestCountryCases(country string) (*CountryCases, error) {
	countryURL := covidTrackerAPIURL + "locations?source=jhu&country_code=" + widget.settings.country
	resp, err := http.Get(countryURL)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var latestCountryCases CountryCases
	err = utils.ParseJSON(&latestCountryCases, resp.Body)
	if err != nil {
		return nil, err
	}

	return &latestCountryCases, nil
}
