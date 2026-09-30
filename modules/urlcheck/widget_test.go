package urlcheck

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- DoRequest tests ---

func TestDoRequest_Success(t *testing.T) {
	tests := []struct {
		name           string
		handler        http.HandlerFunc
		wantCode       int
		wantMsgPrefix  string
	}{
		{
			name: "200 OK",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			},
			wantCode: http.StatusOK,
		},
		{
			name: "404 Not Found",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			wantCode: http.StatusNotFound,
		},
		{
			name: "500 Internal Server Error",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			},
			wantCode: http.StatusInternalServerError,
		},
		{
			name: "301 Redirect (no follow)",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusMovedPermanently)
			},
			wantCode: http.StatusMovedPermanently,
		},
		{
			name: "503 Service Unavailable",
			handler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusServiceUnavailable)
			},
			wantCode: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(tt.handler)
			defer ts.Close()

			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}
			code, _ := DoRequest(ts.URL, 5*time.Second, client)
			assert.Equal(t, tt.wantCode, code)
		})
	}
}

func TestDoRequest_UsesHeadMethod(t *testing.T) {
	var receivedMethod string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		receivedMethod = r.Method
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	client := &http.Client{}
	DoRequest(ts.URL, 5*time.Second, client)
	assert.Equal(t, http.MethodHead, receivedMethod)
}

func TestDoRequest_InvalidURL(t *testing.T) {
	client := &http.Client{}
	code, msg := DoRequest("://invalid-url", 5*time.Second, client)
	assert.Equal(t, InvalidResultCode, code)
	assert.Equal(t, "New Request Error", msg)
}

func TestDoRequest_ConnectionRefused(t *testing.T) {
	client := &http.Client{}
	// Use a port that's not listening
	code, msg := DoRequest("http://127.0.0.1:1", 2*time.Second, client)
	assert.Equal(t, InvalidResultCode, code)
	assert.Equal(t, "Error", msg)
}

func TestDoRequest_ContextTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
	}))
	defer ts.Close()

	client := &http.Client{}
	code, msg := DoRequest(ts.URL, 1*time.Millisecond, client)
	assert.Equal(t, InvalidResultCode, code)
	assert.Equal(t, "Timeout", msg)
}

// --- getResultColor tests ---

func TestGetResultColor(t *testing.T) {
	tests := []struct {
		name      string
		result    urlResult
		wantColor string
	}{
		{
			name:      "valid 200 is green",
			result:    urlResult{IsValid: true, ResultCode: http.StatusOK},
			wantColor: "[green]",
		},
		{
			name:      "valid 404 is green (below 500)",
			result:    urlResult{IsValid: true, ResultCode: http.StatusNotFound},
			wantColor: "[green]",
		},
		{
			name:      "valid 499 is green",
			result:    urlResult{IsValid: true, ResultCode: 499},
			wantColor: "[green]",
		},
		{
			name:      "valid 500 is red",
			result:    urlResult{IsValid: true, ResultCode: http.StatusInternalServerError},
			wantColor: "[red]",
		},
		{
			name:      "valid 503 is red",
			result:    urlResult{IsValid: true, ResultCode: http.StatusServiceUnavailable},
			wantColor: "[red]",
		},
		{
			name:      "invalid is red regardless of code",
			result:    urlResult{IsValid: false, ResultCode: http.StatusOK},
			wantColor: "[red]",
		},
		{
			name:      "invalid with 999 is red",
			result:    urlResult{IsValid: false, ResultCode: InvalidResultCode},
			wantColor: "[red]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getResultColor(tt.result)
			assert.Equal(t, tt.wantColor, got)
		})
	}
}

// --- FormatResult tests ---

func newTestWidget(urls []*urlResult) *Widget {
	w := &Widget{
		urlList:        urls,
		templateString: "",
	}
	// Build a simple template without color settings
	w.templateString = "{{range .}} " +
		"{{. | getResultColor}}" +
		"[{{if eq .ResultCode 999}}---{{else}}{{.ResultCode}}{{end}}]" +
		" [white]{{.Url}}" +
		" [grey]{{.ResultMessage}}" +
		"\n{{end}}"
	w.PreparedTemplate = template.New("tmpl").Funcs(template.FuncMap{"getResultColor": getResultColor})
	return w
}

func TestFormatResult_EmptyList(t *testing.T) {
	w := newTestWidget(nil)
	result := w.FormatResult()
	assert.Equal(t, "empty URL list", result)
}

func TestFormatResult_EmptySlice(t *testing.T) {
	w := newTestWidget([]*urlResult{})
	result := w.FormatResult()
	assert.Equal(t, "empty URL list", result)
}

func TestFormatResult_WithResults(t *testing.T) {
	urls := []*urlResult{
		{Url: "http://example.com", ResultCode: 200, ResultMessage: "200 OK", IsValid: true},
		{Url: "http://bad.com", ResultCode: 500, ResultMessage: "500 Internal Server Error", IsValid: true},
	}
	w := newTestWidget(urls)
	result := w.FormatResult()

	assert.Contains(t, result, "http://example.com")
	assert.Contains(t, result, "http://bad.com")
	assert.Contains(t, result, "[200]")
	assert.Contains(t, result, "[500]")
	assert.Contains(t, result, "[green]")
	assert.Contains(t, result, "[red]")
}

func TestFormatResult_InvalidResultCode_ShowsDashes(t *testing.T) {
	urls := []*urlResult{
		{Url: "http://timeout.com", ResultCode: InvalidResultCode, ResultMessage: "Timeout", IsValid: true},
	}
	w := newTestWidget(urls)
	result := w.FormatResult()

	assert.Contains(t, result, "[---]")
	assert.Contains(t, result, "Timeout")
}

