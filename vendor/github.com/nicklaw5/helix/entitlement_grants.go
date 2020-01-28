package helix

type entitlementUploadURLRequest struct {
	ManifestID string `query:"manifest_id"`
	Type       string `query:"type"`
}

// EntitlementsUploadURL ...
type EntitlementsUploadURL struct {
	URL string `json:"url"`
}

// ManyEntitlementsUploadURLs ...
type ManyEntitlementsUploadURLs struct {
	URLs []EntitlementsUploadURL `json:"data"`
}

// EntitlementsUploadResponse ...
type EntitlementsUploadResponse struct {
	ResponseCommon
	Data ManyEntitlementsUploadURLs
}

// CreateEntitlementsUploadURL return a URL where you can upload a manifest
// file and notify users that they have an entitlement. Entitlements are digital
// items that users are entitled to use. Twitch entitlements are granted to users
// gratis or as part of a purchase on Twitch.
func (c *Client) CreateEntitlementsUploadURL(manifestID, entitlementType string) (*EntitlementsUploadResponse, error) {
	data := &entitlementUploadURLRequest{
		ManifestID: manifestID,
		Type:       entitlementType,
	}

	resp, err := c.post("/entitlements/upload", &ManyEntitlementsUploadURLs{}, data)
	if err != nil {
		return nil, err
	}

	url := &EntitlementsUploadResponse{}
	url.StatusCode = resp.StatusCode
	url.Header = resp.Header
	url.Error = resp.Error
	url.ErrorStatus = resp.ErrorStatus
	url.ErrorMessage = resp.ErrorMessage
	url.Data.URLs = resp.Data.(*ManyEntitlementsUploadURLs).URLs

	return url, nil
}
