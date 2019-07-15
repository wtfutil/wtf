package trello

import "fmt"

// CustomFieldItem represents the custom field items of Trello a trello card.
type CustomFieldItem struct {
	ID            string `json:"id"`
	IDValue       string `json:"idValue"`
	IDCustomField string `json:"idCustomField"`
	IDModel       string `json:"idModel"`
	IDModelType   string `json:"modelType,omitempty"`
}

// CustomField represents Trello's custom fields: "extra bits of structured data
// attached to cards when our users need a bit more than what Trello provides."
// https://developers.trello.com/reference/#custom-fields
type CustomField struct {
	ID          string `json:"id"`
	IDModel     string `json:"idModel"`
	IDModelType string `json:"modelType,omitempty"`
	FieldGroup  string `json:"fieldGroup"`
	Name        string `json:"name"`
	Pos         int    `json:"pos"`
	Display     struct {
		CardFront bool `json:"cardfront"`
	} `json:"display"`
	Type    string               `json:"type"`
	Options []*CustomFieldOption `json:"options"`
}

// CustomFieldOption are nested resources of CustomFields
type CustomFieldOption struct {
	ID            string `json:"id"`
	IDCustomField string `json:"idCustomField"`
	Value         struct {
		Text string `json:"text"`
	} `json:"value"`
	Color string `json:"color,omitempty"`
	Pos   int    `json:"pos"`
}

// GetCustomField takes a field id string and Arguments and returns the matching custom Field.
func (c *Client) GetCustomField(fieldID string, args Arguments) (customField *CustomField, err error) {
	path := fmt.Sprintf("customFields/%s", fieldID)
	err = c.Get(path, args, &customField)
	return
}

// GetCustomFields returns a slice of all receiver board's custom fields.
func (b *Board) GetCustomFields(args Arguments) (customFields []*CustomField, err error) {
	path := fmt.Sprintf("boards/%s/customFields", b.ID)
	err = b.client.Get(path, args, &customFields)
	return
}
