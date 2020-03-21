package main

import (
	"flag"
	"fmt"
	"github/issue"
	"log"
	"os"
)

var get = flag.Bool("get", false, "list issues in a repository\n"+
	"for example: -get [username / organization] [repository]")
var put = flag.Bool("put", false, "put a new issue to a repository\n"+
	"for example: -put [username / organization] [repository] [title] [body] [label]")

func main() {
	flag.Parse()
	invalid := true
	if *get {
		invalid = false
		username := os.Args[2]
		repo := os.Args[3]
		result, err := issue.GetIssues(username, repo)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Issues of repo: ", repo)
		fmt.Printf("id\tcreate time\tupdate time\tuser\t\ttitle\n")
		for _, item := range result.Issues {
			fmt.Printf("#%-5d\t%.10v\t%.10v\t%-15.15s %.55s\n", item.Number, item.CreatedAt, item.UpdatedAt, item.User.Login, item.Title)
		}
	} else if *put {
		invalid = false
		username := os.Args[2]
		repo := os.Args[3]
		title := os.Args[4]
		body := os.Args[5]
		label := os.Args[6]
		result := issue.CreateIssue(username, repo, title, body, label)
		fmt.Println(result)
	}
	if invalid {
		fmt.Println("Invalid input, try -h")
	}
}
