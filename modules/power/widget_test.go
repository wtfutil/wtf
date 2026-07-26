//go:build windows

package power

import (
	"strings"
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

func Test_Widget_content_noBattery(t *testing.T) {
	// Test the logic that content() uses: when battery result equals msgNoBattery,
	// it should not be included. We test this indirectly through battery string check.
	battery := NewBattery()
	battery.result = msgNoBattery

	if battery.String() != msgNoBattery {
		t.Errorf("expected %q, got %q", msgNoBattery, battery.String())
	}

	// Verify the conditional logic
	if battery.String() != msgNoBattery {
		t.Error("battery string should equal msgNoBattery when no battery")
	}
}

func Test_Widget_content_withBattery(t *testing.T) {
	battery := NewBattery()
	battery.result = " some battery info\n"

	if battery.String() == msgNoBattery {
		t.Error("battery string should not equal msgNoBattery when battery present")
	}
	if !strings.Contains(battery.String(), "some battery info") {
		t.Error("battery string should contain the set content")
	}
}

func Test_Widget_productNameTruncation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "short name not truncated",
			input:    "Mouse",
			expected: "Mouse",
		},
		{
			name:     "exactly 14 chars not truncated",
			input:    "12345678901234",
			expected: "12345678901234",
		},
		{
			name:     "longer than 14 is truncated",
			input:    "Very Long Product Name",
			expected: "Very Long Prod",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prodName := tt.input
			if len(prodName) > productNameTrimLen {
				prodName = prodName[:productNameTrimLen]
			}
			if prodName != tt.expected {
				t.Errorf("truncated name = %q, want %q", prodName, tt.expected)
			}
		})
	}
}

func Test_Widget_deviceFiltering(t *testing.T) {
	// Test the logic used in content() for filtering devices
	tests := []struct {
		name       string
		hasBattery string
		shouldShow bool
	}{
		{name: "device with battery shows", hasBattery: "Yes", shouldShow: true},
		{name: "device without battery hidden", hasBattery: "No", shouldShow: false},
		{name: "device with empty battery hidden", hasBattery: "", shouldShow: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dev := NewManagedDevice()
			dev.Attributes["HasBattery"] = tt.hasBattery
			dev.Attributes["BatteryPercent"] = "50"
			dev.Attributes["Product"] = "Test Device"

			if dev.HasBattery() != tt.shouldShow {
				t.Errorf("HasBattery() = %v, want %v", dev.HasBattery(), tt.shouldShow)
			}
		})
	}
}

func Test_Constants(t *testing.T) {
	if msgNoBattery != " no battery found" {
		t.Errorf("msgNoBattery = %q, want %q", msgNoBattery, " no battery found")
	}
	if productNameTrimLen != 14 {
		t.Errorf("productNameTrimLen = %d, want 14", productNameTrimLen)
	}
}

func Test_Settings_defaults(t *testing.T) {
	if defaultFocusable != false {
		t.Error("defaultFocusable should be false")
	}
	if defaultTitle != "Power" {
		t.Errorf("defaultTitle = %q, want %q", defaultTitle, "Power")
	}
}

func Test_NewSettingsFromYAML(t *testing.T) {
	ymlConfig, err := config.ParseYaml("power: {}")
	if err != nil {
		t.Fatalf("failed to parse module config: %v", err)
	}
	moduleConfig, _ := ymlConfig.Get("power")

	globalConfig, err := config.ParseYaml("wtf: {}")
	if err != nil {
		t.Fatalf("failed to parse global config: %v", err)
	}

	settings := NewSettingsFromYAML("power", moduleConfig, globalConfig)
	if settings == nil {
		t.Fatal("NewSettingsFromYAML returned nil")
	}
	if settings.Common == nil {
		t.Fatal("settings.Common is nil")
	}
}

func Test_Widget_content_integration(t *testing.T) {
	ymlConfig, _ := config.ParseYaml("power: {}")
	moduleConfig, _ := ymlConfig.Get("power")
	globalConfig, _ := config.ParseYaml("wtf: {}")

	settings := NewSettingsFromYAML("power", moduleConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool)
	widget := NewWidget(app, redrawChan, settings)

	// Set battery to no battery
	widget.Battery.result = msgNoBattery
	widget.ManagedDevices.Devices = []*ManagedDevice{}

	_, content, wrap := widget.content()
	if !wrap {
		t.Error("content should return wrap=true")
	}
	if !strings.Contains(content, "Source") {
		t.Error("content should contain Source")
	}
	// When no battery, battery info should not appear
	if strings.Contains(content, "Charge") {
		t.Error("should not show battery info when no battery")
	}
}

func Test_Widget_content_withBatteryInfo(t *testing.T) {
	ymlConfig, _ := config.ParseYaml("power: {}")
	moduleConfig, _ := ymlConfig.Get("power")
	globalConfig, _ := config.ParseYaml("wtf: {}")

	settings := NewSettingsFromYAML("power", moduleConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool)
	widget := NewWidget(app, redrawChan, settings)

	widget.Battery.result = " Charge: 85%\n"
	widget.ManagedDevices.Devices = []*ManagedDevice{}

	_, content, _ := widget.content()
	if !strings.Contains(content, "Charge: 85%") {
		t.Errorf("content should include battery info, got:\n%s", content)
	}
}

func Test_Widget_content_withDevices(t *testing.T) {
	ymlConfig, _ := config.ParseYaml("power: {}")
	moduleConfig, _ := ymlConfig.Get("power")
	globalConfig, _ := config.ParseYaml("wtf: {}")

	settings := NewSettingsFromYAML("power", moduleConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool)
	widget := NewWidget(app, redrawChan, settings)

	widget.Battery.result = msgNoBattery

	dev := NewManagedDevice()
	dev.Attributes["HasBattery"] = "Yes"
	dev.Attributes["BatteryPercent"] = "75"
	dev.Attributes["Product"] = "Magic Mouse"
	widget.ManagedDevices.Devices = []*ManagedDevice{dev}

	_, content, _ := widget.content()
	if !strings.Contains(content, "Magic Mouse") {
		t.Errorf("should show managed device, got:\n%s", content)
	}
}

func Test_Widget_content_truncatesLongDeviceName(t *testing.T) {
	ymlConfig, _ := config.ParseYaml("power: {}")
	moduleConfig, _ := ymlConfig.Get("power")
	globalConfig, _ := config.ParseYaml("wtf: {}")

	settings := NewSettingsFromYAML("power", moduleConfig, globalConfig)

	app := tview.NewApplication()
	redrawChan := make(chan bool)
	widget := NewWidget(app, redrawChan, settings)

	widget.Battery.result = msgNoBattery

	dev := NewManagedDevice()
	dev.Attributes["HasBattery"] = "Yes"
	dev.Attributes["BatteryPercent"] = "50"
	dev.Attributes["Product"] = "A Very Long Device Name Here"
	widget.ManagedDevices.Devices = []*ManagedDevice{dev}

	_, content, _ := widget.content()
	if strings.Contains(content, "A Very Long Device Name Here") {
		t.Error("should truncate long device names")
	}
	if !strings.Contains(content, "A Very Long De") {
		t.Errorf("should truncate to 14 chars, got:\n%s", content)
	}
}

