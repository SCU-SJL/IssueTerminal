package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssues(username, repo string) (*IssuesList, error) {
	queryUrl := "/" + username + "/" + repo + "/" + "issues"
	resp, err := http.Get(IssueURL + queryUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("query failed: %s", resp.Status)
	}

	var list IssuesList
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		_ = resp.Body.Close()
		return nil, err
	}
	_ = resp.Body.Close()
	list.Issues = result
	return &list, nil
}
