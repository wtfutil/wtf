// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package core

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/identity"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

type ConnectedServiceKind string

type connectedServiceKindValuesType struct {
	Custom            ConnectedServiceKind
	AzureSubscription ConnectedServiceKind
	Chef              ConnectedServiceKind
	Generic           ConnectedServiceKind
}

var ConnectedServiceKindValues = connectedServiceKindValuesType{
	// Custom or unknown service
	Custom: "custom",
	// Azure Subscription
	AzureSubscription: "azureSubscription",
	// Chef Connection
	Chef: "chef",
	// Generic Connection
	Generic: "generic",
}

type IdentityData struct {
	IdentityIds *[]uuid.UUID `json:"identityIds,omitempty"`
}

type Process struct {
	Name        *string      `json:"name,omitempty"`
	Url         *string      `json:"url,omitempty"`
	Links       interface{}  `json:"_links,omitempty"`
	Description *string      `json:"description,omitempty"`
	Id          *uuid.UUID   `json:"id,omitempty"`
	IsDefault   *bool        `json:"isDefault,omitempty"`
	Type        *ProcessType `json:"type,omitempty"`
}

// Type of process customization on a collection.
type ProcessCustomizationType string

type processCustomizationTypeValuesType struct {
	Xml       ProcessCustomizationType
	Inherited ProcessCustomizationType
}

var ProcessCustomizationTypeValues = processCustomizationTypeValuesType{
	// Customization based on project-scoped xml customization
	Xml: "xml",
	// Customization based on process inheritance
	Inherited: "inherited",
}

type ProcessReference struct {
	Name *string `json:"name,omitempty"`
	Url  *string `json:"url,omitempty"`
}

type ProcessType string

type processTypeValuesType struct {
	System    ProcessType
	Custom    ProcessType
	Inherited ProcessType
}

var ProcessTypeValues = processTypeValuesType{
	System:    "system",
	Custom:    "custom",
	Inherited: "inherited",
}

// Contains the image data for project avatar.
type ProjectAvatar struct {
	// The avatar image represented as a byte array.
	Image *[]byte `json:"image,omitempty"`
}

type ProjectChangeType string

type projectChangeTypeValuesType struct {
	Modified ProjectChangeType
	Deleted  ProjectChangeType
	Added    ProjectChangeType
}

var ProjectChangeTypeValues = projectChangeTypeValuesType{
	Modified: "modified",
	Deleted:  "deleted",
	Added:    "added",
}

