package newrelic

import (
	"net/url"
	"time"
)

const (
	testAPIKey        = "test_api_key"
	testTimeRawString = "2016-01-20T20:29:38Z"
)

var (
	testTime, _           = time.Parse(time.RFC3339, testTimeRawString)
	testTimeString        = testTime.String()
	testTimeStringEscaped = url.QueryEscape(testTimeString)
)
