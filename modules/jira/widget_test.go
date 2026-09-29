package jira

import (
	"testing"

	"gotest.tools/assert"
)

func issue(key, issueType, status string) Issue {
	return Issue{
		Key: key,
		IssueFields: &IssueFields{
			IssueType:   &IssueType{Name: issueType},
			IssueStatus: &IssueStatus{IName: status},
		},
	}
}

func TestGetLongestColumnLengths_UsesLongestIssueKey(t *testing.T) {
	issues := []Issue{
		issue("PROD-1901", "Task", "In Progress"),
		issue("GRID-11658", "Bug", "Ready"),
	}

	_, longestKeyLength, _ := getLongestColumnLengths(issues)

	// Must track the actual longest key ("GRID-11658"), otherwise shorter keys
	// are padded to a different width and every column after them misaligns.
	assert.Equal(t, len("GRID-11658"), longestKeyLength)
}

func TestGetLongestColumnLengths_KeysOfEqualLength(t *testing.T) {
	issues := []Issue{
		issue("WTF-1", "Bug", "Ready"),
		issue("WTF-2", "Bug", "Ready"),
	}

	_, longestKeyLength, _ := getLongestColumnLengths(issues)

	assert.Equal(t, len("WTF-1"), longestKeyLength)
}

func TestGetLongestColumnLengths_ClampsTypeAndStatus(t *testing.T) {
	issues := []Issue{
		issue("WTF-1", "Improvement", "Waiting for customer"),
	}

	longestIssueTypeLength, _, longestStatusNameLength := getLongestColumnLengths(issues)

	assert.Equal(t, MaxIssueTypeLength, longestIssueTypeLength)
	assert.Equal(t, MaxStatusNameLength, longestStatusNameLength)
}
