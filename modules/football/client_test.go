package football

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

// TestFootballRequestBodyReadable is a regression test for #1874: the response
// body must still be open and fully readable by the caller. Previously
// footballRequest deferred resp.Body.Close() before returning, so callers
// read an already-closed body and got "response body closed" errors.
func TestFootballRequestBodyReadable(t *testing.T) {
	const want = `{"standings":[]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(want))
	}))
	defer ts.Close()

	origURL := footballAPIUrl
	footballAPIUrl = ts.URL
	defer func() { footballAPIUrl = origURL }()

	client := NewClient("test-api-key")
	resp, err := client.footballRequest("standings", 2021)
	assert.NilError(t, err)
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	assert.NilError(t, err)
	assert.Equal(t, string(data), want)
}

// TestFootballRequestErrorStatus verifies that non-2xx responses (e.g. an
// invalid/expired API key or rate limiting) surface a clear error instead of
// returning a response whose body may be unusable.
func TestFootballRequestErrorStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"message":"invalid api key"}`))
	}))
	defer ts.Close()

	origURL := footballAPIUrl
	footballAPIUrl = ts.URL
	defer func() { footballAPIUrl = origURL }()

	client := NewClient("bad-api-key")
	resp, err := client.footballRequest("standings", 2021)
	assert.Assert(t, err != nil)
	assert.Assert(t, resp == nil)
}