// --- newUrlResult tests (additional) ---

func TestNewUrlResult_ValidURLs(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"http", "http://example.com"},
		{"https", "https://example.com"},
		{"with port", "http://localhost:8080"},
		{"with path", "https://example.com/path/to/resource"},
		{"with query", "https://example.com?key=value&foo=bar"},
		{"with fragment", "https://example.com/page#section"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newUrlResult(tt.url)
			assert.True(t, result.IsValid)
			assert.Equal(t, tt.url, result.Url)
			assert.Equal(t, 0, result.ResultCode)
			assert.Empty(t, result.ResultMessage)
		})
	}
}

func TestNewUrlResult_InvalidURLs(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"empty string", ""},
		{"just text", "not-a-url"},
		{"missing protocol", "example.com"},
		{"spaces", "http://ex ample.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newUrlResult(tt.url)
			assert.False(t, result.IsValid)
			assert.Equal(t, InvalidResultCode, result.ResultCode)
			assert.NotEmpty(t, result.ResultMessage)
		})
	}
}

// --- Settings tests ---

func TestNewSettingsFromYAML(t *testing.T) {
	// Use olebedev/config directly to create test configs
	tests := []struct {
		name            string
		yaml            string
		wantTimeout     int
		wantURLCount    int
	}{
		{
			name:         "default timeout",
			yaml:         "urls:\n  - http://example.com\n",
			wantTimeout:  30,
			wantURLCount: 1,
		},
		{
			name:         "custom timeout",
			yaml:         "timeout: 10\nurls:\n  - http://a.com\n  - http://b.com\n",
			wantTimeout:  10,
			wantURLCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := makeConfig(tt.yaml)
			require.NoError(t, err)

			globalCfg, err := makeConfig("wtf:\n  colors:\n    border:\n      focusable: darkslateblue\n      focused: orange\n      normal: gray\n")
			require.NoError(t, err)

			settings := NewSettingsFromYAML("urlcheck", cfg, globalCfg)
			assert.Equal(t, tt.wantTimeout, settings.requestTimeout)
			assert.Len(t, settings.urls, tt.wantURLCount)
		})
	}
}

// --- Widget.check integration test ---

func TestWidgetCheck_IntegrationWithHTTPTest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	urlResults := []*urlResult{
		{Url: ts.URL, IsValid: true},
		{Url: "not-a-valid-url", IsValid: false, ResultCode: InvalidResultCode, ResultMessage: "parse error"},
	}

	w := &Widget{
		urlList: urlResults,
		client:  &http.Client{},
		timeout: 5 * time.Second,
	}

	w.check()

	// Valid URL should have been checked
	assert.Equal(t, http.StatusOK, urlResults[0].ResultCode)
	// Invalid URL should remain unchanged
	assert.Equal(t, InvalidResultCode, urlResults[1].ResultCode)
	assert.False(t, urlResults[1].IsValid)
}

func TestWidgetCheck_MultipleURLs(t *testing.T) {
	callCount := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	urlResults := []*urlResult{
		{Url: ts.URL + "/a", IsValid: true},
		{Url: ts.URL + "/b", IsValid: true},
		{Url: ts.URL + "/c", IsValid: true},
	}

	w := &Widget{
		urlList: urlResults,
		client:  &http.Client{},
		timeout: 5 * time.Second,
	}

	w.check()

	assert.Equal(t, 3, callCount)
	for _, ur := range urlResults {
		assert.Equal(t, http.StatusOK, ur.ResultCode)
	}
}

func TestWidgetCheck_SkipsInvalidURLs(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	urlResults := []*urlResult{
		{Url: "invalid", IsValid: false, ResultCode: InvalidResultCode},
	}

	w := &Widget{
		urlList: urlResults,
		client:  &http.Client{},
		timeout: 5 * time.Second,
	}

	w.check()

	// Should not have been modified
	assert.Equal(t, InvalidResultCode, urlResults[0].ResultCode)
}

// --- ConfigText test ---

func TestConfigText(t *testing.T) {
	w := &Widget{
		settings: &Settings{},
	}
	text := w.ConfigText()
	// ConfigText uses HelpFromInterface which returns field help tags
	assert.NotEmpty(t, text)
}

// --- InvalidResultCode constant ---

func TestInvalidResultCode(t *testing.T) {
	assert.Equal(t, 999, InvalidResultCode)
}

// --- FormatResult error path ---

func TestFormatResult_ExecuteError(t *testing.T) {
	urls := []*urlResult{
		{Url: "http://example.com", ResultCode: 200, IsValid: true},
	}
	// A template that calls a function which always errors
	errorFunc := func(ur urlResult) (string, error) {
		return "", fmt.Errorf("deliberate error")
	}
	w := &Widget{
		urlList:        urls,
		templateString: "{{range .}}{{. | errorFunc}}{{end}}",
	}
	w.PreparedTemplate = template.New("tmpl").Funcs(template.FuncMap{
		"getResultColor": getResultColor,
		"errorFunc":      errorFunc,
	})
	result := w.FormatResult()
	assert.Contains(t, result, "deliberate error")
}

// --- PrepareTemplate test ---

func TestPrepareTemplate(t *testing.T) {
	w := &Widget{
		settings: &Settings{
			Common: makeMinimalCommon(),
		},
	}
	w.PrepareTemplate()

	assert.NotNil(t, w.PreparedTemplate)
	assert.NotEmpty(t, w.templateString)
	assert.Contains(t, w.templateString, "getResultColor")
	assert.Contains(t, w.templateString, "ResultCode")
}
