package gerrit

import (
	"fmt"
)

// ConfigService contains Config related REST endpoints
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html
type ConfigService struct {
	client *Client
}

// TopMenuItemInfo entity contains information about a menu item in a top menu entry.
type TopMenuItemInfo struct {
	URL    string `json:"url"`
	Name   string `json:"name"`
	Target string `json:"target"`
	ID     string `json:"id,omitempty"`
}

// AuthInfo entity contains information about the authentication configuration of the Gerrit server.
type AuthInfo struct {
	Type                     string   `json:"type"`
	UseContributorAgreements bool     `json:"use_contributor_agreements,omitempty"`
	EditableAccountFields    []string `json:"editable_account_fields"`
	LoginURL                 string   `json:"login_url,omitempty"`
	LoginText                string   `json:"login_text,omitempty"`
	SwitchAccountURL         string   `json:"switch_account_url,omitempty"`
	RegisterURL              string   `json:"register_url,omitempty"`
	RegisterText             string   `json:"register_text,omitempty"`
	EditFullNameURL          string   `json:"edit_full_name_url,omitempty"`
	HTTPPasswordURL          string   `json:"http_password_url,omitempty"`
	IsGitBasicAuth           bool     `json:"is_git_basic_auth,omitempty"`
}

// CacheInfo entity contains information about a cache.
type CacheInfo struct {
	Name       string       `json:"name,omitempty"`
	Type       string       `json:"type"`
	Entries    EntriesInfo  `json:"entries"`
	AverageGet string       `json:"average_get,omitempty"`
	HitRatio   HitRatioInfo `json:"hit_ratio"`
}

// CacheOperationInput entity contains information about an operation that should be executed on caches.
type CacheOperationInput struct {
	Operation string   `json:"operation"`
	Caches    []string `json:"caches,omitempty"`
}

// ConfigCapabilityInfo entity contains information about a capability.type
type ConfigCapabilityInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HitRatioInfo entity contains information about the hit ratio of a cache.
type HitRatioInfo struct {
	Mem  int `json:"mem"`
	Disk int `json:"disk,omitempty"`
}

// EntriesInfo entity contains information about the entries in a cache.
type EntriesInfo struct {
	Mem   int    `json:"mem,omitempty"`
	Disk  int    `json:"disk,omitempty"`
	Space string `json:"space,omitempty"`
}

// UserConfigInfo entity contains information about Gerrit configuration from the user section.
type UserConfigInfo struct {
	AnonymousCowardName string `json:"anonymous_coward_name"`
}

// TopMenuEntryInfo entity contains information about a top menu entry.
type TopMenuEntryInfo struct {
	Name  string            `json:"name"`
	Items []TopMenuItemInfo `json:"items"`
}

// ThreadSummaryInfo entity contains information about the current threads.
type ThreadSummaryInfo struct {
	CPUs    int                       `json:"cpus"`
	Threads int                       `json:"threads"`
	Counts  map[string]map[string]int `json:"counts"`
}

// TaskSummaryInfo entity contains information about the current tasks.
type TaskSummaryInfo struct {
	Total    int `json:"total,omitempty"`
	Running  int `json:"running,omitempty"`
	Ready    int `json:"ready,omitempty"`
	Sleeping int `json:"sleeping,omitempty"`
}

// TaskInfo entity contains information about a task in a background work queue.
type TaskInfo struct {
	ID         string `json:"id"`
	State      string `json:"state"`
	StartTime  string `json:"start_time"`
	Delay      int    `json:"delay"`
	Command    string `json:"command"`
	RemoteName string `json:"remote_name,omitempty"`
	Project    string `json:"project,omitempty"`
}

// SummaryInfo entity contains information about the current state of the server.
type SummaryInfo struct {
	TaskSummary   TaskSummaryInfo `json:"task_summary"`
	MemSummary    MemSummaryInfo  `json:"mem_summary"`
	ThreadSummary ThemeInfo       `json:"thread_summary"`
	JVMSummary    JvmSummaryInfo  `json:"jvm_summary,omitempty"`
}

// SuggestInfo entity contains information about Gerrit configuration from the suggest section.
type SuggestInfo struct {
	From int `json:"from"`
}

// SSHdInfo entity contains information about Gerrit configuration from the sshd section.
type SSHdInfo struct{}

// ServerInfo entity contains information about the configuration of the Gerrit server.
type ServerInfo struct {
	Auth       AuthInfo          `json:"auth"`
	Change     ChangeConfigInfo  `json:"change"`
	Download   DownloadInfo      `json:"download"`
	Gerrit     Info              `json:"gerrit"`
	Gitweb     map[string]string `json:"gitweb,omitempty"`
	Plugin     PluginConfigInfo  `json:"plugin"`
	Receive    ReceiveInfo       `json:"receive,omitempty"`
	SSHd       SSHdInfo          `json:"sshd,omitempty"`
	Suggest    SuggestInfo       `json:"suggest"`
	URLAliases map[string]string `json:"url_aliases,omitempty"`
	User       UserConfigInfo    `json:"user"`
}

