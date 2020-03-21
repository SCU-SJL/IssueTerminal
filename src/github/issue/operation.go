package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIssues(username, repo string) (*IssuesList, error) {
	queryUrl := "/" + username + "/" + repo + "/" + "issues"
	resp, err := http.Get(GitHubIssueAPI + queryUrl)
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

func CreateIssue(username, repo, title, body, label string) string {
	queryUrl := "/" + username + "/" + repo + "/" + "issues" + "?" + "access_token=" + ssh
	jsonStr := "{" +
		"\"title\": " + "\"" + title + "\"," +
		"\"body\": " + "\"" + body + "\"," +
		"\"labels\": [" + "\"" + label + "\"]" +
		"}"
	var postJson = []byte(jsonStr)

	req, err := http.NewRequest("POST", GitHubIssueAPI+queryUrl, bytes.NewBuffer(postJson))
	if err != nil {
		return "bad json format"
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		_ = resp.Body.Close()
		return resp.Status
	}

	return resp.Status
}
