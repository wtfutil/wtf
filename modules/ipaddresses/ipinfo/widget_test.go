package ipinfo

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/olebedev/config"
)

// testSettings creates a minimal Settings for testing.
func testSettings(apiToken string, pv protocolVersion) *Settings {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
refreshInterval: 300
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, _ := config.ParseYaml(yamlStr)
	globalConfig, _ := config.ParseYaml(globalStr)

	s := NewSettingsFromYAML("ipinfo", yamlConfig, globalConfig)
	s.apiToken = apiToken
	s.protocolVersion = pv
	return s
}

// testWidget creates a Widget with a mock HTTP server for testing.
func testWidget(srv *httptest.Server, apiToken string) *Widget {
	settings := testSettings(apiToken, auto)
	w := &Widget{
		settings:   settings,
		httpClient: srv.Client(),
		baseURL:    srv.URL,
	}
	return w
}

func TestIpinfo_SuccessfulParse(t *testing.T) {
	expected := ipinfo{
		Ip:           "8.8.8.8",
		Hostname:     "dns.google",
		City:         "Mountain View",
		Region:       "California",
		Country:      "US",
		Coordinates:  "37.4056,-122.0775",
		PostalCode:   "94043",
		Organization: "AS15169 Google LLC",
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(expected)
	}))
	defer srv.Close()

	widget := testWidget(srv, "")
	widget.ipinfo()

	if widget.result == "" {
		t.Fatal("expected non-empty result")
	}
	if !strings.Contains(widget.result, "8.8.8.8") {
		t.Errorf("result should contain IP, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "dns.google") {
		t.Errorf("result should contain hostname, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "Mountain View") {
		t.Errorf("result should contain city, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "California") {
		t.Errorf("result should contain region, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "US") {
		t.Errorf("result should contain country, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "37.4056,-122.0775") {
		t.Errorf("result should contain coordinates, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "AS15169 Google LLC") {
		t.Errorf("result should contain organization, got: %s", widget.result)
	}
}

func TestIpinfo_AuthorizationHeader(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		wantAuth  bool
		wantValue string
	}{
		{
			name:     "no token",
			token:    "",
			wantAuth: false,
		},
		{
			name:      "with token",
			token:     "test-token-123",
			wantAuth:  true,
			wantValue: "Bearer test-token-123",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				auth := r.Header.Get("Authorization")
				if tc.wantAuth {
					if auth != tc.wantValue {
						t.Errorf("expected Authorization %q, got %q", tc.wantValue, auth)
					}
				} else {
					if auth != "" {
						t.Errorf("expected no Authorization header, got %q", auth)
					}
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"ip":"1.2.3.4"}`))
			}))
			defer srv.Close()

			widget := testWidget(srv, tc.token)
			widget.ipinfo()
		})
	}
}

func TestIpinfo_UserAgentHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := r.Header.Get("User-Agent")
		if ua != "curl" {
			t.Errorf("expected User-Agent 'curl', got %q", ua)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ip":"1.2.3.4"}`))
	}))
	defer srv.Close()

	widget := testWidget(srv, "")
	widget.ipinfo()
}

func TestIpinfo_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	}))
	defer srv.Close()

	widget := testWidget(srv, "")
	widget.ipinfo()

	if widget.result == "" {
		t.Fatal("expected error in result")
	}
	if !strings.Contains(widget.result, "invalid") {
		t.Errorf("expected JSON parse error, got: %s", widget.result)
	}
}

func TestIpinfo_ConnectionError(t *testing.T) {
	settings := testSettings("", auto)
	widget := &Widget{
		settings:   settings,
		httpClient: &http.Client{},
		baseURL:    "http://127.0.0.1:1", // port 1 should refuse
	}

	widget.ipinfo()

	if widget.result == "" {
		t.Fatal("expected error in result")
	}
}

func TestSetResult_FormatsAllFields(t *testing.T) {
	settings := testSettings("", auto)
	widget := &Widget{settings: settings}

	info := &ipinfo{
		Ip:           "192.168.1.1",
		Hostname:     "myhost.local",
		City:         "Seattle",
		Region:       "Washington",
		Country:      "US",
		Coordinates:  "47.6,-122.3",
		PostalCode:   "98101",
		Organization: "AS1234 Test Corp",
	}

	widget.setResult(info)

	fields := []struct {
		label string
		value string
	}{
		{"IP", "192.168.1.1"},
		{"Hostname", "myhost.local"},
		{"City", "Seattle"},
		{"Region", "Washington"},
		{"Country", "US"},
		{"Loc", "47.6,-122.3"},
		{"Org", "AS1234 Test Corp"},
	}

	for _, f := range fields {
		if !strings.Contains(widget.result, f.value) {
			t.Errorf("result missing %s value %q, got: %s", f.label, f.value, widget.result)
		}
	}
}

func TestSetResult_EmptyFields(t *testing.T) {
	settings := testSettings("", auto)
	widget := &Widget{settings: settings}

	info := &ipinfo{}
	widget.setResult(info)

	// Should still produce output (template runs with empty strings)
	if widget.result == "" {
		t.Error("expected non-empty result even with empty fields")
	}
	if !strings.Contains(widget.result, "IP") {
		t.Errorf("result should contain label 'IP', got: %s", widget.result)
	}
}

