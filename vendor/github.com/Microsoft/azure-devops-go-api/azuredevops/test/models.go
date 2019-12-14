// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package test

import (
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/system"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
)

type AbortTestRunRequest struct {
	Options     *int    `json:"options,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
	Revision    *int    `json:"revision,omitempty"`
	TestRunId   *int    `json:"testRunId,omitempty"`
}

type AfnStrip struct {
	// Auxiliary Url to be consumed by MTM
	AuxiliaryUrl *string `json:"auxiliaryUrl,omitempty"`
	// Creation date of the AfnStrip
	CreationDate *azuredevops.Time `json:"creationDate,omitempty"`
	// File name of the attachment created
	FileName *string `json:"fileName,omitempty"`
	// ID of AfnStrip. This is same as the attachment ID.
	Id *int `json:"id,omitempty"`
	// Project identifier which contains AfnStrip
	Project *string `json:"project,omitempty"`
	// Service in which this attachment is stored in
	StoredIn *string `json:"storedIn,omitempty"`
	// Afn strip stream.
	Stream *string `json:"stream,omitempty"`
	// ID of the testcase.
	TestCaseId *int `json:"testCaseId,omitempty"`
	// Backing test result id.
	TestResultId *int `json:"testResultId,omitempty"`
	// Backing test run id.
	TestRunId *int `json:"testRunId,omitempty"`
	// Byte stream (uncompressed) length of Afn strip.
	UnCompressedStreamLength *uint64 `json:"unCompressedStreamLength,omitempty"`
	// Url of the attachment created.
	Url *string `json:"url,omitempty"`
}

type AggregatedDataForResultTrend struct {
	// This is tests execution duration.
	Duration           interface{}                                 `json:"duration,omitempty"`
	ResultsByOutcome   *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
	RunSummaryByState  *map[TestRunState]AggregatedRunsByState     `json:"runSummaryByState,omitempty"`
	TestResultsContext *TestResultsContext                         `json:"testResultsContext,omitempty"`
	TotalTests         *int                                        `json:"totalTests,omitempty"`
}

type AggregatedResultsAnalysis struct {
	Duration                    interface{}                                 `json:"duration,omitempty"`
	NotReportedResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"notReportedResultsByOutcome,omitempty"`
	PreviousContext             *TestResultsContext                         `json:"previousContext,omitempty"`
	ResultsByOutcome            *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
	ResultsDifference           *AggregatedResultsDifference                `json:"resultsDifference,omitempty"`
	RunSummaryByOutcome         *map[TestRunOutcome]AggregatedRunsByOutcome `json:"runSummaryByOutcome,omitempty"`
	RunSummaryByState           *map[TestRunState]AggregatedRunsByState     `json:"runSummaryByState,omitempty"`
	TotalTests                  *int                                        `json:"totalTests,omitempty"`
}

type AggregatedResultsByOutcome struct {
	Count            *int         `json:"count,omitempty"`
	Duration         interface{}  `json:"duration,omitempty"`
	GroupByField     *string      `json:"groupByField,omitempty"`
	GroupByValue     interface{}  `json:"groupByValue,omitempty"`
	Outcome          *TestOutcome `json:"outcome,omitempty"`
	RerunResultCount *int         `json:"rerunResultCount,omitempty"`
}

type AggregatedResultsDifference struct {
	IncreaseInDuration    interface{} `json:"increaseInDuration,omitempty"`
	IncreaseInFailures    *int        `json:"increaseInFailures,omitempty"`
	IncreaseInOtherTests  *int        `json:"increaseInOtherTests,omitempty"`
	IncreaseInPassedTests *int        `json:"increaseInPassedTests,omitempty"`
	IncreaseInTotalTests  *int        `json:"increaseInTotalTests,omitempty"`
}

type AggregatedRunsByOutcome struct {
	Outcome   *TestRunOutcome `json:"outcome,omitempty"`
	RunsCount *int            `json:"runsCount,omitempty"`
}

type AggregatedRunsByState struct {
	ResultsByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsByOutcome,omitempty"`
	RunsCount        *int                                        `json:"runsCount,omitempty"`
	State            *TestRunState                               `json:"state,omitempty"`
}

// The types of test attachments.
type AttachmentType string

type attachmentTypeValuesType struct {
	GeneralAttachment AttachmentType
	CodeCoverage      AttachmentType
	ConsoleLog        AttachmentType
}

var AttachmentTypeValues = attachmentTypeValuesType{
	// Attachment type GeneralAttachment , use this as default type unless you have other type.
	GeneralAttachment: "generalAttachment",
	// Attachment type CodeCoverage.
	CodeCoverage: "codeCoverage",
	// Attachment type ConsoleLog.
	ConsoleLog: "consoleLog",
}

type BatchResponse struct {
	Error     *string     `json:"error,omitempty"`
	Responses *[]Response `json:"responses,omitempty"`
	Status    *string     `json:"status,omitempty"`
}

// BuildConfiguration Details.
type BuildConfiguration struct {
	// Branch name for which build is generated.
	BranchName *string `json:"branchName,omitempty"`
	// BuildDefinitionId for build.
	BuildDefinitionId *int `json:"buildDefinitionId,omitempty"`
	// Build system.
	BuildSystem *string `json:"buildSystem,omitempty"`
	// Build Creation Date.
	CreationDate *azuredevops.Time `json:"creationDate,omitempty"`
	// Build flavor (eg Build/Release).
	Flavor *string `json:"flavor,omitempty"`
	// BuildConfiguration Id.
	Id *int `json:"id,omitempty"`
	// Build Number.
	Number *string `json:"number,omitempty"`
	// BuildConfiguration Platform.
	Platform *string `json:"platform,omitempty"`
	// Project associated with this BuildConfiguration.
	Project *ShallowReference `json:"project,omitempty"`
	// Repository Guid for the Build.
	RepositoryGuid *string `json:"repositoryGuid,omitempty"`
	// Deprecated: Use RepositoryGuid instead
	RepositoryId *int `json:"repositoryId,omitempty"`
	// Repository Type (eg. TFSGit).
	RepositoryType *string `json:"repositoryType,omitempty"`
	// Source Version(/first commit) for the build was triggered.
	SourceVersion *string `json:"sourceVersion,omitempty"`
	// Target BranchName.
	TargetBranchName *string `json:"targetBranchName,omitempty"`
	// Build Uri.
	Uri *string `json:"uri,omitempty"`
}

// Build Coverage Detail
type BuildCoverage struct {
	// Code Coverage File Url
	CodeCoverageFileUrl *string `json:"codeCoverageFileUrl,omitempty"`
	// Build Configuration
	Configuration *BuildConfiguration `json:"configuration,omitempty"`
	// Last Error
	LastError *string `json:"lastError,omitempty"`
	// List of Modules
	Modules *[]ModuleCoverage `json:"modules,omitempty"`
	// State
	State *string `json:"state,omitempty"`
}

// Reference to a build.
type BuildReference struct {
	// Branch name.
	BranchName *string `json:"branchName,omitempty"`
	// Build system.
	BuildSystem *string `json:"buildSystem,omitempty"`
	// Build Definition ID.
	DefinitionId *int `json:"definitionId,omitempty"`
	// Build ID.
	Id *int `json:"id,omitempty"`
	// Build Number.
	Number *string `json:"number,omitempty"`
	// Repository ID.
	RepositoryId *string `json:"repositoryId,omitempty"`
	// Build URI.
	Uri *string `json:"uri,omitempty"`
}

type BuildReference2 struct {
	BranchName           *string           `json:"branchName,omitempty"`
	BuildConfigurationId *int              `json:"buildConfigurationId,omitempty"`
	BuildDefinitionId    *int              `json:"buildDefinitionId,omitempty"`
	BuildDeleted         *bool             `json:"buildDeleted,omitempty"`
	BuildFlavor          *string           `json:"buildFlavor,omitempty"`
	BuildId              *int              `json:"buildId,omitempty"`
	BuildNumber          *string           `json:"buildNumber,omitempty"`
	BuildPlatform        *string           `json:"buildPlatform,omitempty"`
	BuildSystem          *string           `json:"buildSystem,omitempty"`
	BuildUri             *string           `json:"buildUri,omitempty"`
	CoverageId           *int              `json:"coverageId,omitempty"`
	CreatedDate          *azuredevops.Time `json:"createdDate,omitempty"`
	ProjectId            *uuid.UUID        `json:"projectId,omitempty"`
	RepoId               *string           `json:"repoId,omitempty"`
	RepoType             *string           `json:"repoType,omitempty"`
	SourceVersion        *string           `json:"sourceVersion,omitempty"`
}

type BulkResultUpdateRequest struct {
	ProjectName *string                `json:"projectName,omitempty"`
	Requests    *[]ResultUpdateRequest `json:"requests,omitempty"`
}

// Detail About Clone Operation.
type CloneOperationInformation struct {
	// Clone Statistics
	CloneStatistics *CloneStatistics `json:"cloneStatistics,omitempty"`
	// If the operation is complete, the DateTime of completion. If operation is not complete, this is DateTime.MaxValue
	CompletionDate *azuredevops.Time `json:"completionDate,omitempty"`
	// DateTime when the operation was started
	CreationDate *azuredevops.Time `json:"creationDate,omitempty"`
	// Shallow reference of the destination
	DestinationObject *ShallowReference `json:"destinationObject,omitempty"`
	// Shallow reference of the destination
	DestinationPlan *ShallowReference `json:"destinationPlan,omitempty"`
	// Shallow reference of the destination
	DestinationProject *ShallowReference `json:"destinationProject,omitempty"`
	// If the operation has Failed, Message contains the reason for failure. Null otherwise.
	Message *string `json:"message,omitempty"`
	// The ID of the operation
	OpId *int `json:"opId,omitempty"`
	// The type of the object generated as a result of the Clone operation
	ResultObjectType *ResultObjectType `json:"resultObjectType,omitempty"`
	// Shallow reference of the source
	SourceObject *ShallowReference `json:"sourceObject,omitempty"`
	// Shallow reference of the source
	SourcePlan *ShallowReference `json:"sourcePlan,omitempty"`
	// Shallow reference of the source
	SourceProject *ShallowReference `json:"sourceProject,omitempty"`
	// Current state of the operation. When State reaches Succeeded or Failed, the operation is complete
	State *CloneOperationState `json:"state,omitempty"`
	// Url for getting the clone information
	Url *string `json:"url,omitempty"`
}

// Enum of type Clone Operation Type.
type CloneOperationState string

type cloneOperationStateValuesType struct {
	Failed     CloneOperationState
	InProgress CloneOperationState
	Queued     CloneOperationState
	Succeeded  CloneOperationState
}

var CloneOperationStateValues = cloneOperationStateValuesType{
	// value for Failed State
	Failed: "failed",
	// value for Inprogress state
	InProgress: "inProgress",
	// Value for Queued State
	Queued: "queued",
	// value for Success state
	Succeeded: "succeeded",
}

// Clone options for cloning the test suite.
type CloneOptions struct {
	// If set to true requirements will be cloned
	CloneRequirements *bool `json:"cloneRequirements,omitempty"`
	// copy all suites from a source plan
	CopyAllSuites *bool `json:"copyAllSuites,omitempty"`
	// copy ancestor hierarchy
	CopyAncestorHierarchy *bool `json:"copyAncestorHierarchy,omitempty"`
	// Name of the workitem type of the clone
	DestinationWorkItemType *string `json:"destinationWorkItemType,omitempty"`
	// Key value pairs where the key value is overridden by the value.
	OverrideParameters *map[string]string `json:"overrideParameters,omitempty"`
	// Comment on the link that will link the new clone  test case to the original Set null for no comment
	RelatedLinkComment *string `json:"relatedLinkComment,omitempty"`
}

// Clone Statistics Details.
type CloneStatistics struct {
	// Number of requirements cloned so far.
	ClonedRequirementsCount *int `json:"clonedRequirementsCount,omitempty"`
	// Number of shared steps cloned so far.
	ClonedSharedStepsCount *int `json:"clonedSharedStepsCount,omitempty"`
	// Number of test cases cloned so far
	ClonedTestCasesCount *int `json:"clonedTestCasesCount,omitempty"`
	// Total number of requirements to be cloned
	TotalRequirementsCount *int `json:"totalRequirementsCount,omitempty"`
	// Total number of test cases to be cloned
	TotalTestCasesCount *int `json:"totalTestCasesCount,omitempty"`
}

// Represents the build configuration (platform, flavor) and coverage data for the build
type CodeCoverageData struct {
	// Flavor of build for which data is retrieved/published
	BuildFlavor *string `json:"buildFlavor,omitempty"`
	// Platform of build for which data is retrieved/published
	BuildPlatform *string `json:"buildPlatform,omitempty"`
	// List of coverage data for the build
	CoverageStats *[]CodeCoverageStatistics `json:"coverageStats,omitempty"`
}

// Represents the code coverage statistics for a particular coverage label (modules, statements, blocks, etc.)
type CodeCoverageStatistics struct {
	// Covered units
	Covered *int `json:"covered,omitempty"`
	// Delta of coverage
	Delta *float64 `json:"delta,omitempty"`
	// Is delta valid
	IsDeltaAvailable *bool `json:"isDeltaAvailable,omitempty"`
	// Label of coverage data ("Blocks", "Statements", "Modules", etc.)
	Label *string `json:"label,omitempty"`
	// Position of label
	Position *int `json:"position,omitempty"`
	// Total units
	Total *int `json:"total,omitempty"`
}

// Represents the code coverage summary results Used to publish or retrieve code coverage summary against a build
type CodeCoverageSummary struct {
	// Uri of build for which data is retrieved/published
	Build *ShallowReference `json:"build,omitempty"`
	// List of coverage data and details for the build
	CoverageData *[]CodeCoverageData `json:"coverageData,omitempty"`
	// Uri of build against which difference in coverage is computed
	DeltaBuild *ShallowReference `json:"deltaBuild,omitempty"`
	// Uri of build against which difference in coverage is computed
	Status *CoverageSummaryStatus `json:"status,omitempty"`
}

type CodeCoverageSummary2 struct {
	BuildConfigurationId *int       `json:"buildConfigurationId,omitempty"`
	Covered              *int       `json:"covered,omitempty"`
	Label                *string    `json:"label,omitempty"`
	Position             *int       `json:"position,omitempty"`
	ProjectId            *uuid.UUID `json:"projectId,omitempty"`
	Total                *int       `json:"total,omitempty"`
}

type Coverage2 struct {
	CoverageId   *int              `json:"coverageId,omitempty"`
	DateCreated  *azuredevops.Time `json:"dateCreated,omitempty"`
	DateModified *azuredevops.Time `json:"dateModified,omitempty"`
	LastError    *string           `json:"lastError,omitempty"`
	State        *byte             `json:"state,omitempty"`
}

// [Flags] Used to choose which coverage data is returned by a QueryXXXCoverage() call.
type CoverageQueryFlags string

type coverageQueryFlagsValuesType struct {
	Modules   CoverageQueryFlags
	Functions CoverageQueryFlags
	BlockData CoverageQueryFlags
}

var CoverageQueryFlagsValues = coverageQueryFlagsValuesType{
	// If set, the Coverage.Modules property will be populated.
	Modules: "modules",
	// If set, the ModuleCoverage.Functions properties will be populated.
	Functions: "functions",
	// If set, the ModuleCoverage.CoverageData field will be populated.
	BlockData: "blockData",
}

type CoverageStatistics struct {
	BlocksCovered         *int `json:"blocksCovered,omitempty"`
	BlocksNotCovered      *int `json:"blocksNotCovered,omitempty"`
	LinesCovered          *int `json:"linesCovered,omitempty"`
	LinesNotCovered       *int `json:"linesNotCovered,omitempty"`
	LinesPartiallyCovered *int `json:"linesPartiallyCovered,omitempty"`
}

type CoverageStatus string

type coverageStatusValuesType struct {
	Covered          CoverageStatus
	NotCovered       CoverageStatus
	PartiallyCovered CoverageStatus
}

var CoverageStatusValues = coverageStatusValuesType{
	Covered:          "covered",
	NotCovered:       "notCovered",
	PartiallyCovered: "partiallyCovered",
}

// Represents status of code coverage summary for a build
type CoverageSummaryStatus string

type coverageSummaryStatusValuesType struct {
	None       CoverageSummaryStatus
	InProgress CoverageSummaryStatus
	Completed  CoverageSummaryStatus
	Finalized  CoverageSummaryStatus
	Pending    CoverageSummaryStatus
}

var CoverageSummaryStatusValues = coverageSummaryStatusValuesType{
	// No coverage status
	None: "none",
	// The summary evaluation is in progress
	InProgress: "inProgress",
	// The summary evaluation for the previous request is completed. Summary can change in future
	Completed: "completed",
	// The summary evaluation is finalized and won't change
	Finalized: "finalized",
	// The summary evaluation is pending
	Pending: "pending",
}

type CreateTestMessageLogEntryRequest struct {
	ProjectName         *string                `json:"projectName,omitempty"`
	TestMessageLogEntry *[]TestMessageLogEntry `json:"testMessageLogEntry,omitempty"`
	TestRunId           *int                   `json:"testRunId,omitempty"`
}

type CreateTestResultsRequest struct {
	ProjectName *string                 `json:"projectName,omitempty"`
	Results     *[]LegacyTestCaseResult `json:"results,omitempty"`
}

type CreateTestRunRequest struct {
	ProjectName  *string                 `json:"projectName,omitempty"`
	Results      *[]LegacyTestCaseResult `json:"results,omitempty"`
	TestRun      *LegacyTestRun          `json:"testRun,omitempty"`
	TestSettings *LegacyTestSettings     `json:"testSettings,omitempty"`
}

// A custom field information. Allowed Key : Value pairs - ( AttemptId: int value, IsTestResultFlaky: bool)
type CustomTestField struct {
	// Field Name.
	FieldName *string `json:"fieldName,omitempty"`
	// Field value.
	Value interface{} `json:"value,omitempty"`
}

type CustomTestFieldDefinition struct {
	FieldId   *int                  `json:"fieldId,omitempty"`
	FieldName *string               `json:"fieldName,omitempty"`
	FieldType *CustomTestFieldType  `json:"fieldType,omitempty"`
	Scope     *CustomTestFieldScope `json:"scope,omitempty"`
}

// [Flags]
type CustomTestFieldScope string

type customTestFieldScopeValuesType struct {
	None       CustomTestFieldScope
	TestRun    CustomTestFieldScope
	TestResult CustomTestFieldScope
	System     CustomTestFieldScope
	All        CustomTestFieldScope
}

var CustomTestFieldScopeValues = customTestFieldScopeValuesType{
	None:       "none",
	TestRun:    "testRun",
	TestResult: "testResult",
	System:     "system",
	All:        "all",
}

type CustomTestFieldType string

type customTestFieldTypeValuesType struct {
	Bit      CustomTestFieldType
	DateTime CustomTestFieldType
	Int      CustomTestFieldType
	Float    CustomTestFieldType
	String   CustomTestFieldType
	Guid     CustomTestFieldType
}

var CustomTestFieldTypeValues = customTestFieldTypeValuesType{
	Bit:      "bit",
	DateTime: "dateTime",
	Int:      "int",
	Float:    "float",
	String:   "string",
	Guid:     "guid",
}

