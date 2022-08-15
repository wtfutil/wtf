package app

import (
	"testing"
	"time"

	"github.com/olebedev/config"
)

const (
	configExample = `
  wtf:
    mods:
      clocks:
        enabled: true
        position:
          top: 0
          left: 0
          height: 1
          width: 1
        refreshInterval: 2`

	new = `
  wtf:
    mods:
      clocks:
        enabled: true
        position:
          top: 0
          left: 0
          height: 1
          width: 1
        refreshInterval: 100ms`
)

func Test_RefreshInterval(t *testing.T) {
	t.Skip() // slow running test because a ticker is tested
	tests := []struct {
		name         string
		moduleName   string
		config       *config.Config
		testAttempts int
		expected     time.Duration
	}{
		{
			name:       "slow ticking module",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(configExample)
				return cfg
			}(),
			testAttempts: 10,
			expected:     2 * time.Second,
		},
		{
			name:       "fast ticking module",
			moduleName: "clocks",
			config: func() *config.Config {
				cfg, _ := config.ParseYaml(new)
				return cfg
			}(),
			testAttempts: 10,
			expected:     100 * time.Millisecond,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := MakeWidget(nil, nil, tt.moduleName, tt.config, make(chan bool))

			interval := widget.CommonSettings().RefreshInterval // same declaration as in scheduler.go#Schedule
			timer := time.NewTicker(interval)

			attempts := 0
			for {
				select {
				case <-timer.C:
					attempts++
					if attempts == tt.testAttempts {
						return
					}
				// allow for small window (50ms) where a timeout is not triggered
				case <-time.After(tt.expected + 50*time.Millisecond):
					t.Error("Timeout")
				}
			}

		})
	}
}
