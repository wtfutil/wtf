package bamboohr

import (
	"fmt"
	"time"
)

// DateFormat defines the format we expect to receive dates from BambooHR in
const DateFormat = "2006-01-02"

type Item struct {
	Employee Employee `xml:"employee"`
	End      string   `xml:"end"`
	Holiday  string   `xml:"holiday"`
	Start    string   `xml:"start"`
	Type     string   `xml:"type,attr"`
}

func (item *Item) String() string {
	return fmt.Sprintf("Item: %s, %s, %s, %s", item.Type, item.Employee.Name, item.Start, item.End)
}

/* -------------------- Public Functions -------------------- */

func (item *Item) Name() string {
	if (item.Employee != Employee{}) {
		return item.Employee.Name
	}

	return item.Holiday
}

func (item *Item) PrettyStart() string {
	newTime, _ := time.Parse(DateFormat, item.Start)
	return fmt.Sprint(newTime.Format("Jan 2, 2006"))
}

func (item *Item) PrettyEnd() string {
	newTime, _ := time.Parse(DateFormat, item.End)
	end := fmt.Sprint(newTime.Format("Jan 2, 2006"))
	return end
}
