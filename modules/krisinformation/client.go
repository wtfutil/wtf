package krisinformation

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/wtfutil/wtf/logger"
	"github.com/wtfutil/wtf/utils"
)

const (
	krisinformationAPI = "https://api.krisinformation.se/v2/feed?format=json"
)

type Krisinformation []struct {
	Identifier  string    `json:"Identifier"`
	PushMessage string    `json:"PushMessage"`
	Updated     time.Time `json:"Updated"`
	Published   time.Time `json:"Published"`
	Headline    string    `json:"Headline"`
	Preamble    string    `json:"Preamble"`
	BodyText    string    `json:"BodyText"`
	Area        []struct {
		Type                string      `json:"Type"`
		Description         string      `json:"Description"`
		Coordinate          string      `json:"Coordinate"`
		GeometryInformation interface{} `json:"GeometryInformation"`
	} `json:"Area"`
	Web        string        `json:"Web"`
	Language   string        `json:"Language"`
	Event      string        `json:"Event"`
	SenderName string        `json:"SenderName"`
	Push       bool          `json:"Push"`
	BodyLinks  []interface{} `json:"BodyLinks"`
	SourceID   int           `json:"SourceID"`
	IsVma      bool          `json:"IsVma"`
	IsTestVma  bool          `json:"IsTestVma"`
}

// Client holds or configuration
type Client struct {
	latitude  float64
	longitude float64
	radius    int
	county    string
	country   bool
}

// Item holds the interesting parts
type Item struct {
	PushMessage string
	HeadLine    string
	SenderName  string
	Country     bool
	County      bool
	Distance    float64
	Updated     time.Time
}

// NewClient returns a new Client
func NewClient(latitude, longitude float64, radius int, county string, country bool) *Client {
	return &Client{
		latitude:  latitude,
		longitude: longitude,
		radius:    radius,
		county:    county,
		country:   country,
	}

}

// getKrisinformation - return items that match either country, county or a radius
// Priority:
//   - Country
//   - County
//   - Region
func (c *Client) getKrisinformation() (items []Item, err error) {
	resp, err := http.Get(krisinformationAPI)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	var data Krisinformation
	err = utils.ParseJSON(&data, resp.Body)
	if err != nil {
		return nil, err
	}

	for i := range data {
		for a := range data[i].Area {
			// Country wide events
			if c.country && data[i].Area[a].Type == "Country" {
				item := Item{
					PushMessage: data[i].PushMessage,
					HeadLine:    data[i].Headline,
					SenderName:  data[i].SenderName,
					Country:     true,
					Updated:     data[i].Updated,
				}
				items = append(items, item)
				continue
			}

			// County specific events
			if c.county != "" && data[i].Area[a].Type == "County" {
				// We look for county in description
				if strings.Contains(
					strings.ToLower(data[i].Area[a].Description),
					strings.ToLower(c.county),
				) {
					item := Item{
						PushMessage: data[i].PushMessage,
						HeadLine:    data[i].Headline,
						SenderName:  data[i].SenderName,
						County:      true,
						Updated:     data[i].Updated,
					}
					items = append(items, item)
					continue
				}
			}

			if c.radius != -1 {
				coords := data[i].Area[a].Coordinate
				if coords == "" {
					continue
				}
				buf := strings.Split(coords, " ")
				latlon := strings.Split(buf[0], ",")
				kris_latitude, err := strconv.ParseFloat(latlon[0], 32)
				if err != nil {
					return nil, err
				}

				kris_longitude, err := strconv.ParseFloat(latlon[1], 32)
				if err != nil {
					return nil, err
				}

				distance := DistanceInMeters(kris_latitude, kris_longitude, c.latitude, c.longitude)
				logger.Log(fmt.Sprintf("Distance: %f", distance/1000)) // KM
				if distance < float64(c.radius) {
					item := Item{
						PushMessage: data[i].PushMessage,
						HeadLine:    data[i].Headline,
						SenderName:  data[i].SenderName,
						Distance:    distance,
						Updated:     data[i].Updated,
					}
					items = append(items, item)
				}

			}
		}
	}

	return items, nil
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//
//	a given longitude and latitude relatively accurately (using a spherical
//	approximation of the Earth) through the Haversin Distance Formula for
//	great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// http://en.wikipedia.org/wiki/Haversine_formula
func DistanceInMeters(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
