package jira

type Issue struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`

	IssueFields *IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`

	IssueType *IssueType `json:"issuetype"`
}

type IssueType struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
}