// ReceiveInfo entity contains information about the configuration of git-receive-pack behavior on the server.
type ReceiveInfo struct {
	EnableSignedPush bool `json:"enableSignedPush,omitempty"`
}

// PluginConfigInfo entity contains information about Gerrit extensions by plugins.
type PluginConfigInfo struct {
	// HasAvatars reports whether an avatar provider is registered.
	HasAvatars bool `json:"has_avatars,omitempty"`
}

// MemSummaryInfo entity contains information about the current memory usage.
type MemSummaryInfo struct {
	Total     string `json:"total"`
	Used      string `json:"used"`
	Free      string `json:"free"`
	Buffers   string `json:"buffers"`
	Max       string `json:"max"`
	OpenFiles int    `json:"open_files,omitempty"`
}

// JvmSummaryInfo entity contains information about the JVM.
type JvmSummaryInfo struct {
	VMVendor                string `json:"vm_vendor"`
	VMName                  string `json:"vm_name"`
	VMVersion               string `json:"vm_version"`
	OSName                  string `json:"os_name"`
	OSVersion               string `json:"os_version"`
	OSArch                  string `json:"os_arch"`
	User                    string `json:"user"`
	Host                    string `json:"host,omitempty"`
	CurrentWorkingDirectory string `json:"current_working_directory"`
	Site                    string `json:"site"`
}

// Info entity contains information about Gerrit configuration from the gerrit section.
type Info struct {
	AllProjectsName string `json:"all_projects_name"`
	AllUsersName    string `json:"all_users_name"`
	DocURL          string `json:"doc_url,omitempty"`
	ReportBugURL    string `json:"report_bug_url,omitempty"`
	ReportBugText   string `json:"report_bug_text,omitempty"`
}

// GitwebInfo entity contains information about the gitweb configuration.
type GitwebInfo struct {
	URL  string         `json:"url"`
	Type GitwebTypeInfo `json:"type"`
}

// GitwebTypeInfo entity contains information about the gitweb configuration.
type GitwebTypeInfo struct {
	Name          string `json:"name"`
	Revision      string `json:"revision,omitempty"`
	Project       string `json:"project,omitempty"`
	Branch        string `json:"branch,omitempty"`
	RootTree      string `json:"root_tree,omitempty"`
	File          string `json:"file,omitempty"`
	FileHistory   string `json:"file_history,omitempty"`
	PathSeparator string `json:"path_separator"`
	LinkDrafts    bool   `json:"link_drafts,omitempty"`
	URLEncode     bool   `json:"url_encode,omitempty"`
}

// EmailConfirmationInput entity contains information for confirming an email address.
type EmailConfirmationInput struct {
	Token string `json:"token"`
}

// DownloadSchemeInfo entity contains information about a supported download scheme and its commands.
type DownloadSchemeInfo struct {
	URL             string            `json:"url"`
	IsAuthRequired  bool              `json:"is_auth_required,omitempty"`
	IsAuthSupported bool              `json:"is_auth_supported,omitempty"`
	Commands        map[string]string `json:"commands"`
	CloneCommands   map[string]string `json:"clone_commands"`
}

// DownloadInfo entity contains information about supported download options.
type DownloadInfo struct {
	Schemes  map[string]DownloadSchemeInfo `json:"schemes"`
	Archives []string                      `json:"archives"`
}

// ChangeConfigInfo entity contains information about Gerrit configuration from the change section.
type ChangeConfigInfo struct {
	AllowDrafts      bool   `json:"allow_drafts,omitempty"`
	LargeChange      int    `json:"large_change"`
	ReplyLabel       string `json:"reply_label"`
	ReplyTooltip     string `json:"reply_tooltip"`
	UpdateDelay      int    `json:"update_delay"`
	SubmitWholeTopic bool   `json:"submit_whole_topic"`
}

// ListCachesOptions specifies the different output formats.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#list-caches
type ListCachesOptions struct {
	// Format specifies the different output formats.
	Format string `url:"format,omitempty"`
}

// SummaryOptions specifies the different options for the GetSummary call.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-summary
type SummaryOptions struct {
	// JVM includes a JVM summary.
	JVM bool `url:"jvm,omitempty"`
	// GC requests a Java garbage collection before computing the information about the Java memory heap.
	GC bool `url:"gc,omitempty"`
}

// GetVersion returns the version of the Gerrit server.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-version
func (s *ConfigService) GetVersion() (string, *Response, error) {
	u := "config/server/version"
	return getStringResponseWithoutOptions(s.client, u)
}

