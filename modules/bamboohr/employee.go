package bamboohr

/*
* Note: this currently implements the minimum number of fields to fulfill the Away functionality.
* Undoubtedly there are more fields than this to an employee
 */
type Employee struct {
	ID   int    `xml:"id,attr"`
	Name string `xml:",chardata"`
}
