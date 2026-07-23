package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
	"github.com/wtfutil/wtf/view"
)

func Test_Widget_display(t *testing.T) {
	tests := []struct {
		name          string
		title         string
		displayBuffer string
	}{
		{
			name:          "returns the configured title and buffered content",
			title:         "docker",
			displayBuffer: "[red] System[white]\nsome info\n",
		},
		{
			name:          "empty buffer still returns title and ok",
			title:         "My Docker Widget",
			displayBuffer: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			common := &cfg.Common{Title: tt.title}
			widget := &Widget{
				TextWidget: view.TextWidget{
					Base: view.NewBase(nil, make(chan bool), nil, common),
				},
				settings: &Settings{
					Common: common,
				},
				displayBuffer: tt.displayBuffer,
			}

			title, content, wrap := widget.display()

			assert.Equal(t, tt.title, title)
			assert.Equal(t, tt.displayBuffer, content)
			assert.True(t, wrap)
		})
	}
}
