package docker

import (
	"testing"

	"github.com/moby/moby/api/types/container"
	"github.com/stretchr/testify/assert"
)

func Test_formatContainerStates(t *testing.T) {
	tests := []struct {
		name  string
		input []container.Summary
		want  string
	}{
		{
			name:  "no containers",
			input: []container.Summary{},
			want:  " no containers",
		},
		{
			name: "single running container",
			input: []container.Summary{
				{Names: []string{"/web"}, State: "running"},
			},
			want: "[white]web [lime]running\n",
		},
		{
			name: "strips leading slash from container name",
			input: []container.Summary{
				{Names: []string{"/my-app"}, State: "exited"},
			},
			want: "[white]my-app [red]exited\n",
		},
		{
			name: "sorts containers alphabetically by name",
			input: []container.Summary{
				{Names: []string{"/zeta"}, State: "running"},
				{Names: []string{"/alpha"}, State: "paused"},
			},
			want: "[white]alpha [yellow]paused\n[white]zeta  [lime]running\n",
		},
		{
			name: "pads names to equal width, left-aligned",
			input: []container.Summary{
				{Names: []string{"/a"}, State: "created"},
				{Names: []string{"/bb"}, State: "dead"},
				{Names: []string{"/ccc"}, State: "restarting"},
			},
			want: "[white]a   [green]created\n" +
				"[white]bb  [red]dead\n" +
				"[white]ccc [yellow]restarting\n",
		},
		{
			name: "unknown state has no color mapping",
			input: []container.Summary{
				{Names: []string{"/weird"}, State: "totally-unknown"},
			},
			want: "[white]weird []totally-unknown\n",
		},
		{
			name: "container with no names renders empty name",
			input: []container.Summary{
				{Names: []string{}, State: "running"},
			},
			want: "[white] [lime]running\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, formatContainerStates(tt.input))
		})
	}
}
