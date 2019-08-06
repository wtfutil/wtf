package hibp

// Status represents the status of an account in the HIBP system
type Status struct {
	Account  string
	Breaches []Breach
}

// NewStatus creates and returns an instance of Status
func NewStatus(acct string, breaches []Breach) *Status {
	stat := Status{
		Account:  acct,
		Breaches: breaches,
	}

	return &stat
}

// HasBeenCompromised returns TRUE if the specified account has any breaches associated
// with it, FALSE if no breaches are associated with it
func (stat *Status) HasBeenCompromised() bool {
	return stat.Len() > 0
}

// Len returns the number of breaches found for the specified account
func (stat *Status) Len() int {
	if stat == nil || stat.Breaches == nil {
		return 0
	}

	return len(stat.Breaches)
}
