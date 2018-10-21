package gerrit

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ChangesService contains Change related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html
type ChangesService struct {
	client *Client
}

// WebLinkInfo entity describes a link to an external site.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#web-link-info
type WebLinkInfo struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	ImageURL string `json:"image_url"`
}

// GitPersonInfo entity contains information about the author/committer of a commit.
type GitPersonInfo struct {
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Date  Timestamp `json:"date"`
	TZ    int       `json:"tz"`
}

// NotifyInfo entity contains detailed information about who should be
// notified about an update
type NotifyInfo struct {
	Accounts []AccountInfo `json:"accounts"`
}

// AbandonInput entity contains information for abandoning a change.
type AbandonInput struct {
	Message       string       `json:"message,omitempty"`
	Notify        string       `json:"notify"`
	NotifyDetails []NotifyInfo `json:"notify_details"`
}

// ApprovalInfo entity contains information about an approval from a user for a label on a change.
type ApprovalInfo struct {
	AccountInfo
	Value int    `json:"value,omitempty"`
	Date  string `json:"date,omitempty"`
}

// ChangeEditInput entity contains information for restoring a path within change edit.
type ChangeEditInput struct {
	RestorePath string `json:"restore_path,omitempty"`
	OldPath     string `json:"old_path,omitempty"`
	NewPath     string `json:"new_path,omitempty"`
}

// ChangeEditMessageInput entity contains information for changing the commit message within a change edit.
type ChangeEditMessageInput struct {
	Message string `json:"message"`
}

// ChangeMessageInfo entity contains information about a message attached to a change.
type ChangeMessageInfo struct {
	ID             string      `json:"id"`
	Author         AccountInfo `json:"author,omitempty"`
	Date           Timestamp   `json:"date"`
	Message        string      `json:"message"`
	Tag            string      `json:"tag,omitempty"`
	RevisionNumber int         `json:"_revision_number,omitempty"`
}

// CherryPickInput entity contains information for cherry-picking a change to a new branch.
type CherryPickInput struct {
	Message     string `json:"message"`
	Destination string `json:"destination"`
}

// CommentRange entity describes the range of an inline comment.
type CommentRange struct {
	StartLine      int `json:"start_line"`
	StartCharacter int `json:"start_character"`
	EndLine        int `json:"end_line"`
	EndCharacter   int `json:"end_character"`
}

// DiffFileMetaInfo entity contains meta information about a file diff
type DiffFileMetaInfo struct {
	Name        string        `json:"name"`
	ContentType string        `json:"content_type"`
	Lines       int           `json:"lines"`
	WebLinks    []WebLinkInfo `json:"web_links,omitempty"`
}

// DiffWebLinkInfo entity describes a link on a diff screen to an external site.
type DiffWebLinkInfo struct {
	Name                     string `json:"name"`
	URL                      string `json:"url"`
	ImageURL                 string `json:"image_url"`
	ShowOnSideBySideDiffView bool   `json:"show_on_side_by_side_diff_view"`
	ShowOnUnifiedDiffView    bool   `json:"show_on_unified_diff_view"`
}

// FetchInfo entity contains information about how to fetch a patch set via a certain protocol.
type FetchInfo struct {
	URL      string            `json:"url"`
	Ref      string            `json:"ref"`
	Commands map[string]string `json:"commands,omitempty"`
}

// FixInput entity contains options for fixing commits using the fix change endpoint.
type FixInput struct {
	DeletePatchSetIfCommitMissing bool   `json:"delete_patch_set_if_commit_missing"`
	ExpectMergedAs                string `json:"expect_merged_as"`
}

// GroupBaseInfo entity contains base information about the group.
type GroupBaseInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IncludedInInfo entity contains information about the branches a change was merged into and tags it was tagged with.
type IncludedInInfo struct {
	Branches []string          `json:"branches"`
	Tags     []string          `json:"tags"`
	External map[string]string `json:"external,omitempty"`
}

// ProblemInfo entity contains a description of a potential consistency problem with a change.
// These are not related to the code review process, but rather indicate some inconsistency in Gerrit’s database or repository metadata related to the enclosing change.
type ProblemInfo struct {
	Message string `json:"message"`
	Status  string `json:"status,omitempty"`
	Outcome string `json:"outcome,omitempty"`
}