// Contains information describing a project.
type ProjectInfo struct {
	// The abbreviated name of the project.
	Abbreviation *string `json:"abbreviation,omitempty"`
	// The description of the project.
	Description *string `json:"description,omitempty"`
	// The id of the project.
	Id *uuid.UUID `json:"id,omitempty"`
	// The time that this project was last updated.
	LastUpdateTime *azuredevops.Time `json:"lastUpdateTime,omitempty"`
	// The name of the project.
	Name *string `json:"name,omitempty"`
	// A set of name-value pairs storing additional property data related to the project.
	Properties *[]ProjectProperty `json:"properties,omitempty"`
	// The current revision of the project.
	Revision *uint64 `json:"revision,omitempty"`
	// The current state of the project.
	State *ProjectState `json:"state,omitempty"`
	// A Uri that can be used to refer to this project.
	Uri *string `json:"uri,omitempty"`
	// The version number of the project.
	Version *uint64 `json:"version,omitempty"`
	// Indicates whom the project is visible to.
	Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

type ProjectMessage struct {
	Project                     *ProjectInfo       `json:"project,omitempty"`
	ProjectChangeType           *ProjectChangeType `json:"projectChangeType,omitempty"`
	ShouldInvalidateSystemStore *bool              `json:"shouldInvalidateSystemStore,omitempty"`
}

type ProjectProperties struct {
	// The team project Id
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
	// The collection of team project properties
	Properties *[]ProjectProperty `json:"properties,omitempty"`
}

// A named value associated with a project.
type ProjectProperty struct {
	// The name of the property.
	Name *string `json:"name,omitempty"`
	// The value of the property.
	Value interface{} `json:"value,omitempty"`
}

type ProjectState string

type projectStateValuesType struct {
	Deleting      ProjectState
	New           ProjectState
	WellFormed    ProjectState
	CreatePending ProjectState
	All           ProjectState
	Unchanged     ProjectState
	Deleted       ProjectState
}

var ProjectStateValues = projectStateValuesType{
	// Project is in the process of being deleted.
	Deleting: "deleting",
	// Project is in the process of being created.
	New: "new",
	// Project is completely created and ready to use.
	WellFormed: "wellFormed",
	// Project has been queued for creation, but the process has not yet started.
	CreatePending: "createPending",
	// All projects regardless of state.
	All: "all",
	// Project has not been changed.
	Unchanged: "unchanged",
	// Project has been deleted.
	Deleted: "deleted",
}

type ProjectVisibility string

type projectVisibilityValuesType struct {
	Private ProjectVisibility
	Public  ProjectVisibility
}

var ProjectVisibilityValues = projectVisibilityValuesType{
	// The project is only visible to users with explicit access.
	Private: "private",
	// The project is visible to all.
	Public: "public",
}

type Proxy struct {
	Authorization *ProxyAuthorization `json:"authorization,omitempty"`
	// This is a description string
	Description *string `json:"description,omitempty"`
	// The friendly name of the server
	FriendlyName  *string `json:"friendlyName,omitempty"`
	GlobalDefault *bool   `json:"globalDefault,omitempty"`
	// This is a string representation of the site that the proxy server is located in (e.g. "NA-WA-RED")
	Site        *string `json:"site,omitempty"`
	SiteDefault *bool   `json:"siteDefault,omitempty"`
	// The URL of the proxy server
	Url *string `json:"url,omitempty"`
}

type ProxyAuthorization struct {
	// Gets or sets the endpoint used to obtain access tokens from the configured token service.
	AuthorizationUrl *string `json:"authorizationUrl,omitempty"`
	// Gets or sets the client identifier for this proxy.
	ClientId *uuid.UUID `json:"clientId,omitempty"`
	// Gets or sets the user identity to authorize for on-prem.
	Identity *string `json:"identity,omitempty"`
	// Gets or sets the public key used to verify the identity of this proxy. Only specify on hosted.
	PublicKey *webapi.PublicKey `json:"publicKey,omitempty"`
}

type SourceControlTypes string

type sourceControlTypesValuesType struct {
	Tfvc SourceControlTypes
	Git  SourceControlTypes
}

var SourceControlTypesValues = sourceControlTypesValuesType{
	Tfvc: "tfvc",
	Git:  "git",
}

// The Team Context for an operation.
type TeamContext struct {
	// The team project Id or name.  Ignored if ProjectId is set.
	Project *string `json:"project,omitempty"`
	// The Team Project ID.  Required if Project is not set.
	ProjectId *uuid.UUID `json:"projectId,omitempty"`
	// The Team Id or name.  Ignored if TeamId is set.
	Team *string `json:"team,omitempty"`
	// The Team Id
	TeamId *uuid.UUID `json:"teamId,omitempty"`
}

// Represents a Team Project object.
type TeamProject struct {
	// Project abbreviation.
	Abbreviation *string `json:"abbreviation,omitempty"`
	// Url to default team identity image.
	DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
	// The project's description (if any).
	Description *string `json:"description,omitempty"`
	// Project identifier.
	Id *uuid.UUID `json:"id,omitempty"`
	// Project last update time.
	LastUpdateTime *azuredevops.Time `json:"lastUpdateTime,omitempty"`
	// Project name.
	Name *string `json:"name,omitempty"`
	// Project revision.
	Revision *uint64 `json:"revision,omitempty"`
	// Project state.
	State *ProjectState `json:"state,omitempty"`
	// Url to the full version of the object.
	Url *string `json:"url,omitempty"`
	// Project visibility.
	Visibility *ProjectVisibility `json:"visibility,omitempty"`
	// The links to other objects related to this object.
	Links interface{} `json:"_links,omitempty"`
	// Set of capabilities this project has (such as process template & version control).
	Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
	// The shallow ref to the default team.
	DefaultTeam *WebApiTeamRef `json:"defaultTeam,omitempty"`
}

// Data contract for a TeamProjectCollection.
type TeamProjectCollection struct {
	// Collection Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// Collection Name.
	Name *string `json:"name,omitempty"`
	// Collection REST Url.
	Url *string `json:"url,omitempty"`
	// The links to other objects related to this object.
	Links interface{} `json:"_links,omitempty"`
	// Project collection description.
	Description *string `json:"description,omitempty"`
	// Process customization type on this collection. It can be Xml or Inherited.
	ProcessCustomizationType *ProcessCustomizationType `json:"processCustomizationType,omitempty"`
	// Project collection state.
	State *string `json:"state,omitempty"`
}

// Reference object for a TeamProjectCollection.
type TeamProjectCollectionReference struct {
	// Collection Id.
	Id *uuid.UUID `json:"id,omitempty"`
	// Collection Name.
	Name *string `json:"name,omitempty"`
	// Collection REST Url.
	Url *string `json:"url,omitempty"`
}

// Represents a shallow reference to a TeamProject.
type TeamProjectReference struct {
	// Project abbreviation.
	Abbreviation *string `json:"abbreviation,omitempty"`
	// Url to default team identity image.
	DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
	// The project's description (if any).
	Description *string `json:"description,omitempty"`
	// Project identifier.
	Id *uuid.UUID `json:"id,omitempty"`
	// Project last update time.
	LastUpdateTime *azuredevops.Time `json:"lastUpdateTime,omitempty"`
	// Project name.
	Name *string `json:"name,omitempty"`
	// Project revision.
	Revision *uint64 `json:"revision,omitempty"`
	// Project state.
	State *ProjectState `json:"state,omitempty"`
	// Url to the full version of the object.
	Url *string `json:"url,omitempty"`
	// Project visibility.
	Visibility *ProjectVisibility `json:"visibility,omitempty"`
}

// A data transfer object that stores the metadata associated with the creation of temporary data.
type TemporaryDataCreatedDTO struct {
	ExpirationSeconds *int              `json:"expirationSeconds,omitempty"`
	Origin            *string           `json:"origin,omitempty"`
	Value             interface{}       `json:"value,omitempty"`
	ExpirationDate    *azuredevops.Time `json:"expirationDate,omitempty"`
	Id                *uuid.UUID        `json:"id,omitempty"`
	Url               *string           `json:"url,omitempty"`
}

// A data transfer object that stores the metadata associated with the temporary data.
type TemporaryDataDTO struct {
	ExpirationSeconds *int        `json:"expirationSeconds,omitempty"`
	Origin            *string     `json:"origin,omitempty"`
	Value             interface{} `json:"value,omitempty"`
}

// Updateable properties for a WebApiTeam.
type UpdateTeam struct {
	// New description for the team.
	Description *string `json:"description,omitempty"`
	// New name for the team.
	Name *string `json:"name,omitempty"`
}

type WebApiConnectedService struct {
	Url *string `json:"url,omitempty"`
	// The user who did the OAuth authentication to created this service
	AuthenticatedBy *webapi.IdentityRef `json:"authenticatedBy,omitempty"`
	// Extra description on the service.
	Description *string `json:"description,omitempty"`
	// Friendly Name of service connection
	FriendlyName *string `json:"friendlyName,omitempty"`
	// Id/Name of the connection service. For Ex: Subscription Id for Azure Connection
	Id *string `json:"id,omitempty"`
	// The kind of service.
	Kind *string `json:"kind,omitempty"`
	// The project associated with this service
	Project *TeamProjectReference `json:"project,omitempty"`
	// Optional uri to connect directly to the service such as https://windows.azure.com
	ServiceUri *string `json:"serviceUri,omitempty"`
}

type WebApiConnectedServiceDetails struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
	// Meta data for service connection
	ConnectedServiceMetaData *WebApiConnectedService `json:"connectedServiceMetaData,omitempty"`
	// Credential info
	CredentialsXml *string `json:"credentialsXml,omitempty"`
	// Optional uri to connect directly to the service such as https://windows.azure.com
	EndPoint *string `json:"endPoint,omitempty"`
}

