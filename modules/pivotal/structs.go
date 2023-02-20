package pivotal

import (
	"time"
)

type User struct {
	ID int `json:"id"`
}

type Story struct {
	Kind       string    `json:"kind"`
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	AcceptedAt time.Time `json:"accepted_at"`
	//    CreatedAt       int64    `json:"created_at"`
	//    UpdatedAt       int64    `json:"updated_at"`
	//    AcceptedAt      int64    `json:"accepted_at"`
	Estimate        int                `json:"estimate"`
	StoryType       string             `json:"story_type"`
	StoryPriority   string             `json:"story_priority"`
	Name            string             `json:"name"`
	Description     string             `json:"description"`
	CurrentState    string             `json:"current_state"`
	RequestedByID   int                `json:"requested_by_id"`
	URL             string             `json:"url"`
	ProjectID       int                `json:"project_id"`
	OwnerIDs        []int              `json:"owner_ids"`
	OwnedByID       int                `json:"owned_by_id"`
	Labels          []Label            `json:"labels"`
	Tasks           []interface{}      `json:"tasks"`
	PullRequests    []StoryPullRequest `json:"pull_requests"`
	CicdEvents      []interface{}      `json:"cicd_events"`
	Branches        []StoryBranch      `json:"branches"`
	Blockers        []interface{}      `json:"blockers"`
	FollowerIDs     []int              `json:"follower_ids"`
	Comments        []StoryComment     `json:"comments"`
	BlockedStoryIDs []int              `json:"blocked_story_ids"`
	Reviews         []StoryReview      `json:"reviews"`
	Project         StoryProject       `json:"project"`
}

type PivotalTrackerResponse struct {
	Stories *StoryResponse `json:"stories"`
	Epics   *EpicResponse  `json:"epics"`
	Query   string         `json:"query"`
}

type StoryResponse struct {
	Stories              []Story `json:"stories"`
	TotalPoints          int     `json:"total_points"`
	TotalPointsCompleted int     `json:"total_points_completed"`
	TotalHits            int     `json:"total_hits"`
	TotalHitsWithDone    int     `json:"total_hits_with_done"`
}

type EpicResponse struct {
	Epics             []Epic `json:"epics"`
	TotalHits         int    `json:"total_hits"`
	TotalHitsWithDone int    `json:"total_hits_with_done"`
}

type Epic struct {
	ID        int       `json:"id"`
	Kind      string    `json:"kind"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//    CreatedAt   int64    `json:"created_at"`
	//    UpdatedAt   int64    `json:"updated_at"`
	ProjectID int    `json:"project_id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Label     Label  `json:"label"`
}

type Label struct {
	ID        int       `json:"id"`
	ProjectID int       `json:"project_id"`
	Kind      string    `json:"kind"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//    CreatedAt   int64    `json:"created_at"`
	//    UpdatedAt   int64    `json:"updated_at"`
}

type StoryLabel struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type StoryPullRequest struct {
	ID        int       `json:"id"`
	Kind      string    `json:"kind"`
	StoryID   int       `json:"story_id"`
	Owner     string    `json:"owner"`
	Repo      string    `json:"repo"`
	HostURL   string    `json:"host_url"`
	Status    string    `json:"status"`
	Number    int       `json:"number"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//    CreatedAt   int64  `json:"created_at"`
	//    UpdatedAt   int64  `json:"updated_at"`
}

type StoryTask struct {
	ID          int    `json:"id"`
	Kind        string `json:"kind"`
	Description string `json:"description"`
	Complete    bool   `json:"complete"`
	Position    int    `json:"position"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
	StoryID     int    `json:"story_id"`
}

type StoryBranch struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	CommitHash    string `json:"commit_hash,omitempty"`
	CommitMessage string `json:"commit_message,omitempty"`
	AuthorName    string `json:"author_name,omitempty"`
	AuthorEmail   string `json:"author_email,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
}

type StoryComment struct {
	Kind        string        `json:"kind"`
	ID          int64         `json:"id"`
	Text        string        `json:"text"`
	PersonID    int64         `json:"person_id"`
	CreatedAt   int64         `json:"created_at"`
	UpdatedAt   int64         `json:"updated_at"`
	StoryID     int64         `json:"story_id"`
	Attachments []interface{} `json:"attachments"`
	Reactions   []interface{} `json:"reactions"`
}

type StoryReview struct {
	ID           int    `json:"id"`
	ReviewerID   int    `json:"reviewer_id"`
	Kind         string `json:"kind"`
	StoryID      int    `json:"story_id"`
	ReviewTypeID int    `json:"review_type_id"`
	Status       string `json:"status"`
	CreatedAt    int    `json:"created_at"`
	UpdatedAt    int    `json:"updated_at"`
}

type StoryProject struct {
	ID                          int    `json:"id"`
	Kind                        string `json:"kind"`
	Name                        string `json:"name"`
	Version                     int    `json:"version"`
	IterationLength             int    `json:"iteration_length"`
	WeekStartDay                string `json:"week_start_day"`
	PointScale                  string `json:"point_scale"`
	PointScaleIsCustom          bool   `json:"point_scale_is_custom"`
	BugsAndChoresAreEstimatable bool   `json:"bugs_and_chores_are_estimatable"`
	AutomaticPlanning           bool   `json:"automatic_planning"`
	EnableTasks                 bool   `json:"enable_tasks"`
	TimeZone                    struct {
		Kind      string `json:"kind"`
		OlsonName string `json:"olson_name"`
		Offset    string `json:"offset"`
	} `json:"time_zone"`
	VelocityAveragedOver         int    `json:"velocity_averaged_over"`
	NumberOfDoneIterationsToShow int    `json:"number_of_done_iterations_to_show"`
	HasGoogleDomain              bool   `json:"has_google_domain"`
	EnableIncomingEmails         bool   `json:"enable_incoming_emails"`
	InitialVelocity              int    `json:"initial_velocity"`
	Public                       bool   `json:"public"`
	AtomEnabled                  bool   `json:"atom_enabled"`
	ProjectType                  string `json:"project_type"`
	HasCICDIntegration           bool   `json:"has_cicd_integration"`
	Capabilities                 struct {
		PrioritySupport         bool `json:"priority_support"`
		LabelsPanel             bool `json:"labels_panel"`
		LabelsPanelBulkActions  bool `json:"labels_panel_bulk_actions"`
		DigitalRiverIntegration bool `json:"digital_river_integration"`
		DigitalRiverDebug       bool `json:"digital_river_debug"`
		StartSendingDRNotices   bool `json:"start_sending_dr_notices"`
		EnableEAPEvents         bool `json:"enable_eap_events"`
	} `json:"capabilities"`
	StartDate                   string `json:"start_date"`
	StartTime                   int64  `json:"start_time"`
	ShownIterationsStartTime    int64  `json:"shown_iterations_start_time"`
	CreatedAt                   int64  `json:"created_at"`
	UpdatedAt                   int64  `json:"updated_at"`
	ShowStoryPriority           bool   `json:"show_story_priority"`
	ShowPriorityIcon            bool   `json:"show_priority_icon"`
	ShowPriorityIconInAllPanels bool   `json:"show_priority_icon_in_all_panels"`
	Epics                       []struct {
		ID      int    `json:"id"`
		Name    string `json:"name"`
		LabelID int    `json:"label_id"`
	} `json:"epics"`
}
