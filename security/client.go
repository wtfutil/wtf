package security

import ()

func Fetch() map[string]string {
	data := make(map[string]string)

	data["Enabled"] = FirewallState()
	data["Stealth"] = FirewallStealthState()
	data["Encryption"] = WifiEncryption()
	data["Network"] = WifiName()

	return data
}