type WebApiConnectedServiceRef struct {
	Id  *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// The representation of data needed to create a tag definition which is sent across the wire.
type WebApiCreateTagRequestData struct {
	// Name of the tag definition that will be created.
	Name *string `json:"name,omitempty"`
}

type WebApiProject struct {
	// Project abbreviation.
	Abbreviation *string `json:"abbreviation,omitempty"`
	// Url to default team identity image.
	DefaultTeamImageUrl *string `json:"defaultTeamImageUrl,omitempty"`
	// The project's description (if any).
	Description *string `json:"description,omitempty"`
	// Project identifier.
	Id *uuid.UUID `json:"id,omitempty"`
	// Project last update time.
	LastUpdateTime *azuredevops.Time `json:"lastUpdateTime,omitempty"`
	// Project name.
	Name *string `json:"name,omitempty"`
	// Project revision.
	Revision *uint64 `json:"revision,omitempty"`
	// Project state.
	State *ProjectState `json:"state,omitempty"`
	// Url to the full version of the object.
	Url *string `json:"url,omitempty"`
	// Project visibility.
	Visibility *ProjectVisibility `json:"visibility,omitempty"`
	// Set of capabilities this project has
	Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
	// Reference to collection which contains this project
	Collection *WebApiProjectCollectionRef `json:"collection,omitempty"`
	// Default team for this project
	DefaultTeam *WebApiTeamRef `json:"defaultTeam,omitempty"`
}

type WebApiProjectCollection struct {
	// Collection Tfs Url (Host Url)
	CollectionUrl *string `json:"collectionUrl,omitempty"`
	// Collection Guid
	Id *uuid.UUID `json:"id,omitempty"`
	// Collection Name
	Name *string `json:"name,omitempty"`
	// Collection REST Url
	Url *string `json:"url,omitempty"`
	// Project collection description
	Description *string `json:"description,omitempty"`
	// Project collection state
	State *string `json:"state,omitempty"`
}

type WebApiProjectCollectionRef struct {
	// Collection Tfs Url (Host Url)
	CollectionUrl *string `json:"collectionUrl,omitempty"`
	// Collection Guid
	Id *uuid.UUID `json:"id,omitempty"`
	// Collection Name
	Name *string `json:"name,omitempty"`
	// Collection REST Url
	Url *string `json:"url,omitempty"`
}

// The representation of a tag definition which is sent across the wire.
type WebApiTagDefinition struct {
	// Whether or not the tag definition is active.
	Active *bool `json:"active,omitempty"`
	// ID of the tag definition.
	Id *uuid.UUID `json:"id,omitempty"`
	// The name of the tag definition.
	Name *string `json:"name,omitempty"`
	// Resource URL for the Tag Definition.
	Url *string `json:"url,omitempty"`
}

type WebApiTeam struct {
	// Team (Identity) Guid. A Team Foundation ID.
	Id *uuid.UUID `json:"id,omitempty"`
	// Team name
	Name *string `json:"name,omitempty"`
	// Team REST API Url
	Url *string `json:"url,omitempty"`
	// Team description
	Description *string `json:"description,omitempty"`
	// Team identity.
	Identity *identity.Identity `json:"identity,omitempty"`
	// Identity REST API Url to this team
	IdentityUrl *string    `json:"identityUrl,omitempty"`
	ProjectId   *uuid.UUID `json:"projectId,omitempty"`
	ProjectName *string    `json:"projectName,omitempty"`
}

type WebApiTeamRef struct {
	// Team (Identity) Guid. A Team Foundation ID.
	Id *uuid.UUID `json:"id,omitempty"`
	// Team name
	Name *string `json:"name,omitempty"`
	// Team REST API Url
	Url *string `json:"url,omitempty"`
}
