package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github/util"
	"os"
	"strings"
)

func Interactive() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Username / Organization: ")
	username := util.ReadParam("Username/ Organization", in)

	fmt.Print("Repository: ")
	repo := util.ReadParam("Repository", in)

	Delete(username, repo)
}

func Create() {
	in := bufio.NewReader(os.Stdin)
	var params = RepoParam{}

	fmt.Print("Repository name [required]: ")
	repoName := util.ReadParam("Repository name", in)

	fmt.Print("Description [\"n\" for null]: ")
	description := util.ReadParam("Description", in)

	params.Name = repoName
	if description != "n" {
		params.Description = description
	}

	jsonStr, _ := json.Marshal(params)
	result := CreateRepo(jsonStr)

	fmt.Println(result)
}

func Delete(username, repo string) {
	result := DeleteRepo(username, repo)
	fmt.Println(result)
	if strings.HasPrefix(result, "204") {
		fmt.Println("delete successfully.")
	}
}
