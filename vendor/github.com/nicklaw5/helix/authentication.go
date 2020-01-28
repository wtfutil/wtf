package helix

import (
	"strings"
)

var authPaths = map[string]string{
	"token":  "/token",
	"revoke": "/revoke",
}

// GetAuthorizationURL ...
func (c *Client) GetAuthorizationURL(state string, forceVerify bool) string {
	opts := c.opts

	url := AuthBaseURL + "/authorize?response_type=code"
	url += "&client_id=" + opts.ClientID
	url += "&redirect_uri=" + opts.RedirectURI

	if state != "" {
		url += "&state=" + state
	}

	if forceVerify {
		url += "&force_verify=true"
	}

	if len(opts.Scopes) > 0 {
		url += "&scope=" + strings.Join(opts.Scopes, "%20")
	}

	return url
}

// AppAccessCredentials ...
type AppAccessCredentials struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// AppAccessTokenResponse ...
type AppAccessTokenResponse struct {
	ResponseCommon
	Data AppAccessCredentials
}

type appAccessTokenRequestData struct {
	ClientID     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
	RedirectURI  string `query:"redirect_uri"`
	GrantType    string `query:"grant_type"`
}

// GetAppAccessToken ...
func (c *Client) GetAppAccessToken() (*AppAccessTokenResponse, error) {
	opts := c.opts
	data := &accessTokenRequestData{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		RedirectURI:  opts.RedirectURI,
		GrantType:    "client_credentials",
	}

	resp, err := c.post(authPaths["token"], &AppAccessCredentials{}, data)
	if err != nil {
		return nil, err
	}

	token := &AppAccessTokenResponse{}
	token.StatusCode = resp.StatusCode
	token.Error = resp.Error
	token.ErrorStatus = resp.ErrorStatus
	token.ErrorMessage = resp.ErrorMessage
	token.Data.AccessToken = resp.Data.(*AppAccessCredentials).AccessToken
	token.Data.ExpiresIn = resp.Data.(*AppAccessCredentials).ExpiresIn

	return token, nil
}

// UserAccessCredentials ...
type UserAccessCredentials struct {
	AccessToken  string   `json:"access_token"`
	RefreshToken string   `json:"refresh_token"`
	ExpiresIn    int      `json:"expires_in"`
	Scopes       []string `json:"scope"`
}

// UserAccessTokenResponse ...
type UserAccessTokenResponse struct {
	ResponseCommon
	Data UserAccessCredentials
}

type accessTokenRequestData struct {
	Code         string `query:"code"`
	ClientID     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
	RedirectURI  string `query:"redirect_uri"`
	GrantType    string `query:"grant_type"`
}

// GetUserAccessToken ...
func (c *Client) GetUserAccessToken(code string) (*UserAccessTokenResponse, error) {
	opts := c.opts
	data := &accessTokenRequestData{
		Code:         code,
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		RedirectURI:  opts.RedirectURI,
		GrantType:    "authorization_code",
	}

	resp, err := c.post(authPaths["token"], &UserAccessCredentials{}, data)
	if err != nil {
		return nil, err
	}

	token := &UserAccessTokenResponse{}
	token.StatusCode = resp.StatusCode
	token.Error = resp.Error
	token.ErrorStatus = resp.ErrorStatus
	token.ErrorMessage = resp.ErrorMessage
	token.Data.AccessToken = resp.Data.(*UserAccessCredentials).AccessToken
	token.Data.RefreshToken = resp.Data.(*UserAccessCredentials).RefreshToken
	token.Data.ExpiresIn = resp.Data.(*UserAccessCredentials).ExpiresIn
	token.Data.Scopes = resp.Data.(*UserAccessCredentials).Scopes

	return token, nil
}

// RefreshTokenResponse ...
type RefreshTokenResponse struct {
	ResponseCommon
	Data UserAccessCredentials
}

type refreshTokenRequestData struct {
	ClientID     string `query:"client_id"`
	ClientSecret string `query:"client_secret"`
	GrantType    string `query:"grant_type"`
	RefreshToken string `query:"refresh_token"`
}

// RefreshUserAccessToken submits a request to have the longevity of an
// access token extended. Twitch OAuth2 access tokens have expirations.
// Token-expiration periods vary in length. You should build your applications
// in such a way that they are resilient to token authentication failures.
func (c *Client) RefreshUserAccessToken(refreshToken string) (*RefreshTokenResponse, error) {
	opts := c.opts
	data := &refreshTokenRequestData{
		ClientID:     opts.ClientID,
		ClientSecret: opts.ClientSecret,
		GrantType:    "refresh_token",
		RefreshToken: refreshToken,
	}

	resp, err := c.post(authPaths["token"], &UserAccessCredentials{}, data)
	if err != nil {
		return nil, err
	}

	refresh := &RefreshTokenResponse{}
	refresh.StatusCode = resp.StatusCode
	refresh.Error = resp.Error
	refresh.ErrorStatus = resp.ErrorStatus
	refresh.ErrorMessage = resp.ErrorMessage
	refresh.Data.AccessToken = resp.Data.(*UserAccessCredentials).AccessToken
	refresh.Data.RefreshToken = resp.Data.(*UserAccessCredentials).RefreshToken
	refresh.Data.ExpiresIn = resp.Data.(*UserAccessCredentials).ExpiresIn
	refresh.Data.Scopes = resp.Data.(*UserAccessCredentials).Scopes

	return refresh, nil
}

// RevokeAccessTokenResponse ...
type RevokeAccessTokenResponse struct {
	ResponseCommon
}

type revokeAccessTokenRequestData struct {
	ClientID    string `query:"client_id"`
	AccessToken string `query:"token"`
}

// RevokeUserAccessToken submits a request to Twitch to have an access token revoked.
//
// Both successful requests and requests with bad tokens return 200 OK with
// no body. Requests with bad tokens return the same response, as there is no
// meaningful action a client can take after sending a bad token.
func (c *Client) RevokeUserAccessToken(accessToken string) (*RevokeAccessTokenResponse, error) {
	data := &revokeAccessTokenRequestData{
		ClientID:    c.opts.ClientID,
		AccessToken: accessToken,
	}

	resp, err := c.post(authPaths["revoke"], nil, data)
	if err != nil {
		return nil, err
	}

	revoke := &RevokeAccessTokenResponse{}
	revoke.StatusCode = resp.StatusCode
	revoke.Error = resp.Error
	revoke.ErrorStatus = resp.ErrorStatus
	revoke.ErrorMessage = resp.ErrorMessage

	return revoke, nil
}
