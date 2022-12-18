package flighty

type FlightData struct {
	Aircraft string
}

func NewFlightData(aircraft string) (*FlightData, error) {
	fd := &FlightData{
		Aircraft: aircraft,
	}

	return fd, nil
}
