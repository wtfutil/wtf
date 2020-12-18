package covid

// Cases holds the latest cases
type Cases struct {
	Latest Latest `json:"latest"`
}

// Latest holds the number of confirmed, recovered cases and deaths due to Covid
type Latest struct {
	Confirmed int `json:"confirmed"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
}
