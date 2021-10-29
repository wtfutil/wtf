package power

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func Test_Refresh(t *testing.T) {
	// ioreg -c AppleDeviceManagementHIDEventService -r -l
	data := `
+-o AppleDeviceManagementHIDEventService  <class AppleDeviceManagementHIDEventService, id 0x100000892, registered, matched, active, busy 0 (0 ms), retain 8>
    {
      "LowBatteryNotificationPercentage" = 2
      "PrimaryUsagePage" = 65333
      "BatteryFaultNotificationType" = "TPBatteryFault"
      "HasBattery" = Yes
      "VendorID" = 76
      "VersionNumber" = 0
      "Built-In" = No
      "DeviceAddress" = "3c-9b"
      "WakeReason" = "Button (0x03)"
      "Product" = "Magic Trackpad 2"
      "SerialNumber" = "3c-9b"
      "Transport" = "Bluetooth"
      "BatteryLowNotificationType" = "TPLowBattery"
      "ProductID" = 613
      "DeviceUsagePairs" = ({"DeviceUsagePage"=65333,"DeviceUsage"=11},{"DeviceUsagePage"=65333,"DeviceUsage"=20})
      "IOPersonalityPublisher" = "com.apple.driver.AppleTopCaseHIDEventDriver"
      "BatteryPercent" = 81
      "MTFW Version" = 944
      "BD_ADDR" = <3ca6f6cccc9b>
      "BatteryStatusNotificationType" = "BatteryStatusChanged"
      "CriticallyLowBatteryNotificationPercentage" = 1
      "ReportInterval" = 11250
      "RadioFW Version" = 272
      "VendorIDSource" = 1
      "STFW Version" = 2144
      "CFBundleIdentifier" = "com.apple.driver.AppleTopCaseHIDEventDriver"
      "IOProviderClass" = "IOHIDInterface"
      "LocationID" = 1642556667
      "BluetoothDevice" = Yes
      "IOClass" = "AppleDeviceManagementHIDEventService"
      "HIDServiceSupport" = No
      "CFBundleIdentifierKernel" = "com.apple.driver.AppleTopCaseHIDEventDriver"
      "ProductIDArray" = (613)
      "BatteryStatusFlags" = 0
      "ColorID" = 33
      "IOMatchCategory" = "IODefaultMatchCategory"
      "CountryCode" = 0
      "IOProbeScore" = 7175
      "PrimaryUsage" = 11
      "IOGeneralInterest" = "IOCommand is not serializable"
      "BTFW Version" = 272
    }
    

+-o AppleDeviceManagementHIDEventService  <class AppleDeviceManagementHIDEventService, id 0x10000091b, registered, matched, active, busy 0 (0 ms), retain 8>
    {
		"LowBatteryNotificationPercentage" = 2
		"PrimaryUsagePage" = 65666
		"BatteryFaultNotificationType" = "KBBatteryFault"
		"HasBattery" = Yes
		"VendorID" = 76
		"TrustedAccessoryFW Version" = 5666
		"Built-In" = No
		"DeviceAddress" = "ac-c5"
		"VersionNumber" = 0
		"WakeReason" = "Host (0x01)"
		"Product" = "Magic Keyboard with Touch ID"
		"SerialNumber" = "ac-c5"
		"Transport" = "Bluetooth"
		"BatteryLowNotificationType" = "KB2LowBattery"
		"ProductID" = 666
		"DeviceUsagePairs" = ({"DeviceUsagePage"=65666,"DeviceUsage"=11},{"DeviceUsagePage"=65666,"DeviceUsage"=20})
		"IOPersonalityPublisher" = "com.apple.driver.AppleTopCaseDriverV2"
		"BatteryPercent" = 93
		"BD_ADDR" = <ac49dbbbbbc5>
		"BatteryStatusNotificationType" = "BatteryStatusChanged"
		"CriticallyLowBatteryNotificationPercentage" = 1
		"ReportInterval" = 11250
		"RadioFW Version" = 328
		"VendorIDSource" = 1
		"STFW Version" = 1024
		"CFBundleIdentifier" = "com.apple.driver.AppleTopCaseHIDEventDriver"
		"IOProviderClass" = "IOHIDInterface"
		"LocationID" = 1642556667
		"BluetoothDevice" = Yes
		"IOClass" = "AppleDeviceManagementHIDEventService"
		"HIDServiceSupport" = No
		"CFBundleIdentifierKernel" = "com.apple.driver.AppleTopCaseHIDEventDriver"
		"ProductIDArray" = (666)
		"BatteryStatusFlags" = 0
		"ColorID" = 32
		"IOMatchCategory" = "IODefaultMatchCategory"
		"CountryCode" = 2
		"IOProbeScore" = 7175
		"PrimaryUsage" = 11
		"IOGeneralInterest" = "IOCommand is not serializable"
		"BTFW Version" = 328
    }
`

	manDevices := NewManagedDevices()
	manDevices.Devices = manDevices.parse(data)

	assert.Equal(t, 2, len(manDevices.Devices))

	first := manDevices.Devices[0]
	assert.Equal(t, "Magic Trackpad 2", first.Product())
	assert.Equal(t, int64(81), first.BatteryPercent())
	assert.Equal(t, true, first.BluetoothDevice())
	assert.Equal(t, true, first.HasBattery())
}

