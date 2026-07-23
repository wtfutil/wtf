package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

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
