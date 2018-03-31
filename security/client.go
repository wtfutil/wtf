package security

import ()

func Fetch() map[string]string {
	data := make(map[string]string)

	data["Wifi Network"] = WifiName()
	data["Wifi Encryption"] = WifiEncryption()
	data["Firewall Enabled"] = FirewallState()
	data["Firewall Stealth"] = FirewallStealthState()

	return data
}
