package digitalocean

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/digitalocean/godo"
	"github.com/olekukonko/tablewriter"
	"github.com/wtfutil/wtf/utils"
)

const (
	infoTableBodyHeight = 16
	infoTableColWidth0  = 24
	infoTableColWidth1  = 47
)

type infoTable struct {
	droplet     *godo.Droplet
	propertyMap map[string]string
}

func newInfoTable(droplet *godo.Droplet) *infoTable {
	infoTbl := &infoTable{
		droplet: droplet,
	}

	infoTbl.propertyMap = infoTbl.buildPropertyMap()

	return infoTbl
}

/* -------------------- Unexported Functions -------------------- */

// buildPropertyMap creates a mapping of droplet property names to droplet property values
func (infoTbl *infoTable) buildPropertyMap() map[string]string {
	propMap := map[string]string{}

	if infoTbl.droplet == nil {
		return propMap
	}

	publicV4, _ := infoTbl.droplet.PublicIPv4()
	publicV6, _ := infoTbl.droplet.PublicIPv6()

	propMap["CPUs"] = strconv.Itoa(infoTbl.droplet.Vcpus)
	propMap["Created"] = infoTbl.droplet.Created
	propMap["Disk"] = strconv.Itoa(infoTbl.droplet.Disk)
	propMap["Image"] = fmt.Sprintf("%s (%s)", infoTbl.droplet.Image.Name, infoTbl.droplet.Image.Distribution)
	propMap["Memory"] = strconv.Itoa(infoTbl.droplet.Memory)
	propMap["Public IP v4"] = publicV4
	propMap["Public IP v6"] = publicV6
	propMap["Region"] = fmt.Sprintf("%s (%s)", infoTbl.droplet.Region.Name, infoTbl.droplet.Region.Slug)
	propMap["Size"] = infoTbl.droplet.SizeSlug
	propMap["Status"] = infoTbl.droplet.Status
	propMap["Tags"] = utils.Truncate(strings.Join(infoTbl.droplet.Tags, ","), infoTableColWidth1, true)
	propMap["URN"] = utils.Truncate(infoTbl.droplet.URN(), infoTableColWidth1, true)
	propMap["VPC"] = infoTbl.droplet.VPCUUID

	return propMap
}

// render creates a new Table and returns it as a displayable string
func (infoTbl *infoTable) render() string {
	if infoTbl.droplet == nil {
		return "no droplet selected"
	}

	buf := new(bytes.Buffer)
	tbl := tablewriter.NewWriter(buf)

	tbl.SetHeader([]string{"Property", "Value"})
	tbl.SetBorder(true)
	tbl.SetCenterSeparator(" ")
	tbl.SetColumnSeparator(" ")
	tbl.SetRowSeparator("-")
	tbl.SetAlignment(tablewriter.ALIGN_LEFT)
	tbl.SetColMinWidth(0, infoTableColWidth0)
	tbl.SetColMinWidth(1, infoTableColWidth1)

	keys := []string{}
	for key := range infoTbl.propertyMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// Enumerate over the alphabetically-sorted keys to render the property values
	for _, key := range keys {
		tbl.Append([]string{key, infoTbl.propertyMap[key]})
	}

	// Pad the table with extra rows to push it to the bottom
	paddingAmt := infoTableBodyHeight - len(infoTbl.propertyMap) - 1
	if paddingAmt > 0 {
		for i := 0; i < paddingAmt; i++ {
			tbl.Append([]string{"", ""})
		}
	}

	tbl.Render()

	return buf.String()
}