// RebaseInput entity contains information for changing parent when rebasing.
type RebaseInput struct {
	Base string `json:"base,omitempty"`
}

// RestoreInput entity contains information for restoring a change.
type RestoreInput struct {
	Message string `json:"message,omitempty"`
}

// RevertInput entity contains information for reverting a change.
type RevertInput struct {
	Message string `json:"message,omitempty"`
}

// ReviewInfo entity contains information about a review.
type ReviewInfo struct {
	Labels map[string]int `json:"labels"`
}

// ReviewResult entity contains information regarding the updates that were
// made to a review.
type ReviewResult struct {
	ReviewInfo
	Reviewers map[string]AddReviewerResult `json:"reviewers,omitempty"`
	Ready     bool                         `json:"ready,omitempty"`
}

// TopicInput entity contains information for setting a topic.
type TopicInput struct {
	Topic string `json:"topic,omitempty"`
}

// SubmitRecord entity describes results from a submit_rule.
type SubmitRecord struct {
	Status       string                            `json:"status"`
	Ok           map[string]map[string]AccountInfo `json:"ok,omitempty"`
	Reject       map[string]map[string]AccountInfo `json:"reject,omitempty"`
	Need         map[string]interface{}            `json:"need,omitempty"`
	May          map[string]map[string]AccountInfo `json:"may,omitempty"`
	Impossible   map[string]interface{}            `json:"impossible,omitempty"`
	ErrorMessage string                            `json:"error_message,omitempty"`
}

// SubmitInput entity contains information for submitting a change.
type SubmitInput struct {
	WaitForMerge bool `json:"wait_for_merge"`
}

// SubmitInfo entity contains information about the change status after submitting.
type SubmitInfo struct {
	Status     string `json:"status"`
	OnBehalfOf string `json:"on_behalf_of,omitempty"`
}

// RuleInput entity contains information to test a Prolog rule.
type RuleInput struct {
	Rule    string `json:"rule"`
	Filters string `json:"filters,omitempty"`
}

// ReviewerInput entity contains information for adding a reviewer to a change.
type ReviewerInput struct {
	Reviewer  string `json:"reviewer"`
	Confirmed bool   `json:"confirmed,omitempty"`
}

// ReviewInput entity contains information for adding a review to a revision.
type ReviewInput struct {
	Message               string                         `json:"message,omitempty"`
	Tag                   string                         `json:"tag,omitempty"`
	Labels                map[string]string              `json:"labels,omitempty"`
	Comments              map[string][]CommentInput      `json:"comments,omitempty"`
	RobotComments         map[string][]RobotCommentInput `json:"robot_comments,omitempty"`
	StrictLabels          bool                           `json:"strict_labels,omitempty"`
	Drafts                string                         `json:"drafts,omitempty"`
	Notify                string                         `json:"notify,omitempty"`
	OmitDuplicateComments bool                           `json:"omit_duplicate_comments,omitempty"`
	OnBehalfOf            string                         `json:"on_behalf_of,omitempty"`
}

// RelatedChangeAndCommitInfo entity contains information about a related change and commit.
type RelatedChangeAndCommitInfo struct {
	ChangeID              string     `json:"change_id,omitempty"`
	Commit                CommitInfo `json:"commit"`
	ChangeNumber          int        `json:"_change_number,omitempty"`
	RevisionNumber        int        `json:"_revision_number,omitempty"`
	CurrentRevisionNumber int        `json:"_current_revision_number,omitempty"`
	Status                string     `json:"status,omitempty"`
}

// DiffContent entity contains information about the content differences in a file.
type DiffContent struct {
	A      []string          `json:"a,omitempty"`
	B      []string          `json:"b,omitempty"`
	AB     []string          `json:"ab,omitempty"`
	EditA  DiffIntralineInfo `json:"edit_a,omitempty"`
	EditB  DiffIntralineInfo `json:"edit_b,omitempty"`
	Skip   int               `json:"skip,omitempty"`
	Common bool              `json:"common,omitempty"`
}

