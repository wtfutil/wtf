package gerrit

import (
	"fmt"
)

// AccountsService contains Account related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html
type AccountsService struct {
	client *Client
}

// AccountInfo entity contains information about an account.
type AccountInfo struct {
	AccountID int    `json:"_account_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`

	// Avatars lists avatars of various sizes for the account.
	// This field is only populated if the avatars plugin is enabled.
	Avatars []struct {
		URL    string `json:"url,omitempty"`
		Height int    `json:"height,omitempty"`
	} `json:"avatars,omitempty"`
}

// SSHKeyInfo entity contains information about an SSH key of a user.
type SSHKeyInfo struct {
	Seq          int    `json:"seq"`
	SSHPublicKey string `json:"ssh_public_key"`
	EncodedKey   string `json:"encoded_key"`
	Algorithm    string `json:"algorithm"`
	Comment      string `json:"comment,omitempty"`
	Valid        bool   `json:"valid"`
}

// UsernameInput entity contains information for setting the username for an account.
type UsernameInput struct {
	Username string `json:"username"`
}

// QueryLimitInfo entity contains information about the Query Limit of a user.
type QueryLimitInfo struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// HTTPPasswordInput entity contains information for setting/generating an HTTP password.
type HTTPPasswordInput struct {
	Generate     bool   `json:"generate,omitempty"`
	HTTPPassword string `json:"http_password,omitempty"`
}

// GpgKeysInput entity contains information for adding/deleting GPG keys.
type GpgKeysInput struct {
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
}

// GpgKeyInfo entity contains information about a GPG public key.
type GpgKeyInfo struct {
	ID          string   `json:"id,omitempty"`
	Fingerprint string   `json:"fingerprint,omitempty"`
	UserIDs     []string `json:"user_ids,omitempty"`
	Key         string   `json:"key,omitempty"`
}

// EmailInput entity contains information for registering a new email address.
type EmailInput struct {
	Email          string `json:"email"`
	Preferred      bool   `json:"preferred,omitempty"`
	NoConfirmation bool   `json:"no_confirmation,omitempty"`
}

// EmailInfo entity contains information about an email address of a user.
type EmailInfo struct {
	Email               string `json:"email"`
	Preferred           bool   `json:"preferred,omitempty"`
	PendingConfirmation bool   `json:"pending_confirmation,omitempty"`
}

// AccountInput entity contains information for the creation of a new account.
type AccountInput struct {
	Username     string   `json:"username,omitempty"`
	Name         string   `json:"name,omitempty"`
	Email        string   `json:"email,omitempty"`
	SSHKey       string   `json:"ssh_key,omitempty"`
	HTTPPassword string   `json:"http_password,omitempty"`
	Groups       []string `json:"groups,omitempty"`
}

// AccountDetailInfo entity contains detailed information about an account.
type AccountDetailInfo struct {
	AccountInfo
	RegisteredOn Timestamp `json:"registered_on"`
}

// AccountNameInput entity contains information for setting a name for an account.
type AccountNameInput struct {
	Name string `json:"name,omitempty"`
}

// AccountCapabilityInfo entity contains information about the global capabilities of a user.
type AccountCapabilityInfo struct {
	AccessDatabase     bool           `json:"accessDatabase,omitempty"`
	AdministrateServer bool           `json:"administrateServer,omitempty"`
	CreateAccount      bool           `json:"createAccount,omitempty"`
	CreateGroup        bool           `json:"createGroup,omitempty"`
	CreateProject      bool           `json:"createProject,omitempty"`
	EmailReviewers     bool           `json:"emailReviewers,omitempty"`
	FlushCaches        bool           `json:"flushCaches,omitempty"`
	KillTask           bool           `json:"killTask,omitempty"`
	MaintainServer     bool           `json:"maintainServer,omitempty"`
	Priority           string         `json:"priority,omitempty"`
	QueryLimit         QueryLimitInfo `json:"queryLimit"`
	RunAs              bool           `json:"runAs,omitempty"`
	RunGC              bool           `json:"runGC,omitempty"`
	StreamEvents       bool           `json:"streamEvents,omitempty"`
	ViewAllAccounts    bool           `json:"viewAllAccounts,omitempty"`
	ViewCaches         bool           `json:"viewCaches,omitempty"`
	ViewConnections    bool           `json:"viewConnections,omitempty"`
	ViewPlugins        bool           `json:"viewPlugins,omitempty"`
	ViewQueue          bool           `json:"viewQueue,omitempty"`
}

