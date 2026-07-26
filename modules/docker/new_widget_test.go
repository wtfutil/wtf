package docker

import (
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
	"github.com/olebedev/config"
	"github.com/stretchr/testify/assert"
)

func Test_NewWidget(t *testing.T) {
	moduleConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)
	globalConfig, err := config.ParseYaml("{}")
	assert.NoError(t, err)
	settings := NewSettingsFromYAML("docker", moduleConfig, globalConfig)

	t.Run("real docker client constructor does not panic", func(t *testing.T) {
		// Exercises the actual client.NewClientWithOpts(client.FromEnv) call.
		// This only builds a client object from local environment variables;
		// it does not require a running docker daemon.
		cli, err := newDockerClient()

		if err != nil {
			assert.Nil(t, cli)
		} else {
			assert.NotNil(t, cli)
		}
	})

	t.Run("client construction failure surfaces in display buffer", func(t *testing.T) {
		original := newDockerClient
		newDockerClient = func() (dockerAPIClient, error) {
			return nil, errBoom
		}
		defer func() { newDockerClient = original }()

		widget := NewWidget(nil, make(chan bool, 1), nil, settings)

		assert.Nil(t, widget.cli)
		assert.Contains(t, widget.displayBuffer, "could not create client: boom")
	})

	t.Run("client construction success populates display buffer", func(t *testing.T) {
		original := newDockerClient
		newDockerClient = func() (dockerAPIClient, error) {
			return &fakeDockerClient{
				info:  system.Info{Name: "host"},
				cntrs: []container.Summary{},
			}, nil
		}
		defer func() { newDockerClient = original }()

		widget := NewWidget(nil, make(chan bool, 1), nil, settings)

		assert.NotNil(t, widget.cli)
		assert.Contains(t, widget.displayBuffer, "host")
	})
}
