package gerrit

import (
	"crypto/md5" // nolint: gosec
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	// ErrWWWAuthenticateHeaderMissing is returned by digestAuthHeader when the WWW-Authenticate header is missing
	ErrWWWAuthenticateHeaderMissing = errors.New("WWW-Authenticate header is missing")

	// ErrWWWAuthenticateHeaderInvalid is returned by digestAuthHeader when the WWW-Authenticate invalid
	ErrWWWAuthenticateHeaderInvalid = errors.New("WWW-Authenticate header is invalid")

	// ErrWWWAuthenticateHeaderNotDigest is returned by digestAuthHeader when the WWW-Authenticate header is not 'Digest'
	ErrWWWAuthenticateHeaderNotDigest = errors.New("WWW-Authenticate header type is not Digest")
)

const (
	// HTTP Basic Authentication
	authTypeBasic = 1
	// HTTP Digest Authentication
	authTypeDigest = 2
	// HTTP Cookie Authentication
	authTypeCookie = 3
)

// AuthenticationService contains Authentication related functions.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api.html#authentication
type AuthenticationService struct {
	client *Client

	// Storage for authentication
	// Username or name of cookie
	name string
	// Password or value of cookie
	secret   string
	authType int
}

// SetBasicAuth sets basic parameters for HTTP Basic auth
func (s *AuthenticationService) SetBasicAuth(username, password string) {
	s.name = username
	s.secret = password
	s.authType = authTypeBasic
}

// SetDigestAuth sets digest parameters for HTTP Digest auth.
func (s *AuthenticationService) SetDigestAuth(username, password string) {
	s.name = username
	s.secret = password
	s.authType = authTypeDigest
}

// digestAuthHeader is called by gerrit.Client.Do in the event the server
// returns 401 Unauthorized and authType was set to authTypeDigest. The
// resulting string is used to set the Authorization header before retrying
// the request.
func (s *AuthenticationService) digestAuthHeader(response *http.Response) (string, error) {
	authenticateHeader := response.Header.Get("WWW-Authenticate")
	if authenticateHeader == "" {
		return "", ErrWWWAuthenticateHeaderMissing
	}

	split := strings.SplitN(authenticateHeader, " ", 2)
	if len(split) != 2 {
		return "", ErrWWWAuthenticateHeaderInvalid
	}

	if split[0] != "Digest" {
		return "", ErrWWWAuthenticateHeaderNotDigest
	}

	// Iterate over all the fields from the WWW-Authenticate header
	// and create a map of keys and values.
	authenticate := map[string]string{}
	for _, value := range strings.Split(split[1], ",") {
		kv := strings.SplitN(value, "=", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.Trim(strings.Trim(kv[0], " "), "\"")
		value := strings.Trim(strings.Trim(kv[1], " "), "\"")
		authenticate[key] = value
	}

	// Gerrit usually responds without providing the algorithm.  According
	// to RFC2617 if no algorithm is provided then the default is to use
	// MD5. At the time this code was implemented Gerrit did not appear
	// to support other algorithms or provide a means of changing the
	// algorithm.
	if value, ok := authenticate["algorithm"]; ok {
		if value != "MD5" {
			return "", fmt.Errorf(
				"algorithm not implemented: %s", value)
		}
	}

	realmHeader := authenticate["realm"]
	qopHeader := authenticate["qop"]
	nonceHeader := authenticate["nonce"]

	// If the server does not inform us what the uri is supposed
	// to be then use the last requests's uri instead.
	if _, ok := authenticate["uri"]; !ok {
		authenticate["uri"] = response.Request.URL.Path
	}

	uriHeader := authenticate["uri"]

	// A1
	h := md5.New() // nolint: gosec
	A1 := fmt.Sprintf("%s:%s:%s", s.name, realmHeader, s.secret)
	if _, err := io.WriteString(h, A1); err != nil {
		return "", err
	}
	HA1 := fmt.Sprintf("%x", h.Sum(nil))

	// A2
	h = md5.New() // nolint: gosec
	A2 := fmt.Sprintf("%s:%s", response.Request.Method, uriHeader)
	if _, err := io.WriteString(h, A2); err != nil {
		return "", err
	}
	HA2 := fmt.Sprintf("%x", h.Sum(nil))

	k := make([]byte, 12)
	for bytes := 0; bytes < len(k); {
		n, err := rand.Read(k[bytes:])
		if err != nil {
			return "", fmt.Errorf("cnonce generation failed: %s", err)
		}
		bytes += n
	}
	cnonce := base64.StdEncoding.EncodeToString(k)
	digest := md5.New() // nolint: gosec
	if _, err := digest.Write([]byte(strings.Join([]string{HA1, nonceHeader, "00000001", cnonce, qopHeader, HA2}, ":"))); err != nil {
		return "", err
	}
	responseField := fmt.Sprintf("%x", digest.Sum(nil))

	return fmt.Sprintf(
		`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=00000001, qop=%s, response="%s"`,
		s.name, realmHeader, nonceHeader, uriHeader, cnonce, qopHeader, responseField), nil
}

// SetCookieAuth sets basic parameters for HTTP Cookie
func (s *AuthenticationService) SetCookieAuth(name, value string) {
	s.name = name
	s.secret = value
	s.authType = authTypeCookie
}

// HasBasicAuth checks if the auth type is HTTP Basic auth
func (s *AuthenticationService) HasBasicAuth() bool {
	return s.authType == authTypeBasic
}

// HasDigestAuth checks if the auth type is HTTP Digest based
func (s *AuthenticationService) HasDigestAuth() bool {
	return s.authType == authTypeDigest
}

// HasCookieAuth checks if the auth type is HTTP Cookie based
func (s *AuthenticationService) HasCookieAuth() bool {
	return s.authType == authTypeCookie
}

// HasAuth checks if an auth type is used
func (s *AuthenticationService) HasAuth() bool {
	return s.authType > 0
}

// ResetAuth resets all former authentification settings
func (s *AuthenticationService) ResetAuth() {
	s.name = ""
	s.secret = ""
	s.authType = 0
}
