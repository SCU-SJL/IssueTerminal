package repository

import (
	"bytes"
	"github/util"
	"log"
	"net/http"
)

const GithubRepoApi = "https://api.github.com/user/repos"

type RepoParam struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateRepo(jsonStr []byte) string {
	req, err := http.NewRequest("POST", GithubRepoApi, bytes.NewBuffer(jsonStr))
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
