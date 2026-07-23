package docker

import (
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/system"
	"github.com/docker/docker/api/types/volume"
	"github.com/stretchr/testify/assert"
)

func Test_formatSystemInfo(t *testing.T) {
	tests := []struct {
		name           string
		info           system.Info
		diskUsage      types.DiskUsage
		labelColor     string
		evenForeground string
		want           string
	}{
		{
			name: "fully populated system info",
			info: system.Info{
				Name:              "docker-host",
				ServerVersion:     "24.0.5",
				DockerRootDir:     "/var/lib/docker",
				ContainersRunning: 3,
				ContainersPaused:  1,
				ContainersStopped: 2,
				Images:            10,
				MemTotal:          2147483648,
			},
			diskUsage: types.DiskUsage{
				Containers: []*container.Summary{
					{SizeRw: 100000000},
				},
				Images: []*image.Summary{
					{Size: 200000000},
				},
				Volumes: []*volume.Volume{
					{UsageData: &volume.UsageData{Size: 50000000}},
					// A volume with nil UsageData must not panic and
					// contributes zero bytes to the disk usage total.
					{UsageData: nil},
				},
			},
			labelColor:     "white",
			evenForeground: "blue",
			want: "[white]        name: [blue]docker-host\n" +
				"[white]     version: [blue]24.0.5\n" +
				"[white]        root: [blue]/var/lib/docker\n" +
				"[white]  containers: [lime]3[white]/[yellow]1[white]/[red]2\n" +
				"[white]      images: [blue]10\n" +
				"[white]     volumes: [blue]2\n" +
				"[white]memory limit: [blue]2.1 GB\n" +
				"[white]  disk usage: \n" +
				"    [white]* containers: [blue]100 MB\n" +
				"    [white]* images:     [blue]200 MB\n" +
				"    [white]* volumes:    [blue]50 MB\n" +
				"    [white]* [::b]total:      [blue]350 MB[::-]\n" +
				"\n",
		},
		{
			name:           "zero-value info with no disk usage entries",
			info:           system.Info{},
			diskUsage:      types.DiskUsage{},
			labelColor:     "",
			evenForeground: "",
			want: "[]        name: []\n" +
				"[]     version: []\n" +
				"[]        root: []\n" +
				"[]  containers: [lime]0[white]/[yellow]0[white]/[red]0\n" +
				"[]      images: []0\n" +
				"[]     volumes: []0\n" +
				"[]memory limit: []0 B\n" +
				"[]  disk usage: \n" +
				"    []* containers: []0 B\n" +
				"    []* images:     []0 B\n" +
				"    []* volumes:    []0 B\n" +
				"    []* [::b]total:      []0 B[::-]\n" +
				"\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatSystemInfo(tt.info, tt.diskUsage, tt.labelColor, tt.evenForeground)
			assert.Equal(t, tt.want, got)
		})
	}
}