type DatedTestFieldData struct {
	Date  *azuredevops.Time `json:"date,omitempty"`
	Value *TestFieldData    `json:"value,omitempty"`
}

type DefaultAfnStripBinding struct {
	TestCaseId   *int `json:"testCaseId,omitempty"`
	TestResultId *int `json:"testResultId,omitempty"`
	TestRunId    *int `json:"testRunId,omitempty"`
}

type DeleteTestRunRequest struct {
	ProjectName *string `json:"projectName,omitempty"`
	TestRunIds  *[]int  `json:"testRunIds,omitempty"`
}

type DownloadAttachmentsRequest struct {
	Ids     *[]int    `json:"ids,omitempty"`
	Lengths *[]uint64 `json:"lengths,omitempty"`
}

// This is a temporary class to provide the details for the test run environment.
type DtlEnvironmentDetails struct {
	CsmContent       *string `json:"csmContent,omitempty"`
	CsmParameters    *string `json:"csmParameters,omitempty"`
	SubscriptionName *string `json:"subscriptionName,omitempty"`
}

// Failing since information of a test result.
type FailingSince struct {
	// Build reference since failing.
	Build *BuildReference `json:"build,omitempty"`
	// Time since failing.
	Date *azuredevops.Time `json:"date,omitempty"`
	// Release reference since failing.
	Release *ReleaseReference `json:"release,omitempty"`
}

type FetchTestResultsRequest struct {
	IdAndRevs            *[]TestCaseResultIdAndRev `json:"idAndRevs,omitempty"`
	IncludeActionResults *bool                     `json:"includeActionResults,omitempty"`
	ProjectName          *string                   `json:"projectName,omitempty"`
}

type FetchTestResultsResponse struct {
	ActionResults  *[]TestActionResult               `json:"actionResults,omitempty"`
	Attachments    *[]TestResultAttachment           `json:"attachments,omitempty"`
	DeletedIds     *[]LegacyTestCaseResultIdentifier `json:"deletedIds,omitempty"`
	Results        *[]LegacyTestCaseResult           `json:"results,omitempty"`
	TestParameters *[]TestResultParameter            `json:"testParameters,omitempty"`
}

type FieldDetailsForTestResults struct {
	// Group by field name
	FieldName *string `json:"fieldName,omitempty"`
	// Group by field values
	GroupsForField *[]interface{} `json:"groupsForField,omitempty"`
}

type FileCoverage struct {
	// List of line blocks along with their coverage status
	LineBlocksCoverage *[]LineBlockCoverage `json:"lineBlocksCoverage,omitempty"`
	// File path for which coverage information is sought for
	Path *string `json:"path,omitempty"`
}

type FileCoverageRequest struct {
	FilePath                   *string `json:"filePath,omitempty"`
	PullRequestBaseIterationId *int    `json:"pullRequestBaseIterationId,omitempty"`
	PullRequestId              *int    `json:"pullRequestId,omitempty"`
	PullRequestIterationId     *int    `json:"pullRequestIterationId,omitempty"`
	RepoId                     *string `json:"repoId,omitempty"`
}

type FilterPointQuery struct {
	PlanId       *int    `json:"planId,omitempty"`
	PointIds     *[]int  `json:"pointIds,omitempty"`
	PointOutcome *[]byte `json:"pointOutcome,omitempty"`
	ResultState  *[]byte `json:"resultState,omitempty"`
}

type FlakyDetection struct {
	// FlakyDetectionPipelines defines Pipelines for Detection.
	FlakyDetectionPipelines *FlakyDetectionPipelines `json:"flakyDetectionPipelines,omitempty"`
	// FlakyDetectionType defines Detection type i.e. 1. System or 2. Manual.
	FlakyDetectionType *FlakyDetectionType `json:"flakyDetectionType,omitempty"`
}

type FlakyDetectionPipelines struct {
	// AllowedPipelines - List All Pipelines allowed for detection.
	AllowedPipelines *[]int `json:"allowedPipelines,omitempty"`
	// IsAllPipelinesAllowed if users configure all system's pipelines.
	IsAllPipelinesAllowed *bool `json:"isAllPipelinesAllowed,omitempty"`
}

type FlakyDetectionType string

type flakyDetectionTypeValuesType struct {
	Custom FlakyDetectionType
	System FlakyDetectionType
}

var FlakyDetectionTypeValues = flakyDetectionTypeValuesType{
	// Custom defines manual detection type.
	Custom: "custom",
	// Defines System detection type.
	System: "system",
}

type FlakySettings struct {
	// FlakyDetection defines types of detection.
	FlakyDetection *FlakyDetection `json:"flakyDetection,omitempty"`
	// FlakyInSummaryReport defines flaky data should show in summary report or not.
	FlakyInSummaryReport *bool `json:"flakyInSummaryReport,omitempty"`
	// ManualMarkUnmarkFlaky defines manual marking unmarking of flaky testcase.
	ManualMarkUnmarkFlaky *bool `json:"manualMarkUnmarkFlaky,omitempty"`
}

type FunctionCoverage struct {
	Class      *string             `json:"class,omitempty"`
	Name       *string             `json:"name,omitempty"`
	Namespace  *string             `json:"namespace,omitempty"`
	SourceFile *string             `json:"sourceFile,omitempty"`
	Statistics *CoverageStatistics `json:"statistics,omitempty"`
}

type FunctionCoverage2 struct {
	BlocksCovered         *int    `json:"blocksCovered,omitempty"`
	BlocksNotCovered      *int    `json:"blocksNotCovered,omitempty"`
	Class                 *string `json:"class,omitempty"`
	CoverageId            *int    `json:"coverageId,omitempty"`
	FunctionId            *int    `json:"functionId,omitempty"`
	LinesCovered          *int    `json:"linesCovered,omitempty"`
	LinesNotCovered       *int    `json:"linesNotCovered,omitempty"`
	LinesPartiallyCovered *int    `json:"linesPartiallyCovered,omitempty"`
	ModuleId              *int    `json:"moduleId,omitempty"`
	Name                  *string `json:"name,omitempty"`
	Namespace             *string `json:"namespace,omitempty"`
	SourceFile            *string `json:"sourceFile,omitempty"`
}

type HttpPostedTcmAttachment struct {
	AttachmentContent *string `json:"attachmentContent,omitempty"`
	ContentLength     *int    `json:"contentLength,omitempty"`
	ContentType       *string `json:"contentType,omitempty"`
	FileName          *string `json:"fileName,omitempty"`
}

// Job in pipeline. This is related to matrixing in YAML.
type JobReference struct {
	// Attempt number of the job
	Attempt *int `json:"attempt,omitempty"`
	// Matrixing in YAML generates copies of a job with different inputs in matrix. JobName is the name of those input. Maximum supported length for name is 256 character.
	JobName *string `json:"jobName,omitempty"`
}

// Last result details of test point.
type LastResultDetails struct {
	// CompletedDate of LastResult.
	DateCompleted *azuredevops.Time `json:"dateCompleted,omitempty"`
	// Duration of LastResult.
	Duration *uint64 `json:"duration,omitempty"`
	// RunBy.
	RunBy *webapi.IdentityRef `json:"runBy,omitempty"`
}

type LegacyBuildConfiguration struct {
	BranchName              *string           `json:"branchName,omitempty"`
	BuildConfigurationId    *int              `json:"buildConfigurationId,omitempty"`
	BuildDefinitionId       *int              `json:"buildDefinitionId,omitempty"`
	BuildDefinitionName     *string           `json:"buildDefinitionName,omitempty"`
	BuildFlavor             *string           `json:"buildFlavor,omitempty"`
	BuildId                 *int              `json:"buildId,omitempty"`
	BuildNumber             *string           `json:"buildNumber,omitempty"`
	BuildPlatform           *string           `json:"buildPlatform,omitempty"`
	BuildQuality            *string           `json:"buildQuality,omitempty"`
	BuildSystem             *string           `json:"buildSystem,omitempty"`
	BuildUri                *string           `json:"buildUri,omitempty"`
	CompletedDate           *azuredevops.Time `json:"completedDate,omitempty"`
	CreatedDate             *azuredevops.Time `json:"createdDate,omitempty"`
	OldBuildConfigurationId *int              `json:"oldBuildConfigurationId,omitempty"`
	RepositoryId            *string           `json:"repositoryId,omitempty"`
	RepositoryType          *string           `json:"repositoryType,omitempty"`
	SourceVersion           *string           `json:"sourceVersion,omitempty"`
	TeamProjectName         *string           `json:"teamProjectName,omitempty"`
}

type LegacyReleaseReference struct {
	Attempt                  *int              `json:"attempt,omitempty"`
	EnvironmentCreationDate  *azuredevops.Time `json:"environmentCreationDate,omitempty"`
	PrimaryArtifactBuildId   *int              `json:"primaryArtifactBuildId,omitempty"`
	PrimaryArtifactProjectId *string           `json:"primaryArtifactProjectId,omitempty"`
	PrimaryArtifactType      *string           `json:"primaryArtifactType,omitempty"`
	ReleaseCreationDate      *azuredevops.Time `json:"releaseCreationDate,omitempty"`
	ReleaseDefId             *int              `json:"releaseDefId,omitempty"`
	ReleaseEnvDefId          *int              `json:"releaseEnvDefId,omitempty"`
	ReleaseEnvId             *int              `json:"releaseEnvId,omitempty"`
	ReleaseEnvName           *string           `json:"releaseEnvName,omitempty"`
	ReleaseEnvUri            *string           `json:"releaseEnvUri,omitempty"`
	ReleaseId                *int              `json:"releaseId,omitempty"`
	ReleaseName              *string           `json:"releaseName,omitempty"`
	ReleaseRefId             *int              `json:"releaseRefId,omitempty"`
	ReleaseUri               *string           `json:"releaseUri,omitempty"`
}

type LegacyTestCaseResult struct {
	AfnStripId           *int                            `json:"afnStripId,omitempty"`
	AreaId               *int                            `json:"areaId,omitempty"`
	AreaUri              *string                         `json:"areaUri,omitempty"`
	AutomatedTestId      *string                         `json:"automatedTestId,omitempty"`
	AutomatedTestName    *string                         `json:"automatedTestName,omitempty"`
	AutomatedTestStorage *string                         `json:"automatedTestStorage,omitempty"`
	AutomatedTestType    *string                         `json:"automatedTestType,omitempty"`
	AutomatedTestTypeId  *string                         `json:"automatedTestTypeId,omitempty"`
	BuildNumber          *string                         `json:"buildNumber,omitempty"`
	BuildReference       *LegacyBuildConfiguration       `json:"buildReference,omitempty"`
	Comment              *string                         `json:"comment,omitempty"`
	ComputerName         *string                         `json:"computerName,omitempty"`
	ConfigurationId      *int                            `json:"configurationId,omitempty"`
	ConfigurationName    *string                         `json:"configurationName,omitempty"`
	CreationDate         *azuredevops.Time               `json:"creationDate,omitempty"`
	CustomFields         *[]TestExtensionField           `json:"customFields,omitempty"`
	DateCompleted        *azuredevops.Time               `json:"dateCompleted,omitempty"`
	DateStarted          *azuredevops.Time               `json:"dateStarted,omitempty"`
	Duration             *uint64                         `json:"duration,omitempty"`
	ErrorMessage         *string                         `json:"errorMessage,omitempty"`
	FailingSince         *FailingSince                   `json:"failingSince,omitempty"`
	FailureType          *byte                           `json:"failureType,omitempty"`
	Id                   *LegacyTestCaseResultIdentifier `json:"id,omitempty"`
	IsRerun              *bool                           `json:"isRerun,omitempty"`
	LastUpdated          *azuredevops.Time               `json:"lastUpdated,omitempty"`
	LastUpdatedBy        *uuid.UUID                      `json:"lastUpdatedBy,omitempty"`
	LastUpdatedByName    *string                         `json:"lastUpdatedByName,omitempty"`
	Outcome              *byte                           `json:"outcome,omitempty"`
	Owner                *uuid.UUID                      `json:"owner,omitempty"`
	OwnerName            *string                         `json:"ownerName,omitempty"`
	Priority             *byte                           `json:"priority,omitempty"`
	ReleaseReference     *LegacyReleaseReference         `json:"releaseReference,omitempty"`
	ResetCount           *int                            `json:"resetCount,omitempty"`
	ResolutionStateId    *int                            `json:"resolutionStateId,omitempty"`
	ResultGroupType      *ResultGroupType                `json:"resultGroupType,omitempty"`
	Revision             *int                            `json:"revision,omitempty"`
	RunBy                *uuid.UUID                      `json:"runBy,omitempty"`
	RunByName            *string                         `json:"runByName,omitempty"`
	SequenceId           *int                            `json:"sequenceId,omitempty"`
	StackTrace           *TestExtensionField             `json:"stackTrace,omitempty"`
	State                *byte                           `json:"state,omitempty"`
	SubResultCount       *int                            `json:"subResultCount,omitempty"`
	SuiteName            *string                         `json:"suiteName,omitempty"`
	TestCaseArea         *string                         `json:"testCaseArea,omitempty"`
	TestCaseAreaUri      *string                         `json:"testCaseAreaUri,omitempty"`
	TestCaseId           *int                            `json:"testCaseId,omitempty"`
	TestCaseReferenceId  *int                            `json:"testCaseReferenceId,omitempty"`
	TestCaseRevision     *int                            `json:"testCaseRevision,omitempty"`
	TestCaseTitle        *string                         `json:"testCaseTitle,omitempty"`
	TestPlanId           *int                            `json:"testPlanId,omitempty"`
	TestPointId          *int                            `json:"testPointId,omitempty"`
	TestResultId         *int                            `json:"testResultId,omitempty"`
	TestRunId            *int                            `json:"testRunId,omitempty"`
	TestRunTitle         *string                         `json:"testRunTitle,omitempty"`
	TestSuiteId          *int                            `json:"testSuiteId,omitempty"`
}

type LegacyTestCaseResultIdentifier struct {
	AreaUri      *string `json:"areaUri,omitempty"`
	TestResultId *int    `json:"testResultId,omitempty"`
	TestRunId    *int    `json:"testRunId,omitempty"`
}

type LegacyTestRun struct {
	BugsCount                 *int                      `json:"bugsCount,omitempty"`
	BuildConfigurationId      *int                      `json:"buildConfigurationId,omitempty"`
	BuildFlavor               *string                   `json:"buildFlavor,omitempty"`
	BuildNumber               *string                   `json:"buildNumber,omitempty"`
	BuildPlatform             *string                   `json:"buildPlatform,omitempty"`
	BuildReference            *LegacyBuildConfiguration `json:"buildReference,omitempty"`
	BuildUri                  *string                   `json:"buildUri,omitempty"`
	Comment                   *string                   `json:"comment,omitempty"`
	CompleteDate              *azuredevops.Time         `json:"completeDate,omitempty"`
	ConfigurationIds          *[]int                    `json:"configurationIds,omitempty"`
	Controller                *string                   `json:"controller,omitempty"`
	CreationDate              *azuredevops.Time         `json:"creationDate,omitempty"`
	CsmContent                *string                   `json:"csmContent,omitempty"`
	CsmParameters             *string                   `json:"csmParameters,omitempty"`
	CustomFields              *[]TestExtensionField     `json:"customFields,omitempty"`
	DropLocation              *string                   `json:"dropLocation,omitempty"`
	DtlAutEnvironment         *ShallowReference         `json:"dtlAutEnvironment,omitempty"`
	DtlTestEnvironment        *ShallowReference         `json:"dtlTestEnvironment,omitempty"`
	DueDate                   *azuredevops.Time         `json:"dueDate,omitempty"`
	ErrorMessage              *string                   `json:"errorMessage,omitempty"`
	Filter                    *RunFilter                `json:"filter,omitempty"`
	IncompleteTests           *int                      `json:"incompleteTests,omitempty"`
	IsAutomated               *bool                     `json:"isAutomated,omitempty"`
	IsBvt                     *bool                     `json:"isBvt,omitempty"`
	Iteration                 *string                   `json:"iteration,omitempty"`
	IterationId               *int                      `json:"iterationId,omitempty"`
	LastUpdated               *azuredevops.Time         `json:"lastUpdated,omitempty"`
	LastUpdatedBy             *uuid.UUID                `json:"lastUpdatedBy,omitempty"`
	LastUpdatedByName         *string                   `json:"lastUpdatedByName,omitempty"`
	LegacySharePath           *string                   `json:"legacySharePath,omitempty"`
	NotApplicableTests        *int                      `json:"notApplicableTests,omitempty"`
	Owner                     *uuid.UUID                `json:"owner,omitempty"`
	OwnerName                 *string                   `json:"ownerName,omitempty"`
	PassedTests               *int                      `json:"passedTests,omitempty"`
	PostProcessState          *byte                     `json:"postProcessState,omitempty"`
	PublicTestSettingsId      *int                      `json:"publicTestSettingsId,omitempty"`
	ReleaseEnvironmentUri     *string                   `json:"releaseEnvironmentUri,omitempty"`
	ReleaseReference          *LegacyReleaseReference   `json:"releaseReference,omitempty"`
	ReleaseUri                *string                   `json:"releaseUri,omitempty"`
	Revision                  *int                      `json:"revision,omitempty"`
	RowVersion                *[]byte                   `json:"rowVersion,omitempty"`
	RunHasDtlEnvironment      *bool                     `json:"runHasDtlEnvironment,omitempty"`
	RunTimeout                interface{}               `json:"runTimeout,omitempty"`
	ServiceVersion            *string                   `json:"serviceVersion,omitempty"`
	SourceWorkflow            *string                   `json:"sourceWorkflow,omitempty"`
	StartDate                 *azuredevops.Time         `json:"startDate,omitempty"`
	State                     *byte                     `json:"state,omitempty"`
	SubscriptionName          *string                   `json:"subscriptionName,omitempty"`
	Substate                  *byte                     `json:"substate,omitempty"`
	TeamProject               *string                   `json:"teamProject,omitempty"`
	TeamProjectUri            *string                   `json:"teamProjectUri,omitempty"`
	TestConfigurationsMapping *string                   `json:"testConfigurationsMapping,omitempty"`
	TestEnvironmentId         *uuid.UUID                `json:"testEnvironmentId,omitempty"`
	TestMessageLogEntries     *[]TestMessageLogDetails  `json:"testMessageLogEntries,omitempty"`
	TestMessageLogId          *int                      `json:"testMessageLogId,omitempty"`
	TestPlanId                *int                      `json:"testPlanId,omitempty"`
	TestRunId                 *int                      `json:"testRunId,omitempty"`
	TestRunStatistics         *[]LegacyTestRunStatistic `json:"testRunStatistics,omitempty"`
	TestSettingsId            *int                      `json:"testSettingsId,omitempty"`
	Title                     *string                   `json:"title,omitempty"`
	TotalTests                *int                      `json:"totalTests,omitempty"`
	Type                      *byte                     `json:"type,omitempty"`
	UnanalyzedTests           *int                      `json:"unanalyzedTests,omitempty"`
	Version                   *int                      `json:"version,omitempty"`
}

