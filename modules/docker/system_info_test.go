package docker

import (
	"testing"

	"github.com/moby/moby/api/types/system"
	"github.com/stretchr/testify/assert"
)

func Test_formatSystemInfo(t *testing.T) {
	tests := []struct {
		name           string
		info           system.Info
		containersSize int64
		imagesSize     int64
		volumesSize    int64
		volumesCount   int64
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
			containersSize: 100000000,
			imagesSize:     200000000,
			volumesSize:    50000000,
			volumesCount:   2,
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
			name:           "zero-value info with no disk usage",
			info:           system.Info{},
			containersSize: 0,
			imagesSize:     0,
			volumesSize:    0,
			volumesCount:   0,
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
			got := formatSystemInfo(tt.info, tt.containersSize, tt.imagesSize, tt.volumesSize, tt.volumesCount, tt.labelColor, tt.evenForeground)
			assert.Equal(t, tt.want, got)
		})
	}
}
