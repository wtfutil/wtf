// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package webapi

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/system"
)

// Information about the location of a REST API resource
type ApiResourceLocation struct {
	// Area name for this resource
	Area *string `json:"area,omitempty"`
	// Unique Identifier for this location
	Id *uuid.UUID `json:"id,omitempty"`
	// Maximum api version that this resource supports (current server version for this resource)
	MaxVersion *string `json:"maxVersion,omitempty"`
	// Minimum api version that this resource supports
	MinVersion *string `json:"minVersion,omitempty"`
	// The latest version of this resource location that is in "Release" (non-preview) mode
	ReleasedVersion *string `json:"releasedVersion,omitempty"`
	// Resource name
	ResourceName *string `json:"resourceName,omitempty"`
	// The current resource version supported by this resource location
	ResourceVersion *int `json:"resourceVersion,omitempty"`
	// This location's route template (templated relative path)
	RouteTemplate *string `json:"routeTemplate,omitempty"`
}

// [Flags] Enumeration of the options that can be passed in on Connect.
type ConnectOptions string

type connectOptionsValuesType struct {
	None                               ConnectOptions
	IncludeServices                    ConnectOptions
	IncludeLastUserAccess              ConnectOptions
	IncludeInheritedDefinitionsOnly    ConnectOptions
	IncludeNonInheritedDefinitionsOnly ConnectOptions
}

var ConnectOptionsValues = connectOptionsValuesType{
	// Retrieve no optional data.
	None: "none",
	// Includes information about AccessMappings and ServiceDefinitions.
	IncludeServices: "includeServices",
	// Includes the last user access for this host.
	IncludeLastUserAccess: "includeLastUserAccess",
	// This is only valid on the deployment host and when true. Will only return inherited definitions.
	IncludeInheritedDefinitionsOnly: "includeInheritedDefinitionsOnly",
	// When true will only return non inherited definitions. Only valid at non-deployment host.
	IncludeNonInheritedDefinitionsOnly: "includeNonInheritedDefinitionsOnly",
}

// [Flags]
type DeploymentFlags string

type deploymentFlagsValuesType struct {
	None       DeploymentFlags
	Hosted     DeploymentFlags
	OnPremises DeploymentFlags
}

var DeploymentFlagsValues = deploymentFlagsValuesType{
	None:       "none",
	Hosted:     "hosted",
	OnPremises: "onPremises",
}

// Defines an "actor" for an event.
type EventActor struct {
	// Required: This is the identity of the user for the specified role.
	Id *uuid.UUID `json:"id,omitempty"`
	// Required: The event specific name of a role.
	Role *string `json:"role,omitempty"`
}

// Defines a scope for an event.
type EventScope struct {
	// Required: This is the identity of the scope for the type.
	Id *uuid.UUID `json:"id,omitempty"`
	// Optional: The display name of the scope
	Name *string `json:"name,omitempty"`
	// Required: The event specific type of a scope.
	Type *string `json:"type,omitempty"`
}

type IdentityRef struct {
	// Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
	DirectoryAlias *string `json:"directoryAlias,omitempty"`
	Id             *string `json:"id,omitempty"`
	// Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
	Inactive *bool `json:"inactive,omitempty"`
	// Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
	IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
	// Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
	IsContainer       *bool `json:"isContainer,omitempty"`
	IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
	// Deprecated - not in use in most preexisting implementations of ToIdentityRef
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// Deprecated - use Domain+PrincipalName instead
	UniqueName *string `json:"uniqueName,omitempty"`
}

type IdentityRefWithEmail struct {
	// Deprecated - Can be retrieved by querying the Graph user referenced in the "self" entry of the IdentityRef "_links" dictionary
	DirectoryAlias *string `json:"directoryAlias,omitempty"`
	Id             *string `json:"id,omitempty"`
	// Deprecated - Available in the "avatar" entry of the IdentityRef "_links" dictionary
	ImageUrl *string `json:"imageUrl,omitempty"`
	// Deprecated - Can be retrieved by querying the Graph membership state referenced in the "membershipState" entry of the GraphUser "_links" dictionary
	Inactive *bool `json:"inactive,omitempty"`
	// Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsAadUserType/Descriptor.IsAadGroupType)
	IsAadIdentity *bool `json:"isAadIdentity,omitempty"`
	// Deprecated - Can be inferred from the subject type of the descriptor (Descriptor.IsGroupType)
	IsContainer       *bool `json:"isContainer,omitempty"`
	IsDeletedInOrigin *bool `json:"isDeletedInOrigin,omitempty"`
	// Deprecated - not in use in most preexisting implementations of ToIdentityRef
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// Deprecated - use Domain+PrincipalName instead
	UniqueName            *string `json:"uniqueName,omitempty"`
	PreferredEmailAddress *string `json:"preferredEmailAddress,omitempty"`
}