// DiffPreferencesInfo entity contains information about the diff preferences of a user.
type DiffPreferencesInfo struct {
	Context                 int    `json:"context"`
	Theme                   string `json:"theme"`
	ExpandAllComments       bool   `json:"expand_all_comments,omitempty"`
	IgnoreWhitespace        string `json:"ignore_whitespace"`
	IntralineDifference     bool   `json:"intraline_difference,omitempty"`
	LineLength              int    `json:"line_length"`
	ManualReview            bool   `json:"manual_review,omitempty"`
	RetainHeader            bool   `json:"retain_header,omitempty"`
	ShowLineEndings         bool   `json:"show_line_endings,omitempty"`
	ShowTabs                bool   `json:"show_tabs,omitempty"`
	ShowWhitespaceErrors    bool   `json:"show_whitespace_errors,omitempty"`
	SkipDeleted             bool   `json:"skip_deleted,omitempty"`
	SkipUncommented         bool   `json:"skip_uncommented,omitempty"`
	SyntaxHighlighting      bool   `json:"syntax_highlighting,omitempty"`
	HideTopMenu             bool   `json:"hide_top_menu,omitempty"`
	AutoHideDiffTableHeader bool   `json:"auto_hide_diff_table_header,omitempty"`
	HideLineNumbers         bool   `json:"hide_line_numbers,omitempty"`
	TabSize                 int    `json:"tab_size"`
	HideEmptyPane           bool   `json:"hide_empty_pane,omitempty"`
}

// DiffPreferencesInput entity contains information for setting the diff preferences of a user.
// Fields which are not set will not be updated.
type DiffPreferencesInput struct {
	Context                 int    `json:"context,omitempty"`
	ExpandAllComments       bool   `json:"expand_all_comments,omitempty"`
	IgnoreWhitespace        string `json:"ignore_whitespace,omitempty"`
	IntralineDifference     bool   `json:"intraline_difference,omitempty"`
	LineLength              int    `json:"line_length,omitempty"`
	ManualReview            bool   `json:"manual_review,omitempty"`
	RetainHeader            bool   `json:"retain_header,omitempty"`
	ShowLineEndings         bool   `json:"show_line_endings,omitempty"`
	ShowTabs                bool   `json:"show_tabs,omitempty"`
	ShowWhitespaceErrors    bool   `json:"show_whitespace_errors,omitempty"`
	SkipDeleted             bool   `json:"skip_deleted,omitempty"`
	SkipUncommented         bool   `json:"skip_uncommented,omitempty"`
	SyntaxHighlighting      bool   `json:"syntax_highlighting,omitempty"`
	HideTopMenu             bool   `json:"hide_top_menu,omitempty"`
	AutoHideDiffTableHeader bool   `json:"auto_hide_diff_table_header,omitempty"`
	HideLineNumbers         bool   `json:"hide_line_numbers,omitempty"`
	TabSize                 int    `json:"tab_size,omitempty"`
}

// PreferencesInfo entity contains information about a user’s preferences.
type PreferencesInfo struct {
	ChangesPerPage            int               `json:"changes_per_page"`
	ShowSiteHeader            bool              `json:"show_site_header,omitempty"`
	UseFlashClipboard         bool              `json:"use_flash_clipboard,omitempty"`
	DownloadScheme            string            `json:"download_scheme"`
	DownloadCommand           string            `json:"download_command"`
	CopySelfOnEmail           bool              `json:"copy_self_on_email,omitempty"`
	DateFormat                string            `json:"date_format"`
	TimeFormat                string            `json:"time_format"`
	RelativeDateInChangeTable bool              `json:"relative_date_in_change_table,omitempty"`
	SizeBarInChangeTable      bool              `json:"size_bar_in_change_table,omitempty"`
	LegacycidInChangeTable    bool              `json:"legacycid_in_change_table,omitempty"`
	MuteCommonPathPrefixes    bool              `json:"mute_common_path_prefixes,omitempty"`
	ReviewCategoryStrategy    string            `json:"review_category_strategy"`
	DiffView                  string            `json:"diff_view"`
	My                        []TopMenuItemInfo `json:"my"`
	URLAliases                string            `json:"url_aliases,omitempty"`
}

