package krisinformation

import (
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// helper to create a test server returning given JSON data
func newTestServer(t *testing.T, data interface{}) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			t.Fatalf("failed to encode response: %v", err)
		}
	}))
}

func makeKrisData(areas []struct {
	Type        string
	Description string
	Coordinate  string
}, headline, sender, pushMsg string, updated time.Time) Krisinformation {
	areaSlice := make([]struct {
		Type                string      `json:"Type"`
		Description         string      `json:"Description"`
		Coordinate          string      `json:"Coordinate"`
		GeometryInformation interface{} `json:"GeometryInformation"`
	}, len(areas))
	for i, a := range areas {
		areaSlice[i].Type = a.Type
		areaSlice[i].Description = a.Description
		areaSlice[i].Coordinate = a.Coordinate
	}

	data := Krisinformation{
		{
			Identifier:  "test-id",
			PushMessage: pushMsg,
			Updated:     updated,
			Published:   updated,
			Headline:    headline,
			Preamble:    "preamble",
			BodyText:    "body",
			Area:        areaSlice,
			Web:         "https://example.com",
			Language:    "sv",
			Event:       "event",
			SenderName:  sender,
			Push:        true,
			SourceID:    1,
		},
	}
	return data
}

func TestNewClient(t *testing.T) {
	c := NewClient(59.0, 18.0, 100, "Stockholm", true)
	if c.latitude != 59.0 {
		t.Errorf("expected latitude 59.0, got %f", c.latitude)
	}
	if c.longitude != 18.0 {
		t.Errorf("expected longitude 18.0, got %f", c.longitude)
	}
	if c.radius != 100 {
		t.Errorf("expected radius 100, got %d", c.radius)
	}
	if c.county != "Stockholm" {
		t.Errorf("expected county Stockholm, got %s", c.county)
	}
	if !c.country {
		t.Error("expected country true")
	}
	if c.apiURL != krisinformationAPI {
		t.Errorf("expected default API URL, got %s", c.apiURL)
	}
}

func TestDistanceInMeters(t *testing.T) {
	tests := []struct {
		name     string
		lat1     float64
		lon1     float64
		lat2     float64
		lon2     float64
		expected float64
		delta    float64
	}{
		{
			name:     "same point",
			lat1:     59.3293, lon1: 18.0686,
			lat2:     59.3293, lon2: 18.0686,
			expected: 0,
			delta:    0.1,
		},
		{
			name:     "Stockholm to Gothenburg approx 400km",
			lat1:     59.3293, lon1: 18.0686,
			lat2:     57.7089, lon2: 11.9746,
			expected: 398000,
			delta:    5000,
		},
		{
			name:     "short distance ~1km",
			lat1:     59.3293, lon1: 18.0686,
			lat2:     59.3380, lon2: 18.0686,
			expected: 967,
			delta:    50,
		},
		{
			name:     "antipodal points",
			lat1:     0, lon1: 0,
			lat2:     0, lon2: 180,
			expected: math.Pi * 6378100,
			delta:    1000,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := DistanceInMeters(tc.lat1, tc.lon1, tc.lat2, tc.lon2)
			if math.Abs(got-tc.expected) > tc.delta {
				t.Errorf("DistanceInMeters(%f,%f,%f,%f) = %f, want %f (±%f)",
					tc.lat1, tc.lon1, tc.lat2, tc.lon2, got, tc.expected, tc.delta)
			}
		})
	}
}

func TestGetKrisinformation_CountryFilter(t *testing.T) {
	now := time.Now()
	data := makeKrisData([]struct {
		Type        string
		Description string
		Coordinate  string
	}{
		{Type: "Country", Description: "Sverige"},
	}, "Country Alert", "MSB", "Country push", now)

	srv := newTestServer(t, data)
	defer srv.Close()

	tests := []struct {
		name      string
		country   bool
		wantCount int
	}{
		{"country enabled matches Country area", true, 1},
		{"country disabled skips Country area", false, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(0, 0, -1, "", tc.country)
			c.apiURL = srv.URL
			items, err := c.getKrisinformation()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(items) != tc.wantCount {
				t.Fatalf("expected %d items, got %d", tc.wantCount, len(items))
			}
			if tc.wantCount > 0 {
				if !items[0].Country {
					t.Error("expected Country flag to be true")
				}
				if items[0].HeadLine != "Country Alert" {
					t.Errorf("expected headline 'Country Alert', got %q", items[0].HeadLine)
				}
				if items[0].SenderName != "MSB" {
					t.Errorf("expected sender 'MSB', got %q", items[0].SenderName)
				}
			}
		})
	}
}

