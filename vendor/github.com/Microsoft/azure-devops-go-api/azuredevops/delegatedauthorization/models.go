// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package delegatedauthorization

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

type AccessTokenResult struct {
	AccessToken      *webapi.JsonWebToken `json:"accessToken,omitempty"`
	AccessTokenError *TokenError          `json:"accessTokenError,omitempty"`
	AuthorizationId  *uuid.UUID           `json:"authorizationId,omitempty"`
	ErrorDescription *string              `json:"errorDescription,omitempty"`
	HasError         *bool                `json:"hasError,omitempty"`
	RefreshToken     *RefreshTokenGrant   `json:"refreshToken,omitempty"`
	TokenType        *string              `json:"tokenType,omitempty"`
	ValidTo          *azuredevops.Time    `json:"validTo,omitempty"`
}

type Authorization struct {
	AccessIssued    *azuredevops.Time `json:"accessIssued,omitempty"`
	Audience        *string           `json:"audience,omitempty"`
	AuthorizationId *uuid.UUID        `json:"authorizationId,omitempty"`
	IdentityId      *uuid.UUID        `json:"identityId,omitempty"`
	IsAccessUsed    *bool             `json:"isAccessUsed,omitempty"`
	IsValid         *bool             `json:"isValid,omitempty"`
	RedirectUri     *string           `json:"redirectUri,omitempty"`
	RegistrationId  *uuid.UUID        `json:"registrationId,omitempty"`
	Scopes          *string           `json:"scopes,omitempty"`
	Source          *string           `json:"source,omitempty"`
	ValidFrom       *azuredevops.Time `json:"validFrom,omitempty"`
	ValidTo         *azuredevops.Time `json:"validTo,omitempty"`
}

type AuthorizationDecision struct {
	Authorization      *Authorization      `json:"authorization,omitempty"`
	AuthorizationError *AuthorizationError `json:"authorizationError,omitempty"`
	AuthorizationGrant *AuthorizationGrant `json:"authorizationGrant,omitempty"`
	HasError           *bool               `json:"hasError,omitempty"`
	IsAuthorized       *bool               `json:"isAuthorized,omitempty"`
}

type AuthorizationDescription struct {
	ClientRegistration *Registration                    `json:"clientRegistration,omitempty"`
	HasError           *bool                            `json:"hasError,omitempty"`
	InitiationError    *InitiationError                 `json:"initiationError,omitempty"`
	ScopeDescriptions  *[]AuthorizationScopeDescription `json:"scopeDescriptions,omitempty"`
}

type AuthorizationDetails struct {
	Authorization      *Authorization                   `json:"authorization,omitempty"`
	ClientRegistration *Registration                    `json:"clientRegistration,omitempty"`
	ScopeDescriptions  *[]AuthorizationScopeDescription `json:"scopeDescriptions,omitempty"`
}

type AuthorizationError string

type authorizationErrorValuesType struct {
	None                     AuthorizationError
	ClientIdRequired         AuthorizationError
	InvalidClientId          AuthorizationError
	ResponseTypeRequired     AuthorizationError
	ResponseTypeNotSupported AuthorizationError
	ScopeRequired            AuthorizationError
	InvalidScope             AuthorizationError
	RedirectUriRequired      AuthorizationError
	InsecureRedirectUri      AuthorizationError
	InvalidRedirectUri       AuthorizationError
	InvalidUserId            AuthorizationError
	InvalidUserType          AuthorizationError
	AccessDenied             AuthorizationError
}

var AuthorizationErrorValues = authorizationErrorValuesType{
	None:                     "none",
	ClientIdRequired:         "clientIdRequired",
	InvalidClientId:          "invalidClientId",
	ResponseTypeRequired:     "responseTypeRequired",
	ResponseTypeNotSupported: "responseTypeNotSupported",
	ScopeRequired:            "scopeRequired",
	InvalidScope:             "invalidScope",
	RedirectUriRequired:      "redirectUriRequired",
	InsecureRedirectUri:      "insecureRedirectUri",
	InvalidRedirectUri:       "invalidRedirectUri",
	InvalidUserId:            "invalidUserId",
	InvalidUserType:          "invalidUserType",
	AccessDenied:             "accessDenied",
}

type AuthorizationGrant struct {
	GrantType *GrantType `json:"grantType,omitempty"`
}

type AuthorizationScopeDescription struct {
	Description *string `json:"description,omitempty"`
	Market      *string `json:"market,omitempty"`
	Title       *string `json:"title,omitempty"`
}

type ClientType string

type clientTypeValuesType struct {
	Confidential ClientType
	Public       ClientType
	MediumTrust  ClientType
	HighTrust    ClientType
	FullTrust    ClientType
}

var ClientTypeValues = clientTypeValuesType{
	Confidential: "confidential",
	Public:       "public",
	MediumTrust:  "mediumTrust",
	HighTrust:    "highTrust",
	FullTrust:    "fullTrust",
}

type GrantType string

type grantTypeValuesType struct {
	None              GrantType
	JwtBearer         GrantType
	RefreshToken      GrantType
	Implicit          GrantType
	ClientCredentials GrantType
}

