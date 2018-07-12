package gerrit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/go-querystring/query"
)

// TODO Try to reduce the code duplications of a std API req
// Maybe with http://play.golang.org/p/j-667shCCB
// and https://groups.google.com/forum/#!topic/golang-nuts/D-gIr24k5uY

// A Client manages communication with the Gerrit API.
type Client struct {
	// client is the HTTP client used to communicate with the API.
	client *http.Client

	// baseURL is the base URL of the Gerrit instance for API requests.
	// It must have a trailing slash.
	baseURL *url.URL

	// Gerrit service for authentication.
	Authentication *AuthenticationService

	// Services used for talking to different parts of the standard Gerrit API.
	Access   *AccessService
	Accounts *AccountsService
	Changes  *ChangesService
	Config   *ConfigService
	Groups   *GroupsService
	Plugins  *PluginsService
	Projects *ProjectsService

	// Additional services used for talking to non-standard Gerrit APIs.
	EventsLog *EventsLogService
}

// Response is a Gerrit API response.
// This wraps the standard http.Response returned from Gerrit.
type Response struct {
	*http.Response
}

var (
	// ErrNoInstanceGiven is returned by NewClient in the event the
	// gerritURL argument was blank.
	ErrNoInstanceGiven = errors.New("no Gerrit instance given")

	// ErrUserProvidedWithoutPassword is returned by NewClient
	// if a user name is provided without a password.
	ErrUserProvidedWithoutPassword = errors.New("a username was provided without a password")

	// ErrAuthenticationFailed is returned by NewClient in the event the provided
	// credentials didn't allow us to query account information using digest, basic or cookie
	// auth.
	ErrAuthenticationFailed = errors.New("failed to authenticate using the provided credentials")

	// ReParseURL is used to parse the url provided to NewClient(). This
	// regular expression contains five groups which capture the scheme,
	// username, password, hostname and port. If we parse the url with this
	// regular expression
	ReParseURL = regexp.MustCompile(`^(http|https)://(.+):(.+)@(.+):(\d+)(.*)$`)
)

// NewClient returns a new Gerrit API client. gerritURL specifies the
// HTTP endpoint of the Gerrit instance. For example, "http://localhost:8080/".
// If gerritURL does not have a trailing slash, one is added automatically.
// If a nil httpClient is provided, http.DefaultClient will be used.
//
// The url may contain credentials, http://admin:secret@localhost:8081/ for
// example. These credentials may either be a user name and password or
// name and value as in the case of cookie based authentication. If the url contains
// credentials then this function will attempt to validate the credentials before
// returning the client. ErrAuthenticationFailed will be returned if the credentials
// cannot be validated. The process of validating the credentials is relatively simple and
// only requires that the provided user have permission to GET /a/accounts/self.
func NewClient(gerritURL string, httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	endpoint := gerritURL
	if endpoint == "" {
		return nil, ErrNoInstanceGiven
	}

	hasAuth := false
	username := ""
	password := ""

	// Depending on the contents of the username and password the default
	// url.Parse may not work. The below is an example URL that
	// would end up being parsed incorrectly with url.Parse:
	//   http://admin:ZOSOKjgV/kgEkN0bzPJp+oGeJLqpXykqWFJpon/Ckg@localhost:38607
	// So instead of depending on url.Parse we'll try using a regular expression
	// first to match a specific pattern. If that ends up working we modify
	// the incoming endpoint to remove the username and password so the rest
	// of this function will run as expected.
	submatches := ReParseURL.FindAllStringSubmatch(endpoint, -1)
	if len(submatches) > 0 && len(submatches[0]) > 5 {
		submatch := submatches[0]
		username = submatch[2]
		password = submatch[3]
		endpoint = fmt.Sprintf(
			"%s://%s:%s%s", submatch[1], submatch[4], submatch[5], submatch[6])
		hasAuth = true
	}

	baseURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(baseURL.Path, "/") {
		baseURL.Path += "/"
	}

	// Note, if we retrieved the URL and password using the regular
	// expression above then the below code will do nothing.
	if baseURL.User != nil {
		username = baseURL.User.Username()
		parsedPassword, haspassword := baseURL.User.Password()

		// Catches cases like http://user@localhost:8081/ where no password
		// was at all. If a blank password is required
		if !haspassword {
			return nil, ErrUserProvidedWithoutPassword
		}

		password = parsedPassword

		// Reconstruct the url but without the username and password.
		baseURL, err = url.Parse(
			fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, baseURL.RequestURI()))
		if err != nil {
			return nil, err
		}
		hasAuth = true
	}

	c := &Client{
		client:  httpClient,
		baseURL: baseURL,
	}
	c.Authentication = &AuthenticationService{client: c}
	c.Access = &AccessService{client: c}
	c.Accounts = &AccountsService{client: c}
	c.Changes = &ChangesService{client: c}
	c.Config = &ConfigService{client: c}
	c.Groups = &GroupsService{client: c}
	c.Plugins = &PluginsService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.EventsLog = &EventsLogService{client: c}

	if hasAuth {
		// Digest auth (first since that's the default auth type)
		c.Authentication.SetDigestAuth(username, password)
		if success, err := checkAuth(c); success || err != nil {
			return c, err
		}

		// Basic auth
		c.Authentication.SetBasicAuth(username, password)
		if success, err := checkAuth(c); success || err != nil {
			return c, err
		}

		// Cookie auth
		c.Authentication.SetCookieAuth(username, password)
		if success, err := checkAuth(c); success || err != nil {
			return c, err
		}

		// Reset auth in case the consumer needs to do something special.
		c.Authentication.ResetAuth()
		return c, ErrAuthenticationFailed
	}

	return c, nil
}

