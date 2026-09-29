package fxmacrodata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"time"
)

// Release is one scheduled or recent macroeconomic release.
type Release struct {
	Currency             string
	Name                 string
	AnnouncementDatetime int64
	TopTier              bool
	Confirmed            bool
}

type calendarRow struct {
	Release              string `json:"release"`
	Name                 string `json:"name"`
	AnnouncementDatetime *int64 `json:"announcement_datetime"`
	TopTierForCurrency   bool   `json:"top_tier_for_currency"`
	ReleaseDateConfirmed bool   `json:"release_date_confirmed"`
}

type calendarResponse struct {
	Data []calendarRow `json:"data"`
}

// upcoming returns the next releases across every configured currency,
// soonest first.
//
// A currency the key does not cover is skipped rather than failing the widget,
// so a mixed list still renders what it can.
func upcoming(settings *Settings, client *http.Client, now time.Time) ([]Release, error) {
	var (
		releases []Release
		failures []string
	)

	for _, currency := range settings.currencies {
		rows, err := calendar(settings, client, currency)
		if err != nil {
			failures = append(failures, currency)
			continue
		}

		for _, row := range rows {
			if row.AnnouncementDatetime == nil || *row.AnnouncementDatetime == 0 {
				continue
			}
			when := time.Unix(*row.AnnouncementDatetime, 0)
			if when.Before(now) {
				continue
			}
			if settings.topTier && !row.TopTierForCurrency {
				continue
			}

			releases = append(releases, Release{
				Currency:             strings.ToUpper(currency),
				Name:                 label(row),
				AnnouncementDatetime: *row.AnnouncementDatetime,
				TopTier:              row.TopTierForCurrency,
				Confirmed:            row.ReleaseDateConfirmed,
			})
		}
	}

	sort.SliceStable(releases, func(i, j int) bool {
		return releases[i].AnnouncementDatetime < releases[j].AnnouncementDatetime
	})

	if len(releases) > settings.count {
		releases = releases[:settings.count]
	}

	if len(releases) == 0 && len(failures) > 0 {
		return nil, fmt.Errorf("no releases available for %s", strings.Join(failures, ", "))
	}

	return releases, nil
}

func label(row calendarRow) string {
	if row.Name != "" {
		return row.Name
	}
	return strings.ReplaceAll(row.Release, "_", " ")
}

func calendar(settings *Settings, client *http.Client, currency string) ([]calendarRow, error) {
	url := fmt.Sprintf("%s/v1/calendar/%s", settings.baseURL, strings.ToLower(currency))

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Accept", "application/json")
	if settings.apiKey != "" {
		// A header rather than a query parameter, so the key stays out of any
		// URL that might be logged or shown in an error.
		request.Header.Set("X-API-Key", settings.apiKey)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s returned status %d", url, response.StatusCode)
	}

	var payload calendarResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Data, nil
}
