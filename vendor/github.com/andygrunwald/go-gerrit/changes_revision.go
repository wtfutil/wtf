package gerrit

import (
	"fmt"
	"net/url"
)

// DiffInfo entity contains information about the diff of a file in a revision.
type DiffInfo struct {
	MetaA           DiffFileMetaInfo  `json:"meta_a,omitempty"`
	MetaB           DiffFileMetaInfo  `json:"meta_b,omitempty"`
	ChangeType      string            `json:"change_type"`
	IntralineStatus string            `json:"intraline_status,omitempty"`
	DiffHeader      []string          `json:"diff_header"`
	Content         []DiffContent     `json:"content"`
	WebLinks        []DiffWebLinkInfo `json:"web_links,omitempty"`
	Binary          bool              `json:"binary,omitempty"`
}

// RelatedChangesInfo entity contains information about related changes.
type RelatedChangesInfo struct {
	Changes []RelatedChangeAndCommitInfo `json:"changes"`
}

// FileInfo entity contains information about a file in a patch set.
type FileInfo struct {
	Status        string `json:"status,omitempty"`
	Binary        bool   `json:"binary,omitempty"`
	OldPath       string `json:"old_path,omitempty"`
	LinesInserted int    `json:"lines_inserted,omitempty"`
	LinesDeleted  int    `json:"lines_deleted,omitempty"`
	SizeDelta     int    `json:"size_delta"`
	Size          int    `json:"size"`
}