func TestGetKrisinformation_CountyFilter(t *testing.T) {
	now := time.Now()
	data := makeKrisData([]struct {
		Type        string
		Description string
		Coordinate  string
	}{
		{Type: "County", Description: "Stockholms län"},
	}, "County Alert", "Police", "County push", now)

	srv := newTestServer(t, data)
	defer srv.Close()

	tests := []struct {
		name      string
		county    string
		wantCount int
	}{
		{"matching county", "Stockholm", 1},
		{"case insensitive match", "stockholm", 1},
		{"non-matching county", "Göteborg", 0},
		{"empty county config", "", 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(0, 0, -1, tc.county, false)
			c.apiURL = srv.URL
			items, err := c.getKrisinformation()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(items) != tc.wantCount {
				t.Fatalf("expected %d items, got %d", tc.wantCount, len(items))
			}
			if tc.wantCount > 0 {
				if !items[0].County {
					t.Error("expected County flag to be true")
				}
			}
		})
	}
}

func TestGetKrisinformation_RadiusFilter(t *testing.T) {
	now := time.Now()
	// Coordinate near Stockholm: 59.33,18.07
	data := makeKrisData([]struct {
		Type        string
		Description string
		Coordinate  string
	}{
		{Type: "Geocode", Description: "Near Stockholm", Coordinate: "59.33,18.07 59.34,18.08"},
	}, "Nearby Alert", "Fire", "Nearby push", now)

	srv := newTestServer(t, data)
	defer srv.Close()

	tests := []struct {
		name      string
		lat       float64
		lon       float64
		radius    int
		wantCount int
	}{
		{"within radius", 59.33, 18.07, 10000, 1},
		{"outside radius", 57.0, 12.0, 1000, 0},
		{"radius disabled (-1)", 59.33, 18.07, -1, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := NewClient(tc.lat, tc.lon, tc.radius, "", false)
			c.apiURL = srv.URL
			items, err := c.getKrisinformation()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(items) != tc.wantCount {
				t.Fatalf("expected %d items, got %d", tc.wantCount, len(items))
			}
			if tc.wantCount > 0 && items[0].Distance == 0 {
				t.Error("expected non-zero distance for radius match")
			}
		})
	}
}

func TestGetKrisinformation_EmptyCoordinate(t *testing.T) {
	now := time.Now()
	data := makeKrisData([]struct {
		Type        string
		Description string
		Coordinate  string
	}{
		{Type: "Geocode", Description: "No coords", Coordinate: ""},
	}, "No Coord Alert", "MSB", "push", now)

	srv := newTestServer(t, data)
	defer srv.Close()

	c := NewClient(59.33, 18.07, 10000, "", false)
	c.apiURL = srv.URL
	items, err := c.getKrisinformation()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items for empty coordinate, got %d", len(items))
	}
}

func TestGetKrisinformation_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not valid json"))
	}))
	defer srv.Close()

	c := NewClient(0, 0, -1, "", true)
	c.apiURL = srv.URL
	_, err := c.getKrisinformation()
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestGetKrisinformation_ConnectionError(t *testing.T) {
	c := NewClient(0, 0, -1, "", true)
	c.apiURL = "http://127.0.0.1:1" // should refuse connection
	_, err := c.getKrisinformation()
	if err == nil {
		t.Fatal("expected error for connection failure")
	}
}

func TestGetKrisinformation_InvalidCoordinates(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		coordinate string
	}{
		{"bad latitude", "abc,18.07 59.34,18.08"},
		{"bad longitude", "59.33,xyz 59.34,18.08"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			data := makeKrisData([]struct {
				Type        string
				Description string
				Coordinate  string
			}{
				{Type: "Geocode", Description: "Bad coords", Coordinate: tc.coordinate},
			}, "Alert", "MSB", "push", now)

			srv := newTestServer(t, data)
			defer srv.Close()

			c := NewClient(59.33, 18.07, 10000, "", false)
			c.apiURL = srv.URL
			_, err := c.getKrisinformation()
			if err == nil {
				t.Fatal("expected error for invalid coordinates")
			}
		})
	}
}

func TestGetKrisinformation_MultipleAreas(t *testing.T) {
	now := time.Now()
	data := makeKrisData([]struct {
		Type        string
		Description string
		Coordinate  string
	}{
		{Type: "Country", Description: "Sverige"},
		{Type: "County", Description: "Stockholms län"},
	}, "Multi Area Alert", "MSB", "push", now)

	srv := newTestServer(t, data)
	defer srv.Close()

	// With country=true and county="Stockholm", both should match
	c := NewClient(0, 0, -1, "Stockholm", true)
	c.apiURL = srv.URL
	items, err := c.getKrisinformation()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("expected 2 items (country + county), got %d", len(items))
	}
	if !items[0].Country {
		t.Error("first item should be Country")
	}
	if !items[1].County {
		t.Error("second item should be County")
	}
}

func TestGetKrisinformation_EmptyResponse(t *testing.T) {
	srv := newTestServer(t, Krisinformation{})
	defer srv.Close()

	c := NewClient(0, 0, -1, "", true)
	c.apiURL = srv.URL
	items, err := c.getKrisinformation()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("expected 0 items, got %d", len(items))
	}
}