type LegacyTestRunStatistic struct {
	Count           *int                 `json:"count,omitempty"`
	Outcome         *byte                `json:"outcome,omitempty"`
	ResolutionState *TestResolutionState `json:"resolutionState,omitempty"`
	State           *byte                `json:"state,omitempty"`
	TestRunId       *int                 `json:"testRunId,omitempty"`
}

type LegacyTestSettings struct {
	AreaId            *int                       `json:"areaId,omitempty"`
	AreaPath          *string                    `json:"areaPath,omitempty"`
	CreatedBy         *uuid.UUID                 `json:"createdBy,omitempty"`
	CreatedByName     *string                    `json:"createdByName,omitempty"`
	CreatedDate       *azuredevops.Time          `json:"createdDate,omitempty"`
	Description       *string                    `json:"description,omitempty"`
	Id                *int                       `json:"id,omitempty"`
	IsAutomated       *bool                      `json:"isAutomated,omitempty"`
	IsPublic          *bool                      `json:"isPublic,omitempty"`
	LastUpdated       *azuredevops.Time          `json:"lastUpdated,omitempty"`
	LastUpdatedBy     *uuid.UUID                 `json:"lastUpdatedBy,omitempty"`
	LastUpdatedByName *string                    `json:"lastUpdatedByName,omitempty"`
	MachineRoles      *[]TestSettingsMachineRole `json:"machineRoles,omitempty"`
	Name              *string                    `json:"name,omitempty"`
	Revision          *int                       `json:"revision,omitempty"`
	Settings          *string                    `json:"settings,omitempty"`
	TeamProjectUri    *string                    `json:"teamProjectUri,omitempty"`
}

type LineBlockCoverage struct {
	// End of line block
	End *int `json:"end,omitempty"`
	// Start of line block
	Start *int `json:"start,omitempty"`
	// Coverage status. Covered: 0, NotCovered: 1,  PartiallyCovered: 2
	Status *int `json:"status,omitempty"`
}

type LinkedWorkItemsQuery struct {
	AutomatedTestNames *[]string `json:"automatedTestNames,omitempty"`
	PlanId             *int      `json:"planId,omitempty"`
	PointIds           *[]int    `json:"pointIds,omitempty"`
	SuiteIds           *[]int    `json:"suiteIds,omitempty"`
	TestCaseIds        *[]int    `json:"testCaseIds,omitempty"`
	WorkItemCategory   *string   `json:"workItemCategory,omitempty"`
}

type LinkedWorkItemsQueryResult struct {
	AutomatedTestName *string              `json:"automatedTestName,omitempty"`
	PlanId            *int                 `json:"planId,omitempty"`
	PointId           *int                 `json:"pointId,omitempty"`
	SuiteId           *int                 `json:"suiteId,omitempty"`
	TestCaseId        *int                 `json:"testCaseId,omitempty"`
	WorkItems         *[]WorkItemReference `json:"workItems,omitempty"`
}

type ModuleCoverage struct {
	BlockCount *int    `json:"blockCount,omitempty"`
	BlockData  *[]byte `json:"blockData,omitempty"`
	// Code Coverage File Url
	FileUrl      *string             `json:"fileUrl,omitempty"`
	Functions    *[]FunctionCoverage `json:"functions,omitempty"`
	Name         *string             `json:"name,omitempty"`
	Signature    *uuid.UUID          `json:"signature,omitempty"`
	SignatureAge *int                `json:"signatureAge,omitempty"`
	Statistics   *CoverageStatistics `json:"statistics,omitempty"`
}

type ModuleCoverage2 struct {
	BlockCount            *int       `json:"blockCount,omitempty"`
	BlockData             *[]byte    `json:"blockData,omitempty"`
	BlockDataLength       *int       `json:"blockDataLength,omitempty"`
	BlocksCovered         *int       `json:"blocksCovered,omitempty"`
	BlocksNotCovered      *int       `json:"blocksNotCovered,omitempty"`
	CoverageFileUrl       *string    `json:"coverageFileUrl,omitempty"`
	CoverageId            *int       `json:"coverageId,omitempty"`
	LinesCovered          *int       `json:"linesCovered,omitempty"`
	LinesNotCovered       *int       `json:"linesNotCovered,omitempty"`
	LinesPartiallyCovered *int       `json:"linesPartiallyCovered,omitempty"`
	ModuleId              *int       `json:"moduleId,omitempty"`
	Name                  *string    `json:"name,omitempty"`
	Signature             *uuid.UUID `json:"signature,omitempty"`
	SignatureAge          *int       `json:"signatureAge,omitempty"`
}

// Name value pair
type NameValuePair struct {
	// Name
	Name *string `json:"name,omitempty"`
	// Value
	Value *string `json:"value,omitempty"`
}

type OperationType string

type operationTypeValuesType struct {
	Add    OperationType
	Delete OperationType
}

var OperationTypeValues = operationTypeValuesType{
	Add:    "add",
	Delete: "delete",
}

// Phase in pipeline
type PhaseReference struct {
	// Attempt number of the phase
	Attempt *int `json:"attempt,omitempty"`
	// Name of the phase. Maximum supported length for name is 256 character.
	PhaseName *string `json:"phaseName,omitempty"`
}

// Pipeline reference
type PipelineReference struct {
	// Reference of the job
	JobReference *JobReference `json:"jobReference,omitempty"`
	// Reference of the phase.
	PhaseReference *PhaseReference `json:"phaseReference,omitempty"`
	// Reference of the pipeline with which this pipeline instance is related.
	PipelineId *int `json:"pipelineId,omitempty"`
	// Reference of the stage.
	StageReference *StageReference `json:"stageReference,omitempty"`
}

// A model class used for creating and updating test plans.
type PlanUpdateModel struct {
	// Area path to which the test plan belongs. This should be set to area path of the team that works on this test plan.
	Area *ShallowReference `json:"area,omitempty"`
	// Build ID of the build whose quality is tested by the tests in this test plan. For automated testing, this build ID is used to find the test binaries that contain automated test methods.
	Build *ShallowReference `json:"build,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *ShallowReference `json:"buildDefinition,omitempty"`
	// IDs of configurations to be applied when new test suites and test cases are added to the test plan.
	ConfigurationIds *[]int `json:"configurationIds,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *string `json:"endDate,omitempty"`
	// Iteration path assigned to the test plan. This indicates when the target iteration by which the testing in this plan is supposed to be complete and the product is ready to be released.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Start date for the test plan.
	StartDate *string `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Test Outcome settings
	TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
}

// Adding test cases to a suite creates one of more test points based on the default configurations and testers assigned to the test suite. PointAssignment is the list of test points that were created for each of the test cases that were added to the test suite.
type PointAssignment struct {
	// Configuration that was assigned to the test case.
	Configuration *ShallowReference `json:"configuration,omitempty"`
	// Tester that was assigned to the test case
	Tester *webapi.IdentityRef `json:"tester,omitempty"`
}

type PointLastResult struct {
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	PointId         *int              `json:"pointId,omitempty"`
}

// Filter class for test point.
type PointsFilter struct {
	// List of Configurations for filtering.
	ConfigurationNames *[]string `json:"configurationNames,omitempty"`
	// List of test case id for filtering.
	TestcaseIds *[]int `json:"testcaseIds,omitempty"`
	// List of tester for filtering.
	Testers *[]webapi.IdentityRef `json:"testers,omitempty"`
}

type PointsReference2 struct {
	PlanId  *int `json:"planId,omitempty"`
	PointId *int `json:"pointId,omitempty"`
}

type PointsResults2 struct {
	ChangeNumber          *int              `json:"changeNumber,omitempty"`
	LastFailureType       *byte             `json:"lastFailureType,omitempty"`
	LastResolutionStateId *int              `json:"lastResolutionStateId,omitempty"`
	LastResultOutcome     *byte             `json:"lastResultOutcome,omitempty"`
	LastResultState       *byte             `json:"lastResultState,omitempty"`
	LastTestResultId      *int              `json:"lastTestResultId,omitempty"`
	LastTestRunId         *int              `json:"lastTestRunId,omitempty"`
	LastUpdated           *azuredevops.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy         *uuid.UUID        `json:"lastUpdatedBy,omitempty"`
	PlanId                *int              `json:"planId,omitempty"`
	PointId               *int              `json:"pointId,omitempty"`
}

// Model to update test point.
type PointUpdateModel struct {
	// Outcome to update.
	Outcome *string `json:"outcome,omitempty"`
	// Reset test point to active.
	ResetToActive *bool `json:"resetToActive,omitempty"`
	// Tester to update. Type IdentityRef.
	Tester *webapi.IdentityRef `json:"tester,omitempty"`
}

// Test point workitem property.
type PointWorkItemProperty struct {
	// key value pair of test point work item property.
	WorkItem *azuredevops.KeyValuePair `json:"workItem,omitempty"`
}

// The class to represent a Generic store for test session data.
type PropertyBag struct {
	// Generic store for test session data
	Bag *map[string]string `json:"bag,omitempty"`
}

type QueryByPointRequest struct {
	ProjectName *string `json:"projectName,omitempty"`
	TestPlanId  *int    `json:"testPlanId,omitempty"`
	TestPointId *int    `json:"testPointId,omitempty"`
}

type QueryByRunRequest struct {
	IncludeActionResults *bool      `json:"includeActionResults,omitempty"`
	Outcome              *byte      `json:"outcome,omitempty"`
	Owner                *uuid.UUID `json:"owner,omitempty"`
	PageSize             *int       `json:"pageSize,omitempty"`
	ProjectName          *string    `json:"projectName,omitempty"`
	State                *byte      `json:"state,omitempty"`
	TestRunId            *int       `json:"testRunId,omitempty"`
}

type QueryModel struct {
	Query *string `json:"query,omitempty"`
}

type QueryTestActionResultRequest struct {
	Identifier  *LegacyTestCaseResultIdentifier `json:"identifier,omitempty"`
	ProjectName *string                         `json:"projectName,omitempty"`
}

type QueryTestActionResultResponse struct {
	TestActionResults    *[]TestActionResult     `json:"testActionResults,omitempty"`
	TestAttachments      *[]TestResultAttachment `json:"testAttachments,omitempty"`
	TestResultParameters *[]TestResultParameter  `json:"testResultParameters,omitempty"`
}

type QueryTestMessageLogEntryRequest struct {
	ProjectName      *string `json:"projectName,omitempty"`
	TestMessageLogId *int    `json:"testMessageLogId,omitempty"`
	TestRunId        *int    `json:"testRunId,omitempty"`
}

type QueryTestRuns2Request struct {
	IncludeStatistics *bool              `json:"includeStatistics,omitempty"`
	Query             *ResultsStoreQuery `json:"query,omitempty"`
}

type QueryTestRunsRequest struct {
	BuildUri        *string    `json:"buildUri,omitempty"`
	Owner           *uuid.UUID `json:"owner,omitempty"`
	PlanId          *int       `json:"planId,omitempty"`
	Skip            *int       `json:"skip,omitempty"`
	TeamProjectName *string    `json:"teamProjectName,omitempty"`
	TestRunId       *int       `json:"testRunId,omitempty"`
	Top             *int       `json:"top,omitempty"`
}

type QueryTestRunStatsRequest struct {
	TeamProjectName *string `json:"teamProjectName,omitempty"`
	TestRunId       *int    `json:"testRunId,omitempty"`
}

// Reference to release environment resource.
type ReleaseEnvironmentDefinitionReference struct {
	// ID of the release definition that contains the release environment definition.
	DefinitionId *int `json:"definitionId,omitempty"`
	// ID of the release environment definition.
	EnvironmentDefinitionId *int `json:"environmentDefinitionId,omitempty"`
}

// Reference to a release.
type ReleaseReference struct {
	// Number of Release Attempt.
	Attempt *int `json:"attempt,omitempty"`
	// Release Creation Date.
	CreationDate *azuredevops.Time `json:"creationDate,omitempty"`
	// Release definition ID.
	DefinitionId *int `json:"definitionId,omitempty"`
	// Environment creation Date.
	EnvironmentCreationDate *azuredevops.Time `json:"environmentCreationDate,omitempty"`
	// Release environment definition ID.
	EnvironmentDefinitionId *int `json:"environmentDefinitionId,omitempty"`
	// Release environment definition name.
	EnvironmentDefinitionName *string `json:"environmentDefinitionName,omitempty"`
	// Release environment ID.
	EnvironmentId *int `json:"environmentId,omitempty"`
	// Release environment name.
	EnvironmentName *string `json:"environmentName,omitempty"`
	// Release ID.
	Id *int `json:"id,omitempty"`
	// Release name.
	Name *string `json:"name,omitempty"`
}

type ReleaseReference2 struct {
	Attempt                 *int              `json:"attempt,omitempty"`
	EnvironmentCreationDate *azuredevops.Time `json:"environmentCreationDate,omitempty"`
	ProjectId               *uuid.UUID        `json:"projectId,omitempty"`
	ReleaseCreationDate     *azuredevops.Time `json:"releaseCreationDate,omitempty"`
	ReleaseDefId            *int              `json:"releaseDefId,omitempty"`
	ReleaseEnvDefId         *int              `json:"releaseEnvDefId,omitempty"`
	ReleaseEnvId            *int              `json:"releaseEnvId,omitempty"`
	ReleaseEnvName          *string           `json:"releaseEnvName,omitempty"`
	ReleaseEnvUri           *string           `json:"releaseEnvUri,omitempty"`
	ReleaseId               *int              `json:"releaseId,omitempty"`
	ReleaseName             *string           `json:"releaseName,omitempty"`
	ReleaseRefId            *int              `json:"releaseRefId,omitempty"`
	ReleaseUri              *string           `json:"releaseUri,omitempty"`
}

type RequirementsToTestsMapping2 struct {
	CreatedBy       *uuid.UUID        `json:"createdBy,omitempty"`
	CreationDate    *azuredevops.Time `json:"creationDate,omitempty"`
	DeletedBy       *uuid.UUID        `json:"deletedBy,omitempty"`
	DeletionDate    *azuredevops.Time `json:"deletionDate,omitempty"`
	IsMigratedToWIT *bool             `json:"isMigratedToWIT,omitempty"`
	ProjectId       *uuid.UUID        `json:"projectId,omitempty"`
	TestMetadataId  *int              `json:"testMetadataId,omitempty"`
	WorkItemId      *int              `json:"workItemId,omitempty"`
}

type ResetTestResultsRequest struct {
	Ids         *[]LegacyTestCaseResultIdentifier `json:"ids,omitempty"`
	ProjectName *string                           `json:"projectName,omitempty"`
}

type Response struct {
	Error  *string    `json:"error,omitempty"`
	Id     *uuid.UUID `json:"id,omitempty"`
	Status *string    `json:"status,omitempty"`
	Url    *string    `json:"url,omitempty"`
}

// Additional details with test result
type ResultDetails string

type resultDetailsValuesType struct {
	None       ResultDetails
	Iterations ResultDetails
	WorkItems  ResultDetails
	SubResults ResultDetails
	Point      ResultDetails
}

var ResultDetailsValues = resultDetailsValuesType{
	// Core fields of test result. Core fields includes State, Outcome, Priority, AutomatedTestName, AutomatedTestStorage, Comments, ErrorMessage etc.
	None: "none",
	// Test iteration details in a test result.
	Iterations: "iterations",
	// Workitems associated with a test result.
	WorkItems: "workItems",
	// Subresults in a test result.
	SubResults: "subResults",
	// Point and plan detail in a test result.
	Point: "point",
}

// Hierarchy type of the result/subresults.
type ResultGroupType string

type resultGroupTypeValuesType struct {
	None        ResultGroupType
	Rerun       ResultGroupType
	DataDriven  ResultGroupType
	OrderedTest ResultGroupType
	Generic     ResultGroupType
}

var ResultGroupTypeValues = resultGroupTypeValuesType{
	// Leaf node of test result.
	None: "none",
	// Hierarchy type of test result.
	Rerun: "rerun",
	// Hierarchy type of test result.
	DataDriven: "dataDriven",
	// Hierarchy type of test result.
	OrderedTest: "orderedTest",
	// Unknown hierarchy type.
	Generic: "generic",
}

// Additional details with test result metadata
type ResultMetaDataDetails string

type resultMetaDataDetailsValuesType struct {
	None             ResultMetaDataDetails
	FlakyIdentifiers ResultMetaDataDetails
}

var ResultMetaDataDetailsValues = resultMetaDataDetailsValuesType{
	// Core fields of test result metadata.
	None: "none",
	// Test FlakyIdentifiers details in test result metadata.
	FlakyIdentifiers: "flakyIdentifiers",
}

// The top level entity that is being cloned as part of a Clone operation
type ResultObjectType string

type resultObjectTypeValuesType struct {
	TestSuite ResultObjectType
	TestPlan  ResultObjectType
}

var ResultObjectTypeValues = resultObjectTypeValuesType{
	// Suite Clone
	TestSuite: "testSuite",
	// Plan Clone
	TestPlan: "testPlan",
}

// Test result retention settings
type ResultRetentionSettings struct {
	// Automated test result retention duration in days
	AutomatedResultsRetentionDuration *int `json:"automatedResultsRetentionDuration,omitempty"`
	// Last Updated by identity
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last updated date
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Manual test result retention duration in days
	ManualResultsRetentionDuration *int `json:"manualResultsRetentionDuration,omitempty"`
}

type ResultsByQueryRequest struct {
	PageSize *int               `json:"pageSize,omitempty"`
	Query    *ResultsStoreQuery `json:"query,omitempty"`
}

type ResultsByQueryResponse struct {
	ExcessIds   *[]LegacyTestCaseResultIdentifier `json:"excessIds,omitempty"`
	TestResults *[]LegacyTestCaseResult           `json:"testResults,omitempty"`
}

type ResultsFilter struct {
	AutomatedTestName    *string             `json:"automatedTestName,omitempty"`
	Branch               *string             `json:"branch,omitempty"`
	ExecutedIn           *Service            `json:"executedIn,omitempty"`
	GroupBy              *string             `json:"groupBy,omitempty"`
	MaxCompleteDate      *azuredevops.Time   `json:"maxCompleteDate,omitempty"`
	ResultsCount         *int                `json:"resultsCount,omitempty"`
	TestCaseId           *int                `json:"testCaseId,omitempty"`
	TestCaseReferenceIds *[]int              `json:"testCaseReferenceIds,omitempty"`
	TestPlanId           *int                `json:"testPlanId,omitempty"`
	TestPointIds         *[]int              `json:"testPointIds,omitempty"`
	TestResultsContext   *TestResultsContext `json:"testResultsContext,omitempty"`
	TrendDays            *int                `json:"trendDays,omitempty"`
}

type ResultsStoreQuery struct {
	DayPrecision    *bool   `json:"dayPrecision,omitempty"`
	QueryText       *string `json:"queryText,omitempty"`
	TeamProjectName *string `json:"teamProjectName,omitempty"`
	TimeZone        *string `json:"timeZone,omitempty"`
}

type ResultUpdateRequest struct {
	ActionResultDeletes *[]TestActionResult             `json:"actionResultDeletes,omitempty"`
	ActionResults       *[]TestActionResult             `json:"actionResults,omitempty"`
	AttachmentDeletes   *[]TestResultAttachmentIdentity `json:"attachmentDeletes,omitempty"`
	Attachments         *[]TestResultAttachment         `json:"attachments,omitempty"`
	ParameterDeletes    *[]TestResultParameter          `json:"parameterDeletes,omitempty"`
	Parameters          *[]TestResultParameter          `json:"parameters,omitempty"`
	TestCaseResult      *LegacyTestCaseResult           `json:"testCaseResult,omitempty"`
	TestResultId        *int                            `json:"testResultId,omitempty"`
	TestRunId           *int                            `json:"testRunId,omitempty"`
}

type ResultUpdateRequestModel struct {
	ActionResultDeletes *[]TestActionResultModel    `json:"actionResultDeletes,omitempty"`
	ActionResults       *[]TestActionResultModel    `json:"actionResults,omitempty"`
	ParameterDeletes    *[]TestResultParameterModel `json:"parameterDeletes,omitempty"`
	Parameters          *[]TestResultParameterModel `json:"parameters,omitempty"`
	TestCaseResult      *TestCaseResultUpdateModel  `json:"testCaseResult,omitempty"`
}

type ResultUpdateResponse struct {
	AttachmentIds          *[]int            `json:"attachmentIds,omitempty"`
	LastUpdated            *azuredevops.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy          *uuid.UUID        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedByName      *string           `json:"lastUpdatedByName,omitempty"`
	MaxReservedSubResultId *int              `json:"maxReservedSubResultId,omitempty"`
	Revision               *int              `json:"revision,omitempty"`
	TestPlanId             *int              `json:"testPlanId,omitempty"`
	TestResultId           *int              `json:"testResultId,omitempty"`
}

type ResultUpdateResponseModel struct {
	Revision *int `json:"revision,omitempty"`
}

// Test run create details.
type RunCreateModel struct {
	// true if test run is automated, false otherwise. By default it will be false.
	Automated *bool `json:"automated,omitempty"`
	// An abstracted reference to the build that it belongs.
	Build *ShallowReference `json:"build,omitempty"`
	// Drop location of the build used for test run.
	BuildDropLocation *string `json:"buildDropLocation,omitempty"`
	// Flavor of the build used for test run. (E.g: Release, Debug)
	BuildFlavor *string `json:"buildFlavor,omitempty"`
	// Platform of the build used for test run. (E.g.: x86, amd64)
	BuildPlatform *string `json:"buildPlatform,omitempty"`
	// BuildReference of the test run.
	BuildReference *BuildConfiguration `json:"buildReference,omitempty"`
	// Comments entered by those analyzing the run.
	Comment *string `json:"comment,omitempty"`
	// Completed date time of the run.
	CompleteDate *string `json:"completeDate,omitempty"`
	// IDs of the test configurations associated with the run.
	ConfigurationIds *[]int `json:"configurationIds,omitempty"`
	// Name of the test controller used for automated run.
	Controller *string `json:"controller,omitempty"`
	// Additional properties of test Run.
	CustomTestFields *[]CustomTestField `json:"customTestFields,omitempty"`
	// An abstracted reference to DtlAutEnvironment.
	DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`
	// An abstracted reference to DtlTestEnvironment.
	DtlTestEnvironment *ShallowReference `json:"dtlTestEnvironment,omitempty"`
	// Due date and time for test run.
	DueDate            *string                `json:"dueDate,omitempty"`
	EnvironmentDetails *DtlEnvironmentDetails `json:"environmentDetails,omitempty"`
	// Error message associated with the run.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Filter used for discovering the Run.
	Filter *RunFilter `json:"filter,omitempty"`
	// The iteration in which to create the run. Root iteration of the team project will be default
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test run.
	Name *string `json:"name,omitempty"`
	// Display name of the owner of the run.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Reference of the pipeline to which this test run belongs. PipelineReference.PipelineId should be equal to RunCreateModel.Build.Id
	PipelineReference *PipelineReference `json:"pipelineReference,omitempty"`
	// An abstracted reference to the plan that it belongs.
	Plan *ShallowReference `json:"plan,omitempty"`
	// IDs of the test points to use in the run.
	PointIds *[]int `json:"pointIds,omitempty"`
	// URI of release environment associated with the run.
	ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
	// Reference to release associated with test run.
	ReleaseReference *ReleaseReference `json:"releaseReference,omitempty"`
	// URI of release associated with the run.
	ReleaseUri *string `json:"releaseUri,omitempty"`
	// Run summary for run Type = NoConfigRun.
	RunSummary *[]RunSummaryModel `json:"runSummary,omitempty"`
	// Timespan till the run times out.
	RunTimeout interface{} `json:"runTimeout,omitempty"`
	// SourceWorkFlow(CI/CD) of the test run.
	SourceWorkflow *string `json:"sourceWorkflow,omitempty"`
	// Start date time of the run.
	StartDate *string `json:"startDate,omitempty"`
	// The state of the run. Type TestRunState Valid states - Unspecified ,NotStarted, InProgress, Completed, Waiting, Aborted, NeedsInvestigation
	State *string `json:"state,omitempty"`
	// Tags to attach with the test run, maximum of 5 tags can be added to run.
	Tags *[]TestTag `json:"tags,omitempty"`
	// TestConfigurationMapping of the test run.
	TestConfigurationsMapping *string `json:"testConfigurationsMapping,omitempty"`
	// ID of the test environment associated with the run.
	TestEnvironmentId *string `json:"testEnvironmentId,omitempty"`
	// An abstracted reference to the test settings resource.
	TestSettings *ShallowReference `json:"testSettings,omitempty"`
	// Type of the run(RunType) Valid Values : (Unspecified, Normal, Blocking, Web, MtrRunInitiatedFromWeb, RunWithDtlEnv, NoConfigRun)
	Type *string `json:"type,omitempty"`
}

