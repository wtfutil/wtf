package arpansagovau

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const sampleXML = `<?xml version="1.0" encoding="UTF-8"?>
<stations>
  <location id="adl">
    <name>adl</name>
    <index>3.2</index>
    <time>12:30 PM</time>
    <date>25/07/2026</date>
    <fulldate>Saturday, 25 July 2026</fulldate>
    <utcdatetime>2026/07/25 03:00</utcdatetime>
    <status>ok</status>
  </location>
  <location id="syd">
    <name>syd</name>
    <index>8.7</index>
    <time>1:00 PM</time>
    <date>25/07/2026</date>
    <fulldate>Saturday, 25 July 2026</fulldate>
    <utcdatetime>2026/07/25 03:30</utcdatetime>
    <status>ok</status>
  </location>
  <location id="mel">
    <name>mel</name>
    <index>0.0</index>
    <time>1:00 PM</time>
    <date>25/07/2026</date>
    <fulldate>Saturday, 25 July 2026</fulldate>
    <utcdatetime>2026/07/25 03:30</utcdatetime>
    <status>unavailable</status>
  </location>
</stations>`

func TestParseXML_ValidData(t *testing.T) {
	stations, err := parseXML(strings.NewReader(sampleXML))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stations.Location) != 3 {
		t.Fatalf("expected 3 locations, got %d", len(stations.Location))
	}

	tests := []struct {
		idx    int
		id     string
		index  float32
		time   string
		date   string
		status string
	}{
		{0, "adl", 3.2, "12:30 PM", "25/07/2026", "ok"},
		{1, "syd", 8.7, "1:00 PM", "25/07/2026", "ok"},
		{2, "mel", 0.0, "1:00 PM", "25/07/2026", "unavailable"},
	}

	for _, tc := range tests {
		loc := stations.Location[tc.idx]
		if loc.ID != tc.id {
			t.Errorf("location[%d]: expected ID %q, got %q", tc.idx, tc.id, loc.ID)
		}
		if loc.Index != tc.index {
			t.Errorf("location[%d]: expected index %v, got %v", tc.idx, tc.index, loc.Index)
		}
		if loc.Time != tc.time {
			t.Errorf("location[%d]: expected time %q, got %q", tc.idx, tc.time, loc.Time)
		}
		if loc.Date != tc.date {
			t.Errorf("location[%d]: expected date %q, got %q", tc.idx, tc.date, loc.Date)
		}
		if loc.Status != tc.status {
			t.Errorf("location[%d]: expected status %q, got %q", tc.idx, tc.status, loc.Status)
		}
	}
}

func TestParseXML_EmptyStations(t *testing.T) {
	xml := `<?xml version="1.0" encoding="UTF-8"?><stations></stations>`
	stations, err := parseXML(strings.NewReader(xml))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(stations.Location) != 0 {
		t.Errorf("expected 0 locations, got %d", len(stations.Location))
	}
}

func TestParseXML_InvalidXML(t *testing.T) {
	_, err := parseXML(strings.NewReader("not xml at all"))
	if err == nil {
		t.Fatal("expected error for invalid XML")
	}
}

func TestParseXML_MalformedXML(t *testing.T) {
	xml := `<?xml version="1.0"?><stations><location id="adl"><name>adl</name>`
	_, err := parseXML(strings.NewReader(xml))
	if err == nil {
		t.Fatal("expected error for malformed/incomplete XML")
	}
}

func TestParseXML_EmptyReader(t *testing.T) {
	_, err := parseXML(strings.NewReader(""))
	if err == nil {
		t.Fatal("expected error for empty input")
	}
}

func TestGetLocationData_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		_, _ = w.Write([]byte(sampleXML))
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	loc, err := getLocationData("syd")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.name != "syd" {
		t.Errorf("expected name 'syd', got %q", loc.name)
	}
	if loc.index != 8.7 {
		t.Errorf("expected index 8.7, got %v", loc.index)
	}
	if loc.time != "1:00 PM" {
		t.Errorf("expected time '1:00 PM', got %q", loc.time)
	}
	if loc.status != "ok" {
		t.Errorf("expected status 'ok', got %q", loc.status)
	}
}

func TestGetLocationData_NotFound(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		_, _ = w.Write([]byte(sampleXML))
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	loc, err := getLocationData("nonexistent")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if loc.name != "" {
		t.Errorf("expected empty name for non-existent city, got %q", loc.name)
	}
}

func TestGetLocationData_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	_, err := getLocationData("syd")
	if err == nil {
		t.Fatal("expected error for server error response")
	}
}

func TestGetLocationData_InvalidXMLResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("not xml"))
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	_, err := getLocationData("syd")
	if err == nil {
		t.Fatal("expected error for invalid XML response")
	}
}

func TestGetLocationData_ConnectionRefused(t *testing.T) {
	oldURL := apiURL
	apiURL = "http://127.0.0.1:1"
	defer func() { apiURL = oldURL }()

	_, err := getLocationData("syd")
	if err == nil {
		t.Fatal("expected error for connection refused")
	}
}

func TestApiRequest_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Errorf("expected GET, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<stations></stations>"))
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	resp, err := apiRequest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != 200 {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestApiRequest_Non200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	}))
	defer srv.Close()

	oldURL := apiURL
	apiURL = srv.URL
	defer func() { apiURL = oldURL }()

	_, err := apiRequest()
	if err == nil {
		t.Fatal("expected error for non-200")
	}
	if !strings.Contains(err.Error(), "403") {
		t.Errorf("expected 403 in error, got %q", err.Error())
	}
}
