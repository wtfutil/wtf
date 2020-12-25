package covid

import (
	"testing"
)

func TestLatestCases(t *testing.T) {
	latestCasesToAssert, err := LatestCases()
	if err != nil {
		t.Error("LatestCases() returned an error")
	}
	if latestCasesToAssert.Latest.Confirmed == 0 {
		t.Error("LatestCases() should return a non 0 integer")
	}
}
