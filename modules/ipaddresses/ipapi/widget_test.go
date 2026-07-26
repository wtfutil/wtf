package ipapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/olebedev/config"
)

// helper to build a Settings from YAML snippets
func testSettings(t *testing.T, yamlStr string) *Settings {
	t.Helper()
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, err := config.ParseYaml(yamlStr)
	if err != nil {
		t.Fatalf("failed to parse yaml: %v", err)
	}
	globalConfig, err := config.ParseYaml(globalStr)
	if err != nil {
		t.Fatalf("failed to parse global yaml: %v", err)
	}
	return NewSettingsFromYAML("ipapi", yamlConfig, globalConfig)
}

// helper to build a widget with a mock server URL
func testWidget(t *testing.T, settings *Settings, serverURL string) *Widget {
	t.Helper()
	w := &Widget{
		settings: settings,
		apiURL:   serverURL,
	}
	return w
}

// sampleIPInfo returns a representative ipinfo struct for testing
func sampleIPInfo() *ipinfo {
	return &ipinfo{
		Query:         "203.0.113.1",
		ISP:           "Example ISP",
		AS:            "AS12345 Example Inc.",
		ASName:        "EXAMPLE-AS",
		District:      "Central",
		City:          "San Francisco",
		Region:        "CA",
		RegionName:    "California",
		Country:       "United States",
		CountryCode:   "US",
		Continent:     "North America",
		ContinentCode: "NA",
		Latitude:      37.774929,
		Longitude:     -122.419418,
		PostalCode:    "94102",
		Currency:      "USD",
		Organization:  "Example Org",
		Timezone:      "America/Los_Angeles",
		ReverseDNS:    "host.example.com",
	}
}

func TestFormatableText(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		wantKey  string
		wantVal  string
	}{
		{
			name:    "ip address",
			key:     "IP Address",
			value:   "ip",
			wantKey: "IP Address",
			wantVal: "{{.ip}}",
		},
		{
			name:    "city",
			key:     "City",
			value:   "city",
			wantKey: "City",
			wantVal: "{{.city}}",
		},
		{
			name:    "coordinates",
			key:     "Coordinates",
			value:   "coordinates",
			wantKey: "Coordinates",
			wantVal: "{{.coordinates}}",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatableText(tt.key, tt.value)
			if !strings.Contains(result, tt.wantKey) {
				t.Errorf("expected result to contain key %q, got %q", tt.wantKey, result)
			}
			if !strings.Contains(result, tt.wantVal) {
				t.Errorf("expected result to contain value template %q, got %q", tt.wantVal, result)
			}
			if !strings.Contains(result, "{{.nameColor}}") {
				t.Error("expected result to contain nameColor template variable")
			}
			if !strings.Contains(result, "{{.valueColor}}") {
				t.Error("expected result to contain valueColor template variable")
			}
		})
	}
}

func TestSetResult_DefaultArgs(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, "")

	info := sampleIPInfo()
	w.setResult(info)

	// Default args include ip, isp, as, city, region, country, coordinates, postalCode, organization, timezone
	expectedContents := []string{
		"203.0.113.1",
		"Example ISP",
		"AS12345 Example Inc.",
		"San Francisco",
		"CA",
		"United States",
		"37.774929,-122.419418",
		"94102",
		"Example Org",
		"America/Los_Angeles",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(w.result, expected) {
			t.Errorf("expected result to contain %q, got:\n%s", expected, w.result)
		}
	}
}

