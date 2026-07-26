package docker

import (
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/system"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func newTestWidget(cli dockerAPIClient) *Widget {
	return &Widget{
		cli: cli,
		settings: &Settings{
			Common:     &cfg.Common{},
			labelColor: "white",
		},
	}
}

func Test_Widget_getSystemInfo(t *testing.T) {
	tests := []struct {
		name    string
		fake    *fakeDockerClient
		wantErr string
		wantOK  bool
	}{
		{
			name:    "returns wrapped error when Info fails",
			fake:    &fakeDockerClient{infoErr: errBoom},
			wantErr: "could not get docker system info: boom",
		},
		{
			name:    "returns wrapped error when DiskUsage fails",
			fake:    &fakeDockerClient{diskErr: errBoom},
			wantErr: "could not get disk usage: boom",
		},
		{
			name: "formats successful info and disk usage",
			fake: &fakeDockerClient{
				info: system.Info{Name: "host", ServerVersion: "1.0"},
			},
			wantOK: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := newTestWidget(tt.fake)

			got := widget.getSystemInfo()

			if tt.wantErr != "" {
				assert.Equal(t, tt.wantErr, got)
			}
			if tt.wantOK {
				assert.Contains(t, got, "host")
				assert.Contains(t, got, "1.0")
			}
		})
	}
}

func Test_Widget_getContainerStates(t *testing.T) {
	tests := []struct {
		name    string
		fake    *fakeDockerClient
		wantErr string
		want    string
	}{
		{
			name:    "returns wrapped error when ContainerList fails",
			fake:    &fakeDockerClient{listErr: errBoom},
			wantErr: "could not get container list: boom",
		},
		{
			name: "returns no containers message when list is empty",
			fake: &fakeDockerClient{cntrs: []container.Summary{}},
			want: " no containers",
		},
		{
			name: "formats returned containers",
			fake: &fakeDockerClient{
				cntrs: []container.Summary{
					{Names: []string{"/svc"}, State: "running"},
				},
			},
			want: "[white]svc [lime]running\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			widget := newTestWidget(tt.fake)

			got := widget.getContainerStates()

			if tt.wantErr != "" {
				assert.Equal(t, tt.wantErr, got)
			} else {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_Widget_refreshDisplayBuffer(t *testing.T) {
	t.Run("nil client leaves display buffer untouched", func(t *testing.T) {
		widget := newTestWidget(nil)
		widget.displayBuffer = "unchanged"

		widget.refreshDisplayBuffer()

		assert.Equal(t, "unchanged", widget.displayBuffer)
	})

	t.Run("populated client renders system and container sections", func(t *testing.T) {
		fake := &fakeDockerClient{
			info: system.Info{Name: "host"},
			cntrs: []container.Summary{
				{Names: []string{"/svc"}, State: "running"},
			},
		}
		widget := newTestWidget(fake)
		widget.settings.Colors.Subheading = "aqua"

		widget.refreshDisplayBuffer()

		assert.Contains(t, widget.displayBuffer, "[aqua] System[white]")
		assert.Contains(t, widget.displayBuffer, "host")
		assert.Contains(t, widget.displayBuffer, "[aqua] Containers[white]")
		assert.Contains(t, widget.displayBuffer, "svc")
	})

	t.Run("error from client surfaces in display buffer", func(t *testing.T) {
		fake := &fakeDockerClient{infoErr: errBoom, listErr: errBoom}
		widget := newTestWidget(fake)

		widget.refreshDisplayBuffer()

		assert.Contains(t, widget.displayBuffer, "could not get docker system info: boom")
		assert.Contains(t, widget.displayBuffer, "could not get container list: boom")
	})
}

func Test_Widget_Refresh(t *testing.T) {
	fake := &fakeDockerClient{
		info:  system.Info{Name: "host"},
		cntrs: []container.Summary{},
	}
	widget := newTestWidget(fake)
	widget.TextWidget = view.NewTextWidget(nil, make(chan bool, 1), nil, widget.settings.Common)

	widget.Refresh()

	assert.Contains(t, widget.displayBuffer, "host")
}
