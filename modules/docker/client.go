package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
)

// dockerAPIClient is the subset of the docker client used by this widget.
// Defining it as an interface (rather than depending on the concrete
// *client.Client type) lets tests inject a fake implementation and exercise
// getSystemInfo/getContainerStates without a running docker daemon.
type dockerAPIClient interface {
	Info(ctx context.Context) (system.Info, error)
	DiskUsage(ctx context.Context, options types.DiskUsageOptions) (types.DiskUsage, error)
	ContainerList(ctx context.Context, options container.ListOptions) ([]container.Summary, error)
}

func (widget *Widget) getSystemInfo() string {
	info, err := widget.cli.Info(context.Background())
	if err != nil {
		return fmt.Errorf("could not get docker system info: %w", err).Error()
	}

	diskUsage, err := widget.cli.DiskUsage(context.Background(), types.DiskUsageOptions{})
	if err != nil {
		return fmt.Errorf("could not get disk usage: %w", err).Error()
	}

	return formatSystemInfo(info, diskUsage, widget.settings.labelColor, widget.settings.Colors.EvenForeground)
}

func (widget *Widget) getContainerStates() string {
	cntrs, err := widget.cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return fmt.Errorf("could not get container list: %w", err).Error()
	}

	return formatContainerStates(cntrs)
}