// CommentInput entity contains information for creating an inline comment.
type CommentInput struct {
	ID        string        `json:"id,omitempty"`
	Path      string        `json:"path,omitempty"`
	Side      string        `json:"side,omitempty"`
	Line      int           `json:"line,omitempty"`
	Range     *CommentRange `json:"range,omitempty"`
	InReplyTo string        `json:"in_reply_to,omitempty"`
	Updated   *Timestamp    `json:"updated,omitempty"`
	Message   string        `json:"message,omitempty"`
}

// RobotCommentInput entity contains information for creating an inline robot comment.
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#robot-comment-input
type RobotCommentInput struct {
	CommentInput

	// The ID of the robot that generated this comment.
	RobotID string `json:"robot_id"`
	// An ID of the run of the robot.
	RobotRunID string `json:"robot_run_id"`
	// URL to more information.
	URL string `json:"url,omitempty"`
	// Robot specific properties as map that maps arbitrary keys to values.
	Properties *map[string]*string `json:"properties,omitempty"`
	// Suggested fixes for this robot comment as a list of FixSuggestionInfo
	// entities.
	FixSuggestions *FixSuggestionInfo `json:"fix_suggestions,omitempty"`
}

// RobotCommentInfo entity contains information about a robot inline comment
// RobotCommentInfo has the same fields as CommentInfo. In addition RobotCommentInfo has the following fields:
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#robot-comment-info
type RobotCommentInfo struct {
	CommentInfo

	// The ID of the robot that generated this comment.
	RobotID string `json:"robot_id"`
	// An ID of the run of the robot.
	RobotRunID string `json:"robot_run_id"`
	// URL to more information.
	URL string `json:"url,omitempty"`
	// Robot specific properties as map that maps arbitrary keys to values.
	Properties map[string]string `json:"properties,omitempty"`
	// Suggested fixes for this robot comment as a list of FixSuggestionInfo
	// entities.
	FixSuggestions *FixSuggestionInfo `json:"fix_suggestions,omitempty"`
}

// FixSuggestionInfo entity represents a suggested fix.
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#fix-suggestion-info
type FixSuggestionInfo struct {
	// The UUID of the suggested fix. It will be generated automatically and
	// hence will be ignored if it’s set for input objects.
	FixID string `json:"fix_id"`
	// A description of the suggested fix.
	Description string `json:"description"`
	// A list of FixReplacementInfo entities indicating how the content of one or
	// several files should be modified. Within a file, they should refer to
	// non-overlapping regions.
	Replacements FixReplacementInfo `json:"replacements"`
}

// FixReplacementInfo entity describes how the content of a file should be replaced by another content.
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#fix-replacement-info
type FixReplacementInfo struct {
	// The path of the file which should be modified. Any file in the repository may be modified.
	Path string `json:"path"`

	// A CommentRange indicating which content of the file should be replaced.
	// Lines in the file are assumed to be separated by the line feed character,
	// the carriage return character, the carriage return followed by the line
	// feed character, or one of the other Unicode linebreak sequences supported
	// by Java.
	Range CommentRange `json:"range"`

	// The content which should be used instead of the current one.
	Replacement string `json:"replacement,omitempty"`
}

// DiffIntralineInfo entity contains information about intraline edits in a file.
//
// The information consists of a list of <skip length, mark length> pairs,
// where the skip length is the number of characters between the end of
// the previous edit and the start of this edit, and the mark length is the
// number of edited characters following the skip. The start of the edits
// is from the beginning of the related diff content lines.
//
// Note that the implied newline character at the end of each line
// is included in the length calculation, and thus it is possible for
// the edits to span newlines.
type DiffIntralineInfo [][2]int