// This class is used to provide the filters used for discovery
type RunFilter struct {
	// filter for the test case sources (test containers)
	SourceFilter *string `json:"sourceFilter,omitempty"`
	// filter for the test cases
	TestCaseFilter *string `json:"testCaseFilter,omitempty"`
}

// Test run statistics per outcome.
type RunStatistic struct {
	// Test result count fo the given outcome.
	Count *int `json:"count,omitempty"`
	// Test result outcome
	Outcome *string `json:"outcome,omitempty"`
	// Test run Resolution State.
	ResolutionState *TestResolutionState `json:"resolutionState,omitempty"`
	// State of the test run
	State *string `json:"state,omitempty"`
}

// Run summary for each output type of test.
type RunSummaryModel struct {
	// Total time taken in milliseconds.
	Duration *uint64 `json:"duration,omitempty"`
	// Number of results for Outcome TestOutcome
	ResultCount *int `json:"resultCount,omitempty"`
	// Summary is based on outcome
	TestOutcome *TestOutcome `json:"testOutcome,omitempty"`
}

type RunType string

type runTypeValuesType struct {
	Unspecified            RunType
	Normal                 RunType
	Blocking               RunType
	Web                    RunType
	MtrRunInitiatedFromWeb RunType
	RunWithDtlEnv          RunType
	NoConfigRun            RunType
}

var RunTypeValues = runTypeValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// Normal test run.
	Normal: "normal",
	// Test run created for the blocked result when a test point is blocked.
	Blocking: "blocking",
	// Test run created from Web.
	Web: "web",
	// Run initiated from web through MTR
	MtrRunInitiatedFromWeb: "mtrRunInitiatedFromWeb",
	// These test run would require DTL environment. These could be either of automated or manual test run.
	RunWithDtlEnv: "runWithDtlEnv",
	// These test run may or may not have published test results but it will have summary like total test, passed test, failed test etc. These are automated tests.
	NoConfigRun: "noConfigRun",
}

type RunUpdateModel struct {
	// An abstracted reference to the build that it belongs.
	Build *ShallowReference `json:"build,omitempty"`
	// Drop location of the build used for test run.
	BuildDropLocation *string `json:"buildDropLocation,omitempty"`
	// Flavor of the build used for test run. (E.g: Release, Debug)
	BuildFlavor *string `json:"buildFlavor,omitempty"`
	// Platform of the build used for test run. (E.g.: x86, amd64)
	BuildPlatform *string `json:"buildPlatform,omitempty"`
	// Comments entered by those analyzing the run.
	Comment *string `json:"comment,omitempty"`
	// Completed date time of the run.
	CompletedDate *string `json:"completedDate,omitempty"`
	// Name of the test controller used for automated run.
	Controller *string `json:"controller,omitempty"`
	// true to delete inProgess Results , false otherwise.
	DeleteInProgressResults *bool `json:"deleteInProgressResults,omitempty"`
	// An abstracted reference to DtlAutEnvironment.
	DtlAutEnvironment *ShallowReference `json:"dtlAutEnvironment,omitempty"`
	// An abstracted reference to DtlEnvironment.
	DtlEnvironment        *ShallowReference      `json:"dtlEnvironment,omitempty"`
	DtlEnvironmentDetails *DtlEnvironmentDetails `json:"dtlEnvironmentDetails,omitempty"`
	// Due date and time for test run.
	DueDate *string `json:"dueDate,omitempty"`
	// Error message associated with the run.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// The iteration in which to create the run.
	Iteration *string `json:"iteration,omitempty"`
	// Log entries associated with the run. Use a comma-separated list of multiple log entry objects. { logEntry }, { logEntry }, ...
	LogEntries *[]TestMessageLogDetails `json:"logEntries,omitempty"`
	// Name of the test run.
	Name *string `json:"name,omitempty"`
	// URI of release environment associated with the run.
	ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
	// URI of release associated with the run.
	ReleaseUri *string `json:"releaseUri,omitempty"`
	// Run summary for run Type = NoConfigRun.
	RunSummary *[]RunSummaryModel `json:"runSummary,omitempty"`
	// SourceWorkFlow(CI/CD) of the test run.
	SourceWorkflow *string `json:"sourceWorkflow,omitempty"`
	// Start date time of the run.
	StartedDate *string `json:"startedDate,omitempty"`
	// The state of the test run Below are the valid values - NotStarted, InProgress, Completed, Aborted, Waiting
	State *string `json:"state,omitempty"`
	// The types of sub states for test run.
	Substate *TestRunSubstate `json:"substate,omitempty"`
	// Tags to attach with the test run.
	Tags *[]TestTag `json:"tags,omitempty"`
	// ID of the test environment associated with the run.
	TestEnvironmentId *string `json:"testEnvironmentId,omitempty"`
	// An abstracted reference to test setting resource.
	TestSettings *ShallowReference `json:"testSettings,omitempty"`
}

type Service string

type serviceValuesType struct {
	Any Service
	Tcm Service
	Tfs Service
}

var ServiceValues = serviceValuesType{
	Any: "any",
	Tcm: "tcm",
	Tfs: "tfs",
}

// An abstracted reference to some other resource. This class is used to provide the build data contracts with a uniform way to reference other resources in a way that provides easy traversal through links.
type ShallowReference struct {
	// ID of the resource
	Id *string `json:"id,omitempty"`
	// Name of the linked resource (definition name, controller name, etc.)
	Name *string `json:"name,omitempty"`
	// Full http link to the resource
	Url *string `json:"url,omitempty"`
}

type ShallowTestCaseResult struct {
	AutomatedTestName    *string   `json:"automatedTestName,omitempty"`
	AutomatedTestStorage *string   `json:"automatedTestStorage,omitempty"`
	DurationInMs         *float64  `json:"durationInMs,omitempty"`
	Id                   *int      `json:"id,omitempty"`
	IsReRun              *bool     `json:"isReRun,omitempty"`
	Outcome              *string   `json:"outcome,omitempty"`
	Owner                *string   `json:"owner,omitempty"`
	Priority             *int      `json:"priority,omitempty"`
	RefId                *int      `json:"refId,omitempty"`
	RunId                *int      `json:"runId,omitempty"`
	Tags                 *[]string `json:"tags,omitempty"`
	TestCaseTitle        *string   `json:"testCaseTitle,omitempty"`
}

// Reference to shared step workitem.
type SharedStepModel struct {
	// WorkItem shared step ID.
	Id *int `json:"id,omitempty"`
	// Shared step workitem revision.
	Revision *int `json:"revision,omitempty"`
}

// Stage in pipeline
type StageReference struct {
	// Attempt number of stage
	Attempt *int `json:"attempt,omitempty"`
	// Name of the stage. Maximum supported length for name is 256 character.
	StageName *string `json:"stageName,omitempty"`
}

// Suite create model
type SuiteCreateModel struct {
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// For query based suites, query string that defines the suite.
	QueryString *string `json:"queryString,omitempty"`
	// For requirements test suites, the IDs of the requirements.
	RequirementIds *[]int `json:"requirementIds,omitempty"`
	// Type of test suite to create. It can have value from DynamicTestSuite, StaticTestSuite and RequirementTestSuite.
	SuiteType *string `json:"suiteType,omitempty"`
}

// A suite entry defines properties for a test suite.
type SuiteEntry struct {
	// Id of child suite in the test suite.
	ChildSuiteId *int `json:"childSuiteId,omitempty"`
	// Sequence number for the test case or child test suite in the test suite.
	SequenceNumber *int `json:"sequenceNumber,omitempty"`
	// Id for the test suite.
	SuiteId *int `json:"suiteId,omitempty"`
	// Id of a test case in the test suite.
	TestCaseId *int `json:"testCaseId,omitempty"`
}

// A model to define sequence of test suite entries in a test suite.
type SuiteEntryUpdateModel struct {
	// Id of the child suite in the test suite.
	ChildSuiteId *int `json:"childSuiteId,omitempty"`
	// Updated sequence number for the test case or child test suite in the test suite.
	SequenceNumber *int `json:"sequenceNumber,omitempty"`
	// Id of the test case in the test suite.
	TestCaseId *int `json:"testCaseId,omitempty"`
}

// [Flags] Option to get details in response
type SuiteExpand string

type suiteExpandValuesType struct {
	Children       SuiteExpand
	DefaultTesters SuiteExpand
}

var SuiteExpandValues = suiteExpandValuesType{
	// Include children in response.
	Children: "children",
	// Include default testers in response.
	DefaultTesters: "defaultTesters",
}

// Test case for the suite.
type SuiteTestCase struct {
	// Point Assignment for test suite's test case.
	PointAssignments *[]PointAssignment `json:"pointAssignments,omitempty"`
	// Test case workItem reference.
	TestCase *WorkItemReference `json:"testCase,omitempty"`
}

// Test suite update model.
type SuiteTestCaseUpdateModel struct {
	// Shallow reference of configurations for the test cases in the suite.
	Configurations *[]ShallowReference `json:"configurations,omitempty"`
}

// Test suite update model.
type SuiteUpdateModel struct {
	// Shallow reference of default configurations for the suite.
	DefaultConfigurations *[]ShallowReference `json:"defaultConfigurations,omitempty"`
	// Shallow reference of test suite.
	DefaultTesters *[]ShallowReference `json:"defaultTesters,omitempty"`
	// Specifies if the default configurations have to be inherited from the parent test suite in which the test suite is created.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Test suite name
	Name *string `json:"name,omitempty"`
	// Shallow reference of the parent.
	Parent *ShallowReference `json:"parent,omitempty"`
	// For query based suites, the new query string.
	QueryString *string `json:"queryString,omitempty"`
}

type TCMPropertyBag2 struct {
	ArtifactId   *int    `json:"artifactId,omitempty"`
	ArtifactType *int    `json:"artifactType,omitempty"`
	Name         *string `json:"name,omitempty"`
	Value        *string `json:"value,omitempty"`
}

type TCMServiceDataMigrationStatus string

type tcmServiceDataMigrationStatusValuesType struct {
	NotStarted TCMServiceDataMigrationStatus
	InProgress TCMServiceDataMigrationStatus
	Completed  TCMServiceDataMigrationStatus
	Failed     TCMServiceDataMigrationStatus
}

var TCMServiceDataMigrationStatusValues = tcmServiceDataMigrationStatusValuesType{
	// Migration Not Started
	NotStarted: "notStarted",
	// Migration InProgress
	InProgress: "inProgress",
	// Migration Completed
	Completed: "completed",
	// Migration Failed
	Failed: "failed",
}

type TestActionResult struct {
	ActionPath         *string                         `json:"actionPath,omitempty"`
	Comment            *string                         `json:"comment,omitempty"`
	CreationDate       *azuredevops.Time               `json:"creationDate,omitempty"`
	DateCompleted      *azuredevops.Time               `json:"dateCompleted,omitempty"`
	DateStarted        *azuredevops.Time               `json:"dateStarted,omitempty"`
	Duration           *uint64                         `json:"duration,omitempty"`
	ErrorMessage       *string                         `json:"errorMessage,omitempty"`
	Id                 *LegacyTestCaseResultIdentifier `json:"id,omitempty"`
	IterationId        *int                            `json:"iterationId,omitempty"`
	LastUpdated        *azuredevops.Time               `json:"lastUpdated,omitempty"`
	LastUpdatedBy      *uuid.UUID                      `json:"lastUpdatedBy,omitempty"`
	Outcome            *byte                           `json:"outcome,omitempty"`
	SharedStepId       *int                            `json:"sharedStepId,omitempty"`
	SharedStepRevision *int                            `json:"sharedStepRevision,omitempty"`
}

type TestActionResult2 struct {
	ActionPath         *string           `json:"actionPath,omitempty"`
	Comment            *string           `json:"comment,omitempty"`
	CreationDate       *azuredevops.Time `json:"creationDate,omitempty"`
	DateCompleted      *azuredevops.Time `json:"dateCompleted,omitempty"`
	DateStarted        *azuredevops.Time `json:"dateStarted,omitempty"`
	Duration           *uint64           `json:"duration,omitempty"`
	ErrorMessage       *string           `json:"errorMessage,omitempty"`
	IterationId        *int              `json:"iterationId,omitempty"`
	LastUpdated        *azuredevops.Time `json:"lastUpdated,omitempty"`
	Outcome            *byte             `json:"outcome,omitempty"`
	SharedStepId       *int              `json:"sharedStepId,omitempty"`
	SharedStepRevision *int              `json:"sharedStepRevision,omitempty"`
	TestResultId       *int              `json:"testResultId,omitempty"`
	TestRunId          *int              `json:"testRunId,omitempty"`
}

// Represents a test step result.
type TestActionResultModel struct {
	// Comment in result.
	Comment *string `json:"comment,omitempty"`
	// Time when execution completed.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Duration of execution.
	DurationInMs *float64 `json:"durationInMs,omitempty"`
	// Error message in result.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Test outcome of result.
	Outcome *string `json:"outcome,omitempty"`
	// Time when execution started.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// Path identifier test step in test case workitem.
	ActionPath *string `json:"actionPath,omitempty"`
	// Iteration ID of test action result.
	IterationId *int `json:"iterationId,omitempty"`
	// Reference to shared step workitem.
	SharedStepModel *SharedStepModel `json:"sharedStepModel,omitempty"`
	// This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
	StepIdentifier *string `json:"stepIdentifier,omitempty"`
	// Url of test action result.
	Url *string `json:"url,omitempty"`
}

