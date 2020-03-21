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
		for _, item := range result.Issues {
			fmt.Printf("#%-5d %-9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	if invalid {
		fmt.Println("Invalid input, try -h")
	}
}
