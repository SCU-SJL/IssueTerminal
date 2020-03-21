package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
	token := GetAccessToken()
	queryUrl := "/" + username + "/" + repo + "/" + "issues" + "?" + "access_token=" + token
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

func GetAccessToken() string {
	token, err := ioutil.ReadFile("../../../resource/access_token.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Open File Error: %v", err)
	}
	return string(token)
}