// PreferencesInput entity contains information for setting the user preferences.
// Fields which are not set will not be updated.
type PreferencesInput struct {
	ChangesPerPage            int               `json:"changes_per_page,omitempty"`
	ShowSiteHeader            bool              `json:"show_site_header,omitempty"`
	UseFlashClipboard         bool              `json:"use_flash_clipboard,omitempty"`
	DownloadScheme            string            `json:"download_scheme,omitempty"`
	DownloadCommand           string            `json:"download_command,omitempty"`
	CopySelfOnEmail           bool              `json:"copy_self_on_email,omitempty"`
	DateFormat                string            `json:"date_format,omitempty"`
	TimeFormat                string            `json:"time_format,omitempty"`
	RelativeDateInChangeTable bool              `json:"relative_date_in_change_table,omitempty"`
	SizeBarInChangeTable      bool              `json:"size_bar_in_change_table,omitempty"`
	LegacycidInChangeTable    bool              `json:"legacycid_in_change_table,omitempty"`
	MuteCommonPathPrefixes    bool              `json:"mute_common_path_prefixes,omitempty"`
	ReviewCategoryStrategy    string            `json:"review_category_strategy,omitempty"`
	DiffView                  string            `json:"diff_view,omitempty"`
	My                        []TopMenuItemInfo `json:"my,omitempty"`
	URLAliases                string            `json:"url_aliases,omitempty"`
}

// CapabilityOptions specifies the parameters to filter for capabilities.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-account-capabilities
type CapabilityOptions struct {
	// To filter the set of global capabilities the q parameter can be used.
	// Filtering may decrease the response time by avoiding looking at every possible alternative for the caller.
	Filter []string `url:"q,omitempty"`
}

