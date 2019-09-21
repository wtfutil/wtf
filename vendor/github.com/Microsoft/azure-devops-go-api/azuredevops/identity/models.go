// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package identity

import (
	"github.com/google/uuid"
)

// Container class for changed identities
type ChangedIdentities struct {
	// Changed Identities
	Identities *[]Identity `json:"identities,omitempty"`
	// More data available, set to true if pagesize is specified.
	MoreData *bool `json:"moreData,omitempty"`
	// Last Identity SequenceId
	SequenceContext *ChangedIdentitiesContext `json:"sequenceContext,omitempty"`
}

// Context class for changed identities
type ChangedIdentitiesContext struct {
	// Last Group SequenceId
	GroupSequenceId *int `json:"groupSequenceId,omitempty"`
	// Last Identity SequenceId
	IdentitySequenceId *int `json:"identitySequenceId,omitempty"`
	// Last Group OrganizationIdentitySequenceId
	OrganizationIdentitySequenceId *int `json:"organizationIdentitySequenceId,omitempty"`
	// Page size
	PageSize *int `json:"pageSize,omitempty"`
}

type CreateScopeInfo struct {
	AdminGroupDescription *string         `json:"adminGroupDescription,omitempty"`
	AdminGroupName        *string         `json:"adminGroupName,omitempty"`
	CreatorId             *uuid.UUID      `json:"creatorId,omitempty"`
	ParentScopeId         *uuid.UUID      `json:"parentScopeId,omitempty"`
	ScopeName             *string         `json:"scopeName,omitempty"`
	ScopeType             *GroupScopeType `json:"scopeType,omitempty"`
}

type FrameworkIdentityInfo struct {
	DisplayName  *string                `json:"displayName,omitempty"`
	Identifier   *string                `json:"identifier,omitempty"`
	IdentityType *FrameworkIdentityType `json:"identityType,omitempty"`
	Role         *string                `json:"role,omitempty"`
}

type FrameworkIdentityType string

type frameworkIdentityTypeValuesType struct {
	None              FrameworkIdentityType
	ServiceIdentity   FrameworkIdentityType
	AggregateIdentity FrameworkIdentityType
	ImportedIdentity  FrameworkIdentityType
}

var FrameworkIdentityTypeValues = frameworkIdentityTypeValuesType{
	None:              "none",
	ServiceIdentity:   "serviceIdentity",
	AggregateIdentity: "aggregateIdentity",
	ImportedIdentity:  "importedIdentity",
}

type GroupMembership struct {
	Active     *bool      `json:"active,omitempty"`
	Descriptor *string    `json:"descriptor,omitempty"`
	Id         *uuid.UUID `json:"id,omitempty"`
	QueriedId  *uuid.UUID `json:"queriedId,omitempty"`
}

type GroupScopeType string

type groupScopeTypeValuesType struct {
	Generic     GroupScopeType
	ServiceHost GroupScopeType
	TeamProject GroupScopeType
}

var GroupScopeTypeValues = groupScopeTypeValuesType{
	Generic:     "generic",
	ServiceHost: "serviceHost",
	TeamProject: "teamProject",
}

type Identity struct {
	// The custom display name for the identity (if any). Setting this property to an empty string will clear the existing custom display name. Setting this property to null will not affect the existing persisted value (since null values do not get sent over the wire or to the database)
	CustomDisplayName *string      `json:"customDisplayName,omitempty"`
	Descriptor        *string      `json:"descriptor,omitempty"`
	Id                *uuid.UUID   `json:"id,omitempty"`
	IsActive          *bool        `json:"isActive,omitempty"`
	IsContainer       *bool        `json:"isContainer,omitempty"`
	MasterId          *uuid.UUID   `json:"masterId,omitempty"`
	MemberIds         *[]uuid.UUID `json:"memberIds,omitempty"`
	MemberOf          *[]string    `json:"memberOf,omitempty"`
	Members           *[]string    `json:"members,omitempty"`
	MetaTypeId        *int         `json:"metaTypeId,omitempty"`
	Properties        interface{}  `json:"properties,omitempty"`
	// The display name for the identity as specified by the source identity provider.
	ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
	ResourceVersion     *int    `json:"resourceVersion,omitempty"`
	SocialDescriptor    *string `json:"socialDescriptor,omitempty"`
	SubjectDescriptor   *string `json:"subjectDescriptor,omitempty"`
	UniqueUserId        *int    `json:"uniqueUserId,omitempty"`
}

// Base Identity class to allow "trimmed" identity class in the GetConnectionData API Makes sure that on-the-wire representations of the derived classes are compatible with each other (e.g. Server responds with PublicIdentity object while client deserializes it as Identity object) Derived classes should not have additional [DataMember] properties
type IdentityBase struct {
	// The custom display name for the identity (if any). Setting this property to an empty string will clear the existing custom display name. Setting this property to null will not affect the existing persisted value (since null values do not get sent over the wire or to the database)
	CustomDisplayName *string      `json:"customDisplayName,omitempty"`
	Descriptor        *string      `json:"descriptor,omitempty"`
	Id                *uuid.UUID   `json:"id,omitempty"`
	IsActive          *bool        `json:"isActive,omitempty"`
	IsContainer       *bool        `json:"isContainer,omitempty"`
	MasterId          *uuid.UUID   `json:"masterId,omitempty"`
	MemberIds         *[]uuid.UUID `json:"memberIds,omitempty"`
	MemberOf          *[]string    `json:"memberOf,omitempty"`
	Members           *[]string    `json:"members,omitempty"`
	MetaTypeId        *int         `json:"metaTypeId,omitempty"`
	Properties        interface{}  `json:"properties,omitempty"`
	// The display name for the identity as specified by the source identity provider.
	ProviderDisplayName *string `json:"providerDisplayName,omitempty"`
	ResourceVersion     *int    `json:"resourceVersion,omitempty"`
	SocialDescriptor    *string `json:"socialDescriptor,omitempty"`
	SubjectDescriptor   *string `json:"subjectDescriptor,omitempty"`
	UniqueUserId        *int    `json:"uniqueUserId,omitempty"`
}

