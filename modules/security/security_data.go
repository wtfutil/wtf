package security

type SecurityData struct {
	Dns             []string
	FirewallEnabled string
	FirewallStealth string
	LoggedInUsers   []string
	WifiEncryption  string
	WifiName        string
}

func NewSecurityData() *SecurityData {
	return &SecurityData{}
}

func (data SecurityData) DnsAt(idx int) string {
	if len(data.Dns) > idx {
		return data.Dns[idx]
	}
	return ""
}

func (data *SecurityData) Fetch() {
	data.Dns = DnsServers()
	data.FirewallEnabled = FirewallState()
	data.FirewallStealth = FirewallStealthState()
	data.LoggedInUsers = LoggedInUsers()
	data.WifiName = WifiName()
	data.WifiEncryption = WifiEncryption()
}
