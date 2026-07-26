package docker

import (
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/system"
)

// containerStateColors maps a container's state to the color used to display it.
var containerStateColors = map[string]string{
	"created":    "green",
	"running":    "lime",
	"paused":     "yellow",
	"restarting": "yellow",
	"removing":   "yellow",
	"exited":     "red",
	"dead":       "red",
}

func padSlice(padLeft bool, slice interface{}, getter func(i int) string, setter func(i int, newVal string)) {
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	maxLen := 0
	for i := 0; i < length; i++ {
		val := getter(i)
		maxLen = int(math.Max(float64(len(val)), float64(maxLen)))
	}

	sign := "-"
	if padLeft {
		sign = ""
	}

	for i := 0; i < length; i++ {
		val := getter(i)
		val = fmt.Sprintf("%"+sign+strconv.Itoa(maxLen)+"s", val)
		setter(i, val)
	}
}

// formatContainerStates renders the given containers as a sorted, aligned
// list of "<name> <state>" lines, colorized by state. It performs no I/O and
// can be exercised with manually constructed container.Summary values.
func formatContainerStates(cntrs []container.Summary) string {
	if len(cntrs) == 0 {
		return " no containers"
	}

	containers := []struct {
		name  string
		state string
	}{}
	for _, c := range cntrs {
		name := ""
		if len(c.Names) > 0 {
			name = c.Names[0]
		}

		containers = append(containers, struct {
			name  string
			state string
		}{
			name:  strings.ReplaceAll(name, "/", ""),
			state: string(c.State),
		})
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
		result += fmt.Sprintf("[white]%s [%s]%s\n", c.name, containerStateColors[c.state], c.state)
	}

	return result
}

// formatSystemInfo renders docker system/disk-usage information as a
// human-readable summary block. It performs no I/O and can be exercised with
// manually constructed system.Info values and disk-usage totals.
func formatSystemInfo(info system.Info, containersSize, imagesSize, volumesSize int64, volumesCount int64, labelColor, evenForeground string) string {
	sysInfo := []struct {
		name  string
		value string
	}{
		{
			name:  "name:",
			value: fmt.Sprintf("[%s]%s", evenForeground, info.Name),
		}, {
			name:  "version:",
			value: fmt.Sprintf("[%s]%s", evenForeground, info.ServerVersion),
		}, {
			name:  "root:",
			value: fmt.Sprintf("[%s]%s", evenForeground, info.DockerRootDir),
		},
		{
			name: "containers:",
			value: fmt.Sprintf("[lime]%d[white]/[yellow]%d[white]/[red]%d",
				info.ContainersRunning,
				info.ContainersPaused, info.ContainersStopped),
		},
		{
			name:  "images:",
			value: fmt.Sprintf("[%s]%d", evenForeground, info.Images),
		},
		{
			name:  "volumes:",
			value: fmt.Sprintf("[%s]%v", evenForeground, volumesCount),
		},
		{
			name:  "memory limit:",
			value: fmt.Sprintf("[%s]%s", evenForeground, humanize.Bytes(uint64(info.MemTotal))),
		},
		{
			name: "disk usage:",
			value: fmt.Sprintf(`
    [%s]* containers: [%s]%s
    [%s]* images:     [%s]%s
    [%s]* volumes:    [%s]%s
    [%s]* [::b]total:      [%s]%s[::-]
`,
				labelColor,
				evenForeground,
				humanize.Bytes(uint64(containersSize)),

				labelColor,
				evenForeground,
				humanize.Bytes(uint64(imagesSize)),

				labelColor,
				evenForeground,
				humanize.Bytes(uint64(volumesSize)),

				labelColor,
				evenForeground,
				humanize.Bytes(uint64(containersSize+imagesSize+volumesSize))),
		},
	}

	padSlice(true, sysInfo, func(i int) string {
		return sysInfo[i].name
	}, func(i int, newVal string) {
		sysInfo[i].name = newVal
	})

	result := ""
	for _, info := range sysInfo {
		result += fmt.Sprintf("[%s]%s %s\n", labelColor, info.name, info.value)
	}

	return result
}