type IdentityBatchInfo struct {
	Descriptors                 *[]string        `json:"descriptors,omitempty"`
	IdentityIds                 *[]uuid.UUID     `json:"identityIds,omitempty"`
	IncludeRestrictedVisibility *bool            `json:"includeRestrictedVisibility,omitempty"`
	PropertyNames               *[]string        `json:"propertyNames,omitempty"`
	QueryMembership             *QueryMembership `json:"queryMembership,omitempty"`
	SocialDescriptors           *[]string        `json:"socialDescriptors,omitempty"`
	SubjectDescriptors          *[]string        `json:"subjectDescriptors,omitempty"`
}

type IdentityScope struct {
	Administrators    *string         `json:"administrators,omitempty"`
	Id                *uuid.UUID      `json:"id,omitempty"`
	IsActive          *bool           `json:"isActive,omitempty"`
	IsGlobal          *bool           `json:"isGlobal,omitempty"`
	LocalScopeId      *uuid.UUID      `json:"localScopeId,omitempty"`
	Name              *string         `json:"name,omitempty"`
	ParentId          *uuid.UUID      `json:"parentId,omitempty"`
	ScopeType         *GroupScopeType `json:"scopeType,omitempty"`
	SecuringHostId    *uuid.UUID      `json:"securingHostId,omitempty"`
	SubjectDescriptor *string         `json:"subjectDescriptor,omitempty"`
}

// Identity information.
type IdentitySelf struct {
	// The UserPrincipalName (UPN) of the account. This value comes from the source provider.
	AccountName *string `json:"accountName,omitempty"`
	// The display name. For AAD accounts with multiple tenants this is the display name of the profile in the home tenant.
	DisplayName *string `json:"displayName,omitempty"`
	// This represents the name of the container of origin. For AAD accounts this is the tenantID of the home tenant. For MSA accounts this is the string "Windows Live ID".
	Domain *string `json:"domain,omitempty"`
	// This is the VSID of the home tenant profile. If the profile is signed into the home tenant or if the profile has no tenants then this Id is the same as the Id returned by the profile/profiles/me endpoint. Going forward it is recommended that you use the combined values of Origin, OriginId and Domain to uniquely identify a user rather than this Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// The type of source provider for the origin identifier. For MSA accounts this is "msa". For AAD accounts this is "aad".
	Origin *string `json:"origin,omitempty"`
	// The unique identifier from the system of origin. If there are multiple tenants this is the unique identifier of the account in the home tenant. (For MSA this is the PUID in hex notation, for AAD this is the object id.)
	OriginId *string `json:"originId,omitempty"`
	// For AAD accounts this is all of the tenants that this account is a member of.
	Tenants *[]TenantInfo `json:"tenants,omitempty"`
}

type IdentitySnapshot struct {
	Groups      *[]Identity        `json:"groups,omitempty"`
	IdentityIds *[]uuid.UUID       `json:"identityIds,omitempty"`
	Memberships *[]GroupMembership `json:"memberships,omitempty"`
	ScopeId     *uuid.UUID         `json:"scopeId,omitempty"`
	Scopes      *[]IdentityScope   `json:"scopes,omitempty"`
}

type IdentityUpdateData struct {
	Id      *uuid.UUID `json:"id,omitempty"`
	Index   *int       `json:"index,omitempty"`
	Updated *bool      `json:"updated,omitempty"`
}

type QueryMembership string

type queryMembershipValuesType struct {
	None         QueryMembership
	Direct       QueryMembership
	Expanded     QueryMembership
	ExpandedUp   QueryMembership
	ExpandedDown QueryMembership
}

var QueryMembershipValues = queryMembershipValuesType{
	// Query will not return any membership data
	None: "none",
	// Query will return only direct membership data
	Direct: "direct",
	// Query will return expanded membership data
	Expanded: "expanded",
	// Query will return expanded up membership data (parents only)
	ExpandedUp: "expandedUp",
	// Query will return expanded down membership data (children only)
	ExpandedDown: "expandedDown",
}

// [Flags]
type ReadIdentitiesOptions string

type readIdentitiesOptionsValuesType struct {
	None                     ReadIdentitiesOptions
	FilterIllegalMemberships ReadIdentitiesOptions
}

var ReadIdentitiesOptionsValues = readIdentitiesOptionsValuesType{
	None:                     "none",
	FilterIllegalMemberships: "filterIllegalMemberships",
}

type SwapIdentityInfo struct {
	Id1 *uuid.UUID `json:"id1,omitempty"`
	Id2 *uuid.UUID `json:"id2,omitempty"`
}

type TenantInfo struct {
	HomeTenant      *bool      `json:"homeTenant,omitempty"`
	TenantId        *uuid.UUID `json:"tenantId,omitempty"`
	TenantName      *string    `json:"tenantName,omitempty"`
	VerifiedDomains *[]string  `json:"verifiedDomains,omitempty"`
}
