package flighty

import "fmt"

func (widget *Widget) display() {
	widget.Redraw(widget.content)
}

func (widget *Widget) content() (string, string, bool) {
	title := widget.CommonSettings().Title
	setWrap := false

	var content string

	for _, flight := range widget.Data {
		content += widget.flightToString(flight) + "\n"
	}

	return title, content, setWrap
}

func (widget *Widget) flightToString(flight *Flight) string {
	return fmt.Sprintf(
		"%s\t%s\t%s\t%s",
		flight.ICAO24,
		flight.CallSign,
		flight.EstDepartureAirport,
		flight.EstArrivalAirport,
	)
}
