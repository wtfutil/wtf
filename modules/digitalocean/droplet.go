package digitalocean

import (
	"strings"

	"github.com/digitalocean/godo"
	"github.com/wtfutil/wtf/utils"
)

// Droplet represents WTF's view of a DigitalOcean droplet
type Droplet struct {
	godo.Droplet

	Image  godo.Image
	Region godo.Region
}

// NewDroplet creates and returns an instance of Droplet
func NewDroplet(doDroplet godo.Droplet) *Droplet {
	return &Droplet{
		doDroplet,
		*doDroplet.Image,
		*doDroplet.Region,
	}
}

/* -------------------- Exported Functions -------------------- */

// StringValueForProperty returns a string value for the given column
func (drop *Droplet) StringValueForProperty(propName string) (string, error) {
	// Figure out if we should forward this property to a sub-object
	// Lets us support "Region.Name" column definitions
	split := strings.Split(propName, ".")

	switch split[0] {
	case "Image":
		return utils.StringValueForProperty(drop.Image, split[1])
	case "Region":
		return utils.StringValueForProperty(drop.Region, split[1])
	default:
		return utils.StringValueForProperty(drop, propName)
	}
}
