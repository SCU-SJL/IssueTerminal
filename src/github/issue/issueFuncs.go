package issue

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func Interact(action string) {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Username / Organization: ")
	username := util.ReadParam("Username/ Organization", in)

	fmt.Print("Repository: ")
	repo := util.ReadParam("Repository", in)

	if action == "get" {
		Get(username, repo)
	} else if action == "close" {
		fmt.Print("Issue ID: ")
		id := util.ReadParam("Issue ID", in)
		Close(username, repo, id)
	}
}

func Get(username, repo string) {
	result, err := GetIssues(username, repo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issues of repo: ", repo)
	fmt.Printf("id\tcreate time\tupdate time\tuser\t\ttitle\n")
	for _, item := range result.Issues {
		fmt.Printf("#%-5d\t%.10v\t%.10v\t%-15.15s %.55s\n", item.Number, item.CreatedAt, item.UpdatedAt, item.User.Login, item.Title)
	}
}

func Put() {
	in := bufio.NewReader(os.Stdin)
	var params = CreateParam{}

	fmt.Print("Username / organization: ")
	username := util.ReadParam("Username/ Organization", in)

	fmt.Print("Repository: ")
	repo := util.ReadParam("Repository", in)

	fmt.Print("Title [required]: ")
	title := util.ReadParam("Title", in)

	fmt.Print("Body [\"n\" for null]: ")
	body := util.ReadParam("Body", in)

	fmt.Print("Milestone [\"n\" for null]: ")
	milestone := util.ReadParam("Milestone", in)

	fmt.Print("Labels [\"n\" for null]: ")
	labels := util.ReadParam("Labels", in)
	labelArr := strings.Fields(labels)
	for i := range labelArr {
		labelArr[i] = strings.Replace(labelArr[i], "gfi", "good first issue", -1)
		labelArr[i] = strings.Replace(labelArr[i], "hw", "help wanted", -1)
	}

	fmt.Print("Assignees [\"n\" for null]: ")
	assignees := util.ReadParam("Assignees", in)
	assigneeArr := strings.Fields(assignees)

	params.Title = title
	if body != "n" {
		params.Body = body
	}
	if milestone != "n" {
		params.Milestone, _ = strconv.Atoi(milestone)
	}
	if labels != "n" {
		params.Labels = labelArr
	}
	if assignees != "n" {
		params.Assignees = assigneeArr
	}

	jsonStr, _ := json.Marshal(params)
	result := CreateIssue(username, repo, jsonStr)
	fmt.Println(result)
}

func Close(username, repo, id string) {
	params := UpdateParam{}
	params.State = "closed"

	jsonStr, _ := json.Marshal(params)

	result := UpdateIssue(username, repo, id, jsonStr)
	fmt.Println(result)
}

func Update() {
	in := bufio.NewReader(os.Stdin)
	var params = UpdateParam{}

	fmt.Print("Username / organization: ")
	username := util.ReadParam("Username/ Organization", in)

	fmt.Print("Repository: ")
	repo := util.ReadParam("Repository", in)

	fmt.Print("Issue ID: ")
	id := util.ReadParam("Issue ID", in)

	fmt.Print("Title [\"n\" for null]: ")
	title := util.ReadParam("Title", in)

	fmt.Print("Body [\"n\" for null]: ")
	body := util.ReadParam("Body", in)

	fmt.Print("State [closed / open / \"n\" for null]: ")
	state := util.ReadParam("State", in)

	fmt.Print("Milestone [\"n\" for null]: ")
	milestone := util.ReadParam("Milestone", in)

	fmt.Print("Labels [\"n\" for null]: ")
	labels := util.ReadParam("Labels", in)
	labelArr := strings.Fields(labels)
	for i := range labelArr {
		labelArr[i] = strings.Replace(labelArr[i], "gfi", "good first issue", -1)
		labelArr[i] = strings.Replace(labelArr[i], "hw", "help wanted", -1)
	}

	fmt.Print("Assignees [\"n\" for null]: ")
	assignees := util.ReadParam("Assignees", in)
	assigneeArr := strings.Fields(assignees)

	if title != "n" {
		params.Title = title
	}
	if body != "n" {
		params.Body = body
	}
	if state != "n" {
		params.State = state
	}
	if milestone != "n" {
		params.Milestone, _ = strconv.Atoi(milestone)
	}
	if labels != "n" {
		params.Labels = labelArr
	}
	if assignees != "n" {
		params.Assignees = assigneeArr
	}

	jsonStr, _ := json.Marshal(params)
	result := UpdateIssue(username, repo, id, jsonStr)
	fmt.Println(result)
}
