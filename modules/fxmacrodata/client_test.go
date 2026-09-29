package fxmacrodata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var now = time.Unix(1_789_000_000, 0)

func row(release string, offset time.Duration, topTier, confirmed bool) string {
	return fmt.Sprintf(
		`{"release":%q,"announcement_datetime":%d,"top_tier_for_currency":%t,"release_date_confirmed":%t}`,
		release, now.Add(offset).Unix(), topTier, confirmed,
	)
}

func serve(t *testing.T, byCurrency map[string]string) (*Settings, *http.Client) {
	t.Helper()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currency := r.URL.Path[len("/v1/calendar/"):]
		body, ok := byCurrency[currency]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		_, _ = fmt.Fprintf(w, `{"data":[%s]}`, body)
	}))
	t.Cleanup(server.Close)

	return &Settings{baseURL: server.URL, count: 10, currencies: []string{"usd"}}, server.Client()
}

func TestUpcomingSkipsReleasesThatHaveAlreadyHappened(t *testing.T) {
	settings, client := serve(t, map[string]string{
		"usd": row("inflation", -2*time.Hour, false, true) + "," + row("policy_rate", 3*time.Hour, false, true),
	})

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	require.Len(t, releases, 1)
	assert.Equal(t, "policy rate", releases[0].Name)
}

func TestUpcomingSortsAcrossCurrenciesBySoonest(t *testing.T) {
	settings, client := serve(t, map[string]string{
		"usd": row("inflation", 5*time.Hour, false, true),
		"eur": row("policy_rate", 1*time.Hour, false, true),
	})
	settings.currencies = []string{"usd", "eur"}

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	require.Len(t, releases, 2)
	assert.Equal(t, "EUR", releases[0].Currency)
	assert.Equal(t, "USD", releases[1].Currency)
}

func TestUpcomingHonoursTheCount(t *testing.T) {
	settings, client := serve(t, map[string]string{
		"usd": row("a", time.Hour, false, true) + "," + row("b", 2*time.Hour, false, true) + "," + row("c", 3*time.Hour, false, true),
	})
	settings.count = 2

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	assert.Len(t, releases, 2)
}

func TestTopTierFilter(t *testing.T) {
	settings, client := serve(t, map[string]string{
		"usd": row("inflation", time.Hour, true, true) + "," + row("housing_starts", 2*time.Hour, false, true),
	})
	settings.topTier = true

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	require.Len(t, releases, 1)
	assert.Equal(t, "inflation", releases[0].Name)
}

func TestOneUncoveredCurrencyDoesNotLoseTheOthers(t *testing.T) {
	// A currency outside the subscription must not blank the whole widget.
	settings, client := serve(t, map[string]string{
		"usd": row("inflation", time.Hour, false, true),
	})
	settings.currencies = []string{"usd", "eur"}

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	require.Len(t, releases, 1)
	assert.Equal(t, "USD", releases[0].Currency)
}

func TestEveryCurrencyFailingIsReported(t *testing.T) {
	// Silently rendering an empty widget would look like a quiet calendar
	// rather than a broken key.
	settings, client := serve(t, map[string]string{})

	_, err := upcoming(settings, client, now)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "usd")
}

func TestReleasesWithoutATimeAreSkipped(t *testing.T) {
	settings, client := serve(t, map[string]string{
		"usd": `{"release":"inflation","announcement_datetime":null}`,
	})

	releases, err := upcoming(settings, client, now)

	require.NoError(t, err)
	assert.Empty(t, releases)
}

func TestTheAPIKeyTravelsAsAHeader(t *testing.T) {
	var header, query string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header = r.Header.Get("X-API-Key")
		query = r.URL.RawQuery
		_, _ = fmt.Fprint(w, `{"data":[]}`)
	}))
	defer server.Close()

	settings := &Settings{baseURL: server.URL, count: 5, currencies: []string{"usd"}, apiKey: "test-key"}
	_, err := upcoming(settings, server.Client(), now)

	require.NoError(t, err)
	assert.Equal(t, "test-key", header)
	assert.NotContains(t, query, "test-key")
}

func TestNameFallsBackToTheReleaseSlug(t *testing.T) {
	assert.Equal(t, "Consumer Price Index", label(calendarRow{Release: "inflation", Name: "Consumer Price Index"}))
	assert.Equal(t, "core inflation", label(calendarRow{Release: "core_inflation"}))
}
