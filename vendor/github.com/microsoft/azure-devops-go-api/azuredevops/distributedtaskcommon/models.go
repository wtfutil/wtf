// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package distributedtaskcommon

type AuthorizationHeader struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// Represents binding of data source for the service endpoint request.
type DataSourceBindingBase struct {
	// Pagination format supported by this data source(ContinuationToken/SkipTop).
	CallbackContextTemplate *string `json:"callbackContextTemplate,omitempty"`
	// Subsequent calls needed?
	CallbackRequiredTemplate *string `json:"callbackRequiredTemplate,omitempty"`
	// Gets or sets the name of the data source.
	DataSourceName *string `json:"dataSourceName,omitempty"`
	// Gets or sets the endpoint Id.
	EndpointId *string `json:"endpointId,omitempty"`
	// Gets or sets the url of the service endpoint.
	EndpointUrl *string `json:"endpointUrl,omitempty"`
	// Gets or sets the authorization headers.
	Headers *[]AuthorizationHeader `json:"headers,omitempty"`
	// Defines the initial value of the query params
	InitialContextTemplate *string `json:"initialContextTemplate,omitempty"`
	// Gets or sets the parameters for the data source.
	Parameters *map[string]string `json:"parameters,omitempty"`
	// Gets or sets http request body
	RequestContent *string `json:"requestContent,omitempty"`
	// Gets or sets http request verb
	RequestVerb *string `json:"requestVerb,omitempty"`
	// Gets or sets the result selector.
	ResultSelector *string `json:"resultSelector,omitempty"`
	// Gets or sets the result template.
	ResultTemplate *string `json:"resultTemplate,omitempty"`
	// Gets or sets the target of the data source.
	Target *string `json:"target,omitempty"`
}

type ProcessParameters struct {
	DataSourceBindings *[]DataSourceBindingBase    `json:"dataSourceBindings,omitempty"`
	Inputs             *[]TaskInputDefinitionBase  `json:"inputs,omitempty"`
	SourceDefinitions  *[]TaskSourceDefinitionBase `json:"sourceDefinitions,omitempty"`
}

type TaskInputDefinitionBase struct {
	Aliases      *[]string            `json:"aliases,omitempty"`
	DefaultValue *string              `json:"defaultValue,omitempty"`
	GroupName    *string              `json:"groupName,omitempty"`
	HelpMarkDown *string              `json:"helpMarkDown,omitempty"`
	Label        *string              `json:"label,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Options      *map[string]string   `json:"options,omitempty"`
	Properties   *map[string]string   `json:"properties,omitempty"`
	Required     *bool                `json:"required,omitempty"`
	Type         *string              `json:"type,omitempty"`
	Validation   *TaskInputValidation `json:"validation,omitempty"`
	VisibleRule  *string              `json:"visibleRule,omitempty"`
}

type TaskInputValidation struct {
	// Conditional expression
	Expression *string `json:"expression,omitempty"`
	// Message explaining how user can correct if validation fails
	Message *string `json:"message,omitempty"`
}

type TaskSourceDefinitionBase struct {
	AuthKey     *string `json:"authKey,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty"`
	KeySelector *string `json:"keySelector,omitempty"`
	Selector    *string `json:"selector,omitempty"`
	Target      *string `json:"target,omitempty"`
}