type TestAttachment struct {
	// Attachment type.
	AttachmentType *AttachmentType `json:"attachmentType,omitempty"`
	// Comment associated with attachment.
	Comment *string `json:"comment,omitempty"`
	// Attachment created date.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Attachment file name
	FileName *string `json:"fileName,omitempty"`
	// ID of the attachment.
	Id *int `json:"id,omitempty"`
	// Attachment size.
	Size *uint64 `json:"size,omitempty"`
	// Attachment Url.
	Url *string `json:"url,omitempty"`
}

// Reference to test attachment.
type TestAttachmentReference struct {
	// ID of the attachment.
	Id *int `json:"id,omitempty"`
	// Url to download the attachment.
	Url *string `json:"url,omitempty"`
}

// Test attachment request model
type TestAttachmentRequestModel struct {
	// Attachment type By Default it will be GeneralAttachment. It can be one of the following type. { GeneralAttachment, AfnStrip, BugFilingData, CodeCoverage, IntermediateCollectorData, RunConfig, TestImpactDetails, TmiTestRunDeploymentFiles, TmiTestRunReverseDeploymentFiles, TmiTestResultDetail, TmiTestRunSummary }
	AttachmentType *string `json:"attachmentType,omitempty"`
	// Comment associated with attachment
	Comment *string `json:"comment,omitempty"`
	// Attachment filename
	FileName *string `json:"fileName,omitempty"`
	// Base64 encoded file stream
	Stream *string `json:"stream,omitempty"`
}

type TestAuthoringDetails struct {
	ConfigurationId *int              `json:"configurationId,omitempty"`
	IsAutomated     *bool             `json:"isAutomated,omitempty"`
	LastUpdated     *azuredevops.Time `json:"lastUpdated,omitempty"`
	PointId         *int              `json:"pointId,omitempty"`
	Priority        *byte             `json:"priority,omitempty"`
	RunBy           *uuid.UUID        `json:"runBy,omitempty"`
	State           *TestPointState   `json:"state,omitempty"`
	SuiteId         *int              `json:"suiteId,omitempty"`
	TesterId        *uuid.UUID        `json:"testerId,omitempty"`
}

type TestCaseMetadata2 struct {
	Container      *string    `json:"container,omitempty"`
	Name           *string    `json:"name,omitempty"`
	ProjectId      *uuid.UUID `json:"projectId,omitempty"`
	TestMetadataId *int       `json:"testMetadataId,omitempty"`
}

type TestCaseReference2 struct {
	AreaId                   *int              `json:"areaId,omitempty"`
	AutomatedTestId          *string           `json:"automatedTestId,omitempty"`
	AutomatedTestName        *string           `json:"automatedTestName,omitempty"`
	AutomatedTestNameHash    *[]byte           `json:"automatedTestNameHash,omitempty"`
	AutomatedTestStorage     *string           `json:"automatedTestStorage,omitempty"`
	AutomatedTestStorageHash *[]byte           `json:"automatedTestStorageHash,omitempty"`
	AutomatedTestType        *string           `json:"automatedTestType,omitempty"`
	ConfigurationId          *int              `json:"configurationId,omitempty"`
	CreatedBy                *uuid.UUID        `json:"createdBy,omitempty"`
	CreationDate             *azuredevops.Time `json:"creationDate,omitempty"`
	LastRefTestRunDate       *azuredevops.Time `json:"lastRefTestRunDate,omitempty"`
	Owner                    *string           `json:"owner,omitempty"`
	Priority                 *byte             `json:"priority,omitempty"`
	ProjectId                *uuid.UUID        `json:"projectId,omitempty"`
	TestCaseId               *int              `json:"testCaseId,omitempty"`
	TestCaseRefId            *int              `json:"testCaseRefId,omitempty"`
	TestCaseRevision         *int              `json:"testCaseRevision,omitempty"`
	TestCaseTitle            *string           `json:"testCaseTitle,omitempty"`
	TestPointId              *int              `json:"testPointId,omitempty"`
}

// Represents a test result.
type TestCaseResult struct {
	// Test attachment ID of action recording.
	AfnStripId *int `json:"afnStripId,omitempty"`
	// Reference to area path of test.
	Area *ShallowReference `json:"area,omitempty"`
	// Reference to bugs linked to test result.
	AssociatedBugs *[]ShallowReference `json:"associatedBugs,omitempty"`
	// ID representing test method in a dll.
	AutomatedTestId *string `json:"automatedTestId,omitempty"`
	// Fully qualified name of test executed.
	AutomatedTestName *string `json:"automatedTestName,omitempty"`
	// Container to which test belongs.
	AutomatedTestStorage *string `json:"automatedTestStorage,omitempty"`
	// Type of automated test.
	AutomatedTestType *string `json:"automatedTestType,omitempty"`
	// TypeId of automated test.
	AutomatedTestTypeId *string `json:"automatedTestTypeId,omitempty"`
	// Shallow reference to build associated with test result.
	Build *ShallowReference `json:"build,omitempty"`
	// Reference to build associated with test result.
	BuildReference *BuildReference `json:"buildReference,omitempty"`
	// Comment in a test result with maxSize= 1000 chars.
	Comment *string `json:"comment,omitempty"`
	// Time when test execution completed. Completed date should be greater than StartedDate.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Machine name where test executed.
	ComputerName *string `json:"computerName,omitempty"`
	// Reference to test configuration. Type ShallowReference.
	Configuration *ShallowReference `json:"configuration,omitempty"`
	// Timestamp when test result created.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// Additional properties of test result.
	CustomFields *[]CustomTestField `json:"customFields,omitempty"`
	// Duration of test execution in milliseconds. If not provided value will be set as CompletedDate - StartedDate
	DurationInMs *float64 `json:"durationInMs,omitempty"`
	// Error message in test execution.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Information when test results started failing.
	FailingSince *FailingSince `json:"failingSince,omitempty"`
	// Failure type of test result. Valid Value= (Known Issue, New Issue, Regression, Unknown, None)
	FailureType *string `json:"failureType,omitempty"`
	// ID of a test result.
	Id *int `json:"id,omitempty"`
	// Test result details of test iterations used only for Manual Testing.
	IterationDetails *[]TestIterationDetailsModel `json:"iterationDetails,omitempty"`
	// Reference to identity last updated test result.
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last updated datetime of test result.
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Test outcome of test result. Valid values = (Unspecified, None, Passed, Failed, Inconclusive, Timeout, Aborted, Blocked, NotExecuted, Warning, Error, NotApplicable, Paused, InProgress, NotImpacted)
	Outcome *string `json:"outcome,omitempty"`
	// Reference to test owner.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Priority of test executed.
	Priority *int `json:"priority,omitempty"`
	// Reference to team project.
	Project *ShallowReference `json:"project,omitempty"`
	// Shallow reference to release associated with test result.
	Release *ShallowReference `json:"release,omitempty"`
	// Reference to release associated with test result.
	ReleaseReference *ReleaseReference `json:"releaseReference,omitempty"`
	// ResetCount.
	ResetCount *int `json:"resetCount,omitempty"`
	// Resolution state of test result.
	ResolutionState *string `json:"resolutionState,omitempty"`
	// ID of resolution state.
	ResolutionStateId *int `json:"resolutionStateId,omitempty"`
	// Hierarchy type of the result, default value of None means its leaf node.
	ResultGroupType *ResultGroupType `json:"resultGroupType,omitempty"`
	// Revision number of test result.
	Revision *int `json:"revision,omitempty"`
	// Reference to identity executed the test.
	RunBy *webapi.IdentityRef `json:"runBy,omitempty"`
	// Stacktrace with maxSize= 1000 chars.
	StackTrace *string `json:"stackTrace,omitempty"`
	// Time when test execution started.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// State of test result. Type TestRunState.
	State *string `json:"state,omitempty"`
	// List of sub results inside a test result, if ResultGroupType is not None, it holds corresponding type sub results.
	SubResults *[]TestSubResult `json:"subResults,omitempty"`
	// Reference to the test executed.
	TestCase *ShallowReference `json:"testCase,omitempty"`
	// Reference ID of test used by test result. Type TestResultMetaData
	TestCaseReferenceId *int `json:"testCaseReferenceId,omitempty"`
	// TestCaseRevision Number.
	TestCaseRevision *int `json:"testCaseRevision,omitempty"`
	// Name of test.
	TestCaseTitle *string `json:"testCaseTitle,omitempty"`
	// Reference to test plan test case workitem is part of.
	TestPlan *ShallowReference `json:"testPlan,omitempty"`
	// Reference to the test point executed.
	TestPoint *ShallowReference `json:"testPoint,omitempty"`
	// Reference to test run.
	TestRun *ShallowReference `json:"testRun,omitempty"`
	// Reference to test suite test case workitem is part of.
	TestSuite *ShallowReference `json:"testSuite,omitempty"`
	// Url of test result.
	Url *string `json:"url,omitempty"`
}

// Test attachment information in a test iteration.
type TestCaseResultAttachmentModel struct {
	// Path identifier test step in test case workitem.
	ActionPath *string `json:"actionPath,omitempty"`
	// Attachment ID.
	Id *int `json:"id,omitempty"`
	// Iteration ID.
	IterationId *int `json:"iterationId,omitempty"`
	// Name of attachment.
	Name *string `json:"name,omitempty"`
	// Attachment size.
	Size *uint64 `json:"size,omitempty"`
	// Url to attachment.
	Url *string `json:"url,omitempty"`
}

type TestCaseResultIdAndRev struct {
	Id       *LegacyTestCaseResultIdentifier `json:"id,omitempty"`
	Revision *int                            `json:"revision,omitempty"`
}

// Reference to a test result.
type TestCaseResultIdentifier struct {
	// Test result ID.
	TestResultId *int `json:"testResultId,omitempty"`
	// Test run ID.
	TestRunId *int `json:"testRunId,omitempty"`
}

type TestCaseResultUpdateModel struct {
	AssociatedWorkItems *[]int              `json:"associatedWorkItems,omitempty"`
	AutomatedTestTypeId *string             `json:"automatedTestTypeId,omitempty"`
	Comment             *string             `json:"comment,omitempty"`
	CompletedDate       *string             `json:"completedDate,omitempty"`
	ComputerName        *string             `json:"computerName,omitempty"`
	CustomFields        *[]CustomTestField  `json:"customFields,omitempty"`
	DurationInMs        *string             `json:"durationInMs,omitempty"`
	ErrorMessage        *string             `json:"errorMessage,omitempty"`
	FailureType         *string             `json:"failureType,omitempty"`
	Outcome             *string             `json:"outcome,omitempty"`
	Owner               *webapi.IdentityRef `json:"owner,omitempty"`
	ResolutionState     *string             `json:"resolutionState,omitempty"`
	RunBy               *webapi.IdentityRef `json:"runBy,omitempty"`
	StackTrace          *string             `json:"stackTrace,omitempty"`
	StartedDate         *string             `json:"startedDate,omitempty"`
	State               *string             `json:"state,omitempty"`
	TestCasePriority    *string             `json:"testCasePriority,omitempty"`
	TestResult          *ShallowReference   `json:"testResult,omitempty"`
}

// Test configuration
type TestConfiguration struct {
	// Area of the configuration
	Area *ShallowReference `json:"area,omitempty"`
	// Description of the configuration
	Description *string `json:"description,omitempty"`
	// Id of the configuration
	Id *int `json:"id,omitempty"`
	// Is the configuration a default for the test plans
	IsDefault *bool `json:"isDefault,omitempty"`
	// Last Updated By  Reference
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last Updated Data
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Name of the configuration
	Name *string `json:"name,omitempty"`
	// Project to which the configuration belongs
	Project *ShallowReference `json:"project,omitempty"`
	// Revision of the the configuration
	Revision *int `json:"revision,omitempty"`
	// State of the configuration
	State *TestConfigurationState `json:"state,omitempty"`
	// Url of Configuration Resource
	Url *string `json:"url,omitempty"`
	// Dictionary of Test Variable, Selected Value
	Values *[]NameValuePair `json:"values,omitempty"`
}

// Represents the state of an ITestConfiguration object.
type TestConfigurationState string

type testConfigurationStateValuesType struct {
	Active   TestConfigurationState
	Inactive TestConfigurationState
}

var TestConfigurationStateValues = testConfigurationStateValuesType{
	// The configuration can be used for new test runs.
	Active: "active",
	// The configuration has been retired and should not be used for new test runs.
	Inactive: "inactive",
}

type TestExecutionReportData struct {
	ReportData *[]DatedTestFieldData `json:"reportData,omitempty"`
}

type TestExtensionField struct {
	Field *TestExtensionFieldDetails `json:"field,omitempty"`
	Value interface{}                `json:"value,omitempty"`
}

type TestExtensionFieldDetails struct {
	Id             *int              `json:"id,omitempty"`
	IsResultScoped *bool             `json:"isResultScoped,omitempty"`
	IsRunScoped    *bool             `json:"isRunScoped,omitempty"`
	IsSystemField  *bool             `json:"isSystemField,omitempty"`
	Name           *string           `json:"name,omitempty"`
	Type           *system.SqlDbType `json:"type,omitempty"`
}

type TestFailureDetails struct {
	Count       *int                        `json:"count,omitempty"`
	TestResults *[]TestCaseResultIdentifier `json:"testResults,omitempty"`
}

type TestFailuresAnalysis struct {
	ExistingFailures *TestFailureDetails `json:"existingFailures,omitempty"`
	FixedTests       *TestFailureDetails `json:"fixedTests,omitempty"`
	NewFailures      *TestFailureDetails `json:"newFailures,omitempty"`
	PreviousContext  *TestResultsContext `json:"previousContext,omitempty"`
}

type TestFailureType struct {
	Id      *int              `json:"id,omitempty"`
	Name    *string           `json:"name,omitempty"`
	Project *ShallowReference `json:"project,omitempty"`
}

type TestFieldData struct {
	Dimensions *map[string]interface{} `json:"dimensions,omitempty"`
	Measure    *uint64                 `json:"measure,omitempty"`
}

type TestFieldsEx2 struct {
	FieldId        *int       `json:"fieldId,omitempty"`
	FieldName      *string    `json:"fieldName,omitempty"`
	FieldType      *byte      `json:"fieldType,omitempty"`
	IsResultScoped *bool      `json:"isResultScoped,omitempty"`
	IsRunScoped    *bool      `json:"isRunScoped,omitempty"`
	IsSystemField  *bool      `json:"isSystemField,omitempty"`
	ProjectId      *uuid.UUID `json:"projectId,omitempty"`
}

// Test Flaky Identifier
type TestFlakyIdentifier struct {
	// Branch Name where Flakiness has to be Marked/Unmarked
	BranchName *string `json:"branchName,omitempty"`
	// State for Flakiness
	IsFlaky *bool `json:"isFlaky,omitempty"`
}

// Filter to get TestCase result history.
type TestHistoryQuery struct {
	// Automated test name of the TestCase.
	AutomatedTestName *string `json:"automatedTestName,omitempty"`
	// Results to be get for a particular branches.
	Branch *string `json:"branch,omitempty"`
	// Get the results history only for this BuildDefinitionId. This to get used in query GroupBy should be Branch. If this is provided, Branch will have no use.
	BuildDefinitionId *int `json:"buildDefinitionId,omitempty"`
	// It will be filled by server. If not null means there are some results still to be get, and we need to call this REST API with this ContinuousToken. It is not supposed to be created (or altered, if received from server in last batch) by user.
	ContinuationToken *string `json:"continuationToken,omitempty"`
	// Group the result on the basis of TestResultGroupBy. This can be Branch, Environment or null(if results are fetched by BuildDefinitionId)
	GroupBy *TestResultGroupBy `json:"groupBy,omitempty"`
	// History to get between time interval MaxCompleteDate and  (MaxCompleteDate - TrendDays). Default is current date time.
	MaxCompleteDate *azuredevops.Time `json:"maxCompleteDate,omitempty"`
	// Get the results history only for this ReleaseEnvDefinitionId. This to get used in query GroupBy should be Environment.
	ReleaseEnvDefinitionId *int `json:"releaseEnvDefinitionId,omitempty"`
	// List of TestResultHistoryForGroup which are grouped by GroupBy
	ResultsForGroup *[]TestResultHistoryForGroup `json:"resultsForGroup,omitempty"`
	// Get the results history only for this testCaseId. This to get used in query to filter the result along with automatedtestname
	TestCaseId *int `json:"testCaseId,omitempty"`
	// Number of days for which history to collect. Maximum supported value is 7 days. Default is 7 days.
	TrendDays *int `json:"trendDays,omitempty"`
}

// Represents a test iteration result.
type TestIterationDetailsModel struct {
	// Test step results in an iteration.
	ActionResults *[]TestActionResultModel `json:"actionResults,omitempty"`
	// Reference to attachments in test iteration result.
	Attachments *[]TestCaseResultAttachmentModel `json:"attachments,omitempty"`
	// Comment in test iteration result.
	Comment *string `json:"comment,omitempty"`
	// Time when execution completed.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Duration of execution.
	DurationInMs *float64 `json:"durationInMs,omitempty"`
	// Error message in test iteration result execution.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// ID of test iteration result.
	Id *int `json:"id,omitempty"`
	// Test outcome if test iteration result.
	Outcome *string `json:"outcome,omitempty"`
	// Test parameters in an iteration.
	Parameters *[]TestResultParameterModel `json:"parameters,omitempty"`
	// Time when execution started.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// Url to test iteration result.
	Url *string `json:"url,omitempty"`
}

// Represents Test Log Result object.
type TestLog struct {
	// Test Log Context run, build
	LogReference *TestLogReference  `json:"logReference,omitempty"`
	MetaData     *map[string]string `json:"metaData,omitempty"`
	// LastUpdatedDate for Log file
	ModifiedOn *azuredevops.Time `json:"modifiedOn,omitempty"`
	// Size in Bytes for Log file
	Size *uint64 `json:"size,omitempty"`
}

type TestLogReference struct {
	// BuildId for test log, if context is build
	BuildId *int `json:"buildId,omitempty"`
	// FileName for log file
	FilePath *string `json:"filePath,omitempty"`
	// ReleaseEnvId for test log, if context is Release
	ReleaseEnvId *int `json:"releaseEnvId,omitempty"`
	// ReleaseId for test log, if context is Release
	ReleaseId *int `json:"releaseId,omitempty"`
	// Resultid for test log, if context is run and log is related to result
	ResultId *int `json:"resultId,omitempty"`
	// runid for test log, if context is run
	RunId *int `json:"runId,omitempty"`
	// Test Log Reference object
	Scope *TestLogScope `json:"scope,omitempty"`
	// SubResultid for test log, if context is run and log is related to subresult
	SubResultId *int `json:"subResultId,omitempty"`
	// Log Type
	Type *TestLogType `json:"type,omitempty"`
}

// Test Log Context
type TestLogScope string

type testLogScopeValuesType struct {
	Run     TestLogScope
	Build   TestLogScope
	Release TestLogScope
}

var TestLogScopeValues = testLogScopeValuesType{
	// Log file is associated with Run, result, subresult
	Run: "run",
	// Log File associated with Build
	Build: "build",
	// Log File associated with Release
	Release: "release",
}

