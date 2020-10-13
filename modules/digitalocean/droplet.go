package digitalocean

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/digitalocean/godo"
)

// Droplet represents WTF's view of a DigitalOcean droplet
type Droplet struct {
	godo.Droplet

	Image  Image
	Region Region
}

// Image represents WTF's view of a DigitalOcean droplet image
type Image struct {
	godo.Image
}

// Region represents WTF's view of a DigitalOcean region
type Region struct {
	godo.Region
}

// NewDroplet creates and returns an instance of Droplet
func NewDroplet(doDroplet godo.Droplet) *Droplet {
	droplet := &Droplet{
		doDroplet,

		Image{
			*doDroplet.Image,
		},

		Region{
			*doDroplet.Region,
		},
	}

	return droplet
}

// ValueForColumn returns a string value for the given column
func (drop *Droplet) ValueForColumn(colName string) string {
	r := reflect.ValueOf(drop)
	f := reflect.Indirect(r).FieldByName(colName)

	var strVal string

	// Figure out if we should forward this property to a sub-object
	// Lets us support "Region.Name" column definitions
	split := strings.Split(colName, ".")

	switch split[0] {
	case "Image":
		strVal = drop.Image.ValueForColumn(split[1])
	case "Region":
		strVal = drop.Region.ValueForColumn(split[1])
	default:
		strVal = fmt.Sprintf("%v", f)
	}

	return strVal
}

// ValueForColumn returns a string value for the given column
func (reg *Image) ValueForColumn(colName string) string {
	r := reflect.ValueOf(reg)
	f := reflect.Indirect(r).FieldByName(colName)

	strVal := fmt.Sprintf("%v", f)

	return strVal
}

// ValueForColumn returns a string value for the given column
func (reg *Region) ValueForColumn(colName string) string {
	r := reflect.ValueOf(reg)
	f := reflect.Indirect(r).FieldByName(colName)

	strVal := fmt.Sprintf("%v", f)

	return strVal
}