func Test_Add(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		expected map[string]string
	}{
		{
			name:     "with empty string",
			src:      "",
			expected: map[string]string{},
		},
		{
			name:     "with no delimiter match",
			src:      "catsdogs",
			expected: map[string]string{},
		},
		{
			name: "with valid src",
			src:  "cats=dogs",
			expected: map[string]string{
				"cats": "dogs",
			},
		},
		{
			name: "with valid multiline src",
			src:  "cats=dogs\nx=y",
			expected: map[string]string{
				"cats": "dogs",
				"x":    "y",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Add(tt.src)

			if !reflect.DeepEqual(tt.expected, manDev.Attributes) {
				t.Errorf("\nexpected %v\n     got %v", tt.expected, manDev.Attributes)
			}
		})
	}
}

func Test_Attributes(t *testing.T) {
	tests := []struct {
		name string
		data string
	}{
		{
			name: "with valid attributes",
			data: `
				"LowBatteryNotificationPercentage" = 2
				"PrimaryUsagePage" = 65666
				"BatteryFaultNotificationType" = "KBBatteryFault"
				"HasBattery" = Yes
				"VendorID" = 76
				"TrustedAccessoryFW Version" = 5666
				"Built-In" = No
				"DeviceAddress" = "ac-c5"
				"VersionNumber" = 0
				"WakeReason" = "Host (0x01)"
				"Product" = "Magic Keyboard with Touch ID"
				"SerialNumber" = "ac-c5"
				"Transport" = "Bluetooth"
				"BatteryLowNotificationType" = "KB2LowBattery"
				"ProductID" = 666
				"DeviceUsagePairs" = ({"DeviceUsagePage"=65666,"DeviceUsage"=11},{"DeviceUsagePage"=65666,"DeviceUsage"=20})
				"IOPersonalityPublisher" = "com.apple.driver.AppleTopCaseDriverV2"
				"BatteryPercent" = 93
				"BD_ADDR" = <ac49dbbbbbc5>
				"BatteryStatusNotificationType" = "BatteryStatusChanged"
				"CriticallyLowBatteryNotificationPercentage" = 1
				"ReportInterval" = 11250
				"RadioFW Version" = 328
				"VendorIDSource" = 1
				"STFW Version" = 1024
				"CFBundleIdentifier" = "com.apple.driver.AppleTopCaseHIDEventDriver"
				"IOProviderClass" = "IOHIDInterface"
				"LocationID" = 1642556667
				"BluetoothDevice" = Yes
				"IOClass" = "AppleDeviceManagementHIDEventService"
				"HIDServiceSupport" = No
				"CFBundleIdentifierKernel" = "com.apple.driver.AppleTopCaseHIDEventDriver"
				"ProductIDArray" = (666)
				"BatteryStatusFlags" = 0
				"ColorID" = 32
				"IOMatchCategory" = "IODefaultMatchCategory"
				"CountryCode" = 2
				"IOProbeScore" = 7175
				"PrimaryUsage" = 11
				"IOGeneralInterest" = "IOCommand is not serializable"
				"BTFW Version" = 328
			`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Add(tt.data)

			assert.Equal(t, manDev.BatteryPercent(), int64(93))
			assert.Equal(t, manDev.BluetoothDevice(), true)
			assert.Equal(t, manDev.BuiltIn(), false)
			assert.Equal(t, manDev.HasBattery(), true)
			assert.Equal(t, manDev.Product(), "Magic Keyboard with Touch ID")
		})
	}
}

func Test_BatteryPercent(t *testing.T) {
	tests := []struct {
		name     string
		percent  string
		expected int64
	}{
		{
			name:     "with empty percent",
			percent:  "",
			expected: -1,
		},
		{
			name:     "with invalid percent",
			percent:  "3a3",
			expected: -1,
		},
		{
			name:     "with negative percent",
			percent:  "-23",
			expected: -23,
		},
		{
			name:     "with valid percent",
			percent:  "23",
			expected: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes["BatteryPercent"] = tt.percent

			actual := manDev.BatteryPercent()
			assert.Equal(t, tt.expected, actual)
		})
	}
}