func TestSetResult_CustomArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     string
		expected []string
		absent   []string
	}{
		{
			name: "only ip and city",
			args: `
args:
  - ip
  - city
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			expected: []string{"203.0.113.1", "San Francisco"},
			absent:   []string{"Example ISP", "America/Los_Angeles"},
		},
		{
			name: "continent and currency",
			args: `
args:
  - continent
  - continentCode
  - currency
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			expected: []string{"North America", "NA", "USD"},
			absent:   []string{"203.0.113.1", "San Francisco"},
		},
		{
			name: "all fields",
			args: `
args:
  - ip
  - isp
  - as
  - asName
  - district
  - city
  - region
  - regionName
  - country
  - countryCode
  - continent
  - continentCode
  - coordinates
  - postalCode
  - currency
  - organization
  - timezone
  - reverseDNS
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			expected: []string{
				"203.0.113.1", "Example ISP", "AS12345 Example Inc.",
				"EXAMPLE-AS", "Central", "San Francisco", "CA",
				"California", "United States", "US", "North America",
				"NA", "37.774929,-122.419418", "94102", "USD",
				"Example Org", "America/Los_Angeles", "host.example.com",
			},
			absent: []string{},
		},
		{
			name: "reverseDNS only",
			args: `
args:
  - reverseDNS
position:
  top: 0
  left: 0
  height: 1
  width: 1
`,
			expected: []string{"host.example.com"},
			absent:   []string{"203.0.113.1", "San Francisco", "Example ISP"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := testSettings(t, tt.args)
			w := testWidget(t, settings, "")
			w.setResult(sampleIPInfo())

			for _, exp := range tt.expected {
				if !strings.Contains(w.result, exp) {
					t.Errorf("expected result to contain %q, got:\n%s", exp, w.result)
				}
			}
			for _, abs := range tt.absent {
				if strings.Contains(w.result, abs) {
					t.Errorf("expected result NOT to contain %q, got:\n%s", abs, w.result)
				}
			}
		})
	}
}

func TestSetResult_InvalidArg(t *testing.T) {
	yamlStr := `
args:
  - nonexistent
  - ip
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, "")
	w.setResult(sampleIPInfo())

	// "nonexistent" should be ignored, but "ip" should still work
	if !strings.Contains(w.result, "203.0.113.1") {
		t.Errorf("expected result to contain IP, got:\n%s", w.result)
	}
}

func TestSetResult_Colors(t *testing.T) {
	yamlStr := `
args:
  - ip
colors:
  name: "green"
  value: "yellow"
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, "")
	w.setResult(sampleIPInfo())

	if !strings.Contains(w.result, "[green]") {
		t.Errorf("expected result to contain color [green], got:\n%s", w.result)
	}
	if !strings.Contains(w.result, "[yellow]") {
		t.Errorf("expected result to contain color [yellow], got:\n%s", w.result)
	}
}

func TestSetResult_DefaultColors(t *testing.T) {
	yamlStr := `
args:
  - ip
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, "")
	w.setResult(sampleIPInfo())

	// Default colors are "red" and "white"
	if !strings.Contains(w.result, "[red]") {
		t.Errorf("expected default name color [red], got:\n%s", w.result)
	}
	if !strings.Contains(w.result, "[white]") {
		t.Errorf("expected default value color [white], got:\n%s", w.result)
	}
}

func TestSetResult_EmptyIPInfo(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, "")
	w.setResult(&ipinfo{})

	// Should still produce output with labels but empty values
	if !strings.Contains(w.result, "IP Address") {
		t.Errorf("expected result to contain 'IP Address' label, got:\n%s", w.result)
	}
	// Coordinates should be "0.000000,0.000000" for zero lat/lon
	if !strings.Contains(w.result, "0.000000,0.000000") {
		t.Errorf("expected zero coordinates, got:\n%s", w.result)
	}
}

func TestIpinfo_Success(t *testing.T) {
	info := sampleIPInfo()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-Agent") != "curl" {
			t.Errorf("expected User-Agent 'curl', got %q", r.Header.Get("User-Agent"))
		}
		if r.Method != "GET" {
			t.Errorf("expected GET request, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(info)
	}))
	defer srv.Close()

	yamlStr := `
args:
  - ip
  - city
  - country
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, srv.URL)
	w.ipinfo()

	if !strings.Contains(w.result, "203.0.113.1") {
		t.Errorf("expected IP in result, got:\n%s", w.result)
	}
	if !strings.Contains(w.result, "San Francisco") {
		t.Errorf("expected city in result, got:\n%s", w.result)
	}
	if !strings.Contains(w.result, "United States") {
		t.Errorf("expected country in result, got:\n%s", w.result)
	}
}

func TestIpinfo_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	}))
	defer srv.Close()

	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, srv.URL)
	w.ipinfo()

	if w.result == "" {
		t.Error("expected error message in result for invalid JSON")
	}
	if !strings.Contains(w.result, "invalid") {
		t.Errorf("expected error about invalid JSON, got: %s", w.result)
	}
}

func TestIpinfo_ConnectionError(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	// Use an unreachable URL
	w := testWidget(t, settings, "http://127.0.0.1:1")
	w.ipinfo()

	if w.result == "" {
		t.Error("expected error message in result for connection error")
	}
}

func TestIpinfo_InvalidURL(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	// Use an invalid URL that will fail at request creation
	w := testWidget(t, settings, "://invalid")
	w.ipinfo()

	if w.result == "" {
		t.Error("expected error message in result for invalid URL")
	}
}

func TestIpinfo_EmptyResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{}"))
	}))
	defer srv.Close()

	yamlStr := `
