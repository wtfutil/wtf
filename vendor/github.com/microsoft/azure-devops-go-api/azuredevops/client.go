// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azuredevops

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"runtime"
	"strings"
	"sync"
)

const (
	// header keys
	headerKeyAccept              = "Accept"
	headerKeyAuthorization       = "Authorization"
	headerKeyContentType         = "Content-Type"
	HeaderKeyContinuationToken   = "X-MS-ContinuationToken"
	headerKeyFedAuthRedirect     = "X-TFS-FedAuthRedirect"
	headerKeyForceMsaPassThrough = "X-VSS-ForceMsaPassThrough"
	headerKeySession             = "X-TFS-Session"
	headerUserAgent              = "User-Agent"

	// media types
	MediaTypeTextPlain       = "text/plain"
	MediaTypeApplicationJson = "application/json"
)

// Unique session id to be used by all requests of this session.
var SessionId = uuid.New().String()

// ApiResourceLocation Cache by Url
var apiResourceLocationCache = make(map[string]*map[uuid.UUID]ApiResourceLocation)
var apiResourceLocationCacheLock = sync.RWMutex{}

// Base user agent string.  The UserAgent set on the connection will be appended to this.
var baseUserAgent = "go/" + runtime.Version() + " (" + runtime.GOOS + " " + runtime.GOARCH + ") azure-devops-go-api/0.0.0" // todo: get real version

func NewClient(connection *Connection, baseUrl string) *Client {
	client := &http.Client{}
	if connection.Timeout != nil {
		client.Timeout = *connection.Timeout
	}
	return &Client{
		baseUrl:                 baseUrl,
		client:                  client,
		authorization:           connection.AuthorizationString,
		suppressFedAuthRedirect: connection.SuppressFedAuthRedirect,
		forceMsaPassThrough:     connection.ForceMsaPassThrough,
		userAgent:               connection.UserAgent,
	}
}

type Client struct {
	baseUrl                 string
	client                  *http.Client
	authorization           string
	suppressFedAuthRedirect bool
	forceMsaPassThrough     bool
	userAgent               string
}

func (client *Client) SendRequest(request *http.Request) (response *http.Response, err error) {
	resp, err := client.client.Do(request) // todo: add retry logic
	if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
		err = client.UnwrapError(resp)
	}
	return resp, err
}

func (client *Client) Send(ctx context.Context,
	httpMethod string,
	locationId uuid.UUID,
	apiVersion string,
	routeValues map[string]string,
	queryParameters url.Values,
	body io.Reader,
	mediaType string,
	acceptMediaType string,
	additionalHeaders map[string]string) (response *http.Response, err error) {
	location, err := client.getResourceLocation(ctx, locationId)
	if err != nil {
		return nil, err
	}
	generatedUrl := client.GenerateUrl(location, routeValues, queryParameters)
	fullUrl := combineUrl(client.baseUrl, generatedUrl)
	negotiatedVersion, err := negotiateRequestVersion(location, apiVersion)
	if err != nil {
		return nil, err
	}

	req, err := client.CreateRequestMessage(ctx, httpMethod, fullUrl, negotiatedVersion, body, mediaType, acceptMediaType, additionalHeaders)
	if err != nil {
		return nil, err
	}

	resp, err := client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	// Set session if one was supplied in the response.
	session, ok := resp.Header[headerKeySession]
	if ok && len(session) > 0 {
		SessionId = session[0]
	}

	return resp, err
}

func (client *Client) GenerateUrl(apiResourceLocation *ApiResourceLocation, routeValues map[string]string, queryParameters url.Values) (request string) {
	builtUrl := *apiResourceLocation.RouteTemplate
	if routeValues == nil {
		routeValues = make(map[string]string)
	}
	routeValues["area"] = *apiResourceLocation.Area
	routeValues["resource"] = *apiResourceLocation.ResourceName
	builtUrl = transformRouteTemplate(builtUrl, routeValues)
	if queryParameters != nil && len(queryParameters) > 0 {
		builtUrl += "?" + queryParameters.Encode()
	}
	return builtUrl
}