// Represents Test Log Status object.
type TestLogStatus struct {
	// Exception message
	Exception *string `json:"exception,omitempty"`
	// Test Log Status code
	Status *TestLogStatusCode `json:"status,omitempty"`
	// Blob Transfer Error code
	TransferFailureType *string `json:"transferFailureType,omitempty"`
}

// Test Log Status codes.
type TestLogStatusCode string

type testLogStatusCodeValuesType struct {
	Success             TestLogStatusCode
	Failed              TestLogStatusCode
	FileAlreadyExists   TestLogStatusCode
	InvalidInput        TestLogStatusCode
	InvalidFileName     TestLogStatusCode
	InvalidContainer    TestLogStatusCode
	TransferFailed      TestLogStatusCode
	FeatureDisabled     TestLogStatusCode
	BuildDoesNotExist   TestLogStatusCode
	RunDoesNotExist     TestLogStatusCode
	ContainerNotCreated TestLogStatusCode
	ApiNotSupported     TestLogStatusCode
	FileSizeExceeds     TestLogStatusCode
	ContainerNotFound   TestLogStatusCode
	FileNotFound        TestLogStatusCode
	DirectoryNotFound   TestLogStatusCode
}

var TestLogStatusCodeValues = testLogStatusCodeValuesType{
	Success:             "success",
	Failed:              "failed",
	FileAlreadyExists:   "fileAlreadyExists",
	InvalidInput:        "invalidInput",
	InvalidFileName:     "invalidFileName",
	InvalidContainer:    "invalidContainer",
	TransferFailed:      "transferFailed",
	FeatureDisabled:     "featureDisabled",
	BuildDoesNotExist:   "buildDoesNotExist",
	RunDoesNotExist:     "runDoesNotExist",
	ContainerNotCreated: "containerNotCreated",
	ApiNotSupported:     "apiNotSupported",
	FileSizeExceeds:     "fileSizeExceeds",
	ContainerNotFound:   "containerNotFound",
	FileNotFound:        "fileNotFound",
	DirectoryNotFound:   "directoryNotFound",
}

// Represents Test Log store endpoint details.
type TestLogStoreEndpointDetails struct {
	// Test log store connection Uri.
	EndpointSASUri *string `json:"endpointSASUri,omitempty"`
	// Test log store endpoint type.
	EndpointType *TestLogStoreEndpointType `json:"endpointType,omitempty"`
	// Test log store status code
	Status *TestLogStatusCode `json:"status,omitempty"`
}

type TestLogStoreEndpointType string

type testLogStoreEndpointTypeValuesType struct {
	Root TestLogStoreEndpointType
	File TestLogStoreEndpointType
}

var TestLogStoreEndpointTypeValues = testLogStoreEndpointTypeValuesType{
	Root: "root",
	File: "file",
}

type TestLogStoreOperationType string

type testLogStoreOperationTypeValuesType struct {
	Read          TestLogStoreOperationType
	Create        TestLogStoreOperationType
	ReadAndCreate TestLogStoreOperationType
}

var TestLogStoreOperationTypeValues = testLogStoreOperationTypeValuesType{
	Read:          "read",
	Create:        "create",
	ReadAndCreate: "readAndCreate",
}

// Test Log Types
type TestLogType string

type testLogTypeValuesType struct {
	GeneralAttachment TestLogType
	CodeCoverage      TestLogType
	TestImpact        TestLogType
	Intermediate      TestLogType
}

var TestLogTypeValues = testLogTypeValuesType{
	// Any gereric attachment.
	GeneralAttachment: "generalAttachment",
	// Code Coverage files
	CodeCoverage: "codeCoverage",
	// Test Impact details.
	TestImpact: "testImpact",
	// Temporary files
	Intermediate: "intermediate",
}

type TestMessageLog2 struct {
	TestMessageLogId *int `json:"testMessageLogId,omitempty"`
}

// An abstracted reference to some other resource. This class is used to provide the build data contracts with a uniform way to reference other resources in a way that provides easy traversal through links.
type TestMessageLogDetails struct {
	// Date when the resource is created
	DateCreated *azuredevops.Time `json:"dateCreated,omitempty"`
	// Id of the resource
	EntryId *int `json:"entryId,omitempty"`
	// Message of the resource
	Message *string `json:"message,omitempty"`
}

type TestMessageLogEntry struct {
	DateCreated      *azuredevops.Time `json:"dateCreated,omitempty"`
	EntryId          *int              `json:"entryId,omitempty"`
	LogLevel         *byte             `json:"logLevel,omitempty"`
	LogUser          *uuid.UUID        `json:"logUser,omitempty"`
	LogUserName      *string           `json:"logUserName,omitempty"`
	Message          *string           `json:"message,omitempty"`
	TestMessageLogId *int              `json:"testMessageLogId,omitempty"`
}

type TestMessageLogEntry2 struct {
	DateCreated      *azuredevops.Time `json:"dateCreated,omitempty"`
	EntryId          *int              `json:"entryId,omitempty"`
	LogLevel         *byte             `json:"logLevel,omitempty"`
	LogUser          *uuid.UUID        `json:"logUser,omitempty"`
	Message          *string           `json:"message,omitempty"`
	TestMessageLogId *int              `json:"testMessageLogId,omitempty"`
}

type TestMethod struct {
	Container *string `json:"container,omitempty"`
	Name      *string `json:"name,omitempty"`
}

// Class representing a reference to an operation.
type TestOperationReference struct {
	Id     *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Url    *string `json:"url,omitempty"`
}

// Valid TestOutcome values.
type TestOutcome string

type testOutcomeValuesType struct {
	Unspecified   TestOutcome
	None          TestOutcome
	Passed        TestOutcome
	Failed        TestOutcome
	Inconclusive  TestOutcome
	Timeout       TestOutcome
	Aborted       TestOutcome
	Blocked       TestOutcome
	NotExecuted   TestOutcome
	Warning       TestOutcome
	Error         TestOutcome
	NotApplicable TestOutcome
	Paused        TestOutcome
	InProgress    TestOutcome
	NotImpacted   TestOutcome
}

var TestOutcomeValues = testOutcomeValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// Test has not been completed, or the test type does not report pass/failure.
	None: "none",
	// Test was executed w/o any issues.
	Passed: "passed",
	// Test was executed, but there were issues. Issues may involve exceptions or failed assertions.
	Failed: "failed",
	// Test has completed, but we can't say if it passed or failed. May be used for aborted tests...
	Inconclusive: "inconclusive",
	// The test timed out
	Timeout: "timeout",
	// Test was aborted. This was not caused by a user gesture, but rather by a framework decision.
	Aborted: "aborted",
	// Test had it chance for been executed but was not, as ITestElement.IsRunnable == false.
	Blocked: "blocked",
	// Test was not executed. This was caused by a user gesture - e.g. user hit stop button.
	NotExecuted: "notExecuted",
	// To be used by Run level results. This is not a failure.
	Warning: "warning",
	// There was a system error while we were trying to execute a test.
	Error: "error",
	// Test is Not Applicable for execution.
	NotApplicable: "notApplicable",
	// Test is paused.
	Paused: "paused",
	// Test is currently executing. Added this for TCM charts
	InProgress: "inProgress",
	// Test is not impacted. Added fot TIA.
	NotImpacted: "notImpacted",
}

// Test outcome settings
type TestOutcomeSettings struct {
	// Value to configure how test outcomes for the same tests across suites are shown
	SyncOutcomeAcrossSuites *bool `json:"syncOutcomeAcrossSuites,omitempty"`
}

type TestParameter2 struct {
	ActionPath    *string           `json:"actionPath,omitempty"`
	Actual        *[]byte           `json:"actual,omitempty"`
	CreationDate  *azuredevops.Time `json:"creationDate,omitempty"`
	DataType      *byte             `json:"dataType,omitempty"`
	DateModified  *azuredevops.Time `json:"dateModified,omitempty"`
	Expected      *[]byte           `json:"expected,omitempty"`
	IterationId   *int              `json:"iterationId,omitempty"`
	ParameterName *string           `json:"parameterName,omitempty"`
	TestResultId  *int              `json:"testResultId,omitempty"`
	TestRunId     *int              `json:"testRunId,omitempty"`
}

// The test plan resource.
type TestPlan struct {
	// Area of the test plan.
	Area *ShallowReference `json:"area,omitempty"`
	// Build to be tested.
	Build *ShallowReference `json:"build,omitempty"`
	// The Build Definition that generates a build associated with this test plan.
	BuildDefinition *ShallowReference `json:"buildDefinition,omitempty"`
	// Description of the test plan.
	Description *string `json:"description,omitempty"`
	// End date for the test plan.
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// ID of the test plan.
	Id *int `json:"id,omitempty"`
	// Iteration path of the test plan.
	Iteration *string `json:"iteration,omitempty"`
	// Name of the test plan.
	Name *string `json:"name,omitempty"`
	// Owner of the test plan.
	Owner         *webapi.IdentityRef `json:"owner,omitempty"`
	PreviousBuild *ShallowReference   `json:"previousBuild,omitempty"`
	// Project which contains the test plan.
	Project *ShallowReference `json:"project,omitempty"`
	// Release Environment to be used to deploy the build and run automated tests from this test plan.
	ReleaseEnvironmentDefinition *ReleaseEnvironmentDefinitionReference `json:"releaseEnvironmentDefinition,omitempty"`
	// Revision of the test plan.
	Revision *int `json:"revision,omitempty"`
	// Root test suite of the test plan.
	RootSuite *ShallowReference `json:"rootSuite,omitempty"`
	// Start date for the test plan.
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test plan.
	State *string `json:"state,omitempty"`
	// Value to configure how same tests across test suites under a test plan need to behave
	TestOutcomeSettings *TestOutcomeSettings `json:"testOutcomeSettings,omitempty"`
	UpdatedBy           *webapi.IdentityRef  `json:"updatedBy,omitempty"`
	UpdatedDate         *azuredevops.Time    `json:"updatedDate,omitempty"`
	// URL of the test plan resource.
	Url *string `json:"url,omitempty"`
}

type TestPlanCloneRequest struct {
	DestinationTestPlan *TestPlan     `json:"destinationTestPlan,omitempty"`
	Options             *CloneOptions `json:"options,omitempty"`
	SuiteIds            *[]int        `json:"suiteIds,omitempty"`
}

type TestPlanHubData struct {
	SelectedSuiteId *int         `json:"selectedSuiteId,omitempty"`
	TestPlan        *TestPlan    `json:"testPlan,omitempty"`
	TestPoints      *[]TestPoint `json:"testPoints,omitempty"`
	TestSuites      *[]TestSuite `json:"testSuites,omitempty"`
	TotalTestPoints *int         `json:"totalTestPoints,omitempty"`
}

type TestPlansWithSelection struct {
	LastSelectedPlan  *int        `json:"lastSelectedPlan,omitempty"`
	LastSelectedSuite *int        `json:"lastSelectedSuite,omitempty"`
	Plans             *[]TestPlan `json:"plans,omitempty"`
}

// Test point.
type TestPoint struct {
	// AssignedTo. Type IdentityRef.
	AssignedTo *webapi.IdentityRef `json:"assignedTo,omitempty"`
	// Automated.
	Automated *bool `json:"automated,omitempty"`
	// Comment associated with test point.
	Comment *string `json:"comment,omitempty"`
	// Configuration. Type ShallowReference.
	Configuration *ShallowReference `json:"configuration,omitempty"`
	// Failure type of test point.
	FailureType *string `json:"failureType,omitempty"`
	// ID of the test point.
	Id *int `json:"id,omitempty"`
	// Last date when test point was reset to Active.
	LastResetToActive *azuredevops.Time `json:"lastResetToActive,omitempty"`
	// Last resolution state id of test point.
	LastResolutionStateId *int `json:"lastResolutionStateId,omitempty"`
	// Last result of test point. Type ShallowReference.
	LastResult *ShallowReference `json:"lastResult,omitempty"`
	// Last result details of test point. Type LastResultDetails.
	LastResultDetails *LastResultDetails `json:"lastResultDetails,omitempty"`
	// Last result state of test point.
	LastResultState *string `json:"lastResultState,omitempty"`
	// LastRun build number of test point.
	LastRunBuildNumber *string `json:"lastRunBuildNumber,omitempty"`
	// Last testRun of test point. Type ShallowReference.
	LastTestRun *ShallowReference `json:"lastTestRun,omitempty"`
	// Test point last updated by. Type IdentityRef.
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last updated date of test point.
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Outcome of test point.
	Outcome *string `json:"outcome,omitempty"`
	// Revision number.
	Revision *int `json:"revision,omitempty"`
	// State of test point.
	State *string `json:"state,omitempty"`
	// Suite of test point. Type ShallowReference.
	Suite *ShallowReference `json:"suite,omitempty"`
	// TestCase associated to test point. Type WorkItemReference.
	TestCase *WorkItemReference `json:"testCase,omitempty"`
	// TestPlan of test point. Type ShallowReference.
	TestPlan *ShallowReference `json:"testPlan,omitempty"`
	// Test point Url.
	Url *string `json:"url,omitempty"`
	// Work item properties of test point.
	WorkItemProperties *[]interface{} `json:"workItemProperties,omitempty"`
}

type TestPointReference struct {
	Id    *int            `json:"id,omitempty"`
	State *TestPointState `json:"state,omitempty"`
}

type TestPointsEvent struct {
	ProjectName *string               `json:"projectName,omitempty"`
	TestPoints  *[]TestPointReference `json:"testPoints,omitempty"`
}

// Test point query class.
type TestPointsQuery struct {
	// Order by results.
	OrderBy *string `json:"orderBy,omitempty"`
	// List of test points
	Points *[]TestPoint `json:"points,omitempty"`
	// Filter
	PointsFilter *PointsFilter `json:"pointsFilter,omitempty"`
	// List of workitem fields to get.
	WitFields *[]string `json:"witFields,omitempty"`
}

type TestPointState string

type testPointStateValuesType struct {
	None       TestPointState
	Ready      TestPointState
	Completed  TestPointState
	NotReady   TestPointState
	InProgress TestPointState
	MaxValue   TestPointState
}

var TestPointStateValues = testPointStateValuesType{
	// Default
	None: "none",
	// The test point needs to be executed in order for the test pass to be considered complete.  Either the test has not been run before or the previous run failed.
	Ready: "ready",
	// The test has passed successfully and does not need to be re-run for the test pass to be considered complete.
	Completed: "completed",
	// The test point needs to be executed but is not able to.
	NotReady: "notReady",
	// The test is being executed.
	InProgress: "inProgress",
	MaxValue:   "maxValue",
}

type TestPointsUpdatedEvent struct {
	ProjectName *string               `json:"projectName,omitempty"`
	TestPoints  *[]TestPointReference `json:"testPoints,omitempty"`
}

// Test Resolution State Details.
type TestResolutionState struct {
	// Test Resolution state Id.
	Id *int `json:"id,omitempty"`
	// Test Resolution State Name.
	Name    *string           `json:"name,omitempty"`
	Project *ShallowReference `json:"project,omitempty"`
}

type TestResult2 struct {
	AfnStripId          *int              `json:"afnStripId,omitempty"`
	ComputerName        *string           `json:"computerName,omitempty"`
	CreationDate        *azuredevops.Time `json:"creationDate,omitempty"`
	DateCompleted       *azuredevops.Time `json:"dateCompleted,omitempty"`
	DateStarted         *azuredevops.Time `json:"dateStarted,omitempty"`
	EffectivePointState *byte             `json:"effectivePointState,omitempty"`
	FailureType         *byte             `json:"failureType,omitempty"`
	LastUpdated         *azuredevops.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy       *uuid.UUID        `json:"lastUpdatedBy,omitempty"`
	Outcome             *byte             `json:"outcome,omitempty"`
	Owner               *uuid.UUID        `json:"owner,omitempty"`
	ProjectId           *uuid.UUID        `json:"projectId,omitempty"`
	ResetCount          *int              `json:"resetCount,omitempty"`
	ResolutionStateId   *int              `json:"resolutionStateId,omitempty"`
	Revision            *int              `json:"revision,omitempty"`
	RunBy               *uuid.UUID        `json:"runBy,omitempty"`
	State               *byte             `json:"state,omitempty"`
	TestCaseRefId       *int              `json:"testCaseRefId,omitempty"`
	TestResultId        *int              `json:"testResultId,omitempty"`
	TestRunId           *int              `json:"testRunId,omitempty"`
}

type TestResultAcrossProjectResponse struct {
	ProjectName *string               `json:"projectName,omitempty"`
	TestResult  *LegacyTestCaseResult `json:"testResult,omitempty"`
}

type TestResultAttachment struct {
	ActionPath          *string           `json:"actionPath,omitempty"`
	AttachmentType      *AttachmentType   `json:"attachmentType,omitempty"`
	Comment             *string           `json:"comment,omitempty"`
	CreationDate        *azuredevops.Time `json:"creationDate,omitempty"`
	DownloadQueryString *string           `json:"downloadQueryString,omitempty"`
	FileName            *string           `json:"fileName,omitempty"`
	Id                  *int              `json:"id,omitempty"`
	IsComplete          *bool             `json:"isComplete,omitempty"`
	IterationId         *int              `json:"iterationId,omitempty"`
	Length              *uint64           `json:"length,omitempty"`
	SessionId           *int              `json:"sessionId,omitempty"`
	TestResultId        *int              `json:"testResultId,omitempty"`
	TestRunId           *int              `json:"testRunId,omitempty"`
	TmiRunId            *uuid.UUID        `json:"tmiRunId,omitempty"`
}

type TestResultAttachmentIdentity struct {
	AttachmentId *int `json:"attachmentId,omitempty"`
	SessionId    *int `json:"sessionId,omitempty"`
	TestResultId *int `json:"testResultId,omitempty"`
	TestRunId    *int `json:"testRunId,omitempty"`
}

type TestResultCreateModel struct {
	Area                 *ShallowReference   `json:"area,omitempty"`
	AssociatedWorkItems  *[]int              `json:"associatedWorkItems,omitempty"`
	AutomatedTestId      *string             `json:"automatedTestId,omitempty"`
	AutomatedTestName    *string             `json:"automatedTestName,omitempty"`
	AutomatedTestStorage *string             `json:"automatedTestStorage,omitempty"`
	AutomatedTestType    *string             `json:"automatedTestType,omitempty"`
	AutomatedTestTypeId  *string             `json:"automatedTestTypeId,omitempty"`
	Comment              *string             `json:"comment,omitempty"`
	CompletedDate        *string             `json:"completedDate,omitempty"`
	ComputerName         *string             `json:"computerName,omitempty"`
	Configuration        *ShallowReference   `json:"configuration,omitempty"`
	CustomFields         *[]CustomTestField  `json:"customFields,omitempty"`
	DurationInMs         *string             `json:"durationInMs,omitempty"`
	ErrorMessage         *string             `json:"errorMessage,omitempty"`
	FailureType          *string             `json:"failureType,omitempty"`
	Outcome              *string             `json:"outcome,omitempty"`
	Owner                *webapi.IdentityRef `json:"owner,omitempty"`
	ResolutionState      *string             `json:"resolutionState,omitempty"`
	RunBy                *webapi.IdentityRef `json:"runBy,omitempty"`
	StackTrace           *string             `json:"stackTrace,omitempty"`
	StartedDate          *string             `json:"startedDate,omitempty"`
	State                *string             `json:"state,omitempty"`
	TestCase             *ShallowReference   `json:"testCase,omitempty"`
	TestCasePriority     *string             `json:"testCasePriority,omitempty"`
	TestCaseTitle        *string             `json:"testCaseTitle,omitempty"`
	TestPoint            *ShallowReference   `json:"testPoint,omitempty"`
}

