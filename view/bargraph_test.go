package view

import (
	"testing"

	"github.com/olebedev/config"
	"github.com/rivo/tview"
	"github.com/stretchr/testify/assert"
	"github.com/wtfutil/wtf/cfg"
)

// MakeData - Create sample data
func makeData() []Bar {
	//this could come from config
	const lineCount = 3
	var stats [lineCount]Bar

	stats[0] = Bar{
		Label:   "Jun 27, 2018",
		Percent: 20,
	}

	stats[1] = Bar{
		Label:      "Jul 09, 2018",
		Percent:    80,
		LabelColor: "red",
	}

	stats[2] = Bar{
		Label:      "Jul 09, 2018",
		Percent:    80,
		LabelColor: "green",
	}

	return stats[:]
}

func newTestGraph(graphStars int, graphIcon string) *BarGraph {
	widget := NewBarGraph(
		tview.NewApplication(),
		make(chan bool),
		"testapp",
		&cfg.Common{
			Config: &config.Config{
				Root: map[string]interface{}{
					"graphStars": graphStars,
					"graphIcon":  graphIcon,
				},
			},
		},
	)
	return &widget
}

func Test_NewBarGraph(t *testing.T) {
	widget := newTestGraph(15, "|")

	assert.NotNil(t, widget.View)
	assert.Equal(t, 15, widget.maxStars)
	assert.Equal(t, "|", widget.starChar)
}

func Test_BuildBars(t *testing.T) {
	widget := newTestGraph(15, "|")

	before := widget.View.GetText(false)
	widget.BuildBars(makeData())
	after := widget.View.GetText(false)

	assert.NotEqual(t, before, after)
}

func Test_TextView(t *testing.T) {
	widget := newTestGraph(15, "|")

	assert.NotNil(t, widget.TextView())
}

func Test_BuildStars(t *testing.T) {
	result := BuildStars(makeData(), 20, "*")
	assert.Equal(t,
		"Jun 27, 2018[[default]****[default]                ] 20\nJul 09, 2018[[red]****************[default]    ] 80\nJul 09, 2018[[green]****************[default]    ] 80\n",
		result,
	)
}