// GetAccount returns an account as an AccountInfo entity.
// If account is "self" the current authenticated account will be returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-account
func (s *AccountsService) GetAccount(account string) (*AccountInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s", account)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetAccountDetails retrieves the details of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-detail
func (s *AccountsService) GetAccountDetails(accountID string) (*AccountDetailInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/detail", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountDetailInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetAccountName retrieves the full name of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-account-name
func (s *AccountsService) GetAccountName(accountID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/name", accountID)
	return getStringResponseWithoutOptions(s.client, u)
}

// GetUsername retrieves the username of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-username
func (s *AccountsService) GetUsername(accountID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/username", accountID)
	return getStringResponseWithoutOptions(s.client, u)
}

// GetHTTPPassword retrieves the HTTP password of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-http-password
func (s *AccountsService) GetHTTPPassword(accountID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/password.http", accountID)
	return getStringResponseWithoutOptions(s.client, u)
}

// ListAccountEmails returns the email addresses that are configured for the specified user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-account-emails
func (s *AccountsService) ListAccountEmails(accountID string) (*[]EmailInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/emails", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]EmailInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetAccountEmail retrieves an email address of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-account-email
func (s *AccountsService) GetAccountEmail(accountID, emailID string) (*EmailInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/emails/%s", accountID, emailID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(EmailInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListSSHKeys returns the SSH keys of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-ssh-keys
func (s *AccountsService) ListSSHKeys(accountID string) (*[]SSHKeyInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/sshkeys", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]SSHKeyInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetSSHKey retrieves an SSH key of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-ssh-key
func (s *AccountsService) GetSSHKey(accountID, sshKeyID string) (*SSHKeyInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/sshkeys/%s", accountID, sshKeyID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(SSHKeyInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListGPGKeys returns the GPG keys of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-gpg-keys
func (s *AccountsService) ListGPGKeys(accountID string) (*map[string]GpgKeyInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gpgkeys", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]GpgKeyInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetGPGKey retrieves a GPG key of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-gpg-key
func (s *AccountsService) GetGPGKey(accountID, gpgKeyID string) (*GpgKeyInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gpgkeys/%s", accountID, gpgKeyID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(GpgKeyInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListAccountCapabilities returns the global capabilities that are enabled for the specified user.
// If the global capabilities for the calling user should be listed, self can be used as account-id.
// This can be used by UI tools to discover if administrative features are available to the caller, so they can hide (or show) relevant UI actions.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-account-capabilities
func (s *AccountsService) ListAccountCapabilities(accountID string, opt *CapabilityOptions) (*AccountCapabilityInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/capabilities", accountID)

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountCapabilityInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListGroups lists all groups that contain the specified user as a member.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#list-groups
func (s *AccountsService) ListGroups(accountID string) (*[]GroupInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/groups", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]GroupInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetUserPreferences retrieves the user’s preferences.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-user-preferences
func (s *AccountsService) GetUserPreferences(accountID string) (*PreferencesInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/preferences", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(PreferencesInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetDiffPreferences retrieves the diff preferences of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-diff-preferences
func (s *AccountsService) GetDiffPreferences(accountID string) (*DiffPreferencesInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/preferences.diff", accountID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(DiffPreferencesInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetStarredChanges gets the changes starred by the identified user account.
// This URL endpoint is functionally identical to the changes query GET /changes/?q=is:starred.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-starred-changes
func (s *AccountsService) GetStarredChanges(accountID string) (*[]ChangeInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/starred.changes", accountID)

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

// SuggestAccount suggests users for a given query q and result limit n.
// If result limit is not passed, then the default 10 is used.
// Returns a list of matching AccountInfo entities.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#query-account
func (s *AccountsService) SuggestAccount(opt *QueryOptions) (*[]AccountInfo, *Response, error) {
	u := "accounts/"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CreateAccount creates a new account.
// In the request body additional data for the account can be provided as AccountInput.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#create-account
func (s *AccountsService) CreateAccount(username string, input *AccountInput) (*AccountInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s", username)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(AccountInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetAccountName sets the full name of an account.
// The new account name must be provided in the request body inside an AccountNameInput entity.
//
// As response the new account name is returned.
// If the name was deleted the response is “204 No Content”.
// Some realms may not allow to modify the account name.
// In this case the request is rejected with “405 Method Not Allowed”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-account-name
func (s *AccountsService) SetAccountName(accountID string, input *AccountNameInput) (*string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/name", accountID)

	// TODO Use here the getStringResponseWithoutOptions (for PUT requests)

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

// DeleteAccountName deletes the name of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-account-name
func (s *AccountsService) DeleteAccountName(accountID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/name", accountID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteActive sets the account state to inactive.
// If the account was already inactive the response is “404 Not Found”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-active
func (s *AccountsService) DeleteActive(accountID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteHTTPPassword deletes the HTTP password of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-http-password
func (s *AccountsService) DeleteHTTPPassword(accountID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/password.http", accountID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteAccountEmail deletes an email address of an account.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-account-email
func (s *AccountsService) DeleteAccountEmail(accountID, emailID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/emails/%s", accountID, emailID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteSSHKey deletes an SSH key of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-ssh-key
func (s *AccountsService) DeleteSSHKey(accountID, sshKeyID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/sshkeys/%s", accountID, sshKeyID)
	return s.client.DeleteRequest(u, nil)
}

// DeleteGPGKey deletes a GPG key of a user.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#delete-gpg-key
func (s *AccountsService) DeleteGPGKey(accountID, gpgKeyID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/gpgkeys/%s", accountID, gpgKeyID)
	return s.client.DeleteRequest(u, nil)
}

// SetUsername sets a new username.
// The new username must be provided in the request body inside a UsernameInput entity.
// Once set, the username cannot be changed or deleted.
// If attempted this fails with “405 Method Not Allowed”.
//
// As response the new username is returned.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-username
func (s *AccountsService) SetUsername(accountID string, input *UsernameInput) (*string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/username", accountID)

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

// GetActive checks if an account is active.
//
// If the account is active the string ok is returned.
// If the account is inactive the response is “204 No Content”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-active
func (s *AccountsService) GetActive(accountID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/active", accountID)
	return getStringResponseWithoutOptions(s.client, u)
}

// SetActive sets the account state to active.
//
// If the account was already active the response is “200 OK”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-active
func (s *AccountsService) SetActive(accountID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/active", accountID)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// SetHTTPPassword sets/Generates the HTTP password of an account.
// The options for setting/generating the HTTP password must be provided in the request body inside a HTTPPasswordInput entity.
//
// As response the new HTTP password is returned.
// If the HTTP password was deleted the response is “204 No Content”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-http-password
func (s *AccountsService) SetHTTPPassword(accountID string, input *HTTPPasswordInput) (*string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/password.http", accountID)

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

// CreateAccountEmail registers a new email address for the user.
// A verification email is sent with a link that needs to be visited to confirm the email address, unless DEVELOPMENT_BECOME_ANY_ACCOUNT is used as authentication type.
// For the development mode email addresses are directly added without confirmation.
// A Gerrit administrator may add an email address without confirmation by setting no_confirmation in the EmailInput.
// In the request body additional data for the email address can be provided as EmailInput.
//
// As response the new email address is returned as EmailInfo entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#create-account-email
func (s *AccountsService) CreateAccountEmail(accountID, emailID string, input *EmailInput) (*EmailInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/emails/%s", accountID, emailID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(EmailInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetPreferredEmail sets an email address as preferred email address for an account.
//
// If the email address was already the preferred email address of the account the response is “200 OK”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-preferred-email
func (s *AccountsService) SetPreferredEmail(accountID, emailID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/emails/%s/preferred", accountID, emailID)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(req, nil)
}

// GetAvatarChangeURL retrieves the URL where the user can change the avatar image.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#get-avatar-change-url
func (s *AccountsService) GetAvatarChangeURL(accountID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/avatar.change.url", accountID)
	return getStringResponseWithoutOptions(s.client, u)
}

// AddGPGKeys Add or delete one or more GPG keys for a user.
// The changes must be provided in the request body as a GpgKeysInput entity.
// Each new GPG key is provided in ASCII armored format, and must contain a self-signed certification matching a registered email or other identity of the user.
//
// As a response, the modified GPG keys are returned as a map of GpgKeyInfo entities, keyed by ID. Deleted keys are represented by an empty object.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#add-delete-gpg-keys
func (s *AccountsService) AddGPGKeys(accountID string, input *GpgKeysInput) (*map[string]GpgKeyInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/gpgkeys", accountID)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]GpgKeyInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// CheckAccountCapability checks if a user has a certain global capability.
//
// If the user has the global capability the string ok is returned.
// If the user doesn’t have the global capability the response is “404 Not Found”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#check-account-capability
func (s *AccountsService) CheckAccountCapability(accountID, capabilityID string) (string, *Response, error) {
	u := fmt.Sprintf("accounts/%s/capabilities/%s", accountID, capabilityID)
	return getStringResponseWithoutOptions(s.client, u)
}

// SetUserPreferences sets the user’s preferences.
// The new preferences must be provided in the request body as a PreferencesInput entity.
//
// As result the new preferences of the user are returned as a PreferencesInfo entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-user-preferences
func (s *AccountsService) SetUserPreferences(accountID string, input *PreferencesInput) (*PreferencesInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/preferences", accountID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(PreferencesInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// SetDiffPreferences sets the diff preferences of a user.
// The new diff preferences must be provided in the request body as a DiffPreferencesInput entity.
//
// As result the new diff preferences of the user are returned as a DiffPreferencesInfo entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#set-diff-preferences
func (s *AccountsService) SetDiffPreferences(accountID string, input *DiffPreferencesInput) (*DiffPreferencesInfo, *Response, error) {
	u := fmt.Sprintf("accounts/%s/preferences.diff", accountID)

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, nil, err
	}

	v := new(DiffPreferencesInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// StarChange star a change.
// Starred changes are returned for the search query is:starred or starredby:USER and automatically notify the user whenever updates are made to the change.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#star-change
func (s *AccountsService) StarChange(accountID, changeID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/starred.changes/%s", accountID, changeID)

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnstarChange nstar a change.
// Removes the starred flag, stopping notifications.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#unstar-change
func (s *AccountsService) UnstarChange(accountID, changeID string) (*Response, error) {
	u := fmt.Sprintf("accounts/%s/starred.changes/%s", accountID, changeID)
	return s.client.DeleteRequest(u, nil)
}

/*
Missing Account Endpoints:
	Add SSH Key
	Get Avatar
*/
