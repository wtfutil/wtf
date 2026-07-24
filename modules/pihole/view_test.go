package pihole

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShorten(t *testing.T) {
	tests := []struct {
		name  string
		input string
		limit int
		want  string
	}{
		{"shorter than limit", "example.com", 20, "example.com"},
		{"equal to limit", "example.com", 11, "example.com"},
		{"longer than limit", "verylongdomainname.example.com", 10, "verylongdo..."},
		{"empty string", "", 5, ""},
		{"zero limit", "abc", 0, "..."},
		{"limit of one", "abcdef", 1, "a..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shorten(tt.input, tt.limit)
			if got != tt.want {
				t.Errorf("shorten(%q, %d) = %q, want %q", tt.input, tt.limit, got, tt.want)
			}
		})
	}
}

func TestSortMapByIntVal(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want [][]string
	}{
		{
			name: "empty map",
			m:    map[string]int{},
			want: nil,
		},
		{
			name: "single entry",
			m:    map[string]int{"a.com": 5},
			want: [][]string{{"a.com", "5"}},
		},
		{
			name: "descending order",
			m:    map[string]int{"low.com": 1, "high.com": 100, "mid.com": 50},
			want: [][]string{{"high.com", "100"}, {"mid.com", "50"}, {"low.com", "1"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortMapByIntVal(tt.m)

			// NOTE: sortMapByIntVal pre-allocates its working slice with
			// make([]kv, len(m)) and then appends onto it, so the result
			// contains len(m) extra zero-value ("", "0") entries after the
			// real ones (for maps with only positive values, since those
			// sort last). Assert on the real, sorted prefix and the shape
			// of the trailing padding rather than exact total length.
			wantLen := len(tt.want) * 2
			if len(got) != wantLen {
				t.Fatalf("sortMapByIntVal(%v) length = %d, want %d (%v)", tt.m, len(got), wantLen, got)
			}

			for i, w := range tt.want {
				if got[i][0] != w[0] || got[i][1] != w[1] {
					t.Errorf("sortMapByIntVal(%v)[%d] = %v, want %v", tt.m, i, got[i], w)
				}
			}

			for i := len(tt.want); i < len(got); i++ {
				if got[i][0] != "" || got[i][1] != "0" {
					t.Errorf("sortMapByIntVal(%v)[%d] = %v, want padding entry [\"\" \"0\"]", tt.m, i, got[i])
				}
			}
		})
	}
}

func TestSortMapByFloatVal(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]float32
		want [][]string
	}{
		{
			name: "empty map",
			m:    map[string]float32{},
			want: nil,
		},
		{
			name: "single entry",
			m:    map[string]float32{"A": 1.5},
			want: [][]string{{"A", "1.50"}},
		},
		{
			name: "descending order",
			m:    map[string]float32{"A": 10.25, "B": 50.5, "C": 5.0},
			want: [][]string{{"B", "50.50"}, {"A", "10.25"}, {"C", "5.00"}},
		},
		{
			name: "skips empty key and zero value",
			m:    map[string]float32{"": 10, "Zero": 0, "Keep": 3.33},
			want: [][]string{{"Keep", "3.33"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortMapByFloatVal(tt.m)

			// NOTE: sortMapByFloatVal pre-allocates its working slice with
			// make([]kv, len(m)) and then appends onto it, so the result
			// contains len(m) extra zero-value ("", "0.00") entries after
			// the real ones (since positive real values sort before the
			// zero padding). Assert on the real, sorted prefix and the
			// shape of the trailing padding rather than exact total length.
			wantLen := len(tt.m) + len(tt.want)
			if len(got) != wantLen {
				t.Fatalf("sortMapByFloatVal(%v) length = %d, want %d (%v)", tt.m, len(got), wantLen, got)
			}

			for i, w := range tt.want {
				if got[i][0] != w[0] || got[i][1] != w[1] {
					t.Errorf("sortMapByFloatVal(%v)[%d] = %v, want %v", tt.m, i, got[i], w)
				}
			}

			for i := len(tt.want); i < len(got); i++ {
				if got[i][0] != "" || got[i][1] != "0.00" {
					t.Errorf("sortMapByFloatVal(%v)[%d] = %v, want padding entry [\"\" \"0.00\"]", tt.m, i, got[i])
				}
			}
		})
	}
}

func TestCreateTable(t *testing.T) {
	tests := []struct {
		name   string
		header []string
	}{
		{"no header", []string{}},
		{"with header", []string{"Col1", "Col2"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			table := createTable(tt.header, &buf)
			if table == nil {
				t.Fatal("createTable() returned nil")
			}
		})
	}
}

func TestGetSummaryView(t *testing.T) {
	tests := []struct {
		name          string
		handler       http.HandlerFunc
		wantContains  string
	}{
		{
			name: "enabled status",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"status":"enabled","domains_being_blocked":"100","dns_queries_today":"200","ads_blocked_today":"10","ads_percentage_today":"5.0","queries_cached":"20","queries_forwarded":"30"}`))
			},
			wantContains: "ENABLED",
		},
		{
			name: "disabled status",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"status":"disabled"}`))
			},
			wantContains: "DISABLED",
		},
		{
			name: "unknown status",
			handler: func(w http.ResponseWriter, r *http.Request) {
				_, _ = w.Write([]byte(`{"status":"weird"}`))
			},
			wantContains: "UNKNOWN",
		},
		{
			name: "server error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			wantContains: "failed to retrieve version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.handler)
			defer ts.Close()

			settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok"}

			got := getSummaryView(*ts.Client(), settings)
			if !strings.Contains(got, tt.wantContains) {
				t.Errorf("getSummaryView() = %q, want it to contain %q", got, tt.wantContains)
			}
		})
	}
}

func TestGetTopItemsView(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"top_queries":{"query1.com":10,"query2.com":5},"top_ads":{"ad1.com":8}}`))
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok", maxDomainWidth: 20}

	got := getTopItemsView(*ts.Client(), settings)
	if !strings.Contains(got, "query1.com") || !strings.Contains(got, "ad1.com") {
		t.Errorf("getTopItemsView() = %q, want it to contain query1.com and ad1.com", got)
	}
}

func TestGetTopItemsView_Error(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL, token: "tok"}

	got := getTopItemsView(*ts.Client(), settings)
	if !strings.Contains(got, "failed to retrieve version") {
		t.Errorf("getTopItemsView() = %q, want error message", got)
	}
}

func TestGetTopClientsView(t *testing.T) {
	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		if strings.Contains(r.URL.RawQuery, "topClients") {
			_, _ = w.Write([]byte(`{"top_sources":{"192.168.1.1":10}}`))
			return
		}
		_, _ = w.Write([]byte(`{"querytypes":{"A":80.5,"AAAA":19.5}}`))
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL + "/admin/api.php", token: "tok", showTopClients: 5}

	got := getTopClientsView(*ts.Client(), settings)
	if !strings.Contains(got, "192.168.1.1") {
		t.Errorf("getTopClientsView() = %q, want it to contain 192.168.1.1", got)
	}
}

func TestGetTopClientsView_TopClientsError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	settings := &Settings{apiUrl: ts.URL, token: "tok", showTopClients: 5}

	got := getTopClientsView(*ts.Client(), settings)
	if !strings.Contains(got, "failed to retrieve version") {
		t.Errorf("getTopClientsView() = %q, want error message", got)
	}
}
