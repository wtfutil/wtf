package power

import (
	"bufio"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/wtfutil/wtf/utils"
)

// ManageDevices are ...
type ManagedDevices struct {
	Devices []*ManagedDevice

	args []string
	cmd  string
}

func NewManagedDevices() *ManagedDevices {
	manDevices := &ManagedDevices{
		Devices: []*ManagedDevice{},

		// This command queries for all managed devices
		args: []string{"-c", "AppleDeviceManagementHIDEventService", "-r", "-l"},
		cmd:  "ioreg",
	}

	return manDevices
}

func (manDevices *ManagedDevices) Refresh() {
	cmd := exec.Command(manDevices.cmd, manDevices.args...)
	data := utils.ExecuteCommand(cmd)

	manDevices.Devices = manDevices.parse(data)
}

/* -------------------- Unexported Functions -------------------- */

// parse takes the output of the command and turns it into ManagedDevice instances
func (manDevices *ManagedDevices) parse(data string) []*ManagedDevice {
	devices := []*ManagedDevice{}

	chunks := utils.FindBetween(data, "{\n", "}\n")

	for _, chunk := range chunks {
		manDev := NewManagedDevice()
		manDev.Add(chunk)

		devices = append(devices, manDev)
	}

	return devices
}

/* -------------------- And Another Thing -------------------- */

// ManagedDevice represents an entry in the output returned by ioreg when
// passed AppleDeviceManagementHIDEventService
type ManagedDevice struct {
	Attributes map[string]string
}

func NewManagedDevice() *ManagedDevice {
	manDev := &ManagedDevice{
		Attributes: map[string]string{},
	}

	return manDev
}

/* -------------------- Exported Functions -------------------- */

// Add takes a chunk of raw text and attempts to parse it as managed device data
// and create an attribute map from it.
/*
	A typical chunk will look like:

		"LowBatteryNotificationPercentage" = 2
      	"BatteryFaultNotificationType" = "TPBatteryFault"
      	...
      	"VersionNumber" = 0

	which should become:

		{
			"LowBatteryNotificationPercentage": 2,
			"BatteryFaultNotificationType": "TPBatteryFault",
			"VersionNumber": 0,
		}
*/
func (manDev *ManagedDevice) Add(chunk string) {
	scanner := bufio.NewScanner(strings.NewReader(chunk))

	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), "\"", "")

		pieces := strings.Split(line, "=")
		if len(pieces) == 2 {
			left := strings.TrimSpace(pieces[0])
			right := strings.TrimSpace(pieces[1])

			manDev.Attributes[left] = right
		}
	}
}

// Dump writes out all the device attributes as a single string
func (manDev *ManagedDevice) Dump() string {
	out := ""

	for attribute, value := range manDev.Attributes {
		out += fmt.Sprintf("%s %s\n", attribute, value)
	}

	return out
}

/* -------------------- Attributes -------------------- */

// BatteryPercent returns the percent of the device battery
func (manDev *ManagedDevice) BatteryPercent() int64 {
	percent, err := strconv.ParseInt(manDev.Attributes["BatteryPercent"], 10, 64)
	if err != nil {
		return -1
	}

	return percent
}

// BluetoothDevice returns whether or not the device supports bluetooth
func (manDev *ManagedDevice) BluetoothDevice() bool {
	return manDev.Attributes["BluetoothDevice"] == "Yes"
}

// BuiltIn returns whether or not the device is built into the computer
func (manDev *ManagedDevice) BuiltIn() bool {
	return manDev.Attributes["BuiltIn"] == "Yes"
}

// HasBattery returns whether or not the device has a battery
func (manDev *ManagedDevice) HasBattery() bool {
	return manDev.Attributes["HasBattery"] == "Yes"
}

// Product returns the name of the device
func (manDev *ManagedDevice) Product() string {
	return manDev.Attributes["Product"]
}
