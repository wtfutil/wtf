package hibp

import "time"

// Breach represents a breach in the HIBP system
type Breach struct {
	Date string `json:"BreachDate"`
	Name string `json:"Name"`
}

// BreachDate returns the date of the breach
func (br *Breach) BreachDate() (time.Time, error) {
	dt, err := time.Parse("2006-01-02", br.Date)
	if err != nil {
		// I would much rather return (nil, err) err but that doesn't seem possible
		// Not sure what a better value would be
		return time.Now(), err
	}

	return dt, nil
}
