package bargraphtests

import (
	"testing"

	. "github.com/senorprogrammer/wtf/wtf"
	. "github.com/stretchr/testify/assert"
)

// MakeData - Create sample data
func makeData() [][2]int64 {

	//this could come from config
	const lineCount = 2
	var stats [lineCount][2]int64

	stats[0][1] = 1530122942
	stats[0][0] = 100

	stats[1][1] = 1530132942
	stats[1][0] = 210

	return stats[:]

}

//TestOutput of the bargraph make string (BuildStars) function
func TestOutput(t *testing.T) {

	result := BuildStars(makeData(), 20, "*")

	Equal(t, result, "Jan 18, 1970 -\t [red]*[white] - (100)\nJan 18, 1970 -\t [red]********************[white] - (210)\n")
}
