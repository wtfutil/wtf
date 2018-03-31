package security

import ()

func Fetch() map[string]string {
	data := make(map[string]string)

	data["Firewall Enabled"] = FirewallState()
	data["Firewall Stealth"] = FirewallStealthState()
	data["Wifi Encryption"] = WifiEncryption()
	data["Wifi Network"] = WifiName()

	return data
}