func (client *Client) CreateRequestMessage(ctx context.Context,
	httpMethod string,
	url string,
	apiVersion string,
	body io.Reader,
	mediaType string,
	acceptMediaType string,
	additionalHeaders map[string]string) (request *http.Request, err error) {
	req, err := http.NewRequest(httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if client.authorization != "" {
		req.Header.Add(headerKeyAuthorization, client.authorization)
	}
	accept := acceptMediaType
	if apiVersion != "" {
		accept += ";api-version=" + apiVersion
	}
	req.Header.Add(headerKeyAccept, accept)
	if mediaType != "" {
		req.Header.Add(headerKeyContentType, mediaType+";charset=utf-8")
	}
	if client.suppressFedAuthRedirect {
		req.Header.Add(headerKeyFedAuthRedirect, "Suppress")
	}
	if client.forceMsaPassThrough {
		req.Header.Add(headerKeyForceMsaPassThrough, "true")
	}

	// set session if it has not already been set
	_, ok := req.Header[headerKeySession]
	if !ok {
		req.Header.Add(headerKeySession, SessionId)
	}

	userAgent := baseUserAgent
	if client.userAgent != "" {
		userAgent += " " + client.userAgent
	}
	req.Header.Add(headerUserAgent, userAgent)

	for key, value := range additionalHeaders {
		req.Header.Add(key, value)
	}

	return req, err
}

func (client *Client) getResourceLocation(ctx context.Context, locationId uuid.UUID) (*ApiResourceLocation, error) {
	locationsMap, ok := getApiResourceLocationCache(client.baseUrl)
	if !ok {
		locations, err := client.getResourceLocationsFromServer(ctx)
		if err != nil {
			return nil, err
		}
		newMap := make(map[uuid.UUID]ApiResourceLocation)
		locationsMap = &newMap
		for _, locationEntry := range locations {
			(*locationsMap)[*locationEntry.Id] = locationEntry
		}

		setApiResourceLocationCache(client.baseUrl, locationsMap)
	}

	location, ok := (*locationsMap)[locationId]
	if ok {
		return &location, nil
	}

	return nil, &LocationIdNotRegisteredError{locationId, client.baseUrl}
}

func getApiResourceLocationCache(url string) (*map[uuid.UUID]ApiResourceLocation, bool) {
	apiResourceLocationCacheLock.RLock()
	defer apiResourceLocationCacheLock.RUnlock()
	locationsMap, ok := apiResourceLocationCache[url]
	return locationsMap, ok
}

func setApiResourceLocationCache(url string, locationsMap *map[uuid.UUID]ApiResourceLocation) {
	apiResourceLocationCacheLock.Lock()
	defer apiResourceLocationCacheLock.Unlock()
	apiResourceLocationCache[url] = locationsMap
}

func (client *Client) getResourceLocationsFromServer(ctx context.Context) ([]ApiResourceLocation, error) {
	optionsUri := combineUrl(client.baseUrl, "_apis")
	request, err := client.CreateRequestMessage(ctx, http.MethodOptions, optionsUri, "", nil, "", MediaTypeApplicationJson, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.SendRequest(request)
	if err != nil {
		return nil, err
	}

	// Set session if one was supplied in the response.
	session, ok := resp.Header[headerKeySession]
	if ok && len(session) > 0 {
		SessionId = session[0]
	}

	if resp != nil && (resp.StatusCode < 200 || resp.StatusCode >= 300) {
		return nil, client.UnwrapError(resp)
	}

	var locations []ApiResourceLocation
	err = client.UnmarshalCollectionBody(resp, &locations)

	return locations, err
}

// Examples of api-version: 5.1, 5.1-preview, 5.1-preview.1
var apiVersionRegEx = regexp.MustCompile(`(\d+(\.\d)?)(-preview(.(\d+))?)?`)

func combineUrl(part1 string, part2 string) string {
	return strings.TrimRight(part1, "/") + "/" + strings.TrimLeft(part2, "/")
}

func transformRouteTemplate(routeTemplate string, routeValues map[string]string) string {
	newTemplate := ""
	routeTemplate = strings.Replace(routeTemplate, "{*", "{", -1)
	segments := strings.Split(routeTemplate, "/")
	for _, segment := range segments {
		length := len(segment)
		if length <= 2 || segment[0] != '{' || segment[length-1] != '}' {
			newTemplate += "/" + segment
		} else {
			value, ok := routeValues[segment[1:length-1]]
			if ok {
				newTemplate += "/" + url.PathEscape(value)
			}
			// else this is an optional parameter that has not been supplied, so don't add it back
		}
	}
	// following covers oddball templates with segments that include the token and additional constants
	for key, value := range routeValues {
		newTemplate = strings.Replace(newTemplate, "{"+key+"}", value, -1)
	}
	return newTemplate
}

func (client *Client) UnmarshalBody(response *http.Response, v interface{}) (err error) {
	if response != nil && response.Body != nil {
		var err error
		defer func() {
			if closeError := response.Body.Close(); closeError != nil {
				err = closeError
			}
		}()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		body = trimByteOrderMark(body)
		return json.Unmarshal(body, &v)
	}
	return nil
}

func (client *Client) UnmarshalCollectionBody(response *http.Response, v interface{}) (err error) {
	if response != nil && response.Body != nil {
		var err error
		defer func() {
			if closeError := response.Body.Close(); closeError != nil {
				err = closeError
			}
		}()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		body = trimByteOrderMark(body)
		err = client.UnmarshalCollectionJson(body, v)
	}
	return nil
}

func (client *Client) UnmarshalCollectionJson(jsonValue []byte, v interface{}) (err error) {
	var wrappedResponse VssJsonCollectionWrapper
	err = json.Unmarshal(jsonValue, &wrappedResponse)
	if err != nil {
		return err
	}

	value, err := json.Marshal(wrappedResponse.Value) // todo: investigate better way to do this.
	if err != nil {
		return err
	}
	return json.Unmarshal(value, &v)
}

// Returns slice of body without utf-8 byte order mark.
// If BOM does not exist body is returned unchanged.
func trimByteOrderMark(body []byte) []byte {
	return bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
}

func (client *Client) UnwrapError(response *http.Response) (err error) {
	if response.ContentLength == 0 {
		message := "Request returned status: " + response.Status
		return &WrappedError{
			Message:    &message,
			StatusCode: &response.StatusCode,
		}
	}

	defer func() {
		if closeError := response.Body.Close(); closeError != nil {
			err = closeError
		}
	}()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	body = trimByteOrderMark(body)

	contentType, ok := response.Header[headerKeyContentType]
	if ok && len(contentType) > 0 && strings.Index(contentType[0], MediaTypeTextPlain) >= 0 {
		message := string(body)
		statusCode := response.StatusCode
		return WrappedError{Message: &message, StatusCode: &statusCode}
	}

	var wrappedError WrappedError
	err = json.Unmarshal(body, &wrappedError)
	if err != nil {
		return err
	}

	if wrappedError.Message == nil {
		var wrappedImproperError WrappedImproperError
		err = json.Unmarshal(body, &wrappedImproperError)
		if err == nil && wrappedImproperError.Value != nil && wrappedImproperError.Value.Message != nil {
			statusCode := response.StatusCode
			return &WrappedError{
				Message:    wrappedImproperError.Value.Message,
				StatusCode: &statusCode,
			}
		}
	}

	return wrappedError
}

func (client *Client) GetResourceAreas(ctx context.Context) (*[]ResourceAreaInfo, error) {
	queryParams := url.Values{}
	locationId, _ := uuid.Parse("e81700f7-3be2-46de-8624-2eb35882fcaa")
	resp, err := client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", nil, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []ResourceAreaInfo
	err = client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

type LocationIdNotRegisteredError struct {
	LocationId uuid.UUID
	Url        string
}

func (e LocationIdNotRegisteredError) Error() string {
	return "API resource location " + e.LocationId.String() + " is not registered on " + e.Url + "."
}

type InvalidApiVersion struct {
	ApiVersion string
}

func (e InvalidApiVersion) Error() string {
	return "The requested api-version is not in a valid format: " + e.ApiVersion
}