// ChangeInfo entity contains information about a change.
type ChangeInfo struct {
	ID                 string                   `json:"id"`
	URL                string                   `json:"url,omitempty"`
	Project            string                   `json:"project"`
	Branch             string                   `json:"branch"`
	Topic              string                   `json:"topic,omitempty"`
	ChangeID           string                   `json:"change_id"`
	Subject            string                   `json:"subject"`
	Status             string                   `json:"status"`
	Created            Timestamp                `json:"created"`
	Updated            Timestamp                `json:"updated"`
	Submitted          *Timestamp               `json:"submitted,omitempty"`
	Starred            bool                     `json:"starred,omitempty"`
	Reviewed           bool                     `json:"reviewed,omitempty"`
	Mergeable          bool                     `json:"mergeable,omitempty"`
	Insertions         int                      `json:"insertions"`
	Deletions          int                      `json:"deletions"`
	Number             int                      `json:"_number"`
	Owner              AccountInfo              `json:"owner"`
	Actions            map[string]ActionInfo    `json:"actions,omitempty"`
	Labels             map[string]LabelInfo     `json:"labels,omitempty"`
	PermittedLabels    map[string][]string      `json:"permitted_labels,omitempty"`
	RemovableReviewers []AccountInfo            `json:"removable_reviewers,omitempty"`
	Reviewers          map[string][]AccountInfo `json:"reviewers,omitempty"`
	Messages           []ChangeMessageInfo      `json:"messages,omitempty"`
	CurrentRevision    string                   `json:"current_revision,omitempty"`
	Revisions          map[string]RevisionInfo  `json:"revisions,omitempty"`
	MoreChanges        bool                     `json:"_more_changes,omitempty"`
	Problems           []ProblemInfo            `json:"problems,omitempty"`
	BaseChange         string                   `json:"base_change,omitempty"`
}

