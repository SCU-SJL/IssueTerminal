package issue

import "time"

const GitHubIssueAPI = "https://api.github.com/repos"
const ssh = "e8e02848161f0e6129b831019e717f53d30b9ba0"

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
