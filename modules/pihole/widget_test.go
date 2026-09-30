package pihole

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/rivo/tview"

	"github.com/wtfutil/wtf/cfg"
)

func testWidget(t *testing.T, apiURL string) *Widget {
	t.Helper()

	app := tview.NewApplication()
	redrawChan := make(chan bool, 10)

	go func() {
		for range redrawChan {
		}
	}()

	settings := &Settings{
		Common: &cfg.Common{
			Title:   "Test Pi-hole",
			Enabled: true,
		},
		apiUrl:         apiURL,
		token:          "tok",
		showSummary:    true,
		showTopItems:   5,
		showTopClients: 5,
		maxDomainWidth: 20,
	}

	return NewWidget(app, redrawChan, nil, settings)
}

func TestNewWidget(t *testing.T) {
	widget := testWidget(t, "http://example.invalid/admin/api.php")

	if widget == nil {
		t.Fatal("NewWidget() returned nil")
	}

	if widget.settings.Title != "Test Pi-hole" {
		t.Errorf("NewWidget() settings.Title = %q, want %q", widget.settings.Title, "Test Pi-hole")
	}

	if widget.settings.RefreshInterval.Seconds() != 30 {
		t.Errorf("NewWidget() RefreshInterval = %v, want 30s", widget.settings.RefreshInterval)
	}
}

func TestWidget_Content_ServerUnreachable(t *testing.T) {
	widget := testWidget(t, "http://%zz")

	title, content, _ := widget.content()

	if title != "Test Pi-hole" {
		t.Errorf("content() title = %q, want %q", title, "Test Pi-hole")
	}

	if content == "" {
		t.Error("content() body is empty, want error message")
	}
}

func TestWidget_Content_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.RawQuery, "version"):
			_, _ = w.Write([]byte(`{"version":3}`))
		case strings.Contains(r.URL.RawQuery, "topItems"):
			_, _ = w.Write([]byte(`{"top_queries":{"q.com":1},"top_ads":{"a.com":1}}`))
		case strings.Contains(r.URL.RawQuery, "topClients"):
			_, _ = w.Write([]byte(`{"top_sources":{"1.2.3.4":1}}`))
		case strings.Contains(r.URL.RawQuery, "getQueryTypes"):
			_, _ = w.Write([]byte(`{"querytypes":{"A":100}}`))
		case strings.Contains(r.URL.RawQuery, "summary"):
			_, _ = w.Write([]byte(`{"status":"enabled"}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	widget := testWidget(t, ts.URL+"/admin/api.php")

	title, content, _ := widget.content()

	if title != "Test Pi-hole" {
		t.Errorf("content() title = %q, want %q", title, "Test Pi-hole")
	}

	if !strings.Contains(content, "ENABLED") {
		t.Errorf("content() = %q, want it to contain ENABLED", content)
	}
}

func TestWidget_Refresh_Disabled(t *testing.T) {
	widget := testWidget(t, "http://example.invalid/admin/api.php")
	widget.settings.Enabled = false

	// Should return early without panicking when disabled.
	widget.Refresh()
}

func TestWidget_Refresh_Enabled(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.URL.RawQuery, "version"):
			_, _ = w.Write([]byte(`{"version":3}`))
		default:
			_, _ = w.Write([]byte(`{"status":"enabled"}`))
		}
	}))
	defer ts.Close()

	widget := testWidget(t, ts.URL+"/admin/api.php")

	// Should complete without hanging or panicking when enabled.
	widget.Refresh()
}

func TestWidget_AdblockSwitch(t *testing.T) {
	var mu sync.Mutex

	var queries []string

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		queries = append(queries, r.URL.RawQuery)
		mu.Unlock()
		_, _ = w.Write([]byte(`{}`))
	}))
	defer ts.Close()

	widget := testWidget(t, ts.URL+"/admin/api.php")

	widget.disable()

	mu.Lock()
	sawDisable := false

	for _, q := range queries {
		if strings.Contains(q, "disable") {
			sawDisable = true
		}
	}

	queries = nil
	mu.Unlock()

	if !sawDisable {
		t.Error("disable(): no request contained 'disable' in its query")
	}

	widget.enable()

	mu.Lock()
	sawEnable := false

	for _, q := range queries {
		if strings.Contains(q, "enable") {
			sawEnable = true
		}
	}
	mu.Unlock()

	if !sawEnable {
		t.Error("enable(): no request contained 'enable' in its query")
	}
}
