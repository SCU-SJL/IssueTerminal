package issue

import (
	"bytes"
	"issue_term/src/github/util"
	"log"
	"net/http"
)

type CreateParam struct {
	Title     string   `json:"title"`
	Body      string   `json:"body,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

func CreateIssue(username, repo string, jsonStr []byte) string {
	queryUrl := util.GetQueryUrl(username, repo)
	req, err := http.NewRequest("POST", GitHubIssueAPI+queryUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", util.Token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp.Status
}
