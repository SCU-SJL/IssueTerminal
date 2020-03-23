package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const GitHubIssueAPI = "https://api.github.com/repos"

type IssuesList struct {
	Issues []Issue
}

type Issue struct {
	Number    int
	Title     string
	HtmlUrl   string `json:"html_url"`
	State     string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string
	User      *User
}

type User struct {
	Login   string
	HtmlUrl string `json:"html_url"`
}

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