// LabelInfo entity contains information about a label on a change, always corresponding to the current patch set.
type LabelInfo struct {
	Optional bool `json:"optional,omitempty"`

	// Fields set by LABELS
	Approved     AccountInfo `json:"approved,omitempty"`
	Rejected     AccountInfo `json:"rejected,omitempty"`
	Recommended  AccountInfo `json:"recommended,omitempty"`
	Disliked     AccountInfo `json:"disliked,omitempty"`
	Blocking     bool        `json:"blocking,omitempty"`
	Value        int         `json:"value,omitempty"`
	DefaultValue int         `json:"default_value,omitempty"`

	// Fields set by DETAILED_LABELS
	All    []ApprovalInfo    `json:"all,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}

// RevisionInfo entity contains information about a patch set.
type RevisionInfo struct {
	Draft             bool                  `json:"draft,omitempty"`
	Number            int                   `json:"_number"`
	Created           Timestamp             `json:"created"`
	Uploader          AccountInfo           `json:"uploader"`
	Ref               string                `json:"ref"`
	Fetch             map[string]FetchInfo  `json:"fetch"`
	Commit            CommitInfo            `json:"commit,omitempty"`
	Files             map[string]FileInfo   `json:"files,omitempty"`
	Actions           map[string]ActionInfo `json:"actions,omitempty"`
	Reviewed          bool                  `json:"reviewed,omitempty"`
	MessageWithFooter string                `json:"messageWithFooter,omitempty"`
}

// CommentInfo entity contains information about an inline comment.
type CommentInfo struct {
	PatchSet  int           `json:"patch_set,omitempty"`
	ID        string        `json:"id"`
	Path      string        `json:"path,omitempty"`
	Side      string        `json:"side,omitempty"`
	Line      int           `json:"line,omitempty"`
	Range     *CommentRange `json:"range,omitempty"`
	InReplyTo string        `json:"in_reply_to,omitempty"`
	Message   string        `json:"message,omitempty"`
	Updated   *Timestamp    `json:"updated"`
	Author    AccountInfo   `json:"author,omitempty"`
}

// QueryOptions specifies global parameters to query changes / reviewers.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
type QueryOptions struct {
	// Query parameter
	// Clients are allowed to specify more than one query by setting the q parameter multiple times.
	// In this case the result is an array of arrays, one per query in the same order the queries were given in.
	//
	// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/user-search.html#_search_operators
	Query []string `url:"q,omitempty"`

	// The n parameter can be used to limit the returned results.
	// If the n query parameter is supplied and additional changes exist that match the query beyond the end, the last change object has a _more_changes: true JSON field set.
	Limit int `url:"n,omitempty"`
}

// QueryChangeOptions specifies the parameters to the ChangesService.QueryChanges.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
type QueryChangeOptions struct {
	QueryOptions

	// The S or start query parameter can be supplied to skip a number of changes from the list.
	Skip  int `url:"S,omitempty"`
	Start int `url:"start,omitempty"`

	ChangeOptions
}

// ChangeOptions specifies the parameters for Query changes.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
type ChangeOptions struct {
	// Additional fields can be obtained by adding o parameters, each option requires more database lookups and slows down the query response time to the client so they are generally disabled by default.
	//
	// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
	AdditionalFields []string `url:"o,omitempty"`
}

// QueryChanges lists changes visible to the caller.
// The query string must be provided by the q parameter.
// The n parameter can be used to limit the returned results.
//
// The change output is sorted by the last update time, most recently updated to oldest updated.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-changes
func (s *ChangesService) QueryChanges(opt *QueryChangeOptions) (*[]ChangeInfo, *Response, error) {
	u := "changes/"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ChangeInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetChange retrieves a change.
// Additional fields can be obtained by adding o parameters, each option requires more database lookups and slows down the query response time to the client so they are generally disabled by default.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-change
func (s *ChangesService) GetChange(changeID string, opt *ChangeOptions) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s", changeID)
	return s.getChangeInfoResponse(u, opt)
}

// GetChangeDetail retrieves a change with labels, detailed labels, detailed accounts, and messages.
// Additional fields can be obtained by adding o parameters, each option requires more database lookups and slows down the query response time to the client so they are generally disabled by default.
//
// This response will contain all votes for each label and include one combined vote.
// The combined label vote is calculated in the following order (from highest to lowest): REJECTED > APPROVED > DISLIKED > RECOMMENDED.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-change-detail
func (s *ChangesService) GetChangeDetail(changeID string, opt *ChangeOptions) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/detail", changeID)
	return s.getChangeInfoResponse(u, opt)
}

// getChangeInfoResponse retrieved a single ChangeInfo Response for a GET request
func (s *ChangesService) getChangeInfoResponse(u string, opt *ChangeOptions) (*ChangeInfo, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
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

// GetTopic retrieves the topic of a change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-topic
func (s *ChangesService) GetTopic(changeID string) (string, *Response, error) {
	u := fmt.Sprintf("changes/%s/topic", changeID)
	return getStringResponseWithoutOptions(s.client, u)
}

// ChangesSubmittedTogether returns a list of all changes which are submitted when {submit} is called for this change, including the current change itself.
// An empty list is returned if this change will be submitted by itself (no other changes).
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#submitted_together
func (s *ChangesService) ChangesSubmittedTogether(changeID string) (*[]ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/submitted_together", changeID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]ChangeInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetIncludedIn retrieves the branches and tags in which a change is included.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-included-in
func (s *ChangesService) GetIncludedIn(changeID string) (*IncludedInInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/in", changeID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(IncludedInInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListChangeComments lists the published comments of all revisions of the change.
// The entries in the map are sorted by file path, and the comments for each path are sorted by patch set number.
// Each comment has the patch_set and author fields set.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-change-comments
func (s *ChangesService) ListChangeComments(changeID string) (*map[string][]CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/comments", changeID)
	return s.getCommentInfoMapResponse(u)
}

// ListChangeDrafts lLists the draft comments of all revisions of the change that belong to the calling user.
// The entries in the map are sorted by file path, and the comments for each path are sorted by patch set number.
// Each comment has the patch_set field set, and no author.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#list-change-drafts
func (s *ChangesService) ListChangeDrafts(changeID string) (*map[string][]CommentInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/drafts", changeID)
	return s.getCommentInfoMapResponse(u)
}

// getCommentInfoMapResponse retrieved a map of CommentInfo Response for a GET request
func (s *ChangesService) getCommentInfoMapResponse(u string) (*map[string][]CommentInfo, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string][]CommentInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CheckChange performs consistency checks on the change, and returns a ChangeInfo entity with the problems field set to a list of ProblemInfo entities.
// Depending on the type of problem, some fields not marked optional may be missing from the result.
// At least id, project, branch, and _number will be present.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#check-change
func (s *ChangesService) CheckChange(changeID string) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/check", changeID)
	return s.getChangeInfoResponse(u, nil)
}

// getCommentInfoResponse retrieved a CommentInfo Response for a GET request
func (s *ChangesService) getCommentInfoResponse(u string) (*CommentInfo, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
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

// getCommentInfoMapSliceResponse retrieved a map with a slice of CommentInfo Response for a GET request
func (s *ChangesService) getCommentInfoMapSliceResponse(u string) (*map[string][]CommentInfo, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string][]CommentInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CreateChange creates a new change.
// The change info ChangeInfo entity must be provided in the request body.
// Only the following attributes are honored: project, branch, subject, status and topic.
// The first three attributes are mandatory.
// Valid values for status are: DRAFT and NEW.
//
// As response a ChangeInfo entity is returned that describes the resulting change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#create-change
func (s *ChangesService) CreateChange(input *ChangeInfo) (*ChangeInfo, *Response, error) {
	u := "changes/"

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

// SetTopic sets the topic of a change.
// The new topic must be provided in the request body inside a TopicInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#set-topic
func (s *ChangesService) SetTopic(changeID string, input *TopicInput) (*string, *Response, error) {
	u := fmt.Sprintf("changes/%s/topic", changeID)

	req, err := s.client.NewRequest("PUT", u, input)
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

// DeleteTopic deletes the topic of a change.
// The request body does not need to include a TopicInput entity if no review comment is added.
//
// Please note that some proxies prohibit request bodies for DELETE requests.
// In this case, if you want to specify a commit message, use PUT to delete the topic.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-topic
func (s *ChangesService) DeleteTopic(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/topic", changeID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteDraftChange deletes a draft change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-draft-change
func (s *ChangesService) DeleteDraftChange(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s", changeID)
	return s.client.DeleteRequest(u, nil)
}

// PublishDraftChange publishes a draft change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#publish-draft-change
func (s *ChangesService) PublishDraftChange(changeID, notify string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/publish", changeID)

	req, err := s.client.NewRequest("POST", u, map[string]string{
		"notify": notify,
	})
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// IndexChange adds or updates the change in the secondary index.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#index-change
func (s *ChangesService) IndexChange(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/index", changeID)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// FixChange performs consistency checks on the change as with GET /check, and additionally fixes any problems that can be fixed automatically.
// The returned field values reflect any fixes.
//
// Some fixes have options controlling their behavior, which can be set in the FixInput entity body.
// Only the change owner, a project owner, or an administrator may fix changes.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#fix-change
func (s *ChangesService) FixChange(changeID string, input *FixInput) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/check", changeID)

	req, err := s.client.NewRequest("PUT", u, input)
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

// change is an internal function to consolidate code used by SubmitChange,
// AbandonChange and other similar functions.
func (s *ChangesService) change(tail string, changeID string, input interface{}) (*ChangeInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/%s", changeID, tail)
	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(ChangeInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	if resp.StatusCode == http.StatusConflict {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return v, resp, err
		}
		return v, resp, errors.New(string(body[:]))
	}
	return v, resp, nil
}

// SubmitChange submits a change.
//
// The request body only needs to include a SubmitInput entity if submitting on behalf of another user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#submit-change
func (s *ChangesService) SubmitChange(changeID string, input *SubmitInput) (*ChangeInfo, *Response, error) {
	return s.change("submit", changeID, input)
}

// AbandonChange abandons a change.
//
// The request body does not need to include a AbandonInput entity if no review
// comment is added.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#abandon-change
func (s *ChangesService) AbandonChange(changeID string, input *AbandonInput) (*ChangeInfo, *Response, error) {
	return s.change("abandon", changeID, input)
}

// RebaseChange rebases a change.
//
// Optionally, the parent revision can be changed to another patch set through
// the RebaseInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#rebase-change
func (s *ChangesService) RebaseChange(changeID string, input *RebaseInput) (*ChangeInfo, *Response, error) {
	return s.change("rebase", changeID, input)
}

// RestoreChange restores a change.
//
// The request body does not need to include a RestoreInput entity if no review
// comment is added.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#restore-change
func (s *ChangesService) RestoreChange(changeID string, input *RestoreInput) (*ChangeInfo, *Response, error) {
	return s.change("restore", changeID, input)
}

// RevertChange reverts a change.
//
// The request body does not need to include a RevertInput entity if no
// review comment is added.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#revert-change
func (s *ChangesService) RevertChange(changeID string, input *RevertInput) (*ChangeInfo, *Response, error) {
	return s.change("revert", changeID, input)
}
