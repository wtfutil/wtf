package datadog

import (
	"fmt"
	"net/url"
)

// SyntheticsTest represents a synthetics test, either api or browser
type SyntheticsTest struct {
	PublicId      *string            `json:"public_id,omitempty"`
	MonitorId     *int               `json:"monitor_id,omitempty"`
	Name          *string            `json:"name,omitempty"`
	Type          *string            `json:"type,omitempty"`
	Tags          []string           `json:"tags"`
	CreatedAt     *string            `json:"created_at,omitempty"`
	ModifiedAt    *string            `json:"modified_at,omitempty"`
	DeletedAt     *string            `json:"deleted_at,omitempty"`
	Config        *SyntheticsConfig  `json:"config,omitempty"`
	Message       *string            `json:"message,omitempty"`
	Options       *SyntheticsOptions `json:"options,omitempty"`
	Locations     []string           `json:"locations,omitempty"`
	CreatedBy     *SyntheticsUser    `json:"created_by,omitempty"`
	ModifiedBy    *SyntheticsUser    `json:"modified_by,omitempty"`
	Status        *string            `json:"status,omitempty"`
	MonitorStatus *string            `json:"monitor_status,omitempty"`
}

type SyntheticsConfig struct {
	Request    *SyntheticsRequest    `json:"request,omitempty"`
	Assertions []SyntheticsAssertion `json:"assertions,omitempty"`
	Variables  []interface{}         `json:"variables,omitempty"`
}

type SyntheticsRequest struct {
	Url     *string           `json:"url,omitempty"`
	Method  *string           `json:"method,omitempty"`
	Timeout *int              `json:"timeout,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    *string           `json:"body,omitempty"`
}

type SyntheticsAssertion struct {
	Operator *string `json:"operator,omitempty"`
	Property *string `json:"property,omitempty"`
	Type     *string `json:"type,omitempty"`
	// sometimes target is string ( like "text/html; charset=UTF-8" for header content-type )
	// and sometimes target is int ( like 1200 for responseTime, 200 for statusCode )
	Target interface{} `json:"target,omitempty"`
}

type SyntheticsOptions struct {
	TickEvery          *int     `json:"tick_every,omitempty"`
	FollowRedirects    *bool    `json:"follow_redirects,omitempty"`
	MinFailureDuration *int     `json:"min_failure_duration,omitempty"`
	MinLocationFailed  *int     `json:"min_location_failed,omitempty"`
	DeviceIds          []string `json:"device_ids,omitempty"`
}

type SyntheticsUser struct {
	Id     *int    `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Email  *string `json:"email,omitempty"`
	Handle *string `json:"handle,omitempty"`
}

type SyntheticsDevice struct {
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Height      *int    `json:"height,omitempty"`
	Width       *int    `json:"width,omitempty"`
	IsLandscape *bool   `json:"isLandscape,omitempty"`
	IsMobile    *bool   `json:"isMobile,omitempty"`
	UserAgent   *string `json:"userAgent,omitempty"`
}

type SyntheticsLocation struct {
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Region      *string `json:"region,omitempty"`
	IsLandscape *bool   `json:"is_active,omitempty"`
}

type ToggleStatus struct {
	NewStatus *string `json:"new_status,omitempty"`
}

// GetSyntheticsTests get all tests of type API
func (client *Client) GetSyntheticsTests() ([]SyntheticsTest, error) {
	var out struct {
		SyntheticsTests []SyntheticsTest `json:"tests,omitempty"`
	}
	if err := client.doJsonRequest("GET", "/v1/synthetics/tests", nil, &out); err != nil {
		return nil, err
	}
	return out.SyntheticsTests, nil
}

// GetSyntheticsTestsByType get all tests by type (e.g. api or browser)
func (client *Client) GetSyntheticsTestsByType(testType string) ([]SyntheticsTest, error) {
	var out struct {
		SyntheticsTests []SyntheticsTest `json:"tests,omitempty"`
	}
	query, err := url.ParseQuery(fmt.Sprintf("type=%v", testType))
	if err != nil {
		return nil, err
	}
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/synthetics/tests?%v", query.Encode()), nil, &out); err != nil {
		return nil, err
	}
	return out.SyntheticsTests, nil
}

// GetSyntheticsTest get test by public id
func (client *Client) GetSyntheticsTest(publicId string) (*SyntheticsTest, error) {
	var out SyntheticsTest
	if err := client.doJsonRequest("GET", "/v1/synthetics/tests/"+publicId, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// CreateSyntheticsTest creates a test
func (client *Client) CreateSyntheticsTest(syntheticsTest *SyntheticsTest) (*SyntheticsTest, error) {
	var out SyntheticsTest
	if err := client.doJsonRequest("POST", "/v1/synthetics/tests", syntheticsTest, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UpdateSyntheticsTest updates a test
func (client *Client) UpdateSyntheticsTest(publicId string, syntheticsTest *SyntheticsTest) (*SyntheticsTest, error) {
	var out SyntheticsTest
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/synthetics/tests/%s", publicId), syntheticsTest, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// PauseSyntheticsTest set a test status to live
func (client *Client) PauseSyntheticsTest(publicId string) (*bool, error) {
	payload := ToggleStatus{NewStatus: String("paused")}
	out := Bool(false)
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/synthetics/tests/%s/status", publicId), &payload, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeSyntheticsTest set a test status to live
func (client *Client) ResumeSyntheticsTest(publicId string) (*bool, error) {
	payload := ToggleStatus{NewStatus: String("live")}
	out := Bool(false)
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v1/synthetics/tests/%s/status", publicId), &payload, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// string array of public_id
type DeleteSyntheticsTestsPayload struct {
	PublicIds []string `json:"public_ids,omitempty"`
}

// DeleteSyntheticsTests deletes tests
func (client *Client) DeleteSyntheticsTests(publicIds []string) error {
	req := DeleteSyntheticsTestsPayload{
		PublicIds: publicIds,
	}
	if err := client.doJsonRequest("POST", "/v1/synthetics/tests/delete", req, nil); err != nil {
		return err
	}
	return nil
}

// GetSyntheticsLocations get all test locations
func (client *Client) GetSyntheticsLocations() ([]SyntheticsLocation, error) {
	var out struct {
		Locations []SyntheticsLocation `json:"locations,omitempty"`
	}
	if err := client.doJsonRequest("GET", "/v1/synthetics/locations", nil, &out); err != nil {
		return nil, err
	}
	return out.Locations, nil
}

// GetSyntheticsBrowserDevices get all test devices (for browser)
func (client *Client) GetSyntheticsBrowserDevices() ([]SyntheticsDevice, error) {
	var out struct {
		Devices []SyntheticsDevice `json:"devices,omitempty"`
	}

	if err := client.doJsonRequest("GET", "/v1/synthetics/browser/devices", nil, &out); err != nil {
		return nil, err
	}
	return out.Devices, nil
}
