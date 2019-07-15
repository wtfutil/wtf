package gerrit

import (
	"fmt"
	"net/url"
)

// EditInfo entity contains information about a change edit.
type EditInfo struct {
	Commit       CommitInfo           `json:"commit"`
	BaseRevision string               `json:"baseRevision"`
	Fetch        map[string]FetchInfo `json:"fetch"`
	Files        map[string]FileInfo  `json:"files,omitempty"`
}

// EditFileInfo entity contains additional information of a file within a change edit.
type EditFileInfo struct {
	WebLinks []WebLinkInfo `json:"web_links,omitempty"`
}

// ChangeEditDetailOptions specifies the parameters to the ChangesService.GetChangeEditDetails.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-detail
type ChangeEditDetailOptions struct {
	// When request parameter list is provided the response also includes the file list.
	List bool `url:"list,omitempty"`
	// When base request parameter is provided the file list is computed against this base revision.
	Base bool `url:"base,omitempty"`
	// When request parameter download-commands is provided fetch info map is also included.
	DownloadCommands bool `url:"download-commands,omitempty"`
}

// GetChangeEditDetails retrieves a change edit details.
// As response an EditInfo entity is returned that describes the change edit, or “204 No Content” when change edit doesn’t exist for this change.
// Change edits are stored on special branches and there can be max one edit per user per change.
// Edits aren’t tracked in the database.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-detail
func (s *ChangesService) GetChangeEditDetails(changeID string, opt *ChangeEditDetailOptions) (*EditInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/edit", changeID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(EditInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// RetrieveMetaDataOfAFileFromChangeEdit retrieves meta data of a file from a change edit.
// Currently only web links are returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-meta-data
func (s *ChangesService) RetrieveMetaDataOfAFileFromChangeEdit(changeID, filePath string) (*EditFileInfo, *Response, error) {
	u := fmt.Sprintf("changes/%s/edit/%s/meta", changeID, filePath)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(EditFileInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// RetrieveCommitMessageFromChangeEdit retrieves commit message from change edit.
// The commit message is returned as base64 encoded string.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-message
func (s *ChangesService) RetrieveCommitMessageFromChangeEdit(changeID string) (string, *Response, error) {
	u := fmt.Sprintf("changes/%s/edit:message", changeID)
	return getStringResponseWithoutOptions(s.client, u)
}

// ChangeFileContentInChangeEdit put content of a file to a change edit.
//
// When change edit doesn’t exist for this change yet it is created.
// When file content isn’t provided, it is wiped out for that file.
// As response “204 No Content” is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#put-edit-file
func (s *ChangesService) ChangeFileContentInChangeEdit(changeID, filePath, content string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit/%s", changeID, url.QueryEscape(filePath))

	req, err := s.client.NewRawPutRequest(u, content)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ChangeCommitMessageInChangeEdit modify commit message.
// The request body needs to include a ChangeEditMessageInput entity.
//
// If a change edit doesn’t exist for this change yet, it is created.
// As response “204 No Content” is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#put-change-edit-message
func (s *ChangesService) ChangeCommitMessageInChangeEdit(changeID string, input *ChangeEditMessageInput) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit:message", changeID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteFileInChangeEdit deletes a file from a change edit.
// This deletes the file from the repository completely.
// This is not the same as reverting or restoring a file to its previous contents.
//
// When change edit doesn’t exist for this change yet it is created.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-edit-file
func (s *ChangesService) DeleteFileInChangeEdit(changeID, filePath string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit/%s", changeID, filePath)
	return s.client.DeleteRequest(u, nil)
}

// DeleteChangeEdit deletes change edit.
//
// As response “204 No Content” is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-edit
func (s *ChangesService) DeleteChangeEdit(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit", changeID)
	return s.client.DeleteRequest(u, nil)
}

// PublishChangeEdit promotes change edit to a regular patch set.
//
// As response “204 No Content” is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#publish-edit
func (s *ChangesService) PublishChangeEdit(changeID, notify string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit:publish", changeID)

	req, err := s.client.NewRequest("POST", u, map[string]string{
		"notify": notify,
	})
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// RebaseChangeEdit rebases change edit on top of latest patch set.
//
// When change was rebased on top of latest patch set, response “204 No Content” is returned.
// When change edit is already based on top of the latest patch set, the response “409 Conflict” is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#rebase-edit
func (s *ChangesService) RebaseChangeEdit(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit:rebase", changeID)

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RetrieveFileContentFromChangeEdit retrieves content of a file from a change edit.
//
// The content of the file is returned as text encoded inside base64.
// The Content-Type header will always be text/plain reflecting the outer base64 encoding.
// A Gerrit-specific X-FYI-Content-Type header can be examined to find the server detected content type of the file.
//
// When the specified file was deleted in the change edit “204 No Content” is returned.
// If only the content type is required, callers should use HEAD to avoid downloading the encoded file contents.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-file
func (s *ChangesService) RetrieveFileContentFromChangeEdit(changeID, filePath string) (*string, *Response, error) {
	u := fmt.Sprintf("changes/%s/edit/%s", changeID, filePath)

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

// RetrieveFileContentTypeFromChangeEdit retrieves content type of a file from a change edit.
// This is nearly the same as RetrieveFileContentFromChangeEdit.
// But if only the content type is required, callers should use HEAD to avoid downloading the encoded file contents.
//
// For further documentation please have a look at RetrieveFileContentFromChangeEdit.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#get-edit-file
func (s *ChangesService) RetrieveFileContentTypeFromChangeEdit(changeID, filePath string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/edit/%s", changeID, filePath)

	req, err := s.client.NewRequest("HEAD", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

/*
Missing Change Edit Endpoints
	Restore file content or rename files in Change Edit
*/
