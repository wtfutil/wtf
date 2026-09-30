package power

import (
	"strings"
	"testing"
)

func Test_ManagedDevice_Dump(t *testing.T) {
	tests := []struct {
		name       string
		attributes map[string]string
		wantParts  []string
	}{
		{
			name:       "empty attributes",
			attributes: map[string]string{},
			wantParts:  []string{},
		},
		{
			name:       "single attribute",
			attributes: map[string]string{"Product": "Magic Mouse"},
			wantParts:  []string{"Product", "Magic Mouse"},
		},
		{
			name: "multiple attributes",
			attributes: map[string]string{
				"Product":        "Keyboard",
				"BatteryPercent": "50",
			},
			wantParts: []string{"Product Keyboard", "BatteryPercent 50"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes = tt.attributes

			result := manDev.Dump()

			for _, part := range tt.wantParts {
				if !strings.Contains(result, part) {
					t.Errorf("Dump() should contain %q, got:\n%s", part, result)
				}
			}

			if len(tt.attributes) == 0 && result != "" {
				t.Errorf("Dump() for empty attributes should be empty, got %q", result)
			}
		})
	}
}

func Test_ManagedDevice_BuiltIn(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{name: "yes is built in", value: "Yes", expected: true},
		{name: "no is not built in", value: "No", expected: false},
		{name: "empty is not built in", value: "", expected: false},
		{name: "case sensitive", value: "yes", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes["BuiltIn"] = tt.value
			if manDev.BuiltIn() != tt.expected {
				t.Errorf("BuiltIn() with %q = %v, want %v", tt.value, manDev.BuiltIn(), tt.expected)
			}
		})
	}
}

func Test_ManagedDevice_HasBattery(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{name: "yes has battery", value: "Yes", expected: true},
		{name: "no does not have battery", value: "No", expected: false},
		{name: "empty has no battery", value: "", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes["HasBattery"] = tt.value
			if manDev.HasBattery() != tt.expected {
				t.Errorf("HasBattery() with %q = %v, want %v", tt.value, manDev.HasBattery(), tt.expected)
			}
		})
	}
}

func Test_ManagedDevice_BluetoothDevice(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected bool
	}{
		{name: "yes is bluetooth", value: "Yes", expected: true},
		{name: "no is not bluetooth", value: "No", expected: false},
		{name: "empty is not bluetooth", value: "", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes["BluetoothDevice"] = tt.value
			if manDev.BluetoothDevice() != tt.expected {
				t.Errorf("BluetoothDevice() with %q = %v, want %v", tt.value, manDev.BluetoothDevice(), tt.expected)
			}
		})
	}
}

func Test_ManagedDevice_Product(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "normal product", value: "Magic Trackpad 2", expected: "Magic Trackpad 2"},
		{name: "empty product", value: "", expected: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			manDev := NewManagedDevice()
			manDev.Attributes["Product"] = tt.value
			if manDev.Product() != tt.expected {
				t.Errorf("Product() = %q, want %q", manDev.Product(), tt.expected)
			}
		})
	}
}

func Test_NewManagedDevices(t *testing.T) {
	md := NewManagedDevices()
	if md == nil {
		t.Fatal("NewManagedDevices() returned nil")
	}
	if len(md.Devices) != 0 {
		t.Errorf("new ManagedDevices should have no devices, got %d", len(md.Devices))
	}
}

func Test_ManagedDevices_parse_empty(t *testing.T) {
	md := NewManagedDevices()
	devices := md.parse("")
	if len(devices) != 0 {
		t.Errorf("parse empty string should return no devices, got %d", len(devices))
	}
}

func Test_ManagedDevices_parse_no_braces(t *testing.T) {
	md := NewManagedDevices()
	devices := md.parse("some random text without braces")
	if len(devices) != 0 {
		t.Errorf("parse without braces should return no devices, got %d", len(devices))
	}
}

func Test_ManagedDevices_parse_single_device(t *testing.T) {
	data := "{\nProduct = Test Device\nBatteryPercent = 75\nHasBattery = Yes\n}\n"
	md := NewManagedDevices()
	devices := md.parse(data)
	if len(devices) != 1 {
		t.Fatalf("expected 1 device, got %d", len(devices))
	}
	if devices[0].Attributes["Product"] != "Test Device" {
		t.Errorf("expected Product 'Test Device', got %q", devices[0].Attributes["Product"])
	}
}

func Test_NewManagedDevice(t *testing.T) {
	md := NewManagedDevice()
	if md == nil {
		t.Fatal("NewManagedDevice() returned nil")
	}
	if md.Attributes == nil {
		t.Error("Attributes should be initialized, not nil")
	}
	if len(md.Attributes) != 0 {
		t.Errorf("Attributes should be empty, got %d entries", len(md.Attributes))
	}
}

func Test_ManagedDevice_Add_multipleEquals(t *testing.T) {
	// Lines with multiple = signs should only split on first
	manDev := NewManagedDevice()
	manDev.Add("key=val=ue")
	// strings.Split on "=" will produce 3 parts, so len != 2, not added
	if _, exists := manDev.Attributes["key"]; exists {
		t.Error("line with multiple = should not be added since Split produces >2 parts")
	}
}

func Test_ManagedDevice_Add_quotesRemoved(t *testing.T) {
	manDev := NewManagedDevice()
	manDev.Add(`"Product" = "Magic Mouse"`)
	if manDev.Attributes["Product"] != "Magic Mouse" {
		t.Errorf("expected quotes removed, got %q", manDev.Attributes["Product"])
	}
}
