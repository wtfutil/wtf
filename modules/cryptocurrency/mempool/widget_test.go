package mempool

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBTCTxFees_Success(t *testing.T) {
	tests := []struct {
		name     string
		fees     feeStruct
		wantFast int
		wantHalf int
		wantHour int
		wantEco  int
	}{
		{
			name:     "typical fees",
			fees:     feeStruct{FastFee: 25, HalfHourFee: 20, HourFee: 15, EcoFee: 10},
			wantFast: 25,
			wantHalf: 20,
			wantHour: 15,
			wantEco:  10,
		},
		{
			name:     "high fees",
			fees:     feeStruct{FastFee: 150, HalfHourFee: 100, HourFee: 80, EcoFee: 50},
			wantFast: 150,
			wantHalf: 100,
			wantHour: 80,
			wantEco:  50,
		},
		{
			name:     "minimum fees",
			fees:     feeStruct{FastFee: 1, HalfHourFee: 1, HourFee: 1, EcoFee: 1},
			wantFast: 1,
			wantHalf: 1,
			wantHour: 1,
			wantEco:  1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_ = json.NewEncoder(w).Encode(tc.fees)
			}))
			defer srv.Close()

			result := getBTCTxFees(srv.URL)

			wantContains := []string{
				"Fast",
				"30 min",
				"60 min",
				"Eco",
				"sat/vB",
			}
			for _, s := range wantContains {
				if !containsStr(result, s) {
					t.Errorf("expected result to contain %q, got:\n%s", s, result)
				}
			}
		})
	}
}

func TestGetBTCTxFees_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("server error"))
	}))
	defer srv.Close()

	result := getBTCTxFees(srv.URL)
	if !containsStr(result, "error") {
		t.Errorf("expected error message, got: %s", result)
	}
}

func TestGetBTCTxFees_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("not json"))
	}))
	defer srv.Close()

	result := getBTCTxFees(srv.URL)
	if !containsStr(result, "error") {
		t.Errorf("expected error message for invalid JSON, got: %s", result)
	}
}

func TestGetBTCTxFees_ConnectionRefused(t *testing.T) {
	result := getBTCTxFees("http://127.0.0.1:1")
	if !containsStr(result, "error") {
		t.Errorf("expected error message for connection failure, got: %s", result)
	}
}

func TestFeeStruct_JSONParsing(t *testing.T) {
	input := `{"fastestFee":42,"halfHourFee":30,"hourFee":20,"economyFee":5}`
	var f feeStruct
	if err := json.Unmarshal([]byte(input), &f); err != nil {
		t.Fatalf("failed to parse JSON: %v", err)
	}
	if f.FastFee != 42 {
		t.Errorf("FastFee: want 42, got %d", f.FastFee)
	}
	if f.HalfHourFee != 30 {
		t.Errorf("HalfHourFee: want 30, got %d", f.HalfHourFee)
	}
	if f.HourFee != 20 {
		t.Errorf("HourFee: want 20, got %d", f.HourFee)
	}
	if f.EcoFee != 5 {
		t.Errorf("EcoFee: want 5, got %d", f.EcoFee)
	}
}

func TestConfigText(t *testing.T) {
	w := &Widget{settings: &Settings{apiURL: "http://example.com"}}
	text := w.ConfigText()
	if text == "" {
		t.Error("ConfigText should return non-empty string")
	}
}

func containsStr(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > 0 && stringContains(s, substr))
}

func stringContains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