// The JSON model for a JSON Patch operation
type JsonPatchOperation struct {
	// The path to copy from for the Move/Copy operation.
	From *string `json:"from,omitempty"`
	// The patch operation
	Op *Operation `json:"op,omitempty"`
	// The path for the operation. In the case of an array, a zero based index can be used to specify the position in the array (e.g. /biscuits/0/name). The "-" character can be used instead of an index to insert at the end of the array (e.g. /biscuits/-).
	Path *string `json:"path,omitempty"`
	// The value for the operation. This is either a primitive or a JToken.
	Value interface{} `json:"value,omitempty"`
}

type JsonWebToken struct {
}

type JWTAlgorithm string

type jwtAlgorithmValuesType struct {
	None  JWTAlgorithm
	HS256 JWTAlgorithm
	RS256 JWTAlgorithm
}

var JWTAlgorithmValues = jwtAlgorithmValuesType{
	None:  "none",
	HS256: "hS256",
	RS256: "rS256",
}

type Operation string

type operationValuesType struct {
	Add     Operation
	Remove  Operation
	Replace Operation
	Move    Operation
	Copy    Operation
	Test    Operation
}

var OperationValues = operationValuesType{
	Add:     "add",
	Remove:  "remove",
	Replace: "replace",
	Move:    "move",
	Copy:    "copy",
	Test:    "test",
}

// Represents the public key portion of an RSA asymmetric key.
type PublicKey struct {
	// Gets or sets the exponent for the public key.
	Exponent *[]byte `json:"exponent,omitempty"`
	// Gets or sets the modulus for the public key.
	Modulus *[]byte `json:"modulus,omitempty"`
}

type Publisher struct {
	// Name of the publishing service.
	Name *string `json:"name,omitempty"`
	// Service Owner Guid Eg. Tfs : 00025394-6065-48CA-87D9-7F5672854EF7
	ServiceOwnerId *uuid.UUID `json:"serviceOwnerId,omitempty"`
}

// The class to represent a REST reference link.  RFC: http://tools.ietf.org/html/draft-kelly-json-hal-06  The RFC is not fully implemented, additional properties are allowed on the reference link but as of yet we don't have a need for them.
type ReferenceLink struct {
	Href *string `json:"href,omitempty"`
}

type ResourceRef struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

type ServiceEvent struct {
	// This is the id of the type. Constants that will be used by subscribers to identify/filter events being published on a topic.
	EventType *string `json:"eventType,omitempty"`
	// This is the service that published this event.
	Publisher *Publisher `json:"publisher,omitempty"`
	// The resource object that carries specific information about the event. The object must have the ServiceEventObject applied for serialization/deserialization to work.
	Resource interface{} `json:"resource,omitempty"`
	// This dictionary carries the context descriptors along with their ids.
	ResourceContainers *map[string]interface{} `json:"resourceContainers,omitempty"`
	// This is the version of the resource.
	ResourceVersion *string `json:"resourceVersion,omitempty"`
}

// A signed url allowing limited-time anonymous access to private resources.
type SignedUrl struct {
	SignatureExpires *azuredevops.Time `json:"signatureExpires,omitempty"`
	Url              *string           `json:"url,omitempty"`
}

type TeamMember struct {
	Identity    *IdentityRef `json:"identity,omitempty"`
	IsTeamAdmin *bool        `json:"isTeamAdmin,omitempty"`
}

// A single secured timing consisting of a duration and start time
type TimingEntry struct {
	// Duration of the entry in ticks
	ElapsedTicks *uint64 `json:"elapsedTicks,omitempty"`
	// Properties to distinguish timings within the same group or to provide data to send with telemetry
	Properties *map[string]interface{} `json:"properties,omitempty"`
	// Offset from Server Request Context start time in microseconds
	StartOffset *uint64 `json:"startOffset,omitempty"`
}

