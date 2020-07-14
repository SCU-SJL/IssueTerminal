package issue

import (
	"bytes"
	"issue_term/src/github/util"
	"log"
	"net/http"
)

type UpdateParam struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

func UpdateIssue(username, repo, id string, jsonStr []byte) string {
	queryUrl := util.GetQueryUrl(username, repo) + "/" + id
	req, err := http.NewRequest("PATCH", GitHubIssueAPI+queryUrl, bytes.NewBuffer(jsonStr))
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
