package arpansagovau

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Stations struct {
	XMLName  xml.Name `xml:"stations"`
	Text     string   `xml:",chardata"`
	Location []struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id,attr"`
		Name        string `xml:"name"`        // adl, ali, bri, can, cas, ...
		Index       float32 `xml:"index"`      // 0.0, 0.0, 0.0, 0.0, 0.0, ...
		Time        string `xml:"time"`        // 7:24 PM, 7:24 PM, 7:54 PM...
		Date        string `xml:"date"`        // 29/08/2019, 29/08/2019, 2...
		Fulldate    string `xml:"fulldate"`    // Thursday, 29 August 2019,...
		Utcdatetime string `xml:"utcdatetime"` // 2019/08/29 09:54, 2019/08...
		Status      string `xml:"status"`      // ok, ok, ok, ok, ok, ok, o...
	} `xml:"location"`
}


type location struct {
	name string
	index float32
	time string
	date string
	status string
}

func GetLocationData(cityname string) (*location, error) {
	var locdata location;
	resp, err := apiRequest()
	if err != nil {
		return nil, err
	}

	stations, err := parseXML(resp.Body);
	if(err != nil) {
		return nil, err
	}

	for _, city := range stations.Location {
		if(city.ID == cityname) {
			locdata = location { name: city.ID, index: city.Index, time: city.Time, date: city.Date, status: city.Status }
			break;
		}
	}
	return &locdata, err
}

/* -------------------- Unexported Functions -------------------- */

func apiRequest() (*http.Response, error) {
	req, err := http.NewRequest("GET", "https://uvdata.arpansa.gov.au/xml/uvvalues.xml", nil)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}

	return resp, nil
}
func parseXML(text io.Reader) (Stations, error) {
	dec := xml.NewDecoder(text)
	dec.Strict = false

	var v Stations;
	err := dec.Decode(&v)
	return v, err;
}
