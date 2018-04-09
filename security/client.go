package security

import ()

func Fetch() map[string]string {
	data := make(map[string]string)

	data["Dns"] = DnsServers()
	data["Enabled"] = FirewallState()
	data["Encryption"] = WifiEncryption()
	data["Network"] = WifiName()
	data["Stealth"] = FirewallStealthState()

	return data
}
