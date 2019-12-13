package digitalocean

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/digitalocean/godo"
	"github.com/wtfutil/wtf/utils"
	"github.com/wtfutil/wtf/view"
)

type dropletPropertiesTable struct {
	droplet     *godo.Droplet
	propertyMap map[string]string

	colWidth0   int
	colWidth1   int
	tableHeight int
}

// newDropletPropertiesTable creates and returns an instance of DropletPropertiesTable
func newDropletPropertiesTable(droplet *godo.Droplet) *dropletPropertiesTable {
	propTable := &dropletPropertiesTable{
		droplet: droplet,

		colWidth0:   24,
		colWidth1:   47,
		tableHeight: 16,
	}

	propTable.propertyMap = propTable.buildPropertyMap()

	return propTable
}

/* -------------------- Unexported Functions -------------------- */

// buildPropertyMap creates a mapping of droplet property names to droplet property values
func (propTable *dropletPropertiesTable) buildPropertyMap() map[string]string {
	propMap := map[string]string{}

	if propTable.droplet == nil {
		return propMap
	}

	publicV4, _ := propTable.droplet.PublicIPv4()
	publicV6, _ := propTable.droplet.PublicIPv6()

	propMap["CPUs"] = strconv.Itoa(propTable.droplet.Vcpus)
	propMap["Created"] = propTable.droplet.Created
	propMap["Disk"] = strconv.Itoa(propTable.droplet.Disk)
	propMap["Features"] = utils.Truncate(strings.Join(propTable.droplet.Features, ","), propTable.colWidth1, true)
	propMap["Image"] = fmt.Sprintf("%s (%s)", propTable.droplet.Image.Name, propTable.droplet.Image.Distribution)
	propMap["Memory"] = strconv.Itoa(propTable.droplet.Memory)
	propMap["Public IP v4"] = publicV4
	propMap["Public IP v6"] = publicV6
	propMap["Region"] = fmt.Sprintf("%s (%s)", propTable.droplet.Region.Name, propTable.droplet.Region.Slug)
	propMap["Size"] = propTable.droplet.SizeSlug
	propMap["Status"] = propTable.droplet.Status
	propMap["Tags"] = utils.Truncate(strings.Join(propTable.droplet.Tags, ","), propTable.colWidth1, true)
	propMap["URN"] = utils.Truncate(propTable.droplet.URN(), propTable.colWidth1, true)
	propMap["VPC"] = propTable.droplet.VPCUUID

	return propMap
}

// render creates a new Table and returns it as a displayable string
func (propTable *dropletPropertiesTable) render() string {
	tbl := view.NewInfoTable(
		[]string{"Property", "Value"},
		propTable.propertyMap,
		propTable.colWidth0,
		propTable.colWidth1,
		propTable.tableHeight,
	)

	return tbl.Render()
}