type TestResultDocument struct {
	OperationReference *TestOperationReference `json:"operationReference,omitempty"`
	Payload            *TestResultPayload      `json:"payload,omitempty"`
}

// Group by for results
type TestResultGroupBy string

type testResultGroupByValuesType struct {
	Branch      TestResultGroupBy
	Environment TestResultGroupBy
}

var TestResultGroupByValues = testResultGroupByValuesType{
	// Group the results by branches
	Branch: "branch",
	// Group the results by environment
	Environment: "environment",
}

type TestResultHistory struct {
	GroupByField    *string                             `json:"groupByField,omitempty"`
	ResultsForGroup *[]TestResultHistoryDetailsForGroup `json:"resultsForGroup,omitempty"`
}

type TestResultHistoryDetailsForGroup struct {
	GroupByValue interface{}     `json:"groupByValue,omitempty"`
	LatestResult *TestCaseResult `json:"latestResult,omitempty"`
}

// List of test results filtered on the basis of GroupByValue
type TestResultHistoryForGroup struct {
	// Display name of the group.
	DisplayName *string `json:"displayName,omitempty"`
	// Name or Id of the group identifier by which results are grouped together.
	GroupByValue *string `json:"groupByValue,omitempty"`
	// List of results for GroupByValue
	Results *[]TestCaseResult `json:"results,omitempty"`
}

// Represents a Meta Data of a test result.
type TestResultMetaData struct {
	// AutomatedTestName of test result.
	AutomatedTestName *string `json:"automatedTestName,omitempty"`
	// AutomatedTestStorage of test result.
	AutomatedTestStorage *string `json:"automatedTestStorage,omitempty"`
	// List of Flaky Identifier for TestCaseReferenceId
	FlakyIdentifiers *[]TestFlakyIdentifier `json:"flakyIdentifiers,omitempty"`
	// Owner of test result.
	Owner *string `json:"owner,omitempty"`
	// Priority of test result.
	Priority *int `json:"priority,omitempty"`
	// ID of TestCaseReference.
	TestCaseReferenceId *int `json:"testCaseReferenceId,omitempty"`
	// TestCaseTitle of test result.
	TestCaseTitle *string `json:"testCaseTitle,omitempty"`
}

// Represents a TestResultMetaData Input
type TestResultMetaDataUpdateInput struct {
	// List of Flaky Identifiers
	FlakyIdentifiers *[]TestFlakyIdentifier `json:"flakyIdentifiers,omitempty"`
}

type TestResultMetaDataUpdateResponse struct {
	Status *string `json:"status,omitempty"`
}

type TestResultModelBase struct {
	// Comment in result.
	Comment *string `json:"comment,omitempty"`
	// Time when execution completed.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Duration of execution.
	DurationInMs *float64 `json:"durationInMs,omitempty"`
	// Error message in result.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Test outcome of result.
	Outcome *string `json:"outcome,omitempty"`
	// Time when execution started.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
}

type TestResultParameter struct {
	ActionPath    *string `json:"actionPath,omitempty"`
	Actual        *[]byte `json:"actual,omitempty"`
	Expected      *[]byte `json:"expected,omitempty"`
	IterationId   *int    `json:"iterationId,omitempty"`
	ParameterName *string `json:"parameterName,omitempty"`
	TestResultId  *int    `json:"testResultId,omitempty"`
	TestRunId     *int    `json:"testRunId,omitempty"`
}

// Test parameter information in a test iteration.
type TestResultParameterModel struct {
	// Test step path where parameter is referenced.
	ActionPath *string `json:"actionPath,omitempty"`
	// Iteration ID.
	IterationId *int `json:"iterationId,omitempty"`
	// Name of parameter.
	ParameterName *string `json:"parameterName,omitempty"`
	// This is step Id of test case. For shared step, it is step Id of shared step in test case workitem; step Id in shared step. Example: TestCase workitem has two steps: 1) Normal step with Id = 1 2) Shared Step with Id = 2. Inside shared step: a) Normal Step with Id = 1 Value for StepIdentifier for First step: "1" Second step: "2;1"
	StepIdentifier *string `json:"stepIdentifier,omitempty"`
	// Url of test parameter.
	Url *string `json:"url,omitempty"`
	// Value of parameter.
	Value *string `json:"value,omitempty"`
}

type TestResultPayload struct {
	Comment *string `json:"comment,omitempty"`
	Name    *string `json:"name,omitempty"`
	Stream  *string `json:"stream,omitempty"`
}

type TestResultReset2 struct {
	AuditIdentity *uuid.UUID        `json:"auditIdentity,omitempty"`
	DateModified  *azuredevops.Time `json:"dateModified,omitempty"`
	ProjectId     *uuid.UUID        `json:"projectId,omitempty"`
	Revision      *int              `json:"revision,omitempty"`
	TestResultId  *int              `json:"testResultId,omitempty"`
	TestResultRV  *[]byte           `json:"testResultRV,omitempty"`
	TestRunId     *int              `json:"testRunId,omitempty"`
}

type TestResultsContext struct {
	Build       *BuildReference         `json:"build,omitempty"`
	ContextType *TestResultsContextType `json:"contextType,omitempty"`
	Release     *ReleaseReference       `json:"release,omitempty"`
}

type TestResultsContextType string

type testResultsContextTypeValuesType struct {
	Build   TestResultsContextType
	Release TestResultsContextType
}

var TestResultsContextTypeValues = testResultsContextTypeValuesType{
	Build:   "build",
	Release: "release",
}

type TestResultsDetails struct {
	GroupByField    *string                       `json:"groupByField,omitempty"`
	ResultsForGroup *[]TestResultsDetailsForGroup `json:"resultsForGroup,omitempty"`
}

type TestResultsDetailsForGroup struct {
	GroupByValue          interface{}                                 `json:"groupByValue,omitempty"`
	Results               *[]TestCaseResult                           `json:"results,omitempty"`
	ResultsCountByOutcome *map[TestOutcome]AggregatedResultsByOutcome `json:"resultsCountByOutcome,omitempty"`
	Tags                  *[]string                                   `json:"tags,omitempty"`
}

type TestResultsEx2 struct {
	BitValue      *bool             `json:"bitValue,omitempty"`
	CreationDate  *azuredevops.Time `json:"creationDate,omitempty"`
	DateTimeValue *azuredevops.Time `json:"dateTimeValue,omitempty"`
	FieldId       *int              `json:"fieldId,omitempty"`
	FieldName     *string           `json:"fieldName,omitempty"`
	FloatValue    *float64          `json:"floatValue,omitempty"`
	GuidValue     *uuid.UUID        `json:"guidValue,omitempty"`
	IntValue      *int              `json:"intValue,omitempty"`
	ProjectId     *uuid.UUID        `json:"projectId,omitempty"`
	StringValue   *string           `json:"stringValue,omitempty"`
	TestResultId  *int              `json:"testResultId,omitempty"`
	TestRunId     *int              `json:"testRunId,omitempty"`
}

type TestResultsGroupsForBuild struct {
	// BuildId for which groupby result is fetched.
	BuildId *int `json:"buildId,omitempty"`
	// The group by results
	Fields *[]FieldDetailsForTestResults `json:"fields,omitempty"`
}

type TestResultsGroupsForRelease struct {
	// The group by results
	Fields *[]FieldDetailsForTestResults `json:"fields,omitempty"`
	// Release Environment Id for which groupby result is fetched.
	ReleaseEnvId *int `json:"releaseEnvId,omitempty"`
	// ReleaseId for which groupby result is fetched.
	ReleaseId *int `json:"releaseId,omitempty"`
}

type TestResultsQuery struct {
	Fields        *[]string         `json:"fields,omitempty"`
	Results       *[]TestCaseResult `json:"results,omitempty"`
	ResultsFilter *ResultsFilter    `json:"resultsFilter,omitempty"`
}

type TestResultsSettings struct {
	// IsRequired and EmitDefaultValue are passed as false as if users doesn't pass anything, should not come for serialisation and deserialisation.
	FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

type TestResultsSettingsType string

type testResultsSettingsTypeValuesType struct {
	All   TestResultsSettingsType
	Flaky TestResultsSettingsType
}

var TestResultsSettingsTypeValues = testResultsSettingsTypeValuesType{
	// Returns All Test Settings.
	All: "all",
	// Returns Flaky Test Settings.
	Flaky: "flaky",
}

type TestResultSummary struct {
	AggregatedResultsAnalysis *AggregatedResultsAnalysis `json:"aggregatedResultsAnalysis,omitempty"`
	NoConfigRunsCount         *int                       `json:"noConfigRunsCount,omitempty"`
	TeamProject               *core.TeamProjectReference `json:"teamProject,omitempty"`
	TestFailures              *TestFailuresAnalysis      `json:"testFailures,omitempty"`
	TestResultsContext        *TestResultsContext        `json:"testResultsContext,omitempty"`
	TotalRunsCount            *int                       `json:"totalRunsCount,omitempty"`
}

type TestResultsUpdateSettings struct {
	// FlakySettings defines Flaky Settings Data.
	FlakySettings *FlakySettings `json:"flakySettings,omitempty"`
}

type TestResultsWithWatermark struct {
	ChangedDate   *azuredevops.Time `json:"changedDate,omitempty"`
	PointsResults *[]PointsResults2 `json:"pointsResults,omitempty"`
	ResultId      *int              `json:"resultId,omitempty"`
	RunId         *int              `json:"runId,omitempty"`
}

type TestResultTrendFilter struct {
	BranchNames      *[]string         `json:"branchNames,omitempty"`
	BuildCount       *int              `json:"buildCount,omitempty"`
	DefinitionIds    *[]int            `json:"definitionIds,omitempty"`
	EnvDefinitionIds *[]int            `json:"envDefinitionIds,omitempty"`
	MaxCompleteDate  *azuredevops.Time `json:"maxCompleteDate,omitempty"`
	PublishContext   *string           `json:"publishContext,omitempty"`
	TestRunTitles    *[]string         `json:"testRunTitles,omitempty"`
	TrendDays        *int              `json:"trendDays,omitempty"`
}

// Test run details.
type TestRun struct {
	// Build associated with this test run.
	Build *ShallowReference `json:"build,omitempty"`
	// Build configuration details associated with this test run.
	BuildConfiguration *BuildConfiguration `json:"buildConfiguration,omitempty"`
	// Comments entered by those analyzing the run.
	Comment *string `json:"comment,omitempty"`
	// Completed date time of the run.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Test Run Controller.
	Controller *string `json:"controller,omitempty"`
	// Test Run CreatedDate.
	CreatedDate *azuredevops.Time `json:"createdDate,omitempty"`
	// List of Custom Fields for TestRun.
	CustomFields *[]CustomTestField `json:"customFields,omitempty"`
	// Drop Location for the test Run.
	DropLocation                  *string                `json:"dropLocation,omitempty"`
	DtlAutEnvironment             *ShallowReference      `json:"dtlAutEnvironment,omitempty"`
	DtlEnvironment                *ShallowReference      `json:"dtlEnvironment,omitempty"`
	DtlEnvironmentCreationDetails *DtlEnvironmentDetails `json:"dtlEnvironmentCreationDetails,omitempty"`
	// Due date and time for test run.
	DueDate *azuredevops.Time `json:"dueDate,omitempty"`
	// Error message associated with the run.
	ErrorMessage *string    `json:"errorMessage,omitempty"`
	Filter       *RunFilter `json:"filter,omitempty"`
	// ID of the test run.
	Id *int `json:"id,omitempty"`
	// Number of Incomplete Tests.
	IncompleteTests *int `json:"incompleteTests,omitempty"`
	// true if test run is automated, false otherwise.
	IsAutomated *bool `json:"isAutomated,omitempty"`
	// The iteration to which the run belongs.
	Iteration *string `json:"iteration,omitempty"`
	// Team foundation ID of the last updated the test run.
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last updated date and time
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Name of the test run.
	Name *string `json:"name,omitempty"`
	// Number of Not Applicable Tests.
	NotApplicableTests *int `json:"notApplicableTests,omitempty"`
	// Team Foundation ID of the owner of the runs.
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Number of passed tests in the run
	PassedTests *int `json:"passedTests,omitempty"`
	// Phase/State for the testRun.
	Phase *string `json:"phase,omitempty"`
	// Reference of the pipeline to which this test run belongs.
	PipelineReference *PipelineReference `json:"pipelineReference,omitempty"`
	// Test plan associated with this test run.
	Plan *ShallowReference `json:"plan,omitempty"`
	// Post Process State.
	PostProcessState *string `json:"postProcessState,omitempty"`
	// Project associated with this run.
	Project *ShallowReference `json:"project,omitempty"`
	// Release Reference for the Test Run.
	Release *ReleaseReference `json:"release,omitempty"`
	// Release Environment Uri for TestRun.
	ReleaseEnvironmentUri *string `json:"releaseEnvironmentUri,omitempty"`
	// Release Uri for TestRun.
	ReleaseUri *string `json:"releaseUri,omitempty"`
	Revision   *int    `json:"revision,omitempty"`
	// RunSummary by outcome.
	RunStatistics *[]RunStatistic `json:"runStatistics,omitempty"`
	// Start date time of the run.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// The state of the run. Type TestRunState Valid states - Unspecified ,NotStarted, InProgress, Completed, Waiting, Aborted, NeedsInvestigation
	State *string `json:"state,omitempty"`
	// TestRun Substate.
	Substate *TestRunSubstate `json:"substate,omitempty"`
	// Tags attached with this test run.
	Tags             *[]TestTag        `json:"tags,omitempty"`
	TestMessageLogId *int              `json:"testMessageLogId,omitempty"`
	TestSettings     *ShallowReference `json:"testSettings,omitempty"`
	// Total tests in the run
	TotalTests *int `json:"totalTests,omitempty"`
	// Number of failed tests in the run.
	UnanalyzedTests *int `json:"unanalyzedTests,omitempty"`
	// Url of the test run
	Url *string `json:"url,omitempty"`
	// Web Access Url for TestRun.
	WebAccessUrl *string `json:"webAccessUrl,omitempty"`
}

type TestRun2 struct {
	BuildConfigurationId  *int              `json:"buildConfigurationId,omitempty"`
	BuildNumber           *string           `json:"buildNumber,omitempty"`
	Comment               *string           `json:"comment,omitempty"`
	CompleteDate          *azuredevops.Time `json:"completeDate,omitempty"`
	Controller            *string           `json:"controller,omitempty"`
	CoverageId            *int              `json:"coverageId,omitempty"`
	CreationDate          *azuredevops.Time `json:"creationDate,omitempty"`
	DeletedOn             *azuredevops.Time `json:"deletedOn,omitempty"`
	DropLocation          *string           `json:"dropLocation,omitempty"`
	DueDate               *azuredevops.Time `json:"dueDate,omitempty"`
	ErrorMessage          *string           `json:"errorMessage,omitempty"`
	IncompleteTests       *int              `json:"incompleteTests,omitempty"`
	IsAutomated           *bool             `json:"isAutomated,omitempty"`
	IsBvt                 *bool             `json:"isBvt,omitempty"`
	IsMigrated            *bool             `json:"isMigrated,omitempty"`
	IterationId           *int              `json:"iterationId,omitempty"`
	LastUpdated           *azuredevops.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy         *uuid.UUID        `json:"lastUpdatedBy,omitempty"`
	LegacySharePath       *string           `json:"legacySharePath,omitempty"`
	MaxReservedResultId   *int              `json:"maxReservedResultId,omitempty"`
	NotApplicableTests    *int              `json:"notApplicableTests,omitempty"`
	Owner                 *uuid.UUID        `json:"owner,omitempty"`
	PassedTests           *int              `json:"passedTests,omitempty"`
	PostProcessState      *byte             `json:"postProcessState,omitempty"`
	ProjectId             *uuid.UUID        `json:"projectId,omitempty"`
	PublicTestSettingsId  *int              `json:"publicTestSettingsId,omitempty"`
	ReleaseEnvironmentUri *string           `json:"releaseEnvironmentUri,omitempty"`
	ReleaseUri            *string           `json:"releaseUri,omitempty"`
	Revision              *int              `json:"revision,omitempty"`
	StartDate             *azuredevops.Time `json:"startDate,omitempty"`
	State                 *byte             `json:"state,omitempty"`
	TestEnvironmentId     *uuid.UUID        `json:"testEnvironmentId,omitempty"`
	TestMessageLogId      *int              `json:"testMessageLogId,omitempty"`
	TestPlanId            *int              `json:"testPlanId,omitempty"`
	TestRunContextId      *int              `json:"testRunContextId,omitempty"`
	TestRunId             *int              `json:"testRunId,omitempty"`
	TestSettingsId        *int              `json:"testSettingsId,omitempty"`
	Title                 *string           `json:"title,omitempty"`
	TotalTests            *int              `json:"totalTests,omitempty"`
	Type                  *byte             `json:"type,omitempty"`
	UnanalyzedTests       *int              `json:"unanalyzedTests,omitempty"`
	Version               *int              `json:"version,omitempty"`
}

type TestRunCanceledEvent struct {
	TestRun *TestRun `json:"testRun,omitempty"`
}

type TestRunContext2 struct {
	BuildRefId       *int       `json:"buildRefId,omitempty"`
	ProjectId        *uuid.UUID `json:"projectId,omitempty"`
	ReleaseRefId     *int       `json:"releaseRefId,omitempty"`
	SourceWorkflow   *string    `json:"sourceWorkflow,omitempty"`
	TestRunContextId *int       `json:"testRunContextId,omitempty"`
}

// Test Run Code Coverage Details
type TestRunCoverage struct {
	// Last Error
	LastError *string `json:"lastError,omitempty"`
	// List of Modules Coverage
	Modules *[]ModuleCoverage `json:"modules,omitempty"`
	// State
	State *string `json:"state,omitempty"`
	// Reference of test Run.
	TestRun *ShallowReference `json:"testRun,omitempty"`
}

type TestRunCreatedEvent struct {
	TestRun *TestRun `json:"testRun,omitempty"`
}

type TestRunEvent struct {
	TestRun *TestRun `json:"testRun,omitempty"`
}

type TestRunEx2 struct {
	BitValue      *bool             `json:"bitValue,omitempty"`
	CreatedDate   *azuredevops.Time `json:"createdDate,omitempty"`
	DateTimeValue *azuredevops.Time `json:"dateTimeValue,omitempty"`
	FieldId       *int              `json:"fieldId,omitempty"`
	FieldName     *string           `json:"fieldName,omitempty"`
	FloatValue    *float64          `json:"floatValue,omitempty"`
	GuidValue     *uuid.UUID        `json:"guidValue,omitempty"`
	IntValue      *int              `json:"intValue,omitempty"`
	ProjectId     *uuid.UUID        `json:"projectId,omitempty"`
	StringValue   *string           `json:"stringValue,omitempty"`
	TestRunId     *int              `json:"testRunId,omitempty"`
}

type TestRunExtended2 struct {
	AutEnvironmentUrl  *string    `json:"autEnvironmentUrl,omitempty"`
	CsmContent         *string    `json:"csmContent,omitempty"`
	CsmParameters      *string    `json:"csmParameters,omitempty"`
	ProjectId          *uuid.UUID `json:"projectId,omitempty"`
	SourceFilter       *string    `json:"sourceFilter,omitempty"`
	SubscriptionName   *string    `json:"subscriptionName,omitempty"`
	Substate           *byte      `json:"substate,omitempty"`
	TestCaseFilter     *string    `json:"testCaseFilter,omitempty"`
	TestEnvironmentUrl *string    `json:"testEnvironmentUrl,omitempty"`
	TestRunId          *int       `json:"testRunId,omitempty"`
}

// The types of outcomes for test run.
type TestRunOutcome string

type testRunOutcomeValuesType struct {
	Passed      TestRunOutcome
	Failed      TestRunOutcome
	NotImpacted TestRunOutcome
	Others      TestRunOutcome
}

var TestRunOutcomeValues = testRunOutcomeValuesType{
	// Run with zero failed tests and has at least one impacted test
	Passed: "passed",
	// Run with at-least one failed test.
	Failed: "failed",
	// Run with no impacted tests.
	NotImpacted: "notImpacted",
	// Runs with All tests in other category.
	Others: "others",
}

// The types of publish context for run.
type TestRunPublishContext string

type testRunPublishContextValuesType struct {
	Build   TestRunPublishContext
	Release TestRunPublishContext
	All     TestRunPublishContext
}

var TestRunPublishContextValues = testRunPublishContextValuesType{
	// Run is published for Build Context.
	Build: "build",
	// Run is published for Release Context.
	Release: "release",
	// Run is published for any Context.
	All: "all",
}

type TestRunStartedEvent struct {
	TestRun *TestRun `json:"testRun,omitempty"`
}

// The types of states for test run.
type TestRunState string

type testRunStateValuesType struct {
	Unspecified        TestRunState
	NotStarted         TestRunState
	InProgress         TestRunState
	Completed          TestRunState
	Aborted            TestRunState
	Waiting            TestRunState
	NeedsInvestigation TestRunState
}

var TestRunStateValues = testRunStateValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// The run is still being created.  No tests have started yet.
	NotStarted: "notStarted",
	// Tests are running.
	InProgress: "inProgress",
	// All tests have completed or been skipped.
	Completed: "completed",
	// Run is stopped and remaining tests have been aborted
	Aborted: "aborted",
	// Run is currently initializing This is a legacy state and should not be used any more
	Waiting: "waiting",
	// Run requires investigation because of a test point failure This is a legacy state and should not be used any more
	NeedsInvestigation: "needsInvestigation",
}

