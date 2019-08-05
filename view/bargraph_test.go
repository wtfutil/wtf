package view

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

// MakeData - Create sample data
func makeData() []Bar {

	//this could come from config
	const lineCount = 2
	var stats [lineCount]Bar

	stats[0] = Bar{
		Label:   "Jun 27, 2018",
		Percent: 20,
	}

	stats[1] = Bar{
		Label:   "Jul 09, 2018",
		Percent: 80,
	}

	return stats[:]

}

//TestOutput of the bargraph make string (BuildStars) function
func TestOutput(t *testing.T) {

	result := BuildStars(makeData(), 20, "*")

	Equal(t,
		"Jun 27, 2018[[red]****[white]                ] 20\nJul 09, 2018[[red]****************[white]    ] 80\n",
		result,
	)
}
