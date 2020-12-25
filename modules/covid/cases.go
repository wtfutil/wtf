package covid

// Cases holds the latest cases
type Cases struct {
	Latest Latest `json:"latest"`
}

// Latest holds the number of global confirmed, recovered cases and deaths due to Covid
type Latest struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
}

// CountryCases holds the latest cases for a given country
type CountryCases struct {
	LatestCountryCases LatestCountryCases `json:"latest"`
}

// LatestCountryCases holds the number of confirmed, recovered cases and deaths due to Covid for a given country
type LatestCountryCases struct {
	Confirmed int                    `json:"confirmed"`
	Deaths    int                    `json:"deaths"`
	Locations map[string]interface{} `json:"locations,omitempty"`
}
