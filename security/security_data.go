package security

import (
	"strings"
)

type SecurityData struct {
	Dns             string
	FirewallEnabled string
	FirewallStealth string
	WifiEncryption  string
	WifiName        string
}

func NewSecurityData() *SecurityData {
	return &SecurityData{}
}

func (data *SecurityData) DnsAt(idx int) string {
	records := strings.Split(data.Dns, "\n")

	if len(records) > 0 && len(records) > idx {
		return records[idx]
	} else {
		return ""
	}
}

func (data *SecurityData) Fetch() {
	data.Dns = DnsServers()
	data.FirewallEnabled = FirewallState()
	data.FirewallStealth = FirewallStealthState()
	data.WifiName = WifiName()
	data.WifiEncryption = WifiEncryption()
}
