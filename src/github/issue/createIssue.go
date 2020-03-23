package issue

import (
	"bytes"
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
	queryUrl := GetQueryUrl(username, repo)
	req, err := http.NewRequest("POST", GitHubIssueAPI+queryUrl, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
		return "bad json format"
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	return resp.Status
}
