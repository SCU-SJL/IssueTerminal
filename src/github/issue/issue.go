package issue

import "time"

// https://api.github.com/repos/username/repo-name/issues
const IssueURL = "https://api.github.com/repos"

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