// A set of secured performance timings all keyed off of the same string
type TimingGroup struct {
	// The total number of timing entries associated with this group
	Count *int `json:"count,omitempty"`
	// Overall duration of all entries in this group in ticks
	ElapsedTicks *uint64 `json:"elapsedTicks,omitempty"`
	// A list of timing entries in this group. Only the first few entries in each group are collected.
	Timings *[]TimingEntry `json:"timings,omitempty"`
}

// This class describes a trace filter, i.e. a set of criteria on whether or not a trace event should be emitted
type TraceFilter struct {
	Area          *string            `json:"area,omitempty"`
	ExceptionType *string            `json:"exceptionType,omitempty"`
	IsEnabled     *bool              `json:"isEnabled,omitempty"`
	Layer         *string            `json:"layer,omitempty"`
	Level         *system.TraceLevel `json:"level,omitempty"`
	Method        *string            `json:"method,omitempty"`
	// Used to serialize additional identity information (display name, etc) to clients. Not set by default. Server-side callers should use OwnerId.
	Owner       *IdentityRef      `json:"owner,omitempty"`
	OwnerId     *uuid.UUID        `json:"ownerId,omitempty"`
	Path        *string           `json:"path,omitempty"`
	ProcessName *string           `json:"processName,omitempty"`
	Service     *string           `json:"service,omitempty"`
	ServiceHost *uuid.UUID        `json:"serviceHost,omitempty"`
	TimeCreated *azuredevops.Time `json:"timeCreated,omitempty"`
	TraceId     *uuid.UUID        `json:"traceId,omitempty"`
	Tracepoint  *int              `json:"tracepoint,omitempty"`
	Uri         *string           `json:"uri,omitempty"`
	UserAgent   *string           `json:"userAgent,omitempty"`
	UserLogin   *string           `json:"userLogin,omitempty"`
}

type VssJsonCollectionWrapper struct {
	Count *int           `json:"count,omitempty"`
	Value *[]interface{} `json:"value,omitempty"`
}

type VssJsonCollectionWrapperBase struct {
	Count *int `json:"count,omitempty"`
}

// This is the type used for firing notifications intended for the subsystem in the Notifications SDK. For components that can't take a dependency on the Notifications SDK directly, they can use ITeamFoundationEventService.PublishNotification and the Notifications SDK ISubscriber implementation will get it.
type VssNotificationEvent struct {
	// Optional: A list of actors which are additional identities with corresponding roles that are relevant to the event.
	Actors *[]EventActor `json:"actors,omitempty"`
	// Optional: A list of artifacts referenced or impacted by this event.
	ArtifactUris *[]string `json:"artifactUris,omitempty"`
	// Required: The event payload.  If Data is a string, it must be in Json or XML format.  Otherwise it must have a serialization format attribute.
	Data interface{} `json:"data,omitempty"`
	// Required: The name of the event.  This event must be registered in the context it is being fired.
	EventType *string `json:"eventType,omitempty"`
	// How long before the event expires and will be cleaned up.  The default is to use the system default.
	ExpiresIn interface{} `json:"expiresIn,omitempty"`
	// The id of the item, artifact, extension, project, etc.
	ItemId *string `json:"itemId,omitempty"`
	// How long to wait before processing this event.  The default is to process immediately.
	ProcessDelay interface{} `json:"processDelay,omitempty"`
	// Optional: A list of scopes which are are relevant to the event.
	Scopes *[]EventScope `json:"scopes,omitempty"`
	// This is the time the original source event for this VssNotificationEvent was created.  For example, for something like a build completion notification SourceEventCreatedTime should be the time the build finished not the time this event was raised.
	SourceEventCreatedTime *azuredevops.Time `json:"sourceEventCreatedTime,omitempty"`
}

type WrappedException struct {
	CustomProperties *map[string]interface{} `json:"customProperties,omitempty"`
	ErrorCode        *int                    `json:"errorCode,omitempty"`
	EventId          *int                    `json:"eventId,omitempty"`
	HelpLink         *string                 `json:"helpLink,omitempty"`
	InnerException   *WrappedException       `json:"innerException,omitempty"`
	Message          *string                 `json:"message,omitempty"`
	StackTrace       *string                 `json:"stackTrace,omitempty"`
	TypeKey          *string                 `json:"typeKey,omitempty"`
	TypeName         *string                 `json:"typeName,omitempty"`
}
