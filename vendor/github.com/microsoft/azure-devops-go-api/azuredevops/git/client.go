// --------------------------------------------------------------------------------------------
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
// --------------------------------------------------------------------------------------------
// Generated file, DO NOT EDIT
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// --------------------------------------------------------------------------------------------

package git

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/policy"
	"github.com/microsoft/azure-devops-go-api/azuredevops/webapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var ResourceAreaId, _ = uuid.Parse("4e080c62-fa21-4fbc-8fef-2a10a2b38049")

type Client interface {
	// [Preview API] Create an annotated tag.
	CreateAnnotatedTag(context.Context, CreateAnnotatedTagArgs) (*GitAnnotatedTag, error)
	// [Preview API] Attach a new file to a pull request.
	CreateAttachment(context.Context, CreateAttachmentArgs) (*Attachment, error)
	// [Preview API] Cherry pick a specific commit or commits that are associated to a pull request into a new branch.
	CreateCherryPick(context.Context, CreateCherryPickArgs) (*GitCherryPick, error)
	// Create a comment on a specific thread in a pull request (up to 500 comments can be created per thread).
	CreateComment(context.Context, CreateCommentArgs) (*Comment, error)
	// Create Git commit status.
	CreateCommitStatus(context.Context, CreateCommitStatusArgs) (*GitStatus, error)
	// [Preview API] Creates a ref favorite
	CreateFavorite(context.Context, CreateFavoriteArgs) (*GitRefFavorite, error)
	// [Preview API] Request that another repository's refs be fetched into this one. It syncs two existing forks. To create a fork, please see the <a href="https://docs.microsoft.com/en-us/rest/api/vsts/git/repositories/create?view=azure-devops-rest-5.1"> repositories endpoint</a>
	CreateForkSyncRequest(context.Context, CreateForkSyncRequestArgs) (*GitForkSyncRequest, error)
	// [Preview API] Create an import request.
	CreateImportRequest(context.Context, CreateImportRequestArgs) (*GitImportRequest, error)
	// [Preview API] Add a like on a comment.
	CreateLike(context.Context, CreateLikeArgs) error
	// [Preview API] Request a git merge operation. Currently we support merging only 2 commits.
	CreateMergeRequest(context.Context, CreateMergeRequestArgs) (*GitMerge, error)
	// Create a pull request.
	CreatePullRequest(context.Context, CreatePullRequestArgs) (*GitPullRequest, error)
	// [Preview API] Create a pull request status on the iteration. This operation will have the same result as Create status on pull request with specified iteration ID in the request body.
	CreatePullRequestIterationStatus(context.Context, CreatePullRequestIterationStatusArgs) (*GitPullRequestStatus, error)
	// [Preview API] Create a label for a specified pull request. The only required field is the name of the new label.
	CreatePullRequestLabel(context.Context, CreatePullRequestLabelArgs) (*core.WebApiTagDefinition, error)
	// Add a reviewer to a pull request or cast a vote.
	CreatePullRequestReviewer(context.Context, CreatePullRequestReviewerArgs) (*IdentityRefWithVote, error)
	// Add reviewers to a pull request.
	CreatePullRequestReviewers(context.Context, CreatePullRequestReviewersArgs) (*[]IdentityRefWithVote, error)
	// [Preview API] Create a pull request status.
	CreatePullRequestStatus(context.Context, CreatePullRequestStatusArgs) (*GitPullRequestStatus, error)
	// Push changes to the repository.
	CreatePush(context.Context, CreatePushArgs) (*GitPush, error)
	// Create a git repository in a team project.
	CreateRepository(context.Context, CreateRepositoryArgs) (*GitRepository, error)
	// [Preview API] Starts the operation to create a new branch which reverts changes introduced by either a specific commit or commits that are associated to a pull request.
	CreateRevert(context.Context, CreateRevertArgs) (*GitRevert, error)
	// Create a thread in a pull request.
	CreateThread(context.Context, CreateThreadArgs) (*GitPullRequestCommentThread, error)
	// [Preview API] Delete a pull request attachment.
	DeleteAttachment(context.Context, DeleteAttachmentArgs) error
	// Delete a comment associated with a specific thread in a pull request.
	DeleteComment(context.Context, DeleteCommentArgs) error
	// [Preview API] Delete a like on a comment.
	DeleteLike(context.Context, DeleteLikeArgs) error
	// [Preview API] Delete pull request iteration status.
	DeletePullRequestIterationStatus(context.Context, DeletePullRequestIterationStatusArgs) error
	// [Preview API] Removes a label from the set of those assigned to the pull request.
	DeletePullRequestLabels(context.Context, DeletePullRequestLabelsArgs) error
	// Remove a reviewer from a pull request.
	DeletePullRequestReviewer(context.Context, DeletePullRequestReviewerArgs) error
	// [Preview API] Delete pull request status.
	DeletePullRequestStatus(context.Context, DeletePullRequestStatusArgs) error
	// [Preview API] Deletes the refs favorite specified
	DeleteRefFavorite(context.Context, DeleteRefFavoriteArgs) error
	// Delete a git repository
	DeleteRepository(context.Context, DeleteRepositoryArgs) error
	// [Preview API] Destroy (hard delete) a soft-deleted Git repository.
	DeleteRepositoryFromRecycleBin(context.Context, DeleteRepositoryFromRecycleBinArgs) error
	// [Preview API] Get an annotated tag.
	GetAnnotatedTag(context.Context, GetAnnotatedTagArgs) (*GitAnnotatedTag, error)
	// [Preview API] Get the file content of a pull request attachment.
	GetAttachmentContent(context.Context, GetAttachmentContentArgs) (io.ReadCloser, error)
	// [Preview API] Get a list of files attached to a given pull request.
	GetAttachments(context.Context, GetAttachmentsArgs) (*[]Attachment, error)
	// [Preview API] Get the file content of a pull request attachment.
	GetAttachmentZip(context.Context, GetAttachmentZipArgs) (io.ReadCloser, error)
	// Get a single blob.
	GetBlob(context.Context, GetBlobArgs) (*GitBlobRef, error)
	// Get a single blob.
	GetBlobContent(context.Context, GetBlobContentArgs) (io.ReadCloser, error)
	// Gets one or more blobs in a zip file download.
	GetBlobsZip(context.Context, GetBlobsZipArgs) (io.ReadCloser, error)
	// Get a single blob.
	GetBlobZip(context.Context, GetBlobZipArgs) (io.ReadCloser, error)
	// Retrieve statistics about a single branch.
	GetBranch(context.Context, GetBranchArgs) (*GitBranchStats, error)
	// Retrieve statistics about all branches within a repository.
	GetBranches(context.Context, GetBranchesArgs) (*[]GitBranchStats, error)
	// Retrieve changes for a particular commit.
	GetChanges(context.Context, GetChangesArgs) (*GitCommitChanges, error)
	// [Preview API] Retrieve information about a cherry pick by cherry pick Id.
	GetCherryPick(context.Context, GetCherryPickArgs) (*GitCherryPick, error)
	// [Preview API] Retrieve information about a cherry pick for a specific branch.
	GetCherryPickForRefName(context.Context, GetCherryPickForRefNameArgs) (*GitCherryPick, error)
	// Retrieve a comment associated with a specific thread in a pull request.
	GetComment(context.Context, GetCommentArgs) (*Comment, error)
	// Retrieve all comments associated with a specific thread in a pull request.
	GetComments(context.Context, GetCommentsArgs) (*[]Comment, error)
	// Retrieve a particular commit.
	GetCommit(context.Context, GetCommitArgs) (*GitCommit, error)
	// Find the closest common commit (the merge base) between base and target commits, and get the diff between either the base and target commits or common and target commits.
	GetCommitDiffs(context.Context, GetCommitDiffsArgs) (*GitCommitDiffs, error)
	// Retrieve git commits for a project
	GetCommits(context.Context, GetCommitsArgs) (*[]GitCommitRef, error)
	// Retrieve git commits for a project matching the search criteria
	GetCommitsBatch(context.Context, GetCommitsBatchArgs) (*[]GitCommitRef, error)
	// [Preview API] Retrieve deleted git repositories.
	GetDeletedRepositories(context.Context, GetDeletedRepositoriesArgs) (*[]GitDeletedRepository, error)
	// [Preview API] Retrieve all forks of a repository in the collection.
	GetForks(context.Context, GetForksArgs) (*[]GitRepositoryRef, error)
	// [Preview API] Get a specific fork sync operation's details.
	GetForkSyncRequest(context.Context, GetForkSyncRequestArgs) (*GitForkSyncRequest, error)
	// [Preview API] Retrieve all requested fork sync operations on this repository.
	GetForkSyncRequests(context.Context, GetForkSyncRequestsArgs) (*[]GitForkSyncRequest, error)
	// [Preview API] Retrieve a particular import request.
	GetImportRequest(context.Context, GetImportRequestArgs) (*GitImportRequest, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
	GetItem(context.Context, GetItemArgs) (*GitItem, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
	GetItemContent(context.Context, GetItemContentArgs) (io.ReadCloser, error)
	// Get Item Metadata and/or Content for a collection of items. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
	GetItems(context.Context, GetItemsArgs) (*[]GitItem, error)
	// Post for retrieving a creating a batch out of a set of items in a repo / project given a list of paths or a long path
	GetItemsBatch(context.Context, GetItemsBatchArgs) (*[][]GitItem, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
	GetItemText(context.Context, GetItemTextArgs) (io.ReadCloser, error)
	// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
	GetItemZip(context.Context, GetItemZipArgs) (io.ReadCloser, error)
	// [Preview API] Get likes for a comment.
	GetLikes(context.Context, GetLikesArgs) (*[]webapi.IdentityRef, error)
	// [Preview API] Find the merge bases of two commits, optionally across forks. If otherRepositoryId is not specified, the merge bases will only be calculated within the context of the local repositoryNameOrId.
	GetMergeBases(context.Context, GetMergeBasesArgs) (*[]GitCommitRef, error)
	// [Preview API] Get a specific merge operation's details.
	GetMergeRequest(context.Context, GetMergeRequestArgs) (*GitMerge, error)
	// [Preview API] Retrieve a list of policy configurations by a given set of scope/filtering criteria.
	GetPolicyConfigurations(context.Context, GetPolicyConfigurationsArgs) (*GitPolicyConfigurationResponse, error)
	// Retrieve a pull request.
	GetPullRequest(context.Context, GetPullRequestArgs) (*GitPullRequest, error)
	// Retrieve a pull request.
	GetPullRequestById(context.Context, GetPullRequestByIdArgs) (*GitPullRequest, error)
	// Get the commits for the specified pull request.
	GetPullRequestCommits(context.Context, GetPullRequestCommitsArgs) (*GetPullRequestCommitsResponseValue, error)
	// Get the specified iteration for a pull request.
	GetPullRequestIteration(context.Context, GetPullRequestIterationArgs) (*GitPullRequestIteration, error)
	// Retrieve the changes made in a pull request between two iterations.
	GetPullRequestIterationChanges(context.Context, GetPullRequestIterationChangesArgs) (*GitPullRequestIterationChanges, error)
	// Get the commits for the specified iteration of a pull request.
	GetPullRequestIterationCommits(context.Context, GetPullRequestIterationCommitsArgs) (*[]GitCommitRef, error)
	// Get the list of iterations for the specified pull request.
	GetPullRequestIterations(context.Context, GetPullRequestIterationsArgs) (*[]GitPullRequestIteration, error)
	// [Preview API] Get the specific pull request iteration status by ID. The status ID is unique within the pull request across all iterations.
	GetPullRequestIterationStatus(context.Context, GetPullRequestIterationStatusArgs) (*GitPullRequestStatus, error)
	// [Preview API] Get all the statuses associated with a pull request iteration.
	GetPullRequestIterationStatuses(context.Context, GetPullRequestIterationStatusesArgs) (*[]GitPullRequestStatus, error)
	// [Preview API] Retrieves a single label that has been assigned to a pull request.
	GetPullRequestLabel(context.Context, GetPullRequestLabelArgs) (*core.WebApiTagDefinition, error)
	// [Preview API] Get all the labels assigned to a pull request.
	GetPullRequestLabels(context.Context, GetPullRequestLabelsArgs) (*[]core.WebApiTagDefinition, error)
	// [Preview API] Get external properties of the pull request.
	GetPullRequestProperties(context.Context, GetPullRequestPropertiesArgs) (interface{}, error)
	// This API is used to find what pull requests are related to a given commit.  It can be used to either find the pull request that created a particular merge commit or it can be used to find all pull requests that have ever merged a particular commit.  The input is a list of queries which each contain a list of commits. For each commit that you search against, you will get back a dictionary of commit -> pull requests.
	GetPullRequestQuery(context.Context, GetPullRequestQueryArgs) (*GitPullRequestQuery, error)
	// Retrieve information about a particular reviewer on a pull request
	GetPullRequestReviewer(context.Context, GetPullRequestReviewerArgs) (*IdentityRefWithVote, error)
	// Retrieve the reviewers for a pull request
	GetPullRequestReviewers(context.Context, GetPullRequestReviewersArgs) (*[]IdentityRefWithVote, error)
	// Retrieve all pull requests matching a specified criteria.
	GetPullRequests(context.Context, GetPullRequestsArgs) (*[]GitPullRequest, error)
	// Retrieve all pull requests matching a specified criteria.
	GetPullRequestsByProject(context.Context, GetPullRequestsByProjectArgs) (*[]GitPullRequest, error)
	// [Preview API] Get the specific pull request status by ID. The status ID is unique within the pull request across all iterations.
	GetPullRequestStatus(context.Context, GetPullRequestStatusArgs) (*GitPullRequestStatus, error)
	// [Preview API] Get all the statuses associated with a pull request.
	GetPullRequestStatuses(context.Context, GetPullRequestStatusesArgs) (*[]GitPullRequestStatus, error)
	// Retrieve a thread in a pull request.
	GetPullRequestThread(context.Context, GetPullRequestThreadArgs) (*GitPullRequestCommentThread, error)
	// Retrieve a list of work items associated with a pull request.
	GetPullRequestWorkItemRefs(context.Context, GetPullRequestWorkItemRefsArgs) (*[]webapi.ResourceRef, error)
	// Retrieves a particular push.
	GetPush(context.Context, GetPushArgs) (*GitPush, error)
	// Retrieve a list of commits associated with a particular push.
	GetPushCommits(context.Context, GetPushCommitsArgs) (*[]GitCommitRef, error)
	// Retrieves pushes associated with the specified repository.
	GetPushes(context.Context, GetPushesArgs) (*[]GitPush, error)
	// [Preview API] Retrieve soft-deleted git repositories from the recycle bin.
	GetRecycleBinRepositories(context.Context, GetRecycleBinRepositoriesArgs) (*[]GitDeletedRepository, error)
	// [Preview API] Gets the refs favorite for a favorite Id.
	GetRefFavorite(context.Context, GetRefFavoriteArgs) (*GitRefFavorite, error)
	// [Preview API] Gets the refs favorites for a repo and an identity.
	GetRefFavorites(context.Context, GetRefFavoritesArgs) (*[]GitRefFavorite, error)
	// Queries the provided repository for its refs and returns them.
	GetRefs(context.Context, GetRefsArgs) (*GetRefsResponseValue, error)
	// Retrieve git repositories.
	GetRepositories(context.Context, GetRepositoriesArgs) (*[]GitRepository, error)
	// Retrieve a git repository.
	GetRepository(context.Context, GetRepositoryArgs) (*GitRepository, error)
	// Retrieve a git repository.
	GetRepositoryWithParent(context.Context, GetRepositoryWithParentArgs) (*GitRepository, error)
	// [Preview API] Retrieve information about a revert operation by revert Id.
	GetRevert(context.Context, GetRevertArgs) (*GitRevert, error)
	// [Preview API] Retrieve information about a revert operation for a specific branch.
	GetRevertForRefName(context.Context, GetRevertForRefNameArgs) (*GitRevert, error)
	// Get statuses associated with the Git commit.
	GetStatuses(context.Context, GetStatusesArgs) (*[]GitStatus, error)
	// [Preview API] Retrieve a pull request suggestion for a particular repository or team project.
	GetSuggestions(context.Context, GetSuggestionsArgs) (*[]GitSuggestion, error)
	// Retrieve all threads in a pull request.
	GetThreads(context.Context, GetThreadsArgs) (*[]GitPullRequestCommentThread, error)
	// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
	GetTree(context.Context, GetTreeArgs) (*GitTreeRef, error)
	// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
	GetTreeZip(context.Context, GetTreeZipArgs) (io.ReadCloser, error)
	// [Preview API] Retrieve import requests for a repository.
	QueryImportRequests(context.Context, QueryImportRequestsArgs) (*[]GitImportRequest, error)
	// [Preview API] Recover a soft-deleted Git repository. Recently deleted repositories go into a soft-delete state for a period of time before they are hard deleted and become unrecoverable.
	RestoreRepositoryFromRecycleBin(context.Context, RestoreRepositoryFromRecycleBinArgs) (*GitRepository, error)
	// [Preview API] Sends an e-mail notification about a specific pull request to a set of recipients
	SharePullRequest(context.Context, SharePullRequestArgs) error
	// Update a comment associated with a specific thread in a pull request.
	UpdateComment(context.Context, UpdateCommentArgs) (*Comment, error)
	// [Preview API] Retry or abandon a failed import request.
	UpdateImportRequest(context.Context, UpdateImportRequestArgs) (*GitImportRequest, error)
	// Update a pull request
	UpdatePullRequest(context.Context, UpdatePullRequestArgs) (*GitPullRequest, error)
	// [Preview API] Update pull request iteration statuses collection. The only supported operation type is `remove`.
	UpdatePullRequestIterationStatuses(context.Context, UpdatePullRequestIterationStatusesArgs) error
	// [Preview API] Create or update pull request external properties. The patch operation can be `add`, `replace` or `remove`. For `add` operation, the path can be empty. If the path is empty, the value must be a list of key value pairs. For `replace` operation, the path cannot be empty. If the path does not exist, the property will be added to the collection. For `remove` operation, the path cannot be empty. If the path does not exist, no action will be performed.
	UpdatePullRequestProperties(context.Context, UpdatePullRequestPropertiesArgs) (interface{}, error)
	// Reset the votes of multiple reviewers on a pull request.  NOTE: This endpoint only supports updating votes, but does not support updating required reviewers (use policy) or display names.
	UpdatePullRequestReviewers(context.Context, UpdatePullRequestReviewersArgs) error
	// [Preview API] Update pull request statuses collection. The only supported operation type is `remove`.
	UpdatePullRequestStatuses(context.Context, UpdatePullRequestStatusesArgs) error
	// Lock or Unlock a branch.
	UpdateRef(context.Context, UpdateRefArgs) (*GitRef, error)
	// Creating, updating, or deleting refs(branches).
	UpdateRefs(context.Context, UpdateRefsArgs) (*[]GitRefUpdateResult, error)
	// Updates the Git repository with either a new repo name or a new default branch.
	UpdateRepository(context.Context, UpdateRepositoryArgs) (*GitRepository, error)
	// Update a thread in a pull request.
	UpdateThread(context.Context, UpdateThreadArgs) (*GitPullRequestCommentThread, error)
}

type ClientImpl struct {
	Client azuredevops.Client
}

func NewClient(ctx context.Context, connection *azuredevops.Connection) (Client, error) {
	client, err := connection.GetClientByResourceAreaId(ctx, ResourceAreaId)
	if err != nil {
		return nil, err
	}
	return &ClientImpl{
		Client: *client,
	}, nil
}

// [Preview API] Create an annotated tag.
func (client *ClientImpl) CreateAnnotatedTag(ctx context.Context, args CreateAnnotatedTagArgs) (*GitAnnotatedTag, error) {
	if args.TagObject == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.TagObject"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.TagObject)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("5e8a8081-3851-4626-b677-9891cc04102e")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitAnnotatedTag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateAnnotatedTag function
type CreateAnnotatedTagArgs struct {
	// (required) Object containing details of tag to be created.
	TagObject *GitAnnotatedTag
	// (required) Project ID or project name
	Project *string
	// (required) ID or name of the repository.
	RepositoryId *string
}

// [Preview API] Attach a new file to a pull request.
func (client *ClientImpl) CreateAttachment(ctx context.Context, args CreateAttachmentArgs) (*Attachment, error) {
	if args.UploadStream == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.UploadStream"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FileName == nil || *args.FileName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, args.UploadStream, "application/octet-stream", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Attachment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateAttachment function
type CreateAttachmentArgs struct {
	// (required) Stream to upload
	UploadStream io.Reader
	// (required) The name of the file.
	FileName *string
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Cherry pick a specific commit or commits that are associated to a pull request into a new branch.
func (client *ClientImpl) CreateCherryPick(ctx context.Context, args CreateCherryPickArgs) (*GitCherryPick, error) {
	if args.CherryPickToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CherryPickToCreate"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.CherryPickToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCherryPick
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateCherryPick function
type CreateCherryPickArgs struct {
	// (required)
	CherryPickToCreate *GitAsyncRefOperationParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the repository.
	RepositoryId *string
}

// Create a comment on a specific thread in a pull request (up to 500 comments can be created per thread).
func (client *ClientImpl) CreateComment(ctx context.Context, args CreateCommentArgs) (*Comment, error) {
	if args.Comment == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Comment"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)

	body, marshalErr := json.Marshal(*args.Comment)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Comment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateComment function
type CreateCommentArgs struct {
	// (required) The comment to create. Comments can be up to 150,000 characters.
	Comment *Comment
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread that the desired comment is in.
	ThreadId *int
	// (optional) Project ID or project name
	Project *string
}

// Create Git commit status.
func (client *ClientImpl) CreateCommitStatus(ctx context.Context, args CreateCommitStatusArgs) (*GitStatus, error) {
	if args.GitCommitStatusToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GitCommitStatusToCreate"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.CommitId == nil || *args.CommitId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CommitId"}
	}
	routeValues["commitId"] = *args.CommitId
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.GitCommitStatusToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("428dd4fb-fda5-4722-af02-9313b80305da")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateCommitStatus function
type CreateCommitStatusArgs struct {
	// (required) Git commit status object to create.
	GitCommitStatusToCreate *GitStatus
	// (required) ID of the Git commit.
	CommitId *string
	// (required) ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Creates a ref favorite
func (client *ClientImpl) CreateFavorite(ctx context.Context, args CreateFavoriteArgs) (*GitRefFavorite, error) {
	if args.Favorite == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Favorite"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	body, marshalErr := json.Marshal(*args.Favorite)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRefFavorite
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateFavorite function
type CreateFavoriteArgs struct {
	// (required) The ref favorite to create.
	Favorite *GitRefFavorite
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Request that another repository's refs be fetched into this one. It syncs two existing forks. To create a fork, please see the <a href="https://docs.microsoft.com/en-us/rest/api/vsts/git/repositories/create?view=azure-devops-rest-5.1"> repositories endpoint</a>
func (client *ClientImpl) CreateForkSyncRequest(ctx context.Context, args CreateForkSyncRequestArgs) (*GitForkSyncRequest, error) {
	if args.SyncParams == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SyncParams"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	body, marshalErr := json.Marshal(*args.SyncParams)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitForkSyncRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateForkSyncRequest function
type CreateForkSyncRequestArgs struct {
	// (required) Source repository and ref mapping.
	SyncParams *GitForkSyncRequestParameters
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) True to include links
	IncludeLinks *bool
}

// [Preview API] Create an import request.
func (client *ClientImpl) CreateImportRequest(ctx context.Context, args CreateImportRequestArgs) (*GitImportRequest, error) {
	if args.ImportRequest == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ImportRequest"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.ImportRequest)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitImportRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateImportRequest function
type CreateImportRequestArgs struct {
	// (required) The import request to create.
	ImportRequest *GitImportRequest
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryId *string
}

// [Preview API] Add a like on a comment.
func (client *ClientImpl) CreateLike(ctx context.Context, args CreateLikeArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the CreateLike function
type CreateLikeArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) The ID of the thread that contains the comment.
	ThreadId *int
	// (required) The ID of the comment.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Request a git merge operation. Currently we support merging only 2 commits.
func (client *ClientImpl) CreateMergeRequest(ctx context.Context, args CreateMergeRequestArgs) (*GitMerge, error) {
	if args.MergeParameters == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.MergeParameters"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	body, marshalErr := json.Marshal(*args.MergeParameters)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("985f7ae9-844f-4906-9897-7ef41516c0e2")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitMerge
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateMergeRequest function
type CreateMergeRequestArgs struct {
	// (required) Parents commitIds and merge commit messsage.
	MergeParameters *GitMergeParameters
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (optional) True to include links
	IncludeLinks *bool
}

// Create a pull request.
func (client *ClientImpl) CreatePullRequest(ctx context.Context, args CreatePullRequestArgs) (*GitPullRequest, error) {
	if args.GitPullRequestToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GitPullRequestToCreate"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.SupportsIterations != nil {
		queryParams.Add("supportsIterations", strconv.FormatBool(*args.SupportsIterations))
	}
	body, marshalErr := json.Marshal(*args.GitPullRequestToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequest function
type CreatePullRequestArgs struct {
	// (required) The pull request to create.
	GitPullRequestToCreate *GitPullRequest
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, subsequent pushes to the pull request will be individually reviewable. Set this to false for large pull requests for performance reasons if this functionality is not needed.
	SupportsIterations *bool
}

// [Preview API] Create a pull request status on the iteration. This operation will have the same result as Create status on pull request with specified iteration ID in the request body.
func (client *ClientImpl) CreatePullRequestIterationStatus(ctx context.Context, args CreatePullRequestIterationStatusArgs) (*GitPullRequestStatus, error) {
	if args.Status == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Status"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	body, marshalErr := json.Marshal(*args.Status)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequestIterationStatus function
type CreatePullRequestIterationStatusArgs struct {
	// (required) Pull request status to create.
	Status *GitPullRequestStatus
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Create a label for a specified pull request. The only required field is the name of the new label.
func (client *ClientImpl) CreatePullRequestLabel(ctx context.Context, args CreatePullRequestLabelArgs) (*core.WebApiTagDefinition, error) {
	if args.Label == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Label"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	body, marshalErr := json.Marshal(*args.Label)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue core.WebApiTagDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequestLabel function
type CreatePullRequestLabelArgs struct {
	// (required) Label to assign to the pull request.
	Label *core.WebApiCreateTagRequestData
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Project ID or project name.
	ProjectId *string
}

// Add a reviewer to a pull request or cast a vote.
func (client *ClientImpl) CreatePullRequestReviewer(ctx context.Context, args CreatePullRequestReviewerArgs) (*IdentityRefWithVote, error) {
	if args.Reviewer == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Reviewer"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ReviewerId == nil || *args.ReviewerId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ReviewerId"}
	}
	routeValues["reviewerId"] = *args.ReviewerId

	body, marshalErr := json.Marshal(*args.Reviewer)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	resp, err := client.Client.Send(ctx, http.MethodPut, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue IdentityRefWithVote
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequestReviewer function
type CreatePullRequestReviewerArgs struct {
	// (required) Reviewer's vote.<br />If the reviewer's ID is included here, it must match the reviewerID parameter.<br />Reviewers can set their own vote with this method.  When adding other reviewers, vote must be set to zero.
	Reviewer *IdentityRefWithVote
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the reviewer.
	ReviewerId *string
	// (optional) Project ID or project name
	Project *string
}

// Add reviewers to a pull request.
func (client *ClientImpl) CreatePullRequestReviewers(ctx context.Context, args CreatePullRequestReviewersArgs) (*[]IdentityRefWithVote, error) {
	if args.Reviewers == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Reviewers"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.Reviewers)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []IdentityRefWithVote
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequestReviewers function
type CreatePullRequestReviewersArgs struct {
	// (required) Reviewers to add to the pull request.
	Reviewers *[]webapi.IdentityRef
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Create a pull request status.
func (client *ClientImpl) CreatePullRequestStatus(ctx context.Context, args CreatePullRequestStatusArgs) (*GitPullRequestStatus, error) {
	if args.Status == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Status"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.Status)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePullRequestStatus function
type CreatePullRequestStatusArgs struct {
	// (required) Pull request status to create.
	Status *GitPullRequestStatus
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Push changes to the repository.
func (client *ClientImpl) CreatePush(ctx context.Context, args CreatePushArgs) (*GitPush, error) {
	if args.Push == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Push"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.Push)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPush
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreatePush function
type CreatePushArgs struct {
	// (required)
	Push *GitPush
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// Create a git repository in a team project.
func (client *ClientImpl) CreateRepository(ctx context.Context, args CreateRepositoryArgs) (*GitRepository, error) {
	if args.GitRepositoryToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GitRepositoryToCreate"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.SourceRef != nil {
		queryParams.Add("sourceRef", *args.SourceRef)
	}
	body, marshalErr := json.Marshal(*args.GitRepositoryToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRepository
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateRepository function
type CreateRepositoryArgs struct {
	// (required) Specify the repo name, team project and/or parent repository. Team project information can be omitted from gitRepositoryToCreate if the request is project-scoped (i.e., includes project Id).
	GitRepositoryToCreate *GitRepositoryCreateOptions
	// (optional) Project ID or project name
	Project *string
	// (optional) [optional] Specify the source refs to use while creating a fork repo
	SourceRef *string
}

// [Preview API] Starts the operation to create a new branch which reverts changes introduced by either a specific commit or commits that are associated to a pull request.
func (client *ClientImpl) CreateRevert(ctx context.Context, args CreateRevertArgs) (*GitRevert, error) {
	if args.RevertToCreate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RevertToCreate"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.RevertToCreate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRevert
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateRevert function
type CreateRevertArgs struct {
	// (required)
	RevertToCreate *GitAsyncRefOperationParameters
	// (required) Project ID or project name
	Project *string
	// (required) ID of the repository.
	RepositoryId *string
}

// Create a thread in a pull request.
func (client *ClientImpl) CreateThread(ctx context.Context, args CreateThreadArgs) (*GitPullRequestCommentThread, error) {
	if args.CommentThread == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommentThread"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.CommentThread)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestCommentThread
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the CreateThread function
type CreateThreadArgs struct {
	// (required) The thread to create. Thread must contain at least one comment.
	CommentThread *GitPullRequestCommentThread
	// (required) Repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete a pull request attachment.
func (client *ClientImpl) DeleteAttachment(ctx context.Context, args DeleteAttachmentArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FileName == nil || *args.FileName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteAttachment function
type DeleteAttachmentArgs struct {
	// (required) The name of the attachment to delete.
	FileName *string
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Delete a comment associated with a specific thread in a pull request.
func (client *ClientImpl) DeleteComment(ctx context.Context, args DeleteCommentArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteComment function
type DeleteCommentArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread that the desired comment is in.
	ThreadId *int
	// (required) ID of the comment.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete a like on a comment.
func (client *ClientImpl) DeleteLike(ctx context.Context, args DeleteLikeArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteLike function
type DeleteLikeArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) The ID of the thread that contains the comment.
	ThreadId *int
	// (required) The ID of the comment.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete pull request iteration status.
func (client *ClientImpl) DeletePullRequestIterationStatus(ctx context.Context, args DeletePullRequestIterationStatusArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)
	if args.StatusId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.StatusId"}
	}
	routeValues["statusId"] = strconv.Itoa(*args.StatusId)

	locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePullRequestIterationStatus function
type DeletePullRequestIterationStatusArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration.
	IterationId *int
	// (required) ID of the pull request status.
	StatusId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Removes a label from the set of those assigned to the pull request.
func (client *ClientImpl) DeletePullRequestLabels(ctx context.Context, args DeletePullRequestLabelsArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.LabelIdOrName == nil || *args.LabelIdOrName == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.LabelIdOrName"}
	}
	routeValues["labelIdOrName"] = *args.LabelIdOrName

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePullRequestLabels function
type DeletePullRequestLabelsArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) The name or ID of the label requested.
	LabelIdOrName *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Project ID or project name.
	ProjectId *string
}

// Remove a reviewer from a pull request.
func (client *ClientImpl) DeletePullRequestReviewer(ctx context.Context, args DeletePullRequestReviewerArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ReviewerId == nil || *args.ReviewerId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ReviewerId"}
	}
	routeValues["reviewerId"] = *args.ReviewerId

	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePullRequestReviewer function
type DeletePullRequestReviewerArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the reviewer to remove.
	ReviewerId *string
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Delete pull request status.
func (client *ClientImpl) DeletePullRequestStatus(ctx context.Context, args DeletePullRequestStatusArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.StatusId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.StatusId"}
	}
	routeValues["statusId"] = strconv.Itoa(*args.StatusId)

	locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeletePullRequestStatus function
type DeletePullRequestStatusArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request status.
	StatusId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Deletes the refs favorite specified
func (client *ClientImpl) DeleteRefFavorite(ctx context.Context, args DeleteRefFavoriteArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.FavoriteId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.FavoriteId"}
	}
	routeValues["favoriteId"] = strconv.Itoa(*args.FavoriteId)

	locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteRefFavorite function
type DeleteRefFavoriteArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The Id of the ref favorite to delete.
	FavoriteId *int
}

// Delete a git repository
func (client *ClientImpl) DeleteRepository(ctx context.Context, args DeleteRepositoryArgs) error {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = (*args.RepositoryId).String()

	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteRepository function
type DeleteRepositoryArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *uuid.UUID
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Destroy (hard delete) a soft-deleted Git repository.
func (client *ClientImpl) DeleteRepositoryFromRecycleBin(ctx context.Context, args DeleteRepositoryFromRecycleBinArgs) error {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = (*args.RepositoryId).String()

	locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
	_, err := client.Client.Send(ctx, http.MethodDelete, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the DeleteRepositoryFromRecycleBin function
type DeleteRepositoryFromRecycleBinArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the repository.
	RepositoryId *uuid.UUID
}

// [Preview API] Get an annotated tag.
func (client *ClientImpl) GetAnnotatedTag(ctx context.Context, args GetAnnotatedTagArgs) (*GitAnnotatedTag, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.ObjectId == nil || *args.ObjectId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ObjectId"}
	}
	routeValues["objectId"] = *args.ObjectId

	locationId, _ := uuid.Parse("5e8a8081-3851-4626-b677-9891cc04102e")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitAnnotatedTag
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAnnotatedTag function
type GetAnnotatedTagArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ObjectId (Sha1Id) of tag to get.
	ObjectId *string
}

// [Preview API] Get the file content of a pull request attachment.
func (client *ClientImpl) GetAttachmentContent(ctx context.Context, args GetAttachmentContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FileName == nil || *args.FileName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetAttachmentContent function
type GetAttachmentContentArgs struct {
	// (required) The name of the attachment.
	FileName *string
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get a list of files attached to a given pull request.
func (client *ClientImpl) GetAttachments(ctx context.Context, args GetAttachmentsArgs) (*[]Attachment, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Attachment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetAttachments function
type GetAttachmentsArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get the file content of a pull request attachment.
func (client *ClientImpl) GetAttachmentZip(ctx context.Context, args GetAttachmentZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.FileName == nil || *args.FileName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.FileName"}
	}
	routeValues["fileName"] = *args.FileName
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("965d9361-878b-413b-a494-45d5b5fd8ab7")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetAttachmentZip function
type GetAttachmentZipArgs struct {
	// (required) The name of the attachment.
	FileName *string
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Get a single blob.
func (client *ClientImpl) GetBlob(ctx context.Context, args GetBlobArgs) (*GitBlobRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.Sha1 == nil || *args.Sha1 == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Sha1"}
	}
	routeValues["sha1"] = *args.Sha1

	queryParams := url.Values{}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitBlobRef
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBlob function
type GetBlobArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
	Sha1 *string
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
	Download *bool
	// (optional) Provide a fileName to use for a download.
	FileName *string
	// (optional) If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
	ResolveLfs *bool
}

// Get a single blob.
func (client *ClientImpl) GetBlobContent(ctx context.Context, args GetBlobContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.Sha1 == nil || *args.Sha1 == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Sha1"}
	}
	routeValues["sha1"] = *args.Sha1

	queryParams := url.Values{}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBlobContent function
type GetBlobContentArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
	Sha1 *string
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
	Download *bool
	// (optional) Provide a fileName to use for a download.
	FileName *string
	// (optional) If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
	ResolveLfs *bool
}

// Gets one or more blobs in a zip file download.
func (client *ClientImpl) GetBlobsZip(ctx context.Context, args GetBlobsZipArgs) (io.ReadCloser, error) {
	if args.BlobIds == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.BlobIds"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Filename != nil {
		queryParams.Add("filename", *args.Filename)
	}
	body, marshalErr := json.Marshal(*args.BlobIds)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBlobsZip function
type GetBlobsZipArgs struct {
	// (required) Blob IDs (SHA1 hashes) to be returned in the zip file.
	BlobIds *[]string
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional)
	Filename *string
}

// Get a single blob.
func (client *ClientImpl) GetBlobZip(ctx context.Context, args GetBlobZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.Sha1 == nil || *args.Sha1 == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Sha1"}
	}
	routeValues["sha1"] = *args.Sha1

	queryParams := url.Values{}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("7b28e929-2c99-405d-9c5c-6167a06e6816")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetBlobZip function
type GetBlobZipArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) SHA1 hash of the file. You can get the SHA1 of a file using the "Git/Items/Get Item" endpoint.
	Sha1 *string
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, prompt for a download rather than rendering in a browser. Note: this value defaults to true if $format is zip
	Download *bool
	// (optional) Provide a fileName to use for a download.
	FileName *string
	// (optional) If true, try to resolve a blob to its LFS contents, if it's an LFS pointer file. Only compatible with octet-stream Accept headers or $format types
	ResolveLfs *bool
}

// Retrieve statistics about a single branch.
func (client *ClientImpl) GetBranch(ctx context.Context, args GetBranchArgs) (*GitBranchStats, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Name == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "name"}
	}
	queryParams.Add("name", *args.Name)
	if args.BaseVersionDescriptor != nil {
		if args.BaseVersionDescriptor.VersionType != nil {
			queryParams.Add("baseVersionDescriptor.versionType", string(*args.BaseVersionDescriptor.VersionType))
		}
		if args.BaseVersionDescriptor.Version != nil {
			queryParams.Add("baseVersionDescriptor.version", *args.BaseVersionDescriptor.Version)
		}
		if args.BaseVersionDescriptor.VersionOptions != nil {
			queryParams.Add("baseVersionDescriptor.versionOptions", string(*args.BaseVersionDescriptor.VersionOptions))
		}
	}
	locationId, _ := uuid.Parse("d5b216de-d8d5-4d32-ae76-51df755b16d3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitBranchStats
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBranch function
type GetBranchArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) Name of the branch.
	Name *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Identifies the commit or branch to use as the base.
	BaseVersionDescriptor *GitVersionDescriptor
}

// Retrieve statistics about all branches within a repository.
func (client *ClientImpl) GetBranches(ctx context.Context, args GetBranchesArgs) (*[]GitBranchStats, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.BaseVersionDescriptor != nil {
		if args.BaseVersionDescriptor.VersionType != nil {
			queryParams.Add("baseVersionDescriptor.versionType", string(*args.BaseVersionDescriptor.VersionType))
		}
		if args.BaseVersionDescriptor.Version != nil {
			queryParams.Add("baseVersionDescriptor.version", *args.BaseVersionDescriptor.Version)
		}
		if args.BaseVersionDescriptor.VersionOptions != nil {
			queryParams.Add("baseVersionDescriptor.versionOptions", string(*args.BaseVersionDescriptor.VersionOptions))
		}
	}
	locationId, _ := uuid.Parse("d5b216de-d8d5-4d32-ae76-51df755b16d3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitBranchStats
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetBranches function
type GetBranchesArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Identifies the commit or branch to use as the base.
	BaseVersionDescriptor *GitVersionDescriptor
}

// Retrieve changes for a particular commit.
func (client *ClientImpl) GetChanges(ctx context.Context, args GetChangesArgs) (*GitCommitChanges, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.CommitId == nil || *args.CommitId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CommitId"}
	}
	routeValues["commitId"] = *args.CommitId
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("5bf884f5-3e07-42e9-afb8-1b872267bf16")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCommitChanges
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetChanges function
type GetChangesArgs struct {
	// (required) The id of the commit.
	CommitId *string
	// (required) The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The maximum number of changes to return.
	Top *int
	// (optional) The number of changes to skip.
	Skip *int
}

// [Preview API] Retrieve information about a cherry pick by cherry pick Id.
func (client *ClientImpl) GetCherryPick(ctx context.Context, args GetCherryPickArgs) (*GitCherryPick, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.CherryPickId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CherryPickId"}
	}
	routeValues["cherryPickId"] = strconv.Itoa(*args.CherryPickId)
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCherryPick
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCherryPick function
type GetCherryPickArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the cherry pick.
	CherryPickId *int
	// (required) ID of the repository.
	RepositoryId *string
}

// [Preview API] Retrieve information about a cherry pick for a specific branch.
func (client *ClientImpl) GetCherryPickForRefName(ctx context.Context, args GetCherryPickForRefNameArgs) (*GitCherryPick, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.RefName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "refName"}
	}
	queryParams.Add("refName", *args.RefName)
	locationId, _ := uuid.Parse("033bad68-9a14-43d1-90e0-59cb8856fef6")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCherryPick
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCherryPickForRefName function
type GetCherryPickForRefNameArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the repository.
	RepositoryId *string
	// (required) The GitAsyncRefOperationParameters generatedRefName used for the cherry pick operation.
	RefName *string
}

// Retrieve a comment associated with a specific thread in a pull request.
func (client *ClientImpl) GetComment(ctx context.Context, args GetCommentArgs) (*Comment, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Comment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetComment function
type GetCommentArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread that the desired comment is in.
	ThreadId *int
	// (required) ID of the comment.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieve all comments associated with a specific thread in a pull request.
func (client *ClientImpl) GetComments(ctx context.Context, args GetCommentsArgs) (*[]Comment, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)

	locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []Comment
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetComments function
type GetCommentsArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread.
	ThreadId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieve a particular commit.
func (client *ClientImpl) GetCommit(ctx context.Context, args GetCommitArgs) (*GitCommit, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.CommitId == nil || *args.CommitId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CommitId"}
	}
	routeValues["commitId"] = *args.CommitId
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.ChangeCount != nil {
		queryParams.Add("changeCount", strconv.Itoa(*args.ChangeCount))
	}
	locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCommit
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCommit function
type GetCommitArgs struct {
	// (required) The id of the commit.
	CommitId *string
	// (required) The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The number of changes to include in the result.
	ChangeCount *int
}

// Find the closest common commit (the merge base) between base and target commits, and get the diff between either the base and target commits or common and target commits.
func (client *ClientImpl) GetCommitDiffs(ctx context.Context, args GetCommitDiffsArgs) (*GitCommitDiffs, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.DiffCommonCommit != nil {
		queryParams.Add("diffCommonCommit", strconv.FormatBool(*args.DiffCommonCommit))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.BaseVersionDescriptor != nil {
		if args.BaseVersionDescriptor.BaseVersionType != nil {
			queryParams.Add("baseVersionType", string(*args.BaseVersionDescriptor.BaseVersionType))
		}
		if args.BaseVersionDescriptor.BaseVersion != nil {
			queryParams.Add("baseVersion", *args.BaseVersionDescriptor.BaseVersion)
		}
		if args.BaseVersionDescriptor.BaseVersionOptions != nil {
			queryParams.Add("baseVersionOptions", string(*args.BaseVersionDescriptor.BaseVersionOptions))
		}
	}
	if args.TargetVersionDescriptor != nil {
		if args.TargetVersionDescriptor.TargetVersionType != nil {
			queryParams.Add("targetVersionType", string(*args.TargetVersionDescriptor.TargetVersionType))
		}
		if args.TargetVersionDescriptor.TargetVersion != nil {
			queryParams.Add("targetVersion", *args.TargetVersionDescriptor.TargetVersion)
		}
		if args.TargetVersionDescriptor.TargetVersionOptions != nil {
			queryParams.Add("targetVersionOptions", string(*args.TargetVersionDescriptor.TargetVersionOptions))
		}
	}
	locationId, _ := uuid.Parse("615588d5-c0c7-4b88-88f8-e625306446e8")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitCommitDiffs
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCommitDiffs function
type GetCommitDiffsArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, diff between common and target commits. If false, diff between base and target commits.
	DiffCommonCommit *bool
	// (optional) Maximum number of changes to return. Defaults to 100.
	Top *int
	// (optional) Number of changes to skip
	Skip *int
	// (optional) Descriptor for base commit.
	BaseVersionDescriptor *GitBaseVersionDescriptor
	// (optional) Descriptor for target commit.
	TargetVersionDescriptor *GitTargetVersionDescriptor
}

// Retrieve git commits for a project
func (client *ClientImpl) GetCommits(ctx context.Context, args GetCommitsArgs) (*[]GitCommitRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.SearchCriteria == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "searchCriteria"}
	}
	if args.SearchCriteria.Ids != nil {
		for index, item := range *args.SearchCriteria.Ids {
			queryParams.Add("searchCriteria.ids["+strconv.Itoa(index)+"]", item)
		}
	}
	if args.SearchCriteria.FromDate != nil {
		queryParams.Add("searchCriteria.fromDate", *args.SearchCriteria.FromDate)
	}
	if args.SearchCriteria.ToDate != nil {
		queryParams.Add("searchCriteria.toDate", *args.SearchCriteria.ToDate)
	}
	if args.SearchCriteria.ItemVersion != nil {
		if args.SearchCriteria.ItemVersion.VersionType != nil {
			queryParams.Add("searchCriteria.itemVersion.versionType", string(*args.SearchCriteria.ItemVersion.VersionType))
		}
		if args.SearchCriteria.ItemVersion.Version != nil {
			queryParams.Add("searchCriteria.itemVersion.version", *args.SearchCriteria.ItemVersion.Version)
		}
		if args.SearchCriteria.ItemVersion.VersionOptions != nil {
			queryParams.Add("searchCriteria.itemVersion.versionOptions", string(*args.SearchCriteria.ItemVersion.VersionOptions))
		}
	}
	if args.SearchCriteria.CompareVersion != nil {
		if args.SearchCriteria.CompareVersion.VersionType != nil {
			queryParams.Add("searchCriteria.compareVersion.versionType", string(*args.SearchCriteria.CompareVersion.VersionType))
		}
		if args.SearchCriteria.CompareVersion.Version != nil {
			queryParams.Add("searchCriteria.compareVersion.version", *args.SearchCriteria.CompareVersion.Version)
		}
		if args.SearchCriteria.CompareVersion.VersionOptions != nil {
			queryParams.Add("searchCriteria.compareVersion.versionOptions", string(*args.SearchCriteria.CompareVersion.VersionOptions))
		}
	}
	if args.SearchCriteria.FromCommitId != nil {
		queryParams.Add("searchCriteria.fromCommitId", *args.SearchCriteria.FromCommitId)
	}
	if args.SearchCriteria.ToCommitId != nil {
		queryParams.Add("searchCriteria.toCommitId", *args.SearchCriteria.ToCommitId)
	}
	if args.SearchCriteria.User != nil {
		queryParams.Add("searchCriteria.user", *args.SearchCriteria.User)
	}
	if args.SearchCriteria.Author != nil {
		queryParams.Add("searchCriteria.author", *args.SearchCriteria.Author)
	}
	if args.SearchCriteria.ItemPath != nil {
		queryParams.Add("searchCriteria.itemPath", *args.SearchCriteria.ItemPath)
	}
	if args.SearchCriteria.ExcludeDeletes != nil {
		queryParams.Add("searchCriteria.excludeDeletes", strconv.FormatBool(*args.SearchCriteria.ExcludeDeletes))
	}
	if args.SearchCriteria.Skip != nil {
		queryParams.Add("searchCriteria.$skip", strconv.Itoa(*args.SearchCriteria.Skip))
	}
	if args.SearchCriteria.Top != nil {
		queryParams.Add("searchCriteria.$top", strconv.Itoa(*args.SearchCriteria.Top))
	}
	if args.SearchCriteria.IncludeLinks != nil {
		queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
	}
	if args.SearchCriteria.IncludeWorkItems != nil {
		queryParams.Add("searchCriteria.includeWorkItems", strconv.FormatBool(*args.SearchCriteria.IncludeWorkItems))
	}
	if args.SearchCriteria.IncludeUserImageUrl != nil {
		queryParams.Add("searchCriteria.includeUserImageUrl", strconv.FormatBool(*args.SearchCriteria.IncludeUserImageUrl))
	}
	if args.SearchCriteria.IncludePushData != nil {
		queryParams.Add("searchCriteria.includePushData", strconv.FormatBool(*args.SearchCriteria.IncludePushData))
	}
	if args.SearchCriteria.HistoryMode != nil {
		queryParams.Add("searchCriteria.historyMode", string(*args.SearchCriteria.HistoryMode))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitCommitRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCommits function
type GetCommitsArgs struct {
	// (required) The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
	RepositoryId *string
	// (required)
	SearchCriteria *GitQueryCommitsCriteria
	// (optional) Project ID or project name
	Project *string
	// (optional)
	Skip *int
	// (optional)
	Top *int
}

// Retrieve git commits for a project matching the search criteria
func (client *ClientImpl) GetCommitsBatch(ctx context.Context, args GetCommitsBatchArgs) (*[]GitCommitRef, error) {
	if args.SearchCriteria == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.SearchCriteria"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.IncludeStatuses != nil {
		queryParams.Add("includeStatuses", strconv.FormatBool(*args.IncludeStatuses))
	}
	body, marshalErr := json.Marshal(*args.SearchCriteria)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("6400dfb2-0bcb-462b-b992-5a57f8f1416c")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitCommitRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetCommitsBatch function
type GetCommitsBatchArgs struct {
	// (required) Search options
	SearchCriteria *GitQueryCommitsCriteria
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Number of commits to skip.
	Skip *int
	// (optional) Maximum number of commits to return.
	Top *int
	// (optional) True to include additional commit status information.
	IncludeStatuses *bool
}

// [Preview API] Retrieve deleted git repositories.
func (client *ClientImpl) GetDeletedRepositories(ctx context.Context, args GetDeletedRepositoriesArgs) (*[]GitDeletedRepository, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("2b6869c4-cb25-42b5-b7a3-0d3e6be0a11a")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitDeletedRepository
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetDeletedRepositories function
type GetDeletedRepositoriesArgs struct {
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Retrieve all forks of a repository in the collection.
func (client *ClientImpl) GetForks(ctx context.Context, args GetForksArgs) (*[]GitRepositoryRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId
	if args.CollectionId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CollectionId"}
	}
	routeValues["collectionId"] = (*args.CollectionId).String()

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("158c0340-bf6f-489c-9625-d572a1480d57")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitRepositoryRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetForks function
type GetForksArgs struct {
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (required) Team project collection ID.
	CollectionId *uuid.UUID
	// (optional) Project ID or project name
	Project *string
	// (optional) True to include links.
	IncludeLinks *bool
}

// [Preview API] Get a specific fork sync operation's details.
func (client *ClientImpl) GetForkSyncRequest(ctx context.Context, args GetForkSyncRequestArgs) (*GitForkSyncRequest, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId
	if args.ForkSyncOperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ForkSyncOperationId"}
	}
	routeValues["forkSyncOperationId"] = strconv.Itoa(*args.ForkSyncOperationId)

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitForkSyncRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetForkSyncRequest function
type GetForkSyncRequestArgs struct {
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (required) OperationId of the sync request.
	ForkSyncOperationId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) True to include links.
	IncludeLinks *bool
}

// [Preview API] Retrieve all requested fork sync operations on this repository.
func (client *ClientImpl) GetForkSyncRequests(ctx context.Context, args GetForkSyncRequestsArgs) (*[]GitForkSyncRequest, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId

	queryParams := url.Values{}
	if args.IncludeAbandoned != nil {
		queryParams.Add("includeAbandoned", strconv.FormatBool(*args.IncludeAbandoned))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("1703f858-b9d1-46af-ab62-483e9e1055b5")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitForkSyncRequest
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetForkSyncRequests function
type GetForkSyncRequestsArgs struct {
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) True to include abandoned requests.
	IncludeAbandoned *bool
	// (optional) True to include links.
	IncludeLinks *bool
}

// [Preview API] Retrieve a particular import request.
func (client *ClientImpl) GetImportRequest(ctx context.Context, args GetImportRequestArgs) (*GitImportRequest, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.ImportRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ImportRequestId"}
	}
	routeValues["importRequestId"] = strconv.Itoa(*args.ImportRequestId)

	locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitImportRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetImportRequest function
type GetImportRequestArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The unique identifier for the import request.
	ImportRequestId *int
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
func (client *ClientImpl) GetItem(ctx context.Context, args GetItemArgs) (*GitItem, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeContentMetadata != nil {
		queryParams.Add("includeContentMetadata", strconv.FormatBool(*args.IncludeContentMetadata))
	}
	if args.LatestProcessedChange != nil {
		queryParams.Add("latestProcessedChange", strconv.FormatBool(*args.LatestProcessedChange))
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
		if args.VersionDescriptor.VersionOptions != nil {
			queryParams.Add("versionDescriptor.versionOptions", string(*args.VersionDescriptor.VersionOptions))
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitItem
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItem function
type GetItemArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The item path.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The path scope.  The default is null.
	ScopePath *string
	// (optional) The recursion level of this request. The default is 'none', no recursion.
	RecursionLevel *VersionControlRecursionType
	// (optional) Set to true to include content metadata.  Default is false.
	IncludeContentMetadata *bool
	// (optional) Set to true to include the latest changes.  Default is false.
	LatestProcessedChange *bool
	// (optional) Set to true to download the response as a file.  Default is false.
	Download *bool
	// (optional) Version descriptor.  Default is the default branch for the repository.
	VersionDescriptor *GitVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
	// (optional) Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
	ResolveLfs *bool
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
func (client *ClientImpl) GetItemContent(ctx context.Context, args GetItemContentArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeContentMetadata != nil {
		queryParams.Add("includeContentMetadata", strconv.FormatBool(*args.IncludeContentMetadata))
	}
	if args.LatestProcessedChange != nil {
		queryParams.Add("latestProcessedChange", strconv.FormatBool(*args.LatestProcessedChange))
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
		if args.VersionDescriptor.VersionOptions != nil {
			queryParams.Add("versionDescriptor.versionOptions", string(*args.VersionDescriptor.VersionOptions))
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/octet-stream", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemContent function
type GetItemContentArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The item path.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The path scope.  The default is null.
	ScopePath *string
	// (optional) The recursion level of this request. The default is 'none', no recursion.
	RecursionLevel *VersionControlRecursionType
	// (optional) Set to true to include content metadata.  Default is false.
	IncludeContentMetadata *bool
	// (optional) Set to true to include the latest changes.  Default is false.
	LatestProcessedChange *bool
	// (optional) Set to true to download the response as a file.  Default is false.
	Download *bool
	// (optional) Version descriptor.  Default is the default branch for the repository.
	VersionDescriptor *GitVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
	// (optional) Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
	ResolveLfs *bool
}

// Get Item Metadata and/or Content for a collection of items. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content which is always returned as a download.
func (client *ClientImpl) GetItems(ctx context.Context, args GetItemsArgs) (*[]GitItem, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeContentMetadata != nil {
		queryParams.Add("includeContentMetadata", strconv.FormatBool(*args.IncludeContentMetadata))
	}
	if args.LatestProcessedChange != nil {
		queryParams.Add("latestProcessedChange", strconv.FormatBool(*args.LatestProcessedChange))
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
		if args.VersionDescriptor.VersionOptions != nil {
			queryParams.Add("versionDescriptor.versionOptions", string(*args.VersionDescriptor.VersionOptions))
		}
	}
	locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItems function
type GetItemsArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The path scope.  The default is null.
	ScopePath *string
	// (optional) The recursion level of this request. The default is 'none', no recursion.
	RecursionLevel *VersionControlRecursionType
	// (optional) Set to true to include content metadata.  Default is false.
	IncludeContentMetadata *bool
	// (optional) Set to true to include the latest changes.  Default is false.
	LatestProcessedChange *bool
	// (optional) Set to true to download the response as a file.  Default is false.
	Download *bool
	// (optional) Set to true to include links to items.  Default is false.
	IncludeLinks *bool
	// (optional) Version descriptor.  Default is the default branch for the repository.
	VersionDescriptor *GitVersionDescriptor
}

// Post for retrieving a creating a batch out of a set of items in a repo / project given a list of paths or a long path
func (client *ClientImpl) GetItemsBatch(ctx context.Context, args GetItemsBatchArgs) (*[][]GitItem, error) {
	if args.RequestData == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RequestData"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.RequestData)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("630fd2e4-fb88-4f85-ad21-13f3fd1fbca9")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue [][]GitItem
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetItemsBatch function
type GetItemsBatchArgs struct {
	// (required) Request data attributes: ItemDescriptors, IncludeContentMetadata, LatestProcessedChange, IncludeLinks. ItemDescriptors: Collection of items to fetch, including path, version, and recursion level. IncludeContentMetadata: Whether to include metadata for all items LatestProcessedChange: Whether to include shallow ref to commit that last changed each item. IncludeLinks: Whether to include the _links field on the shallow references.
	RequestData *GitItemRequestData
	// (required) The name or ID of the repository
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
func (client *ClientImpl) GetItemText(ctx context.Context, args GetItemTextArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeContentMetadata != nil {
		queryParams.Add("includeContentMetadata", strconv.FormatBool(*args.IncludeContentMetadata))
	}
	if args.LatestProcessedChange != nil {
		queryParams.Add("latestProcessedChange", strconv.FormatBool(*args.LatestProcessedChange))
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
		if args.VersionDescriptor.VersionOptions != nil {
			queryParams.Add("versionDescriptor.versionOptions", string(*args.VersionDescriptor.VersionOptions))
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "text/plain", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemText function
type GetItemTextArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The item path.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The path scope.  The default is null.
	ScopePath *string
	// (optional) The recursion level of this request. The default is 'none', no recursion.
	RecursionLevel *VersionControlRecursionType
	// (optional) Set to true to include content metadata.  Default is false.
	IncludeContentMetadata *bool
	// (optional) Set to true to include the latest changes.  Default is false.
	LatestProcessedChange *bool
	// (optional) Set to true to download the response as a file.  Default is false.
	Download *bool
	// (optional) Version descriptor.  Default is the default branch for the repository.
	VersionDescriptor *GitVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
	// (optional) Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
	ResolveLfs *bool
}

// Get Item Metadata and/or Content for a single item. The download parameter is to indicate whether the content should be available as a download or just sent as a stream in the response. Doesn't apply to zipped content, which is always returned as a download.
func (client *ClientImpl) GetItemZip(ctx context.Context, args GetItemZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Path == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "path"}
	}
	queryParams.Add("path", *args.Path)
	if args.ScopePath != nil {
		queryParams.Add("scopePath", *args.ScopePath)
	}
	if args.RecursionLevel != nil {
		queryParams.Add("recursionLevel", string(*args.RecursionLevel))
	}
	if args.IncludeContentMetadata != nil {
		queryParams.Add("includeContentMetadata", strconv.FormatBool(*args.IncludeContentMetadata))
	}
	if args.LatestProcessedChange != nil {
		queryParams.Add("latestProcessedChange", strconv.FormatBool(*args.LatestProcessedChange))
	}
	if args.Download != nil {
		queryParams.Add("download", strconv.FormatBool(*args.Download))
	}
	if args.VersionDescriptor != nil {
		if args.VersionDescriptor.VersionType != nil {
			queryParams.Add("versionDescriptor.versionType", string(*args.VersionDescriptor.VersionType))
		}
		if args.VersionDescriptor.Version != nil {
			queryParams.Add("versionDescriptor.version", *args.VersionDescriptor.Version)
		}
		if args.VersionDescriptor.VersionOptions != nil {
			queryParams.Add("versionDescriptor.versionOptions", string(*args.VersionDescriptor.VersionOptions))
		}
	}
	if args.IncludeContent != nil {
		queryParams.Add("includeContent", strconv.FormatBool(*args.IncludeContent))
	}
	if args.ResolveLfs != nil {
		queryParams.Add("resolveLfs", strconv.FormatBool(*args.ResolveLfs))
	}
	locationId, _ := uuid.Parse("fb93c0db-47ed-4a31-8c20-47552878fb44")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetItemZip function
type GetItemZipArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The item path.
	Path *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The path scope.  The default is null.
	ScopePath *string
	// (optional) The recursion level of this request. The default is 'none', no recursion.
	RecursionLevel *VersionControlRecursionType
	// (optional) Set to true to include content metadata.  Default is false.
	IncludeContentMetadata *bool
	// (optional) Set to true to include the latest changes.  Default is false.
	LatestProcessedChange *bool
	// (optional) Set to true to download the response as a file.  Default is false.
	Download *bool
	// (optional) Version descriptor.  Default is the default branch for the repository.
	VersionDescriptor *GitVersionDescriptor
	// (optional) Set to true to include item content when requesting json.  Default is false.
	IncludeContent *bool
	// (optional) Set to true to resolve Git LFS pointer files to return actual content from Git LFS.  Default is false.
	ResolveLfs *bool
}

// [Preview API] Get likes for a comment.
func (client *ClientImpl) GetLikes(ctx context.Context, args GetLikesArgs) (*[]webapi.IdentityRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	locationId, _ := uuid.Parse("5f2e2851-1389-425b-a00b-fb2adb3ef31b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []webapi.IdentityRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetLikes function
type GetLikesArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) The ID of the thread that contains the comment.
	ThreadId *int
	// (required) The ID of the comment.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Find the merge bases of two commits, optionally across forks. If otherRepositoryId is not specified, the merge bases will only be calculated within the context of the local repositoryNameOrId.
func (client *ClientImpl) GetMergeBases(ctx context.Context, args GetMergeBasesArgs) (*[]GitCommitRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId
	if args.CommitId == nil || *args.CommitId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CommitId"}
	}
	routeValues["commitId"] = *args.CommitId

	queryParams := url.Values{}
	if args.OtherCommitId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "otherCommitId"}
	}
	queryParams.Add("otherCommitId", *args.OtherCommitId)
	if args.OtherCollectionId != nil {
		queryParams.Add("otherCollectionId", (*args.OtherCollectionId).String())
	}
	if args.OtherRepositoryId != nil {
		queryParams.Add("otherRepositoryId", (*args.OtherRepositoryId).String())
	}
	locationId, _ := uuid.Parse("7cf2abb6-c964-4f7e-9872-f78c66e72e9c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitCommitRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetMergeBases function
type GetMergeBasesArgs struct {
	// (required) ID or name of the local repository.
	RepositoryNameOrId *string
	// (required) First commit, usually the tip of the target branch of the potential merge.
	CommitId *string
	// (required) Other commit, usually the tip of the source branch of the potential merge.
	OtherCommitId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) The collection ID where otherCommitId lives.
	OtherCollectionId *uuid.UUID
	// (optional) The repository ID where otherCommitId lives.
	OtherRepositoryId *uuid.UUID
}

// [Preview API] Get a specific merge operation's details.
func (client *ClientImpl) GetMergeRequest(ctx context.Context, args GetMergeRequestArgs) (*GitMerge, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryNameOrId == nil || *args.RepositoryNameOrId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryNameOrId"}
	}
	routeValues["repositoryNameOrId"] = *args.RepositoryNameOrId
	if args.MergeOperationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.MergeOperationId"}
	}
	routeValues["mergeOperationId"] = strconv.Itoa(*args.MergeOperationId)

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("985f7ae9-844f-4906-9897-7ef41516c0e2")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitMerge
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetMergeRequest function
type GetMergeRequestArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryNameOrId *string
	// (required) OperationId of the merge request.
	MergeOperationId *int
	// (optional) True to include links
	IncludeLinks *bool
}

// [Preview API] Retrieve a list of policy configurations by a given set of scope/filtering criteria.
func (client *ClientImpl) GetPolicyConfigurations(ctx context.Context, args GetPolicyConfigurationsArgs) (*GitPolicyConfigurationResponse, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", (*args.RepositoryId).String())
	}
	if args.RefName != nil {
		queryParams.Add("refName", *args.RefName)
	}
	if args.PolicyType != nil {
		queryParams.Add("policyType", (*args.PolicyType).String())
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("2c420070-a0a2-49cc-9639-c9f271c5ff07")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseBodyValue []policy.PolicyConfiguration
	err = client.Client.UnmarshalBody(resp, &responseBodyValue)

	var responseValue *GitPolicyConfigurationResponse
	xmsContinuationTokenHeader := resp.Header.Get("x-ms-continuationtoken")
	if err == nil {
		responseValue = &GitPolicyConfigurationResponse{
			PolicyConfigurations: &responseBodyValue,
			ContinuationToken:    &xmsContinuationTokenHeader,
		}
	}

	return responseValue, err
}

// Arguments for the GetPolicyConfigurations function
type GetPolicyConfigurationsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The repository id.
	RepositoryId *uuid.UUID
	// (optional) The fully-qualified Git ref name (e.g. refs/heads/master).
	RefName *string
	// (optional) The policy type filter.
	PolicyType *uuid.UUID
	// (optional) Maximum number of policies to return.
	Top *int
	// (optional) Pass a policy configuration ID to fetch the next page of results, up to top number of results, for this endpoint.
	ContinuationToken *string
}

// Retrieve a pull request.
func (client *ClientImpl) GetPullRequest(ctx context.Context, args GetPullRequestArgs) (*GitPullRequest, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.MaxCommentLength != nil {
		queryParams.Add("maxCommentLength", strconv.Itoa(*args.MaxCommentLength))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.IncludeCommits != nil {
		queryParams.Add("includeCommits", strconv.FormatBool(*args.IncludeCommits))
	}
	if args.IncludeWorkItemRefs != nil {
		queryParams.Add("includeWorkItemRefs", strconv.FormatBool(*args.IncludeWorkItemRefs))
	}
	locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequest function
type GetPullRequestArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) The ID of the pull request to retrieve.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Not used.
	MaxCommentLength *int
	// (optional) Not used.
	Skip *int
	// (optional) Not used.
	Top *int
	// (optional) If true, the pull request will be returned with the associated commits.
	IncludeCommits *bool
	// (optional) If true, the pull request will be returned with the associated work item references.
	IncludeWorkItemRefs *bool
}

// Retrieve a pull request.
func (client *ClientImpl) GetPullRequestById(ctx context.Context, args GetPullRequestByIdArgs) (*GitPullRequest, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("01a46dea-7d46-4d40-bc84-319e7c260d99")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestById function
type GetPullRequestByIdArgs struct {
	// (required) The ID of the pull request to retrieve.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Get the commits for the specified pull request.
func (client *ClientImpl) GetPullRequestCommits(ctx context.Context, args GetPullRequestCommitsArgs) (*GetPullRequestCommitsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("52823034-34a8-4576-922c-8d8b77e9e4c4")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetPullRequestCommitsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetPullRequestCommits function
type GetPullRequestCommitsArgs struct {
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Maximum number of commits to return.
	Top *int
	// (optional) The continuation token used for pagination.
	ContinuationToken *string
}

// Return type for the GetPullRequestCommits function
type GetPullRequestCommitsResponseValue struct {
	Value []GitCommitRef
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// Get the specified iteration for a pull request.
func (client *ClientImpl) GetPullRequestIteration(ctx context.Context, args GetPullRequestIterationArgs) (*GitPullRequestIteration, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	locationId, _ := uuid.Parse("d43911ee-6958-46b0-a42b-8445b8a0d004")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestIteration
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIteration function
type GetPullRequestIterationArgs struct {
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration to return.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieve the changes made in a pull request between two iterations.
func (client *ClientImpl) GetPullRequestIterationChanges(ctx context.Context, args GetPullRequestIterationChangesArgs) (*GitPullRequestIterationChanges, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.CompareTo != nil {
		queryParams.Add("$compareTo", strconv.Itoa(*args.CompareTo))
	}
	locationId, _ := uuid.Parse("4216bdcf-b6b1-4d59-8b82-c34cc183fc8b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestIterationChanges
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIterationChanges function
type GetPullRequestIterationChangesArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration. <br /> Iteration IDs are zero-based with zero indicating the common commit between the source and target branches. Iteration one is the head of the source branch at the time the pull request is created and subsequent iterations are created when there are pushes to the source branch.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Optional. The number of changes to retrieve.  The default value is 100 and the maximum value is 2000.
	Top *int
	// (optional) Optional. The number of changes to ignore.  For example, to retrieve changes 101-150, set top 50 and skip to 100.
	Skip *int
	// (optional) ID of the pull request iteration to compare against.  The default value is zero which indicates the comparison is made against the common commit between the source and target branches
	CompareTo *int
}

// Get the commits for the specified iteration of a pull request.
func (client *ClientImpl) GetPullRequestIterationCommits(ctx context.Context, args GetPullRequestIterationCommitsArgs) (*[]GitCommitRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("skip", strconv.Itoa(*args.Skip))
	}
	locationId, _ := uuid.Parse("e7ea0883-095f-4926-b5fb-f24691c26fb9")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitCommitRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIterationCommits function
type GetPullRequestIterationCommitsArgs struct {
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the iteration from which to get the commits.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Maximum number of commits to return. The maximum number of commits that can be returned per batch is 500.
	Top *int
	// (optional) Number of commits to skip.
	Skip *int
}

// Get the list of iterations for the specified pull request.
func (client *ClientImpl) GetPullRequestIterations(ctx context.Context, args GetPullRequestIterationsArgs) (*[]GitPullRequestIteration, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.IncludeCommits != nil {
		queryParams.Add("includeCommits", strconv.FormatBool(*args.IncludeCommits))
	}
	locationId, _ := uuid.Parse("d43911ee-6958-46b0-a42b-8445b8a0d004")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequestIteration
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIterations function
type GetPullRequestIterationsArgs struct {
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) If true, include the commits associated with each iteration in the response.
	IncludeCommits *bool
}

// [Preview API] Get the specific pull request iteration status by ID. The status ID is unique within the pull request across all iterations.
func (client *ClientImpl) GetPullRequestIterationStatus(ctx context.Context, args GetPullRequestIterationStatusArgs) (*GitPullRequestStatus, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)
	if args.StatusId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.StatusId"}
	}
	routeValues["statusId"] = strconv.Itoa(*args.StatusId)

	locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIterationStatus function
type GetPullRequestIterationStatusArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration.
	IterationId *int
	// (required) ID of the pull request status.
	StatusId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get all the statuses associated with a pull request iteration.
func (client *ClientImpl) GetPullRequestIterationStatuses(ctx context.Context, args GetPullRequestIterationStatusesArgs) (*[]GitPullRequestStatus, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequestStatus
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestIterationStatuses function
type GetPullRequestIterationStatusesArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Retrieves a single label that has been assigned to a pull request.
func (client *ClientImpl) GetPullRequestLabel(ctx context.Context, args GetPullRequestLabelArgs) (*core.WebApiTagDefinition, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.LabelIdOrName == nil || *args.LabelIdOrName == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.LabelIdOrName"}
	}
	routeValues["labelIdOrName"] = *args.LabelIdOrName

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue core.WebApiTagDefinition
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestLabel function
type GetPullRequestLabelArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) The name or ID of the label requested.
	LabelIdOrName *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Project ID or project name.
	ProjectId *string
}

// [Preview API] Get all the labels assigned to a pull request.
func (client *ClientImpl) GetPullRequestLabels(ctx context.Context, args GetPullRequestLabelsArgs) (*[]core.WebApiTagDefinition, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	locationId, _ := uuid.Parse("f22387e3-984e-4c52-9c6d-fbb8f14c812d")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []core.WebApiTagDefinition
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestLabels function
type GetPullRequestLabelsArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) Project ID or project name.
	ProjectId *string
}

// [Preview API] Get external properties of the pull request.
func (client *ClientImpl) GetPullRequestProperties(ctx context.Context, args GetPullRequestPropertiesArgs) (interface{}, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("48a52185-5b9e-4736-9dc1-bb1e2feac80b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the GetPullRequestProperties function
type GetPullRequestPropertiesArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// This API is used to find what pull requests are related to a given commit.  It can be used to either find the pull request that created a particular merge commit or it can be used to find all pull requests that have ever merged a particular commit.  The input is a list of queries which each contain a list of commits. For each commit that you search against, you will get back a dictionary of commit -> pull requests.
func (client *ClientImpl) GetPullRequestQuery(ctx context.Context, args GetPullRequestQueryArgs) (*GitPullRequestQuery, error) {
	if args.Queries == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Queries"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	body, marshalErr := json.Marshal(*args.Queries)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("b3a6eebe-9cf0-49ea-b6cb-1a4c5f5007b0")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestQuery
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestQuery function
type GetPullRequestQueryArgs struct {
	// (required) The list of queries to perform.
	Queries *GitPullRequestQuery
	// (required) ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// Retrieve information about a particular reviewer on a pull request
func (client *ClientImpl) GetPullRequestReviewer(ctx context.Context, args GetPullRequestReviewerArgs) (*IdentityRefWithVote, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ReviewerId == nil || *args.ReviewerId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.ReviewerId"}
	}
	routeValues["reviewerId"] = *args.ReviewerId

	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue IdentityRefWithVote
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestReviewer function
type GetPullRequestReviewerArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the reviewer.
	ReviewerId *string
	// (optional) Project ID or project name
	Project *string
}

// Retrieve the reviewers for a pull request
func (client *ClientImpl) GetPullRequestReviewers(ctx context.Context, args GetPullRequestReviewersArgs) (*[]IdentityRefWithVote, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []IdentityRefWithVote
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestReviewers function
type GetPullRequestReviewersArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieve all pull requests matching a specified criteria.
func (client *ClientImpl) GetPullRequests(ctx context.Context, args GetPullRequestsArgs) (*[]GitPullRequest, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.SearchCriteria == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "searchCriteria"}
	}
	if args.SearchCriteria.RepositoryId != nil {
		queryParams.Add("searchCriteria.repositoryId", (*args.SearchCriteria.RepositoryId).String())
	}
	if args.SearchCriteria.CreatorId != nil {
		queryParams.Add("searchCriteria.creatorId", (*args.SearchCriteria.CreatorId).String())
	}
	if args.SearchCriteria.ReviewerId != nil {
		queryParams.Add("searchCriteria.reviewerId", (*args.SearchCriteria.ReviewerId).String())
	}
	if args.SearchCriteria.Status != nil {
		queryParams.Add("searchCriteria.status", string(*args.SearchCriteria.Status))
	}
	if args.SearchCriteria.TargetRefName != nil {
		queryParams.Add("searchCriteria.targetRefName", *args.SearchCriteria.TargetRefName)
	}
	if args.SearchCriteria.SourceRepositoryId != nil {
		queryParams.Add("searchCriteria.sourceRepositoryId", (*args.SearchCriteria.SourceRepositoryId).String())
	}
	if args.SearchCriteria.SourceRefName != nil {
		queryParams.Add("searchCriteria.sourceRefName", *args.SearchCriteria.SourceRefName)
	}
	if args.SearchCriteria.IncludeLinks != nil {
		queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
	}
	if args.MaxCommentLength != nil {
		queryParams.Add("maxCommentLength", strconv.Itoa(*args.MaxCommentLength))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequest
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequests function
type GetPullRequestsArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) Pull requests will be returned that match this search criteria.
	SearchCriteria *GitPullRequestSearchCriteria
	// (optional) Project ID or project name
	Project *string
	// (optional) Not used.
	MaxCommentLength *int
	// (optional) The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
	Skip *int
	// (optional) The number of pull requests to retrieve.
	Top *int
}

// Retrieve all pull requests matching a specified criteria.
func (client *ClientImpl) GetPullRequestsByProject(ctx context.Context, args GetPullRequestsByProjectArgs) (*[]GitPullRequest, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.SearchCriteria == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "searchCriteria"}
	}
	if args.SearchCriteria.RepositoryId != nil {
		queryParams.Add("searchCriteria.repositoryId", (*args.SearchCriteria.RepositoryId).String())
	}
	if args.SearchCriteria.CreatorId != nil {
		queryParams.Add("searchCriteria.creatorId", (*args.SearchCriteria.CreatorId).String())
	}
	if args.SearchCriteria.ReviewerId != nil {
		queryParams.Add("searchCriteria.reviewerId", (*args.SearchCriteria.ReviewerId).String())
	}
	if args.SearchCriteria.Status != nil {
		queryParams.Add("searchCriteria.status", string(*args.SearchCriteria.Status))
	}
	if args.SearchCriteria.TargetRefName != nil {
		queryParams.Add("searchCriteria.targetRefName", *args.SearchCriteria.TargetRefName)
	}
	if args.SearchCriteria.SourceRepositoryId != nil {
		queryParams.Add("searchCriteria.sourceRepositoryId", (*args.SearchCriteria.SourceRepositoryId).String())
	}
	if args.SearchCriteria.SourceRefName != nil {
		queryParams.Add("searchCriteria.sourceRefName", *args.SearchCriteria.SourceRefName)
	}
	if args.SearchCriteria.IncludeLinks != nil {
		queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
	}
	if args.MaxCommentLength != nil {
		queryParams.Add("maxCommentLength", strconv.Itoa(*args.MaxCommentLength))
	}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	locationId, _ := uuid.Parse("a5d28130-9cd2-40fa-9f08-902e7daa9efb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequest
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestsByProject function
type GetPullRequestsByProjectArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) Pull requests will be returned that match this search criteria.
	SearchCriteria *GitPullRequestSearchCriteria
	// (optional) Not used.
	MaxCommentLength *int
	// (optional) The number of pull requests to ignore. For example, to retrieve results 101-150, set top to 50 and skip to 100.
	Skip *int
	// (optional) The number of pull requests to retrieve.
	Top *int
}

// [Preview API] Get the specific pull request status by ID. The status ID is unique within the pull request across all iterations.
func (client *ClientImpl) GetPullRequestStatus(ctx context.Context, args GetPullRequestStatusArgs) (*GitPullRequestStatus, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.StatusId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.StatusId"}
	}
	routeValues["statusId"] = strconv.Itoa(*args.StatusId)

	locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestStatus
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestStatus function
type GetPullRequestStatusArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request status.
	StatusId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Get all the statuses associated with a pull request.
func (client *ClientImpl) GetPullRequestStatuses(ctx context.Context, args GetPullRequestStatusesArgs) (*[]GitPullRequestStatus, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequestStatus
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestStatuses function
type GetPullRequestStatusesArgs struct {
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieve a thread in a pull request.
func (client *ClientImpl) GetPullRequestThread(ctx context.Context, args GetPullRequestThreadArgs) (*GitPullRequestCommentThread, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)

	queryParams := url.Values{}
	if args.Iteration != nil {
		queryParams.Add("$iteration", strconv.Itoa(*args.Iteration))
	}
	if args.BaseIteration != nil {
		queryParams.Add("$baseIteration", strconv.Itoa(*args.BaseIteration))
	}
	locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestCommentThread
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestThread function
type GetPullRequestThreadArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread.
	ThreadId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) If specified, thread position will be tracked using this iteration as the right side of the diff.
	Iteration *int
	// (optional) If specified, thread position will be tracked using this iteration as the left side of the diff.
	BaseIteration *int
}

// Retrieve a list of work items associated with a pull request.
func (client *ClientImpl) GetPullRequestWorkItemRefs(ctx context.Context, args GetPullRequestWorkItemRefsArgs) (*[]webapi.ResourceRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	locationId, _ := uuid.Parse("0a637fcc-5370-4ce8-b0e8-98091f5f9482")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []webapi.ResourceRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPullRequestWorkItemRefs function
type GetPullRequestWorkItemRefsArgs struct {
	// (required) ID or name of the repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Retrieves a particular push.
func (client *ClientImpl) GetPush(ctx context.Context, args GetPushArgs) (*GitPush, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PushId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PushId"}
	}
	routeValues["pushId"] = strconv.Itoa(*args.PushId)

	queryParams := url.Values{}
	if args.IncludeCommits != nil {
		queryParams.Add("includeCommits", strconv.Itoa(*args.IncludeCommits))
	}
	if args.IncludeRefUpdates != nil {
		queryParams.Add("includeRefUpdates", strconv.FormatBool(*args.IncludeRefUpdates))
	}
	locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPush
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPush function
type GetPushArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) ID of the push.
	PushId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) The number of commits to include in the result.
	IncludeCommits *int
	// (optional) If true, include the list of refs that were updated by the push.
	IncludeRefUpdates *bool
}

// Retrieve a list of commits associated with a particular push.
func (client *ClientImpl) GetPushCommits(ctx context.Context, args GetPushCommitsArgs) (*[]GitCommitRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.PushId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "pushId"}
	}
	queryParams.Add("pushId", strconv.Itoa(*args.PushId))
	if args.Top != nil {
		queryParams.Add("top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("skip", strconv.Itoa(*args.Skip))
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	locationId, _ := uuid.Parse("c2570c3b-5b3f-41b8-98bf-5407bfde8d58")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitCommitRef
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPushCommits function
type GetPushCommitsArgs struct {
	// (required) The id or friendly name of the repository. To use the friendly name, projectId must also be specified.
	RepositoryId *string
	// (required) The id of the push.
	PushId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) The maximum number of commits to return ("get the top x commits").
	Top *int
	// (optional) The number of commits to skip.
	Skip *int
	// (optional) Set to false to avoid including REST Url links for resources. Defaults to true.
	IncludeLinks *bool
}

// Retrieves pushes associated with the specified repository.
func (client *ClientImpl) GetPushes(ctx context.Context, args GetPushesArgs) (*[]GitPush, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Skip != nil {
		queryParams.Add("$skip", strconv.Itoa(*args.Skip))
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.SearchCriteria != nil {
		if args.SearchCriteria.FromDate != nil {
			queryParams.Add("searchCriteria.fromDate", (*args.SearchCriteria.FromDate).String())
		}
		if args.SearchCriteria.ToDate != nil {
			queryParams.Add("searchCriteria.toDate", (*args.SearchCriteria.ToDate).String())
		}
		if args.SearchCriteria.PusherId != nil {
			queryParams.Add("searchCriteria.pusherId", (*args.SearchCriteria.PusherId).String())
		}
		if args.SearchCriteria.RefName != nil {
			queryParams.Add("searchCriteria.refName", *args.SearchCriteria.RefName)
		}
		if args.SearchCriteria.IncludeRefUpdates != nil {
			queryParams.Add("searchCriteria.includeRefUpdates", strconv.FormatBool(*args.SearchCriteria.IncludeRefUpdates))
		}
		if args.SearchCriteria.IncludeLinks != nil {
			queryParams.Add("searchCriteria.includeLinks", strconv.FormatBool(*args.SearchCriteria.IncludeLinks))
		}
	}
	locationId, _ := uuid.Parse("ea98d07b-3c87-4971-8ede-a613694ffb55")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPush
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetPushes function
type GetPushesArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Number of pushes to skip.
	Skip *int
	// (optional) Number of pushes to return.
	Top *int
	// (optional) Search criteria attributes: fromDate, toDate, pusherId, refName, includeRefUpdates or includeLinks. fromDate: Start date to search from. toDate: End date to search to. pusherId: Identity of the person who submitted the push. refName: Branch name to consider. includeRefUpdates: If true, include the list of refs that were updated by the push. includeLinks: Whether to include the _links field on the shallow references.
	SearchCriteria *GitPushSearchCriteria
}

// [Preview API] Retrieve soft-deleted git repositories from the recycle bin.
func (client *ClientImpl) GetRecycleBinRepositories(ctx context.Context, args GetRecycleBinRepositoriesArgs) (*[]GitDeletedRepository, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitDeletedRepository
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRecycleBinRepositories function
type GetRecycleBinRepositoriesArgs struct {
	// (required) Project ID or project name
	Project *string
}

// [Preview API] Gets the refs favorite for a favorite Id.
func (client *ClientImpl) GetRefFavorite(ctx context.Context, args GetRefFavoriteArgs) (*GitRefFavorite, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.FavoriteId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.FavoriteId"}
	}
	routeValues["favoriteId"] = strconv.Itoa(*args.FavoriteId)

	locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRefFavorite
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRefFavorite function
type GetRefFavoriteArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The Id of the requested ref favorite.
	FavoriteId *int
}

// [Preview API] Gets the refs favorites for a repo and an identity.
func (client *ClientImpl) GetRefFavorites(ctx context.Context, args GetRefFavoritesArgs) (*[]GitRefFavorite, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project

	queryParams := url.Values{}
	if args.RepositoryId != nil {
		queryParams.Add("repositoryId", *args.RepositoryId)
	}
	if args.IdentityId != nil {
		queryParams.Add("identityId", *args.IdentityId)
	}
	locationId, _ := uuid.Parse("876f70af-5792-485a-a1c7-d0a7b2f42bbb")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitRefFavorite
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRefFavorites function
type GetRefFavoritesArgs struct {
	// (required) Project ID or project name
	Project *string
	// (optional) The id of the repository.
	RepositoryId *string
	// (optional) The id of the identity whose favorites are to be retrieved. If null, the requesting identity is used.
	IdentityId *string
}

// Queries the provided repository for its refs and returns them.
func (client *ClientImpl) GetRefs(ctx context.Context, args GetRefsArgs) (*GetRefsResponseValue, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Filter != nil {
		queryParams.Add("filter", *args.Filter)
	}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	if args.IncludeStatuses != nil {
		queryParams.Add("includeStatuses", strconv.FormatBool(*args.IncludeStatuses))
	}
	if args.IncludeMyBranches != nil {
		queryParams.Add("includeMyBranches", strconv.FormatBool(*args.IncludeMyBranches))
	}
	if args.LatestStatusesOnly != nil {
		queryParams.Add("latestStatusesOnly", strconv.FormatBool(*args.LatestStatusesOnly))
	}
	if args.PeelTags != nil {
		queryParams.Add("peelTags", strconv.FormatBool(*args.PeelTags))
	}
	if args.FilterContains != nil {
		queryParams.Add("filterContains", *args.FilterContains)
	}
	if args.Top != nil {
		queryParams.Add("$top", strconv.Itoa(*args.Top))
	}
	if args.ContinuationToken != nil {
		queryParams.Add("continuationToken", *args.ContinuationToken)
	}
	locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GetRefsResponseValue
	responseValue.ContinuationToken = resp.Header.Get(azuredevops.HeaderKeyContinuationToken)
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue.Value)
	return &responseValue, err
}

// Arguments for the GetRefs function
type GetRefsArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) [optional] A filter to apply to the refs (starts with).
	Filter *string
	// (optional) [optional] Specifies if referenceLinks should be included in the result. default is false.
	IncludeLinks *bool
	// (optional) [optional] Includes up to the first 1000 commit statuses for each ref. The default value is false.
	IncludeStatuses *bool
	// (optional) [optional] Includes only branches that the user owns, the branches the user favorites, and the default branch. The default value is false. Cannot be combined with the filter parameter.
	IncludeMyBranches *bool
	// (optional) [optional] True to include only the tip commit status for each ref. This option requires `includeStatuses` to be true. The default value is false.
	LatestStatusesOnly *bool
	// (optional) [optional] Annotated tags will populate the PeeledObjectId property. default is false.
	PeelTags *bool
	// (optional) [optional] A filter to apply to the refs (contains).
	FilterContains *string
	// (optional) [optional] Maximum number of refs to return. It cannot be bigger than 1000. If it is not provided but continuationToken is, top will default to 100.
	Top *int
	// (optional) The continuation token used for pagination.
	ContinuationToken *string
}

// Return type for the GetRefs function
type GetRefsResponseValue struct {
	Value []GitRef
	// The continuation token to be used to get the next page of results.
	ContinuationToken string
}

// Retrieve git repositories.
func (client *ClientImpl) GetRepositories(ctx context.Context, args GetRepositoriesArgs) (*[]GitRepository, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}

	queryParams := url.Values{}
	if args.IncludeLinks != nil {
		queryParams.Add("includeLinks", strconv.FormatBool(*args.IncludeLinks))
	}
	if args.IncludeAllUrls != nil {
		queryParams.Add("includeAllUrls", strconv.FormatBool(*args.IncludeAllUrls))
	}
	if args.IncludeHidden != nil {
		queryParams.Add("includeHidden", strconv.FormatBool(*args.IncludeHidden))
	}
	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitRepository
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRepositories function
type GetRepositoriesArgs struct {
	// (optional) Project ID or project name
	Project *string
	// (optional) [optional] True to include reference links. The default value is false.
	IncludeLinks *bool
	// (optional) [optional] True to include all remote URLs. The default value is false.
	IncludeAllUrls *bool
	// (optional) [optional] True to include hidden repositories. The default value is false.
	IncludeHidden *bool
}

// Retrieve a git repository.
func (client *ClientImpl) GetRepository(ctx context.Context, args GetRepositoryArgs) (*GitRepository, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRepository
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRepository function
type GetRepositoryArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// Retrieve a git repository.
func (client *ClientImpl) GetRepositoryWithParent(ctx context.Context, args GetRepositoryWithParentArgs) (*GitRepository, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.IncludeParent == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "includeParent"}
	}
	queryParams.Add("includeParent", strconv.FormatBool(*args.IncludeParent))
	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRepository
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRepositoryWithParent function
type GetRepositoryWithParentArgs struct {
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) True to include parent repository. Only available in authenticated calls.
	IncludeParent *bool
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Retrieve information about a revert operation by revert Id.
func (client *ClientImpl) GetRevert(ctx context.Context, args GetRevertArgs) (*GitRevert, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RevertId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RevertId"}
	}
	routeValues["revertId"] = strconv.Itoa(*args.RevertId)
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRevert
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRevert function
type GetRevertArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the revert operation.
	RevertId *int
	// (required) ID of the repository.
	RepositoryId *string
}

// [Preview API] Retrieve information about a revert operation for a specific branch.
func (client *ClientImpl) GetRevertForRefName(ctx context.Context, args GetRevertForRefNameArgs) (*GitRevert, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.RefName == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "refName"}
	}
	queryParams.Add("refName", *args.RefName)
	locationId, _ := uuid.Parse("bc866058-5449-4715-9cf1-a510b6ff193c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRevert
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetRevertForRefName function
type GetRevertForRefNameArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) ID of the repository.
	RepositoryId *string
	// (required) The GitAsyncRefOperationParameters generatedRefName used for the revert operation.
	RefName *string
}

// Get statuses associated with the Git commit.
func (client *ClientImpl) GetStatuses(ctx context.Context, args GetStatusesArgs) (*[]GitStatus, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.CommitId == nil || *args.CommitId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.CommitId"}
	}
	routeValues["commitId"] = *args.CommitId
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Top != nil {
		queryParams.Add("top", strconv.Itoa(*args.Top))
	}
	if args.Skip != nil {
		queryParams.Add("skip", strconv.Itoa(*args.Skip))
	}
	if args.LatestOnly != nil {
		queryParams.Add("latestOnly", strconv.FormatBool(*args.LatestOnly))
	}
	locationId, _ := uuid.Parse("428dd4fb-fda5-4722-af02-9313b80305da")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitStatus
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetStatuses function
type GetStatusesArgs struct {
	// (required) ID of the Git commit.
	CommitId *string
	// (required) ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Optional. The number of statuses to retrieve. Default is 1000.
	Top *int
	// (optional) Optional. The number of statuses to ignore. Default is 0. For example, to retrieve results 101-150, set top to 50 and skip to 100.
	Skip *int
	// (optional) The flag indicates whether to get only latest statuses grouped by `Context.Name` and `Context.Genre`.
	LatestOnly *bool
}

// [Preview API] Retrieve a pull request suggestion for a particular repository or team project.
func (client *ClientImpl) GetSuggestions(ctx context.Context, args GetSuggestionsArgs) (*[]GitSuggestion, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	locationId, _ := uuid.Parse("9393b4fb-4445-4919-972b-9ad16f442d83")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, nil, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitSuggestion
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetSuggestions function
type GetSuggestionsArgs struct {
	// (required) ID of the git repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
}

// Retrieve all threads in a pull request.
func (client *ClientImpl) GetThreads(ctx context.Context, args GetThreadsArgs) (*[]GitPullRequestCommentThread, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	queryParams := url.Values{}
	if args.Iteration != nil {
		queryParams.Add("$iteration", strconv.Itoa(*args.Iteration))
	}
	if args.BaseIteration != nil {
		queryParams.Add("$baseIteration", strconv.Itoa(*args.BaseIteration))
	}
	locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitPullRequestCommentThread
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetThreads function
type GetThreadsArgs struct {
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
	// (optional) If specified, thread positions will be tracked using this iteration as the right side of the diff.
	Iteration *int
	// (optional) If specified, thread positions will be tracked using this iteration as the left side of the diff.
	BaseIteration *int
}

// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
func (client *ClientImpl) GetTree(ctx context.Context, args GetTreeArgs) (*GitTreeRef, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.Sha1 == nil || *args.Sha1 == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Sha1"}
	}
	routeValues["sha1"] = *args.Sha1

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	if args.Recursive != nil {
		queryParams.Add("recursive", strconv.FormatBool(*args.Recursive))
	}
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	locationId, _ := uuid.Parse("729f6437-6f92-44ec-8bee-273a7111063c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitTreeRef
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the GetTree function
type GetTreeArgs struct {
	// (required) Repository Id.
	RepositoryId *string
	// (required) SHA1 hash of the tree object.
	Sha1 *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Project Id.
	ProjectId *string
	// (optional) Search recursively. Include trees underneath this tree. Default is false.
	Recursive *bool
	// (optional) Name to use if a .zip file is returned. Default is the object ID.
	FileName *string
}

// The Tree endpoint returns the collection of objects underneath the specified tree. Trees are folders in a Git repository.
func (client *ClientImpl) GetTreeZip(ctx context.Context, args GetTreeZipArgs) (io.ReadCloser, error) {
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.Sha1 == nil || *args.Sha1 == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Sha1"}
	}
	routeValues["sha1"] = *args.Sha1

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	if args.Recursive != nil {
		queryParams.Add("recursive", strconv.FormatBool(*args.Recursive))
	}
	if args.FileName != nil {
		queryParams.Add("fileName", *args.FileName)
	}
	locationId, _ := uuid.Parse("729f6437-6f92-44ec-8bee-273a7111063c")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1", routeValues, queryParams, nil, "", "application/zip", nil)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Arguments for the GetTreeZip function
type GetTreeZipArgs struct {
	// (required) Repository Id.
	RepositoryId *string
	// (required) SHA1 hash of the tree object.
	Sha1 *string
	// (optional) Project ID or project name
	Project *string
	// (optional) Project Id.
	ProjectId *string
	// (optional) Search recursively. Include trees underneath this tree. Default is false.
	Recursive *bool
	// (optional) Name to use if a .zip file is returned. Default is the object ID.
	FileName *string
}

// [Preview API] Retrieve import requests for a repository.
func (client *ClientImpl) QueryImportRequests(ctx context.Context, args QueryImportRequestsArgs) (*[]GitImportRequest, error) {
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.IncludeAbandoned != nil {
		queryParams.Add("includeAbandoned", strconv.FormatBool(*args.IncludeAbandoned))
	}
	locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
	resp, err := client.Client.Send(ctx, http.MethodGet, locationId, "5.1-preview.1", routeValues, queryParams, nil, "", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitImportRequest
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the QueryImportRequests function
type QueryImportRequestsArgs struct {
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) True to include abandoned import requests in the results.
	IncludeAbandoned *bool
}

// [Preview API] Recover a soft-deleted Git repository. Recently deleted repositories go into a soft-delete state for a period of time before they are hard deleted and become unrecoverable.
func (client *ClientImpl) RestoreRepositoryFromRecycleBin(ctx context.Context, args RestoreRepositoryFromRecycleBinArgs) (*GitRepository, error) {
	if args.RepositoryDetails == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RepositoryDetails"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = (*args.RepositoryId).String()

	body, marshalErr := json.Marshal(*args.RepositoryDetails)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("a663da97-81db-4eb3-8b83-287670f63073")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRepository
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the RestoreRepositoryFromRecycleBin function
type RestoreRepositoryFromRecycleBinArgs struct {
	// (required)
	RepositoryDetails *GitRecycleBinRepositoryDetails
	// (required) Project ID or project name
	Project *string
	// (required) The ID of the repository.
	RepositoryId *uuid.UUID
}

// [Preview API] Sends an e-mail notification about a specific pull request to a set of recipients
func (client *ClientImpl) SharePullRequest(ctx context.Context, args SharePullRequestArgs) error {
	if args.UserMessage == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.UserMessage"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.UserMessage)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("696f3a82-47c9-487f-9117-b9d00972ca84")
	_, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the SharePullRequest function
type SharePullRequestArgs struct {
	// (required)
	UserMessage *ShareNotificationContext
	// (required) ID of the git repository.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Update a comment associated with a specific thread in a pull request.
func (client *ClientImpl) UpdateComment(ctx context.Context, args UpdateCommentArgs) (*Comment, error) {
	if args.Comment == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.Comment"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)
	if args.CommentId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommentId"}
	}
	routeValues["commentId"] = strconv.Itoa(*args.CommentId)

	body, marshalErr := json.Marshal(*args.Comment)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("965a3ec7-5ed8-455a-bdcb-835a5ea7fe7b")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue Comment
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateComment function
type UpdateCommentArgs struct {
	// (required) The comment content that should be updated. Comments can be up to 150,000 characters.
	Comment *Comment
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread that the desired comment is in.
	ThreadId *int
	// (required) ID of the comment to update.
	CommentId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Retry or abandon a failed import request.
func (client *ClientImpl) UpdateImportRequest(ctx context.Context, args UpdateImportRequestArgs) (*GitImportRequest, error) {
	if args.ImportRequestToUpdate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ImportRequestToUpdate"}
	}
	routeValues := make(map[string]string)
	if args.Project == nil || *args.Project == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.Project"}
	}
	routeValues["project"] = *args.Project
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.ImportRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ImportRequestId"}
	}
	routeValues["importRequestId"] = strconv.Itoa(*args.ImportRequestId)

	body, marshalErr := json.Marshal(*args.ImportRequestToUpdate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("01828ddc-3600-4a41-8633-99b3a73a0eb3")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitImportRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateImportRequest function
type UpdateImportRequestArgs struct {
	// (required) The updated version of the import request. Currently, the only change allowed is setting the Status to Queued or Abandoned.
	ImportRequestToUpdate *GitImportRequest
	// (required) Project ID or project name
	Project *string
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The unique identifier for the import request to update.
	ImportRequestId *int
}

// Update a pull request
func (client *ClientImpl) UpdatePullRequest(ctx context.Context, args UpdatePullRequestArgs) (*GitPullRequest, error) {
	if args.GitPullRequestToUpdate == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.GitPullRequestToUpdate"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.GitPullRequestToUpdate)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("9946fd70-0d40-406e-b686-b4744cbbcc37")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequest
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdatePullRequest function
type UpdatePullRequestArgs struct {
	// (required) The pull request content that should be updated.
	GitPullRequestToUpdate *GitPullRequest
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request to update.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Update pull request iteration statuses collection. The only supported operation type is `remove`.
func (client *ClientImpl) UpdatePullRequestIterationStatuses(ctx context.Context, args UpdatePullRequestIterationStatusesArgs) error {
	if args.PatchDocument == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PatchDocument"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.IterationId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.IterationId"}
	}
	routeValues["iterationId"] = strconv.Itoa(*args.IterationId)

	body, marshalErr := json.Marshal(*args.PatchDocument)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("75cf11c5-979f-4038-a76e-058a06adf2bf")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePullRequestIterationStatuses function
type UpdatePullRequestIterationStatusesArgs struct {
	// (required) Operations to apply to the pull request statuses in JSON Patch format.
	PatchDocument *[]webapi.JsonPatchOperation
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the pull request iteration.
	IterationId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Create or update pull request external properties. The patch operation can be `add`, `replace` or `remove`. For `add` operation, the path can be empty. If the path is empty, the value must be a list of key value pairs. For `replace` operation, the path cannot be empty. If the path does not exist, the property will be added to the collection. For `remove` operation, the path cannot be empty. If the path does not exist, no action will be performed.
func (client *ClientImpl) UpdatePullRequestProperties(ctx context.Context, args UpdatePullRequestPropertiesArgs) (interface{}, error) {
	if args.PatchDocument == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PatchDocument"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.PatchDocument)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("48a52185-5b9e-4736-9dc1-bb1e2feac80b")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue interface{}
	err = client.Client.UnmarshalBody(resp, responseValue)
	return responseValue, err
}

// Arguments for the UpdatePullRequestProperties function
type UpdatePullRequestPropertiesArgs struct {
	// (required) Properties to add, replace or remove in JSON Patch format.
	PatchDocument *[]webapi.JsonPatchOperation
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Reset the votes of multiple reviewers on a pull request.  NOTE: This endpoint only supports updating votes, but does not support updating required reviewers (use policy) or display names.
func (client *ClientImpl) UpdatePullRequestReviewers(ctx context.Context, args UpdatePullRequestReviewersArgs) error {
	if args.PatchVotes == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PatchVotes"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.PatchVotes)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("4b6702c7-aa35-4b89-9c96-b9abf6d3e540")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePullRequestReviewers function
type UpdatePullRequestReviewersArgs struct {
	// (required) IDs of the reviewers whose votes will be reset to zero
	PatchVotes *[]IdentityRefWithVote
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// [Preview API] Update pull request statuses collection. The only supported operation type is `remove`.
func (client *ClientImpl) UpdatePullRequestStatuses(ctx context.Context, args UpdatePullRequestStatusesArgs) error {
	if args.PatchDocument == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PatchDocument"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)

	body, marshalErr := json.Marshal(*args.PatchDocument)
	if marshalErr != nil {
		return marshalErr
	}
	locationId, _ := uuid.Parse("b5f6bb4f-8d1e-4d79-8d11-4c9172c99c35")
	_, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1-preview.1", routeValues, nil, bytes.NewReader(body), "application/json-patch+json", "application/json", nil)
	if err != nil {
		return err
	}

	return nil
}

// Arguments for the UpdatePullRequestStatuses function
type UpdatePullRequestStatusesArgs struct {
	// (required) Operations to apply to the pull request statuses in JSON Patch format.
	PatchDocument *[]webapi.JsonPatchOperation
	// (required) The repository ID of the pull request’s target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (optional) Project ID or project name
	Project *string
}

// Lock or Unlock a branch.
func (client *ClientImpl) UpdateRef(ctx context.Context, args UpdateRefArgs) (*GitRef, error) {
	if args.NewRefInfo == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.NewRefInfo"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.Filter == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "filter"}
	}
	queryParams.Add("filter", *args.Filter)
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	body, marshalErr := json.Marshal(*args.NewRefInfo)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRef
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRef function
type UpdateRefArgs struct {
	// (required) The ref update action (lock/unlock) to perform
	NewRefInfo *GitRefUpdate
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (required) The name of the branch to lock/unlock
	Filter *string
	// (optional) Project ID or project name
	Project *string
	// (optional) ID or name of the team project. Optional if specifying an ID for repository.
	ProjectId *string
}

// Creating, updating, or deleting refs(branches).
func (client *ClientImpl) UpdateRefs(ctx context.Context, args UpdateRefsArgs) (*[]GitRefUpdateResult, error) {
	if args.RefUpdates == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RefUpdates"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId

	queryParams := url.Values{}
	if args.ProjectId != nil {
		queryParams.Add("projectId", *args.ProjectId)
	}
	body, marshalErr := json.Marshal(*args.RefUpdates)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("2d874a60-a811-4f62-9c9f-963a6ea0a55b")
	resp, err := client.Client.Send(ctx, http.MethodPost, locationId, "5.1", routeValues, queryParams, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue []GitRefUpdateResult
	err = client.Client.UnmarshalCollectionBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRefs function
type UpdateRefsArgs struct {
	// (required) List of ref updates to attempt to perform
	RefUpdates *[]GitRefUpdate
	// (required) The name or ID of the repository.
	RepositoryId *string
	// (optional) Project ID or project name
	Project *string
	// (optional) ID or name of the team project. Optional if specifying an ID for repository.
	ProjectId *string
}

// Updates the Git repository with either a new repo name or a new default branch.
func (client *ClientImpl) UpdateRepository(ctx context.Context, args UpdateRepositoryArgs) (*GitRepository, error) {
	if args.NewRepositoryInfo == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.NewRepositoryInfo"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = (*args.RepositoryId).String()

	body, marshalErr := json.Marshal(*args.NewRepositoryInfo)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("225f7195-f9c7-4d14-ab28-a83f7ff77e1f")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitRepository
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateRepository function
type UpdateRepositoryArgs struct {
	// (required) Specify a new repo name or a new default branch of the repository
	NewRepositoryInfo *GitRepository
	// (required) The name or ID of the repository.
	RepositoryId *uuid.UUID
	// (optional) Project ID or project name
	Project *string
}

// Update a thread in a pull request.
func (client *ClientImpl) UpdateThread(ctx context.Context, args UpdateThreadArgs) (*GitPullRequestCommentThread, error) {
	if args.CommentThread == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.CommentThread"}
	}
	routeValues := make(map[string]string)
	if args.Project != nil && *args.Project != "" {
		routeValues["project"] = *args.Project
	}
	if args.RepositoryId == nil || *args.RepositoryId == "" {
		return nil, &azuredevops.ArgumentNilOrEmptyError{ArgumentName: "args.RepositoryId"}
	}
	routeValues["repositoryId"] = *args.RepositoryId
	if args.PullRequestId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.PullRequestId"}
	}
	routeValues["pullRequestId"] = strconv.Itoa(*args.PullRequestId)
	if args.ThreadId == nil {
		return nil, &azuredevops.ArgumentNilError{ArgumentName: "args.ThreadId"}
	}
	routeValues["threadId"] = strconv.Itoa(*args.ThreadId)

	body, marshalErr := json.Marshal(*args.CommentThread)
	if marshalErr != nil {
		return nil, marshalErr
	}
	locationId, _ := uuid.Parse("ab6e2e5d-a0b7-4153-b64a-a4efe0d49449")
	resp, err := client.Client.Send(ctx, http.MethodPatch, locationId, "5.1", routeValues, nil, bytes.NewReader(body), "application/json", "application/json", nil)
	if err != nil {
		return nil, err
	}

	var responseValue GitPullRequestCommentThread
	err = client.Client.UnmarshalBody(resp, &responseValue)
	return &responseValue, err
}

// Arguments for the UpdateThread function
type UpdateThreadArgs struct {
	// (required) The thread content that should be updated.
	CommentThread *GitPullRequestCommentThread
	// (required) The repository ID of the pull request's target branch.
	RepositoryId *string
	// (required) ID of the pull request.
	PullRequestId *int
	// (required) ID of the thread to update.
	ThreadId *int
	// (optional) Project ID or project name
	Project *string
}