// Test run statistics.
type TestRunStatistic struct {
	Run           *ShallowReference `json:"run,omitempty"`
	RunStatistics *[]RunStatistic   `json:"runStatistics,omitempty"`
}

// The types of sub states for test run. It gives the user more info about the test run beyond the high level test run state
type TestRunSubstate string

type testRunSubstateValuesType struct {
	None                   TestRunSubstate
	CreatingEnvironment    TestRunSubstate
	RunningTests           TestRunSubstate
	CanceledByUser         TestRunSubstate
	AbortedBySystem        TestRunSubstate
	TimedOut               TestRunSubstate
	PendingAnalysis        TestRunSubstate
	Analyzed               TestRunSubstate
	CancellationInProgress TestRunSubstate
}

var TestRunSubstateValues = testRunSubstateValuesType{
	// Run with noState.
	None: "none",
	// Run state while Creating Environment.
	CreatingEnvironment: "creatingEnvironment",
	// Run state while Running Tests.
	RunningTests: "runningTests",
	// Run state while Creating Environment.
	CanceledByUser: "canceledByUser",
	// Run state when it is Aborted By the System.
	AbortedBySystem: "abortedBySystem",
	// Run state when run has timedOut.
	TimedOut: "timedOut",
	// Run state while Pending Analysis.
	PendingAnalysis: "pendingAnalysis",
	// Run state after being Analysed.
	Analyzed: "analyzed",
	// Run state when cancellation is in Progress.
	CancellationInProgress: "cancellationInProgress",
}

type TestRunSummary2 struct {
	IsRerun              *bool             `json:"isRerun,omitempty"`
	ProjectId            *uuid.UUID        `json:"projectId,omitempty"`
	ResultCount          *int              `json:"resultCount,omitempty"`
	ResultDuration       *uint64           `json:"resultDuration,omitempty"`
	RunDuration          *uint64           `json:"runDuration,omitempty"`
	TestOutcome          *byte             `json:"testOutcome,omitempty"`
	TestRunCompletedDate *azuredevops.Time `json:"testRunCompletedDate,omitempty"`
	TestRunContextId     *int              `json:"testRunContextId,omitempty"`
	TestRunId            *int              `json:"testRunId,omitempty"`
	TestRunStatsId       *int              `json:"testRunStatsId,omitempty"`
}

type TestRunWithDtlEnvEvent struct {
	TestRun                   *TestRun    `json:"testRun,omitempty"`
	ConfigurationIds          *[]int      `json:"configurationIds,omitempty"`
	MappedTestRunEventType    *string     `json:"mappedTestRunEventType,omitempty"`
	RunTimeout                interface{} `json:"runTimeout,omitempty"`
	TestConfigurationsMapping *string     `json:"testConfigurationsMapping,omitempty"`
}

// Test Session
type TestSession struct {
	// Area path of the test session
	Area *ShallowReference `json:"area,omitempty"`
	// Comments in the test session
	Comment *string `json:"comment,omitempty"`
	// Duration of the session
	EndDate *azuredevops.Time `json:"endDate,omitempty"`
	// Id of the test session
	Id *int `json:"id,omitempty"`
	// Last Updated By  Reference
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last updated date
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Owner of the test session
	Owner *webapi.IdentityRef `json:"owner,omitempty"`
	// Project to which the test session belongs
	Project *ShallowReference `json:"project,omitempty"`
	// Generic store for test session data
	PropertyBag *PropertyBag `json:"propertyBag,omitempty"`
	// Revision of the test session
	Revision *int `json:"revision,omitempty"`
	// Source of the test session
	Source *TestSessionSource `json:"source,omitempty"`
	// Start date
	StartDate *azuredevops.Time `json:"startDate,omitempty"`
	// State of the test session
	State *TestSessionState `json:"state,omitempty"`
	// Title of the test session
	Title *string `json:"title,omitempty"`
	// Url of Test Session Resource
	Url *string `json:"url,omitempty"`
}

type TestSessionExploredWorkItemReference struct {
	// Id of the workitem
	Id *int `json:"id,omitempty"`
	// Type of the workitem
	Type *string `json:"type,omitempty"`
	// Workitem references of workitems filed as a part of the current workitem exploration.
	AssociatedWorkItems *[]TestSessionWorkItemReference `json:"associatedWorkItems,omitempty"`
	// Time when exploration of workitem ended.
	EndTime *azuredevops.Time `json:"endTime,omitempty"`
	// Time when explore of workitem was started.
	StartTime *azuredevops.Time `json:"startTime,omitempty"`
}

// Represents the source from which the test session was created
type TestSessionSource string

type testSessionSourceValuesType struct {
	Unknown               TestSessionSource
	XtDesktop             TestSessionSource
	FeedbackDesktop       TestSessionSource
	XtWeb                 TestSessionSource
	FeedbackWeb           TestSessionSource
	XtDesktop2            TestSessionSource
	SessionInsightsForAll TestSessionSource
}

var TestSessionSourceValues = testSessionSourceValuesType{
	// Source of test session uncertain as it is stale
	Unknown: "unknown",
	// The session was created from Microsoft Test Manager exploratory desktop tool.
	XtDesktop: "xtDesktop",
	// The session was created from feedback client.
	FeedbackDesktop: "feedbackDesktop",
	// The session was created from browser extension.
	XtWeb: "xtWeb",
	// The session was created from browser extension.
	FeedbackWeb: "feedbackWeb",
	// The session was created from web access using Microsoft Test Manager exploratory desktop tool.
	XtDesktop2: "xtDesktop2",
	// To show sessions from all supported sources.
	SessionInsightsForAll: "sessionInsightsForAll",
}

// Represents the state of the test session.
type TestSessionState string

type testSessionStateValuesType struct {
	Unspecified TestSessionState
	NotStarted  TestSessionState
	InProgress  TestSessionState
	Paused      TestSessionState
	Completed   TestSessionState
	Declined    TestSessionState
}

var TestSessionStateValues = testSessionStateValuesType{
	// Only used during an update to preserve the existing value.
	Unspecified: "unspecified",
	// The session is still being created.
	NotStarted: "notStarted",
	// The session is running.
	InProgress: "inProgress",
	// The session has paused.
	Paused: "paused",
	// The session has completed.
	Completed: "completed",
	// This is required for Feedback session which are declined
	Declined: "declined",
}

type TestSessionWorkItemReference struct {
	// Id of the workitem
	Id *int `json:"id,omitempty"`
	// Type of the workitem
	Type *string `json:"type,omitempty"`
}

// Represents the test settings of the run. Used to create test settings and fetch test settings
type TestSettings struct {
	// Area path required to create test settings
	AreaPath *string `json:"areaPath,omitempty"`
	// Description of the test settings. Used in create test settings.
	Description *string `json:"description,omitempty"`
	// Indicates if the tests settings is public or private.Used in create test settings.
	IsPublic *bool `json:"isPublic,omitempty"`
	// Xml string of machine roles. Used in create test settings.
	MachineRoles *string `json:"machineRoles,omitempty"`
	// Test settings content.
	TestSettingsContent *string `json:"testSettingsContent,omitempty"`
	// Test settings id.
	TestSettingsId *int `json:"testSettingsId,omitempty"`
	// Test settings name.
	TestSettingsName *string `json:"testSettingsName,omitempty"`
}

// Represents the test settings of the run. Used to create test settings and fetch test settings
type TestSettings2 struct {
	// Area path required to create test settings
	AreaPath    *string             `json:"areaPath,omitempty"`
	CreatedBy   *webapi.IdentityRef `json:"createdBy,omitempty"`
	CreatedDate *azuredevops.Time   `json:"createdDate,omitempty"`
	// Description of the test settings. Used in create test settings.
	Description *string `json:"description,omitempty"`
	// Indicates if the tests settings is public or private.Used in create test settings.
	IsPublic        *bool               `json:"isPublic,omitempty"`
	LastUpdatedBy   *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate *azuredevops.Time   `json:"lastUpdatedDate,omitempty"`
	// Xml string of machine roles. Used in create test settings.
	MachineRoles *string `json:"machineRoles,omitempty"`
	// Test settings content.
	TestSettingsContent *string `json:"testSettingsContent,omitempty"`
	// Test settings id.
	TestSettingsId *int `json:"testSettingsId,omitempty"`
	// Test settings name.
	TestSettingsName *string `json:"testSettingsName,omitempty"`
}

type TestSettingsMachineRole struct {
	IsExecution *bool   `json:"isExecution,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// Represents a sub result of a test result.
type TestSubResult struct {
	// Comment in sub result.
	Comment *string `json:"comment,omitempty"`
	// Time when test execution completed.
	CompletedDate *azuredevops.Time `json:"completedDate,omitempty"`
	// Machine where test executed.
	ComputerName *string `json:"computerName,omitempty"`
	// Reference to test configuration.
	Configuration *ShallowReference `json:"configuration,omitempty"`
	// Additional properties of sub result.
	CustomFields *[]CustomTestField `json:"customFields,omitempty"`
	// Name of sub result.
	DisplayName *string `json:"displayName,omitempty"`
	// Duration of test execution.
	DurationInMs *uint64 `json:"durationInMs,omitempty"`
	// Error message in sub result.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// ID of sub result.
	Id *int `json:"id,omitempty"`
	// Time when result last updated.
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Outcome of sub result.
	Outcome *string `json:"outcome,omitempty"`
	// Immediate parent ID of sub result.
	ParentId *int `json:"parentId,omitempty"`
	// Hierarchy type of the result, default value of None means its leaf node.
	ResultGroupType *ResultGroupType `json:"resultGroupType,omitempty"`
	// Index number of sub result.
	SequenceId *int `json:"sequenceId,omitempty"`
	// Stacktrace.
	StackTrace *string `json:"stackTrace,omitempty"`
	// Time when test execution started.
	StartedDate *azuredevops.Time `json:"startedDate,omitempty"`
	// List of sub results inside a sub result, if ResultGroupType is not None, it holds corresponding type sub results.
	SubResults *[]TestSubResult `json:"subResults,omitempty"`
	// Reference to test result.
	TestResult *TestCaseResultIdentifier `json:"testResult,omitempty"`
	// Url of sub result.
	Url *string `json:"url,omitempty"`
}

// Test suite
type TestSuite struct {
	// Area uri of the test suite.
	AreaUri *string `json:"areaUri,omitempty"`
	// Child test suites of current test suite.
	Children *[]TestSuite `json:"children,omitempty"`
	// Test suite default configuration.
	DefaultConfigurations *[]ShallowReference `json:"defaultConfigurations,omitempty"`
	// Test suite default testers.
	DefaultTesters *[]ShallowReference `json:"defaultTesters,omitempty"`
	// Id of test suite.
	Id *int `json:"id,omitempty"`
	// Default configuration was inherited or not.
	InheritDefaultConfigurations *bool `json:"inheritDefaultConfigurations,omitempty"`
	// Last error for test suite.
	LastError *string `json:"lastError,omitempty"`
	// Last populated date.
	LastPopulatedDate *azuredevops.Time `json:"lastPopulatedDate,omitempty"`
	// IdentityRef of user who has updated test suite recently.
	LastUpdatedBy *webapi.IdentityRef `json:"lastUpdatedBy,omitempty"`
	// Last update date.
	LastUpdatedDate *azuredevops.Time `json:"lastUpdatedDate,omitempty"`
	// Name of test suite.
	Name *string `json:"name,omitempty"`
	// Test suite parent shallow reference.
	Parent *ShallowReference `json:"parent,omitempty"`
	// Test plan to which the test suite belongs.
	Plan *ShallowReference `json:"plan,omitempty"`
	// Test suite project shallow reference.
	Project *ShallowReference `json:"project,omitempty"`
	// Test suite query string, for dynamic suites.
	QueryString *string `json:"queryString,omitempty"`
	// Test suite requirement id.
	RequirementId *int `json:"requirementId,omitempty"`
	// Test suite revision.
	Revision *int `json:"revision,omitempty"`
	// State of test suite.
	State *string `json:"state,omitempty"`
	// List of shallow reference of suites.
	Suites *[]ShallowReference `json:"suites,omitempty"`
	// Test suite type.
	SuiteType *string `json:"suiteType,omitempty"`
	// Test cases count.
	TestCaseCount *int `json:"testCaseCount,omitempty"`
	// Test case url.
	TestCasesUrl *string `json:"testCasesUrl,omitempty"`
	// Used in tree view. If test suite is root suite then, it is name of plan otherwise title of the suite.
	Text *string `json:"text,omitempty"`
	// Url of test suite.
	Url *string `json:"url,omitempty"`
}

// Test suite clone request
type TestSuiteCloneRequest struct {
	// Clone options for cloning the test suite.
	CloneOptions *CloneOptions `json:"cloneOptions,omitempty"`
	// Suite id under which, we have to clone the suite.
	DestinationSuiteId *int `json:"destinationSuiteId,omitempty"`
	// Destination suite project name.
	DestinationSuiteProjectName *string `json:"destinationSuiteProjectName,omitempty"`
}

type TestSummaryForWorkItem struct {
	Summary  *AggregatedDataForResultTrend `json:"summary,omitempty"`
	WorkItem *WorkItemReference            `json:"workItem,omitempty"`
}

// Tag attached to a run or result.
type TestTag struct {
	// Name of the tag, alphanumeric value less than 30 chars
	Name *string `json:"name,omitempty"`
}

// Test tag summary for build or release grouped by test run.
type TestTagSummary struct {
	// Dictionary which contains tags associated with a test run.
	TagsGroupByTestArtifact *map[int][]TestTag `json:"tagsGroupByTestArtifact,omitempty"`
}

// Tags to update to a run or result.
type TestTagsUpdateModel struct {
	Tags *[]azuredevops.KeyValuePair `json:"tags,omitempty"`
}

type TestToWorkItemLinks struct {
	Test      *TestMethod          `json:"test,omitempty"`
	WorkItems *[]WorkItemReference `json:"workItems,omitempty"`
}

type TestVariable struct {
	// Description of the test variable
	Description *string `json:"description,omitempty"`
	// Id of the test variable
	Id *int `json:"id,omitempty"`
	// Name of the test variable
	Name *string `json:"name,omitempty"`
	// Project to which the test variable belongs
	Project *ShallowReference `json:"project,omitempty"`
	// Revision
	Revision *int `json:"revision,omitempty"`
	// Url of the test variable
	Url *string `json:"url,omitempty"`
	// List of allowed values
	Values *[]string `json:"values,omitempty"`
}

type UpdatedProperties struct {
	Id                *int              `json:"id,omitempty"`
	LastUpdated       *azuredevops.Time `json:"lastUpdated,omitempty"`
	LastUpdatedBy     *uuid.UUID        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedByName *string           `json:"lastUpdatedByName,omitempty"`
	Revision          *int              `json:"revision,omitempty"`
}

type UpdateTestRunRequest struct {
	AttachmentsToAdd    *[]TestResultAttachment         `json:"attachmentsToAdd,omitempty"`
	AttachmentsToDelete *[]TestResultAttachmentIdentity `json:"attachmentsToDelete,omitempty"`
	ProjectName         *string                         `json:"projectName,omitempty"`
	ShouldHyderate      *bool                           `json:"shouldHyderate,omitempty"`
	TestRun             *LegacyTestRun                  `json:"testRun,omitempty"`
}

type UpdateTestRunResponse struct {
	AttachmentIds     *[]int             `json:"attachmentIds,omitempty"`
	UpdatedProperties *UpdatedProperties `json:"updatedProperties,omitempty"`
}

type UploadAttachmentsRequest struct {
	Attachments   *[]HttpPostedTcmAttachment `json:"attachments,omitempty"`
	RequestParams *map[string]string         `json:"requestParams,omitempty"`
}

// WorkItem reference Details.
type WorkItemReference struct {
	// WorkItem Id.
	Id *string `json:"id,omitempty"`
	// WorkItem Name.
	Name *string `json:"name,omitempty"`
	// WorkItem Type.
	Type *string `json:"type,omitempty"`
	// WorkItem Url. Valid Values : (Bug, Task, User Story, Test Case)
	Url *string `json:"url,omitempty"`
	// WorkItem WebUrl.
	WebUrl *string `json:"webUrl,omitempty"`
}

type WorkItemToTestLinks struct {
	ExecutedIn *Service           `json:"executedIn,omitempty"`
	Tests      *[]TestMethod      `json:"tests,omitempty"`
	WorkItem   *WorkItemReference `json:"workItem,omitempty"`
}