func TestFormatableText(t *testing.T) {
	tests := []struct {
		key   string
		value string
	}{
		{"IP", "Ip"},
		{"Hostname", "Hostname"},
		{"City", "City"},
	}

	for _, tc := range tests {
		t.Run(tc.key, func(t *testing.T) {
			result := formatableText(tc.key, tc.value)
			if !strings.Contains(result, tc.key) {
				t.Errorf("expected key %q in result, got: %s", tc.key, result)
			}
			if !strings.Contains(result, tc.value) {
				t.Errorf("expected value %q in result, got: %s", tc.value, result)
			}
			if !strings.Contains(result, "subheadingColor") {
				t.Error("expected template variable subheadingColor")
			}
			if !strings.Contains(result, "valueColor") {
				t.Error("expected template variable valueColor")
			}
		})
	}
}

func TestProtocolVersion_String(t *testing.T) {
	tests := []struct {
		input    protocolVersion
		expected string
	}{
		{ipV4, "v4"},
		{ipV6, "v6"},
		{auto, "auto"},
		{protocolVersion("unknown"), "auto"},
	}

	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			if got := tc.input.String(); got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

func TestProtocolVersion_ToNetwork(t *testing.T) {
	tests := []struct {
		input    protocolVersion
		expected string
	}{
		{ipV4, "tcp4"},
		{ipV6, "tcp6"},
		{auto, "tcp"},
		{protocolVersion("other"), "tcp"},
	}

	for _, tc := range tests {
		t.Run(string(tc.input), func(t *testing.T) {
			if got := tc.input.toNetwork(); got != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, got)
			}
		})
	}
}

func TestNewProtocolVersion(t *testing.T) {
	tests := []struct {
		input   string
		want    protocolVersion
		wantErr bool
	}{
		{"v4", ipV4, false},
		{"v6", ipV6, false},
		{"auto", auto, false},
		{"invalid", "", true},
		{"", "", true},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := newProtocolVersion(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if got != tc.want {
					t.Errorf("expected %q, got %q", tc.want, got)
				}
			}
		})
	}
}

func TestNewSettingsFromYAML_Defaults(t *testing.T) {
	yamlStr := `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, _ := config.ParseYaml(yamlStr)
	globalConfig, _ := config.ParseYaml(globalStr)

	settings := NewSettingsFromYAML("ipinfo", yamlConfig, globalConfig)

	if settings.apiToken != "" {
		t.Errorf("expected empty apiToken, got %q", settings.apiToken)
	}
	if settings.protocolVersion != auto {
		t.Errorf("expected protocol version 'auto', got %q", settings.protocolVersion)
	}
}

func TestNewSettingsFromYAML_WithToken(t *testing.T) {
	yamlStr := `
apiToken: "my-secret-token"
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
	globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
	yamlConfig, _ := config.ParseYaml(yamlStr)
	globalConfig, _ := config.ParseYaml(globalStr)

	settings := NewSettingsFromYAML("ipinfo", yamlConfig, globalConfig)

	if settings.apiToken != "my-secret-token" {
		t.Errorf("expected apiToken 'my-secret-token', got %q", settings.apiToken)
	}
}

func TestNewSettingsFromYAML_ProtocolVersions(t *testing.T) {
	tests := []struct {
		name     string
		pvConfig string
		expected protocolVersion
	}{
		{"v4", `protocolVersion: "v4"`, ipV4},
		{"v6", `protocolVersion: "v6"`, ipV6},
		{"auto", `protocolVersion: "auto"`, auto},
		{"invalid falls back to auto", `protocolVersion: "invalid"`, auto},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			yamlStr := tc.pvConfig + `
position:
  top: 0
  left: 0
  height: 1
  width: 1
`
			globalStr := `
wtf:
  colors:
    border:
      focusable: "darkslateblue"
      focused: "orange"
      normal: "gray"
`
			yamlConfig, _ := config.ParseYaml(yamlStr)
			globalConfig, _ := config.ParseYaml(globalStr)

			settings := NewSettingsFromYAML("ipinfo", yamlConfig, globalConfig)

			if settings.protocolVersion != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, settings.protocolVersion)
			}
		})
	}
}

func TestIpinfo_EmptyResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	widget := testWidget(srv, "")
	widget.ipinfo()

	// Empty JSON should still produce formatted output without error
	if strings.Contains(widget.result, "error") || strings.Contains(widget.result, "Error") {
		t.Errorf("expected no error for empty JSON object, got: %s", widget.result)
	}
}

func TestIpinfo_PartialResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ip":"10.0.0.1","city":"Portland"}`))
	}))
	defer srv.Close()

	widget := testWidget(srv, "")
	widget.ipinfo()

	if !strings.Contains(widget.result, "10.0.0.1") {
		t.Errorf("result should contain IP, got: %s", widget.result)
	}
	if !strings.Contains(widget.result, "Portland") {
		t.Errorf("result should contain city, got: %s", widget.result)
	}
}
