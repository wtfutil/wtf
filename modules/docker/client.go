package docker

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/dustin/go-humanize"
	"github.com/pkg/errors"
)

func (widget *Widget) getSystemInfo() string {
	info, err := widget.cli.Info(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not get docker system info").Error()
	}

	diskUsage, err := widget.cli.DiskUsage(context.Background(), types.DiskUsageOptions{})
	if err != nil {
		return errors.Wrap(err, "could not get disk usage").Error()
	}

	var duContainer int64
	for _, c := range diskUsage.Containers {
		duContainer += c.SizeRw
	}
	var duImg int64
	for _, im := range diskUsage.Images {
		duImg += im.Size
	}
	var duVol int64
	for _, v := range diskUsage.Volumes {
		duVol += v.UsageData.Size
	}

	sysInfo := []struct {
		name  string
		value string
	}{
		{
			name:  "name:",
			value: fmt.Sprintf("[%s]%s", widget.settings.Colors.RowTheme.EvenForeground, info.Name),
		}, {
			name:  "version:",
			value: fmt.Sprintf("[%s]%s", widget.settings.Colors.RowTheme.EvenForeground, info.ServerVersion),
		}, {
			name:  "root:",
			value: fmt.Sprintf("[%s]%s", widget.settings.Colors.RowTheme.EvenForeground, info.DockerRootDir),
		},
		{
			name: "containers:",
			value: fmt.Sprintf("[lime]%d[white]/[yellow]%d[white]/[red]%d",
				info.ContainersRunning,
				info.ContainersPaused, info.ContainersStopped),
		},
		{
			name:  "images:",
			value: fmt.Sprintf("[%s]%d", widget.settings.Colors.RowTheme.EvenForeground, info.Images),
		},
		{
			name:  "volumes:",
			value: fmt.Sprintf("[%s]%v", widget.settings.Colors.RowTheme.EvenForeground, len(diskUsage.Volumes)),
		},
		{
			name:  "memory limit:",
			value: fmt.Sprintf("[%s]%s", widget.settings.Colors.RowTheme.EvenForeground, humanize.Bytes(uint64(info.MemTotal))),
		},
		{
			name: "disk usage:",
			value: fmt.Sprintf(`
    [%s]* containers: [%s]%s
    [%s]* images:     [%s]%s
    [%s]* volumes:    [%s]%s
    [%s]* [::b]total:      [%s]%s[::-]
`,
				widget.settings.labelColor,
				widget.settings.Colors.RowTheme.EvenForeground,
				humanize.Bytes(uint64(duContainer)),

				widget.settings.labelColor,
				widget.settings.Colors.RowTheme.EvenForeground,
				humanize.Bytes(uint64(duImg)),

				widget.settings.labelColor,
				widget.settings.Colors.RowTheme.EvenForeground,
				humanize.Bytes(uint64(duVol)),

				widget.settings.labelColor,
				widget.settings.Colors.RowTheme.EvenForeground,
				humanize.Bytes(uint64(duContainer+duImg+duVol))),
		},
	}

	padSlice(true, sysInfo, func(i int) string {
		return sysInfo[i].name
	}, func(i int, newVal string) {
		sysInfo[i].name = newVal
	})

	result := ""
	for _, info := range sysInfo {
		result += fmt.Sprintf("[%s]%s %s\n", widget.settings.labelColor, info.name, info.value)
	}

	return result
}

func (widget *Widget) getContainerStates() string {
	cntrs, err := widget.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return errors.Wrapf(err, " could not get container list").Error()
	}

	if len(cntrs) == 0 {
		return " no containers"
	}

	colorMap := map[string]string{
		"created":    "green",
		"running":    "lime",
		"paused":     "yellow",
		"restarting": "yellow",
		"removing":   "yellow",
		"exited":     "red",
		"dead":       "red",
	}

	containers := []struct {
		name  string
		state string
	}{}
	for _, c := range cntrs {
		container := struct {
			name  string
			state string
		}{
			name:  c.Names[0],
			state: c.State,
		}

		container.name = strings.ReplaceAll(container.name, "/", "")
		containers = append(containers, container)
	}

	sort.Slice(containers, func(i, j int) bool {
		return containers[i].name < containers[j].name
	})

	padSlice(false, containers, func(i int) string {
		return containers[i].name
	}, func(i int, val string) {
		containers[i].name = val
	})

	result := ""
	for _, c := range containers {
		result += fmt.Sprintf("[white]%s [%s]%s\n", c.name, colorMap[c.state], c.state)
	}

	return result
}