// checkAuth is used by NewClient to check if the current credentials are
// valid. If the response is 401 Unauthorized then the error will be discarded.
func checkAuth(client *Client) (bool, error) {
	_, response, err := client.Accounts.GetAccount("self")
	switch err {
	case ErrWWWAuthenticateHeaderMissing:
		return false, nil
	case ErrWWWAuthenticateHeaderNotDigest:
		return false, nil
	default:
		// Response could be nil if the connection outright failed
		// or some other error occurred before we got a response.
		if response == nil && err != nil {
			return false, err
		}

		if err != nil && response.StatusCode == http.StatusUnauthorized {
			err = nil
		}
		return response.StatusCode == http.StatusOK, err
	}
}

// NewRequest creates an API request.
// A relative URL can be provided in urlStr, in which case it is resolved relative to the baseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	// Build URL for request
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	// Apply Authentication
	if err := c.addAuthentication(req); err != nil {
		return nil, err
	}

	// Request compact JSON
	// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#output
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// TODO: Add gzip encoding
	// Accept-Encoding request header is set to gzip
	// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#output

	return req, nil
}

// NewRawPutRequest creates a raw PUT request and makes no attempt to encode
// or marshal the body. Just passes it straight through.
func (c *Client) NewRawPutRequest(urlStr string, body string) (*http.Request, error) {
	// Build URL for request
	u, err := c.buildURLForRequest(urlStr)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte(body))
	req, err := http.NewRequest("PUT", u, buf)
	if err != nil {
		return nil, err
	}

	// Apply Authentication
	if err := c.addAuthentication(req); err != nil {
		return nil, err
	}

	// Request compact JSON
	// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#output
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// TODO: Add gzip encoding
	// Accept-Encoding request header is set to gzip
	// See https://gerrit-review.googlesource.com/Documentation/rest-api.html#output

	return req, nil
}

// Call is a combine function for Client.NewRequest and Client.Do.
//
// Most API methods are quite the same.
// Get the URL, apply options, make a request, and get the response.
// Without adding special headers or something.
// To avoid a big amount of code duplication you can Client.Call.
//
// method is the HTTP method you want to call.
// u is the URL you want to call.
// body is the HTTP body.
// v is the HTTP response.
//
// For more information read https://github.com/google/go-github/issues/234
func (c *Client) Call(method, u string, body interface{}, v interface{}) (*Response, error) {
	req, err := c.NewRequest(method, u, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req, v)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// buildURLForRequest will build the URL (as string) that will be called.
// We need such a utility method, because the URL.Path needs to be escaped (partly).
//
// E.g. if a project is called via "projects/%s" and the project is named "plugin/delete-project"
// there has to be "projects/plugin%25Fdelete-project" instead of "projects/plugin/delete-project".
// The second url will return nothing.
func (c *Client) buildURLForRequest(urlStr string) (string, error) {
	// If there is a "/" at the start, remove it.
	// TODO: It can be arranged for all callers of buildURLForRequest to never have a "/" prefix,
	//       which can be ensured via tests. This is how it's done in go-github.
	//       Then, this run-time check becomes unnecessary and can be removed.
	urlStr = strings.TrimPrefix(urlStr, "/")

	// If we are authenticated, let's apply the "a/" prefix,
	// but only if it has not already been applied.
	if c.Authentication.HasAuth() && !strings.HasPrefix(urlStr, "a/") {
		urlStr = "a/" + urlStr
	}

	rel, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	return c.baseURL.String() + rel.String(), nil
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred.
// If v implements the io.Writer interface, the raw response body will be written to v,
// without attempting to first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// Wrap response
	response := &Response{Response: resp}

	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	if v != nil {
		defer resp.Body.Close() // nolint: errcheck
		if w, ok := v.(io.Writer); ok {
			if _, err := io.Copy(w, resp.Body); err != nil { // nolint: vetshadow
				return nil, err
			}
		} else {
			var body []byte
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				// even though there was an error, we still return the response
				// in case the caller wants to inspect it further
				return response, err
			}

			body = RemoveMagicPrefixLine(body)
			err = json.Unmarshal(body, v)
		}
	}
	return response, err
}

