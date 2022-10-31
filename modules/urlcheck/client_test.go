package urlcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestTimeout(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second * 1)
	}))
	defer ts.Close()

	client := &http.Client{
		Timeout: time.Millisecond * 10,
	}

	timeout := 1 * time.Microsecond
	statusCode, statusMsg := DoRequest(ts.URL, timeout, client)

	assert.Equal(t, 999, statusCode)
	assert.Equal(t, "Timeout", statusMsg)

}