var GrantTypeValues = grantTypeValuesType{
	None:              "none",
	JwtBearer:         "jwtBearer",
	RefreshToken:      "refreshToken",
	Implicit:          "implicit",
	ClientCredentials: "clientCredentials",
}

type HostAuthorization struct {
	HostId         *uuid.UUID `json:"hostId,omitempty"`
	Id             *uuid.UUID `json:"id,omitempty"`
	IsValid        *bool      `json:"isValid,omitempty"`
	RegistrationId *uuid.UUID `json:"registrationId,omitempty"`
}

type HostAuthorizationDecision struct {
	HasError               *bool                   `json:"hasError,omitempty"`
	HostAuthorizationError *HostAuthorizationError `json:"hostAuthorizationError,omitempty"`
	HostAuthorizationId    *uuid.UUID              `json:"hostAuthorizationId,omitempty"`
}

type HostAuthorizationError string

type hostAuthorizationErrorValuesType struct {
	None                  HostAuthorizationError
	ClientIdRequired      HostAuthorizationError
	AccessDenied          HostAuthorizationError
	FailedToAuthorizeHost HostAuthorizationError
	ClientIdNotFound      HostAuthorizationError
	InvalidClientId       HostAuthorizationError
}

var HostAuthorizationErrorValues = hostAuthorizationErrorValuesType{
	None:                  "none",
	ClientIdRequired:      "clientIdRequired",
	AccessDenied:          "accessDenied",
	FailedToAuthorizeHost: "failedToAuthorizeHost",
	ClientIdNotFound:      "clientIdNotFound",
	InvalidClientId:       "invalidClientId",
}

type InitiationError string

type initiationErrorValuesType struct {
	None                     InitiationError
	ClientIdRequired         InitiationError
	InvalidClientId          InitiationError
	ResponseTypeRequired     InitiationError
	ResponseTypeNotSupported InitiationError
	ScopeRequired            InitiationError
	InvalidScope             InitiationError
	RedirectUriRequired      InitiationError
	InsecureRedirectUri      InitiationError
	InvalidRedirectUri       InitiationError
}

var InitiationErrorValues = initiationErrorValuesType{
	None:                     "none",
	ClientIdRequired:         "clientIdRequired",
	InvalidClientId:          "invalidClientId",
	ResponseTypeRequired:     "responseTypeRequired",
	ResponseTypeNotSupported: "responseTypeNotSupported",
	ScopeRequired:            "scopeRequired",
	InvalidScope:             "invalidScope",
	RedirectUriRequired:      "redirectUriRequired",
	InsecureRedirectUri:      "insecureRedirectUri",
	InvalidRedirectUri:       "invalidRedirectUri",
}

type RefreshTokenGrant struct {
	GrantType *GrantType           `json:"grantType,omitempty"`
	Jwt       *webapi.JsonWebToken `json:"jwt,omitempty"`
}

type Registration struct {
	ClientType           *ClientType `json:"clientType,omitempty"`
	IdentityId           *uuid.UUID  `json:"identityId,omitempty"`
	Issuer               *string     `json:"issuer,omitempty"`
	IsValid              *bool       `json:"isValid,omitempty"`
	IsWellKnown          *bool       `json:"isWellKnown,omitempty"`
	OrganizationLocation *string     `json:"organizationLocation,omitempty"`
	OrganizationName     *string     `json:"organizationName,omitempty"`
	// Raw cert data string from public key. This will be used for authenticating medium trust clients.
	PublicKey                            *string           `json:"publicKey,omitempty"`
	RedirectUris                         *[]string         `json:"redirectUris,omitempty"`
	RegistrationDescription              *string           `json:"registrationDescription,omitempty"`
	RegistrationId                       *uuid.UUID        `json:"registrationId,omitempty"`
	RegistrationLocation                 *string           `json:"registrationLocation,omitempty"`
	RegistrationLogoSecureLocation       *string           `json:"registrationLogoSecureLocation,omitempty"`
	RegistrationName                     *string           `json:"registrationName,omitempty"`
	RegistrationPrivacyStatementLocation *string           `json:"registrationPrivacyStatementLocation,omitempty"`
	RegistrationTermsOfServiceLocation   *string           `json:"registrationTermsOfServiceLocation,omitempty"`
	ResponseTypes                        *string           `json:"responseTypes,omitempty"`
	Scopes                               *string           `json:"scopes,omitempty"`
	Secret                               *string           `json:"secret,omitempty"`
	SecretValidTo                        *azuredevops.Time `json:"secretValidTo,omitempty"`
	SecretVersionId                      *uuid.UUID        `json:"secretVersionId,omitempty"`
	ValidFrom                            *azuredevops.Time `json:"validFrom,omitempty"`
}

type ResponseType string

type responseTypeValuesType struct {
	None         ResponseType
	Assertion    ResponseType
	IdToken      ResponseType
	TenantPicker ResponseType
	SignoutToken ResponseType
	AppToken     ResponseType
	Code         ResponseType
}