// GetServerInfo returns the information about the Gerrit server configuration.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-info
func (s *ConfigService) GetServerInfo() (*ServerInfo, *Response, error) {
	u := "config/server/info"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(ServerInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListCaches lists the caches of the server. Caches defined by plugins are included.
// The caller must be a member of a group that is granted one of the following capabilities:
// * View Caches
// * Maintain Server
// * Administrate Server
// The entries in the map are sorted by cache name.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#list-caches
func (s *ConfigService) ListCaches(opt *ListCachesOptions) (*map[string]CacheInfo, *Response, error) {
	u := "config/server/caches/"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]CacheInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetCache retrieves information about a cache.
// The caller must be a member of a group that is granted one of the following capabilities:
// * View Caches
// * Maintain Server
// * Administrate Server
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-cache
func (s *ConfigService) GetCache(cacheName string) (*CacheInfo, *Response, error) {
	u := fmt.Sprintf("config/server/caches/%s", cacheName)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(CacheInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetSummary retrieves a summary of the current server state.
// The caller must be a member of a group that is granted the Administrate Server capability.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-summary
func (s *ConfigService) GetSummary(opt *SummaryOptions) (*SummaryInfo, *Response, error) {
	u := "config/server/summary"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(SummaryInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListCapabilities lists the capabilities that are available in the system.
// There are two kinds of capabilities: core and plugin-owned capabilities.
// The entries in the map are sorted by capability ID.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#list-capabilities
func (s *ConfigService) ListCapabilities() (*map[string]ConfigCapabilityInfo, *Response, error) {
	u := "config/server/capabilities"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(map[string]ConfigCapabilityInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ListTasks lists the tasks from the background work queues that the Gerrit daemon is currently performing, or will perform in the near future.
// Gerrit contains an internal scheduler, similar to cron, that it uses to queue and dispatch both short and long term tasks.
// Tasks that are completed or canceled exit the queue very quickly once they enter this state, but it can be possible to observe tasks in these states.
// End-users may see a task only if they can also see the project the task is associated with.
// Tasks operating on other projects, or that do not have a specific project, are hidden.
//
// The caller must be a member of a group that is granted one of the following capabilities:
// * View Queue
// * Maintain Server
// * Administrate Server
//
// The entries in the list are sorted by task state, remaining delay and command.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#list-tasks
func (s *ConfigService) ListTasks() (*[]TaskInfo, *Response, error) {
	u := "config/server/tasks"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]TaskInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetTask retrieves a task from the background work queue that the Gerrit daemon is currently performing, or will perform in the near future.
// End-users may see a task only if they can also see the project the task is associated with.
// Tasks operating on other projects, or that do not have a specific project, are hidden.
//
// The caller must be a member of a group that is granted one of the following capabilities:
// * View Queue
// * Maintain Server
// * Administrate Server
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-task
func (s *ConfigService) GetTask(taskID string) (*TaskInfo, *Response, error) {
	u := fmt.Sprintf("config/server/tasks/%s", taskID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new(TaskInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// GetTopMenus returns the list of additional top menu entries.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#get-top-menus
func (s *ConfigService) GetTopMenus() (*[]TopMenuEntryInfo, *Response, error) {
	u := "config/server/top-menus"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := new([]TopMenuEntryInfo)
	resp, err := s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, err
}

// ConfirmEmail confirms that the user owns an email address.
// The email token must be provided in the request body inside an EmailConfirmationInput entity.
//
// The response is “204 No Content”.
// If the token is invalid or if it’s the token of another user the request fails and the response is “422 Unprocessable Entity”.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#confirm-email
func (s *ConfigService) ConfirmEmail(input *EmailConfirmationInput) (*Response, error) {
	u := "config/server/email.confirm"

	req, err := s.client.NewRequest("PUT", u, input)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// CacheOperations executes a cache operation that is specified in the request body in a CacheOperationInput entity.
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#cache-operations
func (s *ConfigService) CacheOperations(input *CacheOperationInput) (*Response, error) {
	u := "config/server/caches/"

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// FlushCache flushes a cache.
// The caller must be a member of a group that is granted one of the following capabilities:
//
// * Flush Caches (any cache except "web_sessions")
// * Maintain Server (any cache including "web_sessions")
// * Administrate Server (any cache including "web_sessions")
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#flush-cache
func (s *ConfigService) FlushCache(cacheName string, input *CacheOperationInput) (*Response, error) {
	u := fmt.Sprintf("config/server/caches/%s/flush", cacheName)

	req, err := s.client.NewRequest("POST", u, input)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DeleteTask kills a task from the background work queue that the Gerrit daemon is currently performing, or will perform in the near future.
// The caller must be a member of a group that is granted one of the following capabilities:
//
// * Kill Task
// * Maintain Server
// * Administrate Server
//
// End-users may see a task only if they can also see the project the task is associated with.
// Tasks operating on other projects, or that do not have a specific project, are hidden.
// Members of a group granted one of the following capabilities may view all tasks:
//
// * View Queue
// * Maintain Server
// * Administrate Server
//
// Gerrit API docs: https://gerrit-review.googlesource.com/Documentation/rest-api-config.html#delete-task
func (s *ConfigService) DeleteTask(taskID string) (*Response, error) {
	u := fmt.Sprintf("config/server/tasks/%s", taskID)
	return s.client.DeleteRequest(u, nil)
}
