package utils

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_VersionTime(t *testing.T) {
	version, time := VersionTime()
	assert.Equal(t, "dev", version)
	assert.Equal(t, "now", time)
}

func Test_extractVersionTime(t *testing.T) {
	settings := []debug.BuildSetting{
		{Key: "vcs.revision", Value: "test-version"},
		{Key: "vcs.time", Value: "test-time"},
	}
	version, time := extractVersionTime(settings)
	assert.Equal(t, "test-version", version)
	assert.Equal(t, "test-time", time)
}

func Test_extractVersionTime_NoVersionTimeInfo(t *testing.T) {
	version, time := extractVersionTime([]debug.BuildSetting{})
	assert.Equal(t, "dev", version)
	assert.Equal(t, "now", time)
}
