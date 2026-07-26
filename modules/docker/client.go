package docker

import (
	"context"
	"fmt"

	"github.com/moby/moby/client"
)

// dockerAPIClient is the subset of the docker client used by this widget.
// Defining it as an interface (rather than depending on the concrete
// *client.Client type) lets tests inject a fake implementation and exercise
// getSystemInfo/getContainerStates without a running docker daemon.
type dockerAPIClient interface {
	Info(ctx context.Context, options client.InfoOptions) (client.SystemInfoResult, error)
	DiskUsage(ctx context.Context, options client.DiskUsageOptions) (client.DiskUsageResult, error)
	ContainerList(ctx context.Context, options client.ContainerListOptions) (client.ContainerListResult, error)
}

func (widget *Widget) getSystemInfo() string {
	info, err := widget.cli.Info(context.Background(), client.InfoOptions{})
	if err != nil {
		return fmt.Errorf("could not get docker system info: %w", err).Error()
	}

	diskUsage, err := widget.cli.DiskUsage(context.Background(), client.DiskUsageOptions{
		Containers: true,
		Images:     true,
		Volumes:    true,
	})
	if err != nil {
		return fmt.Errorf("could not get disk usage: %w", err).Error()
	}

	return formatSystemInfo(
		info.Info,
		diskUsage.Containers.TotalSize,
		diskUsage.Images.TotalSize,
		diskUsage.Volumes.TotalSize,
		diskUsage.Volumes.TotalCount,
		widget.settings.labelColor,
		widget.settings.Colors.EvenForeground,
	)
}

func (widget *Widget) getContainerStates() string {
	result, err := widget.cli.ContainerList(context.Background(), client.ContainerListOptions{All: true})
	if err != nil {
		return fmt.Errorf("could not get container list: %w", err).Error()
	}

	return formatContainerStates(result.Items)
}
