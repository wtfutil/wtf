package victorops

// OnCallResponse object
type OnCallResponse struct {
	TeamsOnCall []struct {
		Team struct {
			Name string `json:"name"`
			Slug string `json:"slug"`
		} `json:"team"`
		OnCallNow []struct {
			EscalationPolicy struct {
				Name string `json:"name"`
				Slug string `json:"slug"`
			} `json:"escalationPolicy"`
			Users []struct {
				OnCallUser struct {
					Username string `json:"username"`
				} `json:"onCalluser"`
			} `json:"users"`
		} `json:"oncallNow"`
	} `json:"teamsOnCall"`
}
