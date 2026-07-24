package pihole

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{"nil error", nil, "unknown error"},
		{"error without token", errors.New("connection refused"), "connection refused"},
		{"error with token redacted", errors.New("request failed: auth=abc123XYZ"), "request failed: auth=<token>"},
		{"error with token in url query", errors.New("Get \"http://host/api.php?auth=secrettoken123&summary\": timeout"), "Get \"http://host/api.php?auth=<token>&summary\": timeout"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseError(tt.err)
			if got != tt.want {
				t.Errorf("parseError(%v) = %q, want %q", tt.err, got, tt.want)
			}
		})
	}
}

func TestGetClient(t *testing.T) {
	c := getClient()

	if c.Timeout.Seconds() != 21 {
		t.Errorf("getClient() Timeout = %v, want 21s", c.Timeout)
	}
}

func TestGetStatus(t *testing.T) {
	tests := []struct {
		name       string
		apiURL     func(ts *httptest.Server) string
		handler    http.HandlerFunc
		wantErr    bool
		wantStatus string
	}{
		{
			name: "success",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"status":"enabled","domains_being_blocked":"100"}`))
			},
			wantErr:    false,
			wantStatus: "enabled",
		},
		{
			name: "server error",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			wantErr: true,
		},
		{
			name: "invalid json",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`not json`))
			},
			wantErr: true,
		},
		{
			name: "invalid url",
			apiURL: func(ts *httptest.Server) string {
				return "http://%zz"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {},
			wantErr: true,
		},
		{
			name: "invalid query string",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php?%zz"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.handler)
			defer ts.Close()

			status, err := getStatus(*ts.Client(), tt.apiURL(ts))

			if tt.wantErr {
				if err == nil {
					t.Fatalf("getStatus() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("getStatus() unexpected error: %v", err)
			}

			if status.Status != tt.wantStatus {
				t.Errorf("getStatus() Status = %q, want %q", status.Status, tt.wantStatus)
			}
		})
	}
}

func TestCheckServer(t *testing.T) {
	tests := []struct {
		name    string
		apiURL  func(ts *httptest.Server) string
		handler http.HandlerFunc
		wantErr bool
	}{
		{
			name: "supported version",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"version":3}`))
			},
			wantErr: false,
		},
		{
			name: "unsupported version",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"version":5}`))
			},
			wantErr: true,
		},
		{
			name: "http error status",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			wantErr: true,
		},
		{
			name: "invalid json",
			apiURL: func(ts *httptest.Server) string {
				return ts.URL + "/admin/api.php"
			},
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`not json`))
			},
			wantErr: true,
		},
		{
			name: "empty host",
			apiURL: func(ts *httptest.Server) string {
				return ""
			},
			handler: func(w http.ResponseWriter, r *http.Request) {},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.handler)
			defer ts.Close()

			err := checkServer(*ts.Client(), tt.apiURL(ts))

			if tt.wantErr && err == nil {
				t.Errorf("checkServer() expected error, got nil")
			}

			if !tt.wantErr && err != nil {
				t.Errorf("checkServer() unexpected error: %v", err)
			}
		})
	}
}

func TestGetTopItems(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"top_queries":{"a.com":5},"top_ads":{"b.com":2}}`))
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok", showTopItems: 5}

	ti, err := getTopItems(*ts.Client(), settings)
	if err != nil {
		t.Fatalf("getTopItems() unexpected error: %v", err)
	}

	if ti.TopQueries["a.com"] != 5 {
		t.Errorf("getTopItems() TopQueries[a.com] = %d, want 5", ti.TopQueries["a.com"])
	}

	if ti.TopAds["b.com"] != 2 {
		t.Errorf("getTopItems() TopAds[b.com] = %d, want 2", ti.TopAds["b.com"])
	}
}

func TestGetTopItems_Errors(t *testing.T) {
	badSettings := &Settings{apiUrl: "http://%zz"}
	if _, err := getTopItems(http.Client{}, badSettings); err == nil {
		t.Error("getTopItems() with bad URL: expected error, got nil")
	}

	badQuerySettings := &Settings{apiUrl: "http://x/admin/api.php?%zz"}
	if _, err := getTopItems(http.Client{}, badQuerySettings); err == nil {
		t.Error("getTopItems() with bad query string: expected error, got nil")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL, token: "tok", showTopItems: 5}
	if _, err := getTopItems(*ts.Client(), settings); err == nil {
		t.Error("getTopItems() with server error: expected error, got nil")
	}
}

func TestGetTopClients(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"top_sources":{"192.168.1.1":10}}`))
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok", showTopClients: 5}

	tc, err := getTopClients(*ts.Client(), settings)
	if err != nil {
		t.Fatalf("getTopClients() unexpected error: %v", err)
	}

	if tc.TopSources["192.168.1.1"] != 10 {
		t.Errorf("getTopClients() TopSources = %v, want 192.168.1.1:10", tc.TopSources)
	}
}

func TestGetTopClients_Errors(t *testing.T) {
	badSettings := &Settings{apiUrl: "http://%zz"}
	if _, err := getTopClients(http.Client{}, badSettings); err == nil {
		t.Error("getTopClients() with bad URL: expected error, got nil")
	}

	badQuerySettings := &Settings{apiUrl: "http://x/admin/api.php?%zz"}
	if _, err := getTopClients(http.Client{}, badQuerySettings); err == nil {
		t.Error("getTopClients() with bad query string: expected error, got nil")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL, token: "tok", showTopClients: 5}
	if _, err := getTopClients(*ts.Client(), settings); err == nil {
		t.Error("getTopClients() with server error: expected error, got nil")
	}
}

func TestGetQueryTypes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"querytypes":{"A":80.5,"AAAA":19.5}}`))
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok", showTopClients: 5}

	qt, err := getQueryTypes(*ts.Client(), settings)
	if err != nil {
		t.Fatalf("getQueryTypes() unexpected error: %v", err)
	}

	if qt.QueryTypes["A"] != 80.5 {
		t.Errorf("getQueryTypes() QueryTypes[A] = %v, want 80.5", qt.QueryTypes["A"])
	}
}

func TestGetQueryTypes_Errors(t *testing.T) {
	badSettings := &Settings{apiUrl: "http://%zz"}
	if _, err := getQueryTypes(http.Client{}, badSettings); err == nil {
		t.Error("getQueryTypes() with bad URL: expected error, got nil")
	}

	badQuerySettings := &Settings{apiUrl: "http://x/admin/api.php?%zz"}
	if _, err := getQueryTypes(http.Client{}, badQuerySettings); err == nil {
		t.Error("getQueryTypes() with bad query string: expected error, got nil")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL, token: "tok", showTopClients: 5}
	if _, err := getQueryTypes(*ts.Client(), settings); err == nil {
		t.Error("getQueryTypes() with server error: expected error, got nil")
	}
}