func (c *Client) addAuthentication(req *http.Request) error {
	// Apply HTTP Basic Authentication
	if c.Authentication.HasBasicAuth() {
		req.SetBasicAuth(c.Authentication.name, c.Authentication.secret)
		return nil
	}

	// Apply HTTP Cookie
	if c.Authentication.HasCookieAuth() {
		req.AddCookie(&http.Cookie{
			Name:  c.Authentication.name,
			Value: c.Authentication.secret,
		})
		return nil
	}

	// Apply Digest Authentication.  If we're using digest based
	// authentication we need to make a request, process the
	// WWW-Authenticate header, then set the Authorization header on the
	// incoming request.  We do not need to send a body along because
	// the request itself should fail first.
	if c.Authentication.HasDigestAuth() {
		uri, err := c.buildURLForRequest(req.URL.RequestURI())
		if err != nil {
			return err
		}

		// WARNING: Don't use c.NewRequest here unless you like
		// infinite recursion.
		digestRequest, err := http.NewRequest(req.Method, uri, nil)
		digestRequest.Header.Set("Accept", "*/*")
		digestRequest.Header.Set("Content-Type", "application/json")
		if err != nil {
			return err
		}

		response, err := c.client.Do(digestRequest)
		if err != nil {
			return err
		}

		// When the function exits discard the rest of the
		// body and close it.  This should cause go to
		// reuse the connection.
		defer io.Copy(ioutil.Discard, response.Body) // nolint: errcheck
		defer response.Body.Close()                  // nolint: errcheck

		if response.StatusCode == http.StatusUnauthorized {
			authorization, err := c.Authentication.digestAuthHeader(response)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", authorization)
		}
	}

	return nil
}

// DeleteRequest sends an DELETE API Request to urlStr with optional body.
// It is a shorthand combination for Client.NewRequest with Client.Do.
//
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by body is JSON encoded and included as the request body.
func (c *Client) DeleteRequest(urlStr string, body interface{}) (*Response, error) {
	req, err := c.NewRequest("DELETE", urlStr, body)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}

// BaseURL returns the client's Gerrit instance HTTP endpoint.
func (c *Client) BaseURL() url.URL {
	return *c.baseURL
}

// RemoveMagicPrefixLine removes the "magic prefix line" of Gerris JSON
// response if present. The JSON response body starts with a magic prefix line
// that must be stripped before feeding the rest of the response body to a JSON
// parser. The reason for this is to prevent against Cross Site Script
// Inclusion (XSSI) attacks.  By default all standard Gerrit APIs include this
// prefix line though some plugins may not.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#output
func RemoveMagicPrefixLine(body []byte) []byte {
	if bytes.HasPrefix(body, magicPrefix) {
		return body[5:]
	}
	return body
}

var magicPrefix = []byte(")]}'\n")

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
// API error responses are expected to have no response body.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#response-codes
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	// Some calls require an authentification
	// In such cases errors like:
	// 		API call to https://review.typo3.org/accounts/self failed: 403 Forbidden
	// will be thrown.

	err := fmt.Errorf("API call to %s failed: %s", r.Request.URL.String(), r.Status)
	return err
}

// queryParameterReplacements are values in a url, specifically the query
// portion of the url, which should not be escaped before being sent to
// Gerrit. Note, Gerrit itself does not escape these values when using the
// search box so we shouldn't escape them either.
var queryParameterReplacements = map[string]string{
	"+": "GOGERRIT_URL_PLACEHOLDER_PLUS",
	":": "GOGERRIT_URL_PLACEHOLDER_COLON"}

// addOptions adds the parameters in opt as URL query parameters to s.
// opt must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	// If the url contained one or more query parameters (q) then we need
	// to do some escaping on these values before Encode() is called.  By
	// doing so we're ensuring that : and + don't get encoded which means
	// they'll be passed along to Gerrit as raw ascii. Without this Gerrit
	// could return 400 Bad Request depending on the query parameters. For
	// more complete information see this issue on GitHub:
	//   https://github.com/andygrunwald/go-gerrit/issues/18
	_, hasQuery := qs["q"]
	if hasQuery {
		values := []string{}
		for _, value := range qs["q"] {
			for key, replacement := range queryParameterReplacements {
				value = strings.Replace(value, key, replacement, -1)
			}
			values = append(values, value)
		}

		qs.Del("q")
		for _, value := range values {
			qs.Add("q", value)
		}
	}
	encoded := qs.Encode()

	if hasQuery {
		for key, replacement := range queryParameterReplacements {
			encoded = strings.Replace(encoded, replacement, key, -1)
		}
	}

	u.RawQuery = encoded
	return u.String(), nil
}

// getStringResponseWithoutOptions retrieved a single string Response for a GET request
func getStringResponseWithoutOptions(client *Client, u string) (string, *Response, error) {
	v := new(string)
	resp, err := client.Call("GET", u, nil, v)
	return *v, resp, err
}
