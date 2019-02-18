package bamboohr

type Calendar struct {
	Items []Item `xml:"item"`
}

/* -------------------- Public Functions -------------------- */

func (calendar *Calendar) Holidays() []Item {
	return calendar.filteredItems("holiday")
}

func (calendar *Calendar) ItemsByType(itemType string) []Item {
	if itemType == "timeOff" {
		return calendar.TimeOffs()
	}

	return calendar.Holidays()
}

func (calendar *Calendar) TimeOffs() []Item {
	return calendar.filteredItems("timeOff")
}

/* -------------------- Private Functions -------------------- */

func (calendar *Calendar) filteredItems(itemType string) []Item {
	items := []Item{}

	for _, item := range calendar.Items {
		if item.Type == itemType {
			items = append(items, item)
		}
	}

	return items
}
