package digitalocean

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/digitalocean/godo"
	"github.com/wtfutil/wtf/utils"
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
	utils.Reflective
}

// Region represents WTF's view of a DigitalOcean region
type Region struct {
	godo.Region
	utils.Reflective
}

// NewDroplet creates and returns an instance of Droplet
func NewDroplet(doDroplet godo.Droplet) *Droplet {
	droplet := &Droplet{
		doDroplet,

		Image{
			*doDroplet.Image,
			utils.Reflective{},
		},

		Region{
			*doDroplet.Region,
			utils.Reflective{},
		},
	}

	return droplet
}

/* -------------------- Exported Functions -------------------- */

// StringValueForProperty returns a string value for the given column
func (drop *Droplet) StringValueForProperty(propName string) (string, error) {
	var strVal string
	var err error

	// Figure out if we should forward this property to a sub-object
	// Lets us support "Region.Name" column definitions
	split := strings.Split(propName, ".")

	switch split[0] {
	case "Image":
		strVal, err = drop.Image.StringValueForProperty(split[1])
	case "Region":
		strVal, err = drop.Region.StringValueForProperty(split[1])
	default:
		v := reflect.ValueOf(drop)
		refVal := reflect.Indirect(v).FieldByName(propName)

		if !refVal.IsValid() {
			err = fmt.Errorf("invalid property name: %s", propName)
		} else {
			strVal = fmt.Sprintf("%v", refVal)
		}

	}

	return strVal, err
}