// ActionInfo entity describes a REST API call the client can make to manipulate a resource.
// These are frequently implemented by plugins and may be discovered at runtime.
type ActionInfo struct {
	Method  string `json:"method,omitempty"`
	Label   string `json:"label,omitempty"`
	Title   string `json:"title,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
}

// CommitInfo entity contains information about a commit.
type CommitInfo struct {
	Commit    string        `json:"commit,omitempty"`
	Parents   []CommitInfo  `json:"parents"`
	Author    GitPersonInfo `json:"author"`
	Committer GitPersonInfo `json:"committer"`
	Subject   string        `json:"subject"`
	Message   string        `json:"message"`
	WebLinks  []WebLinkInfo `json:"web_links,omitempty"`
}

// MergeableInfo entity contains information about the mergeability of a change.
type MergeableInfo struct {
	SubmitType    string   `json:"submit_type"`
	Mergeable     bool     `json:"mergeable"`
	MergeableInto []string `json:"mergeable_into,omitempty"`
}

// DiffOptions specifies the parameters for GetDiff call.
type DiffOptions struct {
	// If the intraline parameter is specified, intraline differences are included in the diff.
	Intraline bool `url:"intraline,omitempty"`

	// The base parameter can be specified to control the base patch set from which the diff
	// should be generated.
	Base string `url:"base,omitempty"`

	// The integer-valued request parameter parent can be specified to control the parent commit number
	// against which the diff should be generated. This is useful for supporting review of merge commits.
	// The value is the 1-based index of the parent’s position in the commit object.
	Parent int `url:"parent,omitempty"`

	// If the weblinks-only parameter is specified, only the diff web links are returned.
	WeblinksOnly bool `url:"weblinks-only,omitempty"`

	// The ignore-whitespace parameter can be specified to control how whitespace differences are reported in the result. Valid values are NONE, TRAILING, CHANGED or ALL.
	IgnoreWhitespace string `url:"ignore-whitespace,omitempty"`

	// The context parameter can be specified to control the number of lines of surrounding context in the diff.
	// Valid values are ALL or number of lines.
	Context string `url:"context,omitempty"`
}

// CommitOptions specifies the parameters for GetCommit call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-commit
type CommitOptions struct {
	// Adding query parameter links (for example /changes/.../commit?links) returns a CommitInfo with the additional field web_links.
	Weblinks bool `url:"links,omitempty"`
}

// MergableOptions specifies the parameters for GetMergable call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-mergeable
type MergableOptions struct {
	// If the other-branches parameter is specified, the mergeability will also be checked for all other branches.
	OtherBranches bool `url:"other-branches,omitempty"`
}

// FilesOptions specifies the parameters for ListFiles and ListFilesReviewed calls.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-files
type FilesOptions struct {
	// The request parameter q changes the response to return a list of all files (modified or unmodified)
	// that contain that substring in the path name. This is useful to implement suggestion services
	// finding a file by partial name.
	Q string `url:"q,omitempty"`

	// The base parameter can be specified to control the base patch set from which the list of files
	// should be generated.
	//
	// Note: This option is undocumented.
	Base string `url:"base,omitempty"`

	// The integer-valued request parameter parent changes the response to return a list of the files
	// which are different in this commit compared to the given parent commit. This is useful for
	// supporting review of merge commits. The value is the 1-based index of the parent’s position
	// in the commit object.
	Parent int `url:"parent,omitempty"`
}

// PatchOptions specifies the parameters for GetPatch call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-patch
type PatchOptions struct {
	// Adding query parameter zip (for example /changes/.../patch?zip) returns the patch as a single file inside of a ZIP archive.
	// Clients can expand the ZIP to obtain the plain text patch, avoiding the need for a base64 decoding step.
	// This option implies download.
	Zip bool `url:"zip,omitempty"`

	// Query parameter download (e.g. /changes/.../patch?download) will suggest the browser save the patch as commitsha1.diff.base64, for later processing by command line tools.
	Download bool `url:"download,omitempty"`

	// If the path parameter is set, the returned content is a diff of the single file that the path refers to.
	Path string `url:"path,omitempty"`
}

// GetDiff gets the diff of a file from a certain revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-diff
func (s *ChangesService) GetDiff(changeID, revisionID, fileID string, opt *DiffOptions) (*DiffInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/%s/diff", changeID, revisionID, url.PathEscape(fileID))

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(DiffInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetRelatedChanges retrieves related changes of a revision.
// Related changes are changes that either depend on, or are dependencies of the revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-related-changes
func (s *ChangesService) GetRelatedChanges(changeID, revisionID string) (*RelatedChangesInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/related", changeID, revisionID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(RelatedChangesInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetDraft retrieves a draft comment of a revision that belongs to the calling user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-draft
func (s *ChangesService) GetDraft(changeID, revisionID, draftID string) (*CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/drafts/%s", changeID, revisionID, draftID)
	return s.getCommentInfoResponse(u)
}

// GetComment retrieves a published comment of a revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-comment
func (s *ChangesService) GetComment(changeID, revisionID, commentID string) (*CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s//comments/%s", changeID, revisionID, commentID)
	return s.getCommentInfoResponse(u)
}

// GetSubmitType gets the method the server will use to submit (merge) the change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-submit-type
func (s *ChangesService) GetSubmitType(changeID, revisionID string) (string, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/submit_type", changeID, revisionID)
	return getStringResponseWithoutOptions(s.client, u)
}

// GetRevisionActions retrieves revision actions of the revision of a change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-revision-actions
func (s *ChangesService) GetRevisionActions(changeID, revisionID string) (*map[string]ActionInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/actions", changeID, revisionID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]ActionInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetCommit retrieves a parsed commit of a revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-commit
func (s *ChangesService) GetCommit(changeID, revisionID string, opt *CommitOptions) (*CommitInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/commit", changeID, revisionID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(CommitInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetReview retrieves a review of a revision.
//
// As response a ChangeInfo entity with detailed labels and detailed accounts is returned that describes the review of the revision.
// The revision for which the review is retrieved is contained in the revisions field.
// In addition the current_revision field is set if the revision for which the review is retrieved is the current revision of the change.
// Please note that the returned labels are always for the current patch set.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-review
func (s *ChangesService) GetReview(changeID, revisionID string) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/review", changeID, revisionID)
	return s.getChangeInfoResponse(u, nil)
}

// GetMergeable gets the method the server will use to submit (merge) the change and an indicator if the change is currently mergeable.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-mergeable
func (s *ChangesService) GetMergeable(changeID, revisionID string, opt *MergableOptions) (*MergeableInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/mergeable", changeID, revisionID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(MergeableInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListRevisionDrafts lists the draft comments of a revision that belong to the calling user.
// Returns a map of file paths to lists of CommentInfo entries.
// The entries in the map are sorted by file path.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-drafts
func (s *ChangesService) ListRevisionDrafts(changeID, revisionID string) (*map[string][]CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/drafts/", changeID, revisionID)
	return s.getCommentInfoMapSliceResponse(u)
}

// ListRevisionComments lists the published comments of a revision.
// As result a map is returned that maps the file path to a list of CommentInfo entries.
// The entries in the map are sorted by file path and only include file (or inline) comments.
// Use the Get Change Detail endpoint to retrieve the general change message (or comment).
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-comments
func (s *ChangesService) ListRevisionComments(changeID, revisionID string) (*map[string][]CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/comments/", changeID, revisionID)
	return s.getCommentInfoMapSliceResponse(u)
}

// ListFiles lists the files that were modified, added or deleted in a revision.
// As result a map is returned that maps the file path to a list of FileInfo entries.
// The entries in the map are sorted by file path.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-files
func (s *ChangesService) ListFiles(changeID, revisionID string, opt *FilesOptions) (map[string]FileInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/", changeID, revisionID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var v map[string]FileInfo
	resp, err := s.client.Do(req, &v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListFilesReviewed lists the files that were modified, added or deleted in a revision.
// Unlike ListFiles, the response of ListFilesReviewed is a list of the paths the caller
// has marked as reviewed. Clients that also need the FileInfo should make two requests.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-files
func (s *ChangesService) ListFilesReviewed(changeID, revisionID string, opt *FilesOptions) ([]string, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/", changeID, revisionID)

	o := struct {
		// The request parameter reviewed changes the response to return a list of the paths the caller has marked as reviewed.
		Reviewed bool `url:"reviewed,omitempty"`

		FilesOptions
	}{
		Reviewed: true,
	}
	if opt != nil {
		o.FilesOptions = *opt
	}
	u, err := addOptions(u, o)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var v []string
	resp, err := s.client.Do(req, &v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetReview sets a review on a revision.
// The review must be provided in the request body as a ReviewInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#set-review
func (s *ChangesService) SetReview(changeID, revisionID string, input *ReviewInput) (*ReviewResult, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/review", changeID, revisionID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(ReviewResult)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// PublishDraftRevision publishes a draft revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#publish-draft-revision
func (s *ChangesService) PublishDraftRevision(changeID, revisionID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/publish", changeID, revisionID)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteDraftRevision deletes a draft revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-draft-revision
func (s *ChangesService) DeleteDraftRevision(changeID, revisionID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s", changeID, revisionID)
	return s.client.DeleteRequest(u, nil)
}

// GetPatch gets the formatted patch for one revision.
//
// The formatted patch is returned as text encoded inside base64.
// Adding query parameter zip (for example /changes/.../patch?zip) returns the patch as a single file inside of a ZIP archive.
// Clients can expand the ZIP to obtain the plain text patch, avoiding the need for a base64 decoding step.
// This option implies download.
//
// Query parameter download (e.g. /changes/.../patch?download) will suggest the browser save the patch as commitsha1.diff.base64, for later processing by command line tools.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-patch
func (s *ChangesService) GetPatch(changeID, revisionID string, opt *PatchOptions) (*string, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/patch", changeID, revisionID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// TestSubmitType tests the submit_type Prolog rule in the project, or the one given.
//
// Request body may be either the Prolog code as text/plain or a RuleInput object.
// The query parameter filters may be set to SKIP to bypass parent project filters while testing a project-specific rule.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#test-submit-type
func (s *ChangesService) TestSubmitType(changeID, revisionID string, input *RuleInput) (*string, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/test.submit_type", changeID, revisionID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// TestSubmitRule tests the submit_rule Prolog rule in the project, or the one given.
//
// Request body may be either the Prolog code as text/plain or a RuleInput object.
// The query parameter filters may be set to SKIP to bypass parent project filters while testing a project-specific rule.
//
// The response is a list of SubmitRecord entries describing the permutations that satisfy the tested submit rule.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#test-submit-rule
func (s *ChangesService) TestSubmitRule(changeID, revisionID string, input *RuleInput) (*[]SubmitRecord, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/test.submit_rule", changeID, revisionID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new([]SubmitRecord)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CreateDraft creates a draft comment on a revision.
// The new draft comment must be provided in the request body inside a CommentInput entity.
//
// As response a CommentInfo entity is returned that describes the draft comment.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#create-draft
func (s *ChangesService) CreateDraft(changeID, revisionID string, input *CommentInput) (*CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/drafts", changeID, revisionID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(CommentInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// UpdateDraft updates a draft comment on a revision.
// The new draft comment must be provided in the request body inside a CommentInput entity.
//
// As response a CommentInfo entity is returned that describes the draft comment.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#update-draft
func (s *ChangesService) UpdateDraft(changeID, revisionID, draftID string, input *CommentInput) (*CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/drafts/%s", changeID, revisionID, draftID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(CommentInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// DeleteDraft deletes a draft comment from a revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-draft
func (s *ChangesService) DeleteDraft(changeID, revisionID, draftID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/drafts/%s", changeID, revisionID, draftID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteReviewed deletes the reviewed flag of the calling user from a file of a revision.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-reviewed
func (s *ChangesService) DeleteReviewed(changeID, revisionID, fileID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/%s/reviewed", changeID, revisionID, url.PathEscape(fileID))
	return s.client.DeleteRequest(u, nil)
}

// GetContent gets the content of a file from a certain revision.
// The content is returned as base64 encoded string.
// The HTTP response Content-Type is always text/plain, reflecting the base64 wrapping.
// A Gerrit-specific X-FYI-Content-Type header is returned describing the server detected content type of the file.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-content
func (s *ChangesService) GetContent(changeID, revisionID, fileID string) (*string, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/%s/content", changeID, revisionID, url.PathEscape(fileID))

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(string)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetContentType gets the content type of a file from a certain revision.
// This is nearly the same as GetContent.
// But if only the content type is required, callers should use HEAD to avoid downloading the encoded file contents.
//
// For further documentation see GetContent.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-content
func (s *ChangesService) GetContentType(changeID, revisionID, fileID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/%s/content", changeID, revisionID, url.PathEscape(fileID))

	req, err := s.client.NewRequest("HEAD", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// SetReviewed marks a file of a revision as reviewed by the calling user.
//
// If the file was already marked as reviewed by the calling user the response is “200 OK”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#set-reviewed
func (s *ChangesService) SetReviewed(changeID, revisionID, fileID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/files/%s/reviewed", changeID, revisionID, url.PathEscape(fileID))

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// CherryPickRevision cherry picks a revision to a destination branch.
// The commit message and destination branch must be provided in the request body inside a CherryPickInput entity.
//
// As response a ChangeInfo entity is returned that describes the resulting cherry picked change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#cherry-pick
func (s *ChangesService) CherryPickRevision(changeID, revisionID string, input *CherryPickInput) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/revisions/%s/cherrypick", changeID, revisionID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(ChangeInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

/*
TODO: Missing Revision Endpoints
	Rebase Revision
	Submit Revision
	DownloadContent (https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-safe-content)
*/
