package victorops

// OnCallTeam object to make
// managing objects easier
type OnCallTeam struct {
	Name   string
	Slug   string
	OnCall []OnCall
}

// OnCall object to handle
// different on call policies
type OnCall struct {
	Policy   string
	Userlist string
}