var ResponseTypeValues = responseTypeValuesType{
	None:         "none",
	Assertion:    "assertion",
	IdToken:      "idToken",
	TenantPicker: "tenantPicker",
	SignoutToken: "signoutToken",
	AppToken:     "appToken",
	Code:         "code",
}

type SessionToken struct {
	AccessId *uuid.UUID `json:"accessId,omitempty"`
	// This is populated when user requests a compact token. The alternate token value is self describing token.
	AlternateToken      *string            `json:"alternateToken,omitempty"`
	AuthorizationId     *uuid.UUID         `json:"authorizationId,omitempty"`
	Claims              *map[string]string `json:"claims,omitempty"`
	ClientId            *uuid.UUID         `json:"clientId,omitempty"`
	DisplayName         *string            `json:"displayName,omitempty"`
	HostAuthorizationId *uuid.UUID         `json:"hostAuthorizationId,omitempty"`
	IsPublic            *bool              `json:"isPublic,omitempty"`
	IsValid             *bool              `json:"isValid,omitempty"`
	PublicData          *string            `json:"publicData,omitempty"`
	Scope               *string            `json:"scope,omitempty"`
	Source              *string            `json:"source,omitempty"`
	TargetAccounts      *[]uuid.UUID       `json:"targetAccounts,omitempty"`
	// This is computed and not returned in Get queries
	Token     *string           `json:"token,omitempty"`
	UserId    *uuid.UUID        `json:"userId,omitempty"`
	ValidFrom *azuredevops.Time `json:"validFrom,omitempty"`
	ValidTo   *azuredevops.Time `json:"validTo,omitempty"`
}

type TokenError string

type tokenErrorValuesType struct {
	None                           TokenError
	GrantTypeRequired              TokenError
	AuthorizationGrantRequired     TokenError
	ClientSecretRequired           TokenError
	RedirectUriRequired            TokenError
	InvalidAuthorizationGrant      TokenError
	InvalidAuthorizationScopes     TokenError
	InvalidRefreshToken            TokenError
	AuthorizationNotFound          TokenError
	AuthorizationGrantExpired      TokenError
	AccessAlreadyIssued            TokenError
	InvalidRedirectUri             TokenError
	AccessTokenNotFound            TokenError
	InvalidAccessToken             TokenError
	AccessTokenAlreadyRefreshed    TokenError
	InvalidClientSecret            TokenError
	ClientSecretExpired            TokenError
	ServerError                    TokenError
	AccessDenied                   TokenError
	AccessTokenKeyRequired         TokenError
	InvalidAccessTokenKey          TokenError
	FailedToGetAccessToken         TokenError
	InvalidClientId                TokenError
	InvalidClient                  TokenError
	InvalidValidTo                 TokenError
	InvalidUserId                  TokenError
	FailedToIssueAccessToken       TokenError
	AuthorizationGrantScopeMissing TokenError
	InvalidPublicAccessTokenKey    TokenError
	InvalidPublicAccessToken       TokenError
	PublicFeatureFlagNotEnabled    TokenError
	SshPolicyDisabled              TokenError
}

var TokenErrorValues = tokenErrorValuesType{
	None:                           "none",
	GrantTypeRequired:              "grantTypeRequired",
	AuthorizationGrantRequired:     "authorizationGrantRequired",
	ClientSecretRequired:           "clientSecretRequired",
	RedirectUriRequired:            "redirectUriRequired",
	InvalidAuthorizationGrant:      "invalidAuthorizationGrant",
	InvalidAuthorizationScopes:     "invalidAuthorizationScopes",
	InvalidRefreshToken:            "invalidRefreshToken",
	AuthorizationNotFound:          "authorizationNotFound",
	AuthorizationGrantExpired:      "authorizationGrantExpired",
	AccessAlreadyIssued:            "accessAlreadyIssued",
	InvalidRedirectUri:             "invalidRedirectUri",
	AccessTokenNotFound:            "accessTokenNotFound",
	InvalidAccessToken:             "invalidAccessToken",
	AccessTokenAlreadyRefreshed:    "accessTokenAlreadyRefreshed",
	InvalidClientSecret:            "invalidClientSecret",
	ClientSecretExpired:            "clientSecretExpired",
	ServerError:                    "serverError",
	AccessDenied:                   "accessDenied",
	AccessTokenKeyRequired:         "accessTokenKeyRequired",
	InvalidAccessTokenKey:          "invalidAccessTokenKey",
	FailedToGetAccessToken:         "failedToGetAccessToken",
	InvalidClientId:                "invalidClientId",
	InvalidClient:                  "invalidClient",
	InvalidValidTo:                 "invalidValidTo",
	InvalidUserId:                  "invalidUserId",
	FailedToIssueAccessToken:       "failedToIssueAccessToken",
	AuthorizationGrantScopeMissing: "authorizationGrantScopeMissing",
	InvalidPublicAccessTokenKey:    "invalidPublicAccessTokenKey",
	InvalidPublicAccessToken:       "invalidPublicAccessToken",
	PublicFeatureFlagNotEnabled:    "publicFeatureFlagNotEnabled",
	SshPolicyDisabled:              "sshPolicyDisabled",
}
