package hibp

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestWidget(since, apiKey string) *Widget {
	return &Widget{
		settings: newTestSettings(since, apiKey, nil),
	}
}

func TestFullURL(t *testing.T) {
	widget := newTestWidget("", "")

	tests := []struct {
		name      string
		account   string
		truncated bool
		expected  string
	}{
		{"truncated", "test@example.com", true, apiURL + "test@example.com?truncateResponse=true"},
		{"not truncated", "test@example.com", false, apiURL + "test@example.com?truncateResponse=false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, widget.fullURL(tt.account, tt.truncated))
		})
	}
}

func TestValidateHTTPResponse(t *testing.T) {
	widget := newTestWidget("", "")

	tests := []struct {
		name       string
		statusCode int
		body       []byte
		expectNil  bool
		expectMsg  string
	}{
		{
			name:       "ok status",
			statusCode: http.StatusOK,
			body:       []byte(`[]`),
			expectNil:  true,
		},
		{
			name:       "not found status",
			statusCode: http.StatusNotFound,
			body:       []byte(``),
			expectNil:  true,
		},
		{
			name:       "unauthorized with valid json",
			statusCode: http.StatusUnauthorized,
			body:       []byte(`{"statusCode":401,"message":"Access denied, invalid API key"}`),
			expectNil:  false,
			expectMsg:  "Access denied, invalid API key",
		},
		{
			name:       "payment required with valid json",
			statusCode: http.StatusPaymentRequired,
			body:       []byte(`{"statusCode":402,"message":"payment required"}`),
			expectNil:  false,
			expectMsg:  "payment required",
		},
		{
			name:       "unauthorized with bad json",
			statusCode: http.StatusUnauthorized,
			body:       []byte(`not-json`),
			expectNil:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := widget.validateHTTPResponse(tt.statusCode, tt.body)

			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectMsg, result.Message)
			}
		})
	}
}

func TestParseResponseBody(t *testing.T) {
	tests := []struct {
		name          string
		body          []byte
		since         string
		expectError   bool
		expectedCount int
	}{
		{
			name:          "empty body means no breaches",
			body:          []byte(``),
			expectedCount: 0,
		},
		{
			name:          "valid json with breaches",
			body:          []byte(`[{"Name":"Adobe","BreachDate":"2013-10-04"},{"Name":"LinkedIn","BreachDate":"2012-05-05"}]`),
			expectedCount: 2,
		},
		{
			name:        "malformed json",
			body:        []byte(`not-json`),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := newTestWidget(tt.since, "")

			stat, err := widget.parseResponseBody("test@example.com", tt.body)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "test@example.com", stat.Account)
				assert.Len(t, stat.Breaches, tt.expectedCount)
			}
		})
	}
}

func TestFilterBreaches(t *testing.T) {
	breaches := []Breach{
		{Name: "Old Breach", Date: "2010-01-01"},
		{Name: "New Breach", Date: "2020-01-01"},
		{Name: "Malformed Date", Date: "not-a-date"},
	}

	tests := []struct {
		name     string
		since    string
		expected []string
	}{
		{
			name:     "no since filters nothing",
			since:    "",
			expected: []string{"Old Breach", "New Breach", "Malformed Date"},
		},
		{
			name:     "since filters older breaches, keeps malformed",
			since:    "2015-01-01",
			expected: []string{"New Breach", "Malformed Date"},
		},
		{
			name:     "invalid since filters nothing",
			since:    "not-a-date",
			expected: []string{"Old Breach", "New Breach", "Malformed Date"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := newTestWidget(tt.since, "")

			result := widget.filterBreaches(breaches)

			names := make([]string, 0, len(result))
			for _, b := range result {
				names = append(names, b.Name)
			}

			assert.Equal(t, tt.expected, names)
		})
	}
}

func TestFetchForAccount(t *testing.T) {
	origURL := apiURL
	defer func() { apiURL = origURL }()

	t.Run("empty account short circuits", func(t *testing.T) {
		widget := newTestWidget("", "")

		stat, err := widget.fetchForAccount("", "")

		assert.NoError(t, err)
		assert.Nil(t, stat)
	})

	t.Run("successful response with breaches", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "WTFUtil", r.Header.Get("User-Agent"))
			assert.Equal(t, "test-key", r.Header.Get("hibp-api-key"))

			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"Name":"Adobe","BreachDate":"2013-10-04"}]`))
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "test-key")

		stat, err := widget.fetchForAccount("test@example.com", "")

		assert.NoError(t, err)
		assert.NotNil(t, stat)
		assert.Equal(t, "test@example.com", stat.Account)
		assert.True(t, stat.HasBeenCompromised())
	})

	t.Run("no breaches found returns 404", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "")

		stat, err := widget.fetchForAccount("safe@example.com", "")

		assert.NoError(t, err)
		assert.NotNil(t, stat)
		assert.False(t, stat.HasBeenCompromised())
	})

	t.Run("rate limited response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusTooManyRequests)
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "")

		stat, err := widget.fetchForAccount("test@example.com", "")

		// The current validateHTTPResponse implementation only special-cases 401/402,
		// so a 429 falls through to the default case and is treated like a normal response.
		assert.NoError(t, err)
		assert.NotNil(t, stat)
	})

	t.Run("unauthorized response returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"statusCode":401,"message":"Access denied, invalid API key"}`))
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "bad-key")

		stat, err := widget.fetchForAccount("test@example.com", "")

		assert.Error(t, err)
		assert.Nil(t, stat)
		assert.Equal(t, "Access denied, invalid API key", err.Error())
	})

	t.Run("malformed json response returns error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`not-json`))
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "")

		stat, err := widget.fetchForAccount("test@example.com", "")

		assert.Error(t, err)
		assert.Nil(t, stat)
	})

	t.Run("network error returns error", func(t *testing.T) {
		// Point at a server that's been closed so the connection is refused.
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("", "")

		stat, err := widget.fetchForAccount("test@example.com", "")

		assert.Error(t, err)
		assert.Nil(t, stat)
	})

	t.Run("breaches filtered by since date", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`[{"Name":"Old","BreachDate":"2010-01-01"},{"Name":"New","BreachDate":"2022-01-01"}]`))
		}))
		defer server.Close()

		apiURL = server.URL + "/"

		widget := newTestWidget("2015-01-01", "")

		stat, err := widget.fetchForAccount("test@example.com", "2015-01-01")

		assert.NoError(t, err)
		assert.NotNil(t, stat)
		assert.Len(t, stat.Breaches, 1)
		assert.Equal(t, "New", stat.Breaches[0].Name)
	})
}
