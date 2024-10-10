package covid

import (
	"fmt"
	"net/http"

	"github.com/wtfutil/wtf/utils"
)

const covidTrackerAPIURL = "https://coronavirus-tracker-api.herokuapp.com/v2/"

// LatestCases queries the /latest endpoint, does not take any query parameters
func LatestCases() (*Cases, error) {
	latestURL := covidTrackerAPIURL + "latest"
	resp, err := http.Get(latestURL)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("resp status code is not 2**. Status code: %v", resp.Status)
	}
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
func (widget *Widget) LatestCountryCases(countries []interface{}) ([]*Cases, error) {
	countriesCovidData := []*Cases{}
	for _, name := range countries {
		countryURL := covidTrackerAPIURL + "locations?source=jhu&country_code=" + name.(string)
		resp, err := http.Get(countryURL)
		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("resp status code is not 2**. Status code: %v", resp.Status)
		}
		if err != nil {
			return nil, err
		}
		defer func() { _ = resp.Body.Close() }()

		var latestCountryCases Cases
		err = utils.ParseJSON(&latestCountryCases, resp.Body)
		if err != nil {
			return nil, err
		}
		// add stats for each country to the slice
		countriesCovidData = append(countriesCovidData, &latestCountryCases)

	}

	return countriesCovidData, nil
}
