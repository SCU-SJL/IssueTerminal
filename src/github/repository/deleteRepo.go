package repository

import (
	"bytes"
	"github/util"
	"log"
	"net/http"
)

const GithubDelRepoApi = "https://api.github.com/repos"

func DeleteRepo(username, repo string) string {
	queryUrl := "/" + username + "/" + repo
	jsonStr := "{\"scopes\":[\"delete_repo\"]}"
	jsonBytes := []byte(jsonStr)

	req, err := http.NewRequest("DELETE", GithubDelRepoApi+queryUrl, bytes.NewBuffer(jsonBytes))
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
