package jira

import ()

type JiraIssue struct {
	Fields JiraIssueFields
	Id     string
	Key    string
	Self   string
}

type JiraIssueFields struct {
	Comment      JiraIssueComment
	Description  string
	Project      JiraIssueProject
	SubTasks     []JiraIssueSubTask
	TimeTracking JiraIssueTimetracking
	Updated      string
	Watcher      JiraIssueWatcher
}

type JiraIssueWatcher struct {
	IsWatching bool
	Self       string
	WatchCount int
	Watchers   []JiraAttribute
}

type JiraIssueSubTask struct {
	Id   string
	Type struct {
		Id      string
		Name    string
		Inward  string
		Outward string
	}
	OutwardIssue struct {
		Id     string
		Key    string
		Self   string
		Fields struct {
			Status struct {
				IconUrl string
				Name    string
			}
		}
	}
}

type JiraIssueProject struct {
	Id              string
	Key             string
	Name            string
	ProjectCategory JiraAttribute
	Self            string
	Simplified      bool
}

type JiraIssueComment struct {
	Author       JiraAttribute
	Body         string
	Created      string
	Id           string
	Self         string
	UpdateAuthor JiraAttribute
	Updated      string
	Visibility   struct {
		Type  string
		Value string
	}
}

type JiraAttribute struct {
	Active      bool
	DisplayName string
	Name        string
	Self        string
}

type JiraIssueTimetracking struct {
	OriginalEstimate         string
	OriginalEstimateSeconds  int
	RemainingEstimate        string
	RemainingEstimateSeconds int
	TimeSpent                string
	TimeSpentSeconds         int
}

type JiraProject struct {
	AssigneeType    string
	Components      []JiraProjectComponent
	Email           string
	IssueTypes      []JiraProjectIssueType
	Lead            JiraProjectMember
	Name            string
	ProjectCategory JiraProjectCategory
	Simplified      bool
	Url             string

	JiraAttribute
}

type JiraProjectCategory struct {
	Description string
	Id          string
	Name        string
	Self        string
}

type JiraProjectComponent struct {
	Assignee            JiraProjectMember
	AssigneeType        string
	IsAssigneeTypeValid bool
	Lead                JiraProjectMember
	Project             string
	ProjectId           int
	RealAssigne         JiraProjectMember
	RealAssigneeType    string

	JiraAttribute
}

type JiraProjectIssueType struct {
	AvatarId    int
	Description string
	IconUrl     string
	Id          string
	Name        string
	Self        string
	SubTask     bool
}

type JiraProjectMember struct {
	AccountId   string
	Active      bool
	DisplayName string
	Key         string
	Name        string
	Self        string
}
