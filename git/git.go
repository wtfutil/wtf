package git

import ()

func Fetch() map[string][]string {
	client := NewClient()

	result := make(map[string][]string)

	result["repo"] = []string{client.Repository}
	result["branch"] = []string{client.CurrentBranch()}
	result["changes"] = client.ChangedFiles()
	result["commits"] = client.Commits()

	return result
}
