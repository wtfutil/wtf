package bamboohr

import (
	"fmt"
	//"time"

	"github.com/senorprogrammer/wtf/wtf"
)

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

/* -------------------- Exported Functions -------------------- */

func (item *Item) IsOneDay() bool {
	return item.Start == item.End
}

func (item *Item) Name() string {
	if (item.Employee != Employee{}) {
		return item.Employee.Name
	}

	return item.Holiday
}

func (item *Item) PrettyStart() string {
	return wtf.PrettyDate(item.Start)
}

func (item *Item) PrettyEnd() string {
	return wtf.PrettyDate(item.End)
}