args:
  - ip
  - city
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, srv.URL)
	w.ipinfo()

	// Should produce output with labels, even if values are empty
	if !strings.Contains(w.result, "IP Address") {
		t.Errorf("expected label 'IP Address' in result, got:\n%s", w.result)
	}
	if !strings.Contains(w.result, "City") {
		t.Errorf("expected label 'City' in result, got:\n%s", w.result)
	}
}

func TestIpinfo_FullResponseParsing(t *testing.T) {
	// Simulate a realistic full JSON response from ip-api.com
	responseJSON := `{
		"query": "24.48.0.1",
		"isp": "Le Groupe Videotron Ltee",
		"as": "AS5769 Videotron Telecom Ltee",
		"asname": "VIDEOTRON",
		"district": "",
		"city": "Montreal",
		"region": "QC",
		"regionName": "Quebec",
		"country": "Canada",
		"countryCode": "CA",
		"continent": "North America",
		"continentCode": "NA",
		"lat": 45.5017,
		"lon": -73.5673,
		"zip": "H3A",
		"currency": "CAD",
		"org": "Videotron Ltee",
		"timezone": "America/Toronto",
		"reverse": "modemcable001.0-48-24.mc.videotron.ca"
	}`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(responseJSON))
	}))
	defer srv.Close()

	yamlStr := `
args:
  - ip
  - isp
  - as
  - asName
  - city
  - region
  - regionName
  - country
  - countryCode
  - continent
  - continentCode
  - coordinates
  - postalCode
  - currency
  - organization
  - timezone
  - reverseDNS
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	settings := testSettings(t, yamlStr)
	w := testWidget(t, settings, srv.URL)
	w.ipinfo()

	expectations := map[string]string{
		"IP":            "24.48.0.1",
		"ISP":           "Le Groupe Videotron Ltee",
		"AS":            "AS5769 Videotron Telecom Ltee",
		"AS Name":       "VIDEOTRON",
		"City":          "Montreal",
		"Region":        "QC",
		"Region Name":   "Quebec",
		"Country":       "Canada",
		"Country Code":  "CA",
		"Continent":     "North America",
		"Coordinates":   "45.501700,-73.567300",
		"Postal Code":   "H3A",
		"Currency":      "CAD",
		"Organization":  "Videotron Ltee",
		"Timezone":      "America/Toronto",
		"Reverse DNS":   "modemcable001.0-48-24.mc.videotron.ca",
	}

	for label, value := range expectations {
		if !strings.Contains(w.result, value) {
			t.Errorf("expected result to contain %s value %q, got:\n%s", label, value, w.result)
		}
	}
}

func TestArgLookup_AllKeysPresent(t *testing.T) {
	expectedKeys := []string{
		"ip", "isp", "as", "asname", "district", "city", "region",
		"regionname", "country", "countrycode", "continent", "continentcode",
		"coordinates", "postalcode", "currency", "organization", "timezone", "reversedns",
	}

	for _, key := range expectedKeys {
		if _, ok := argLookup[key]; !ok {
			t.Errorf("expected argLookup to contain key %q", key)
		}
	}
}

func TestDefaultAPIURL(t *testing.T) {
	if defaultAPIURL != "http://ip-api.com/json?fields=66846719" {
		t.Errorf("unexpected default API URL: %s", defaultAPIURL)
	}
}
