package main

import (
	"flag"
	"fmt"
	"github/issue"
	"os"
)

var tip = flag.Bool("tip", false, "please save your github token into access_token.txt\n"+
	"if you want to update / close / create issues\n"+
	"and make sure issueTerminal.exe and access_token.txt is in the same directory\n")

var get = flag.Bool("get", false, "list issues in a repository\n"+
	"FOR EXAMPLE: -get [username / organization] [repository]\n")

var put = flag.Bool("put", false, "put a new issue to a repository\n"+
	"terminal will enter interactive mode automatically\n"+
	"FOR EXAMPLE: -put\n")

var interact = flag.Bool("i", false, "enter interactive mode\n"+
	"FOR EXAMPLE: -get -i\n")

var closed = flag.Bool("close", false, "close an issue in a repository\n"+
	"FOR EXAMPLE: -close [username / organization] [repository] [issue id]\n")

var update = flag.Bool("update", false, "update an issue in a repository\n"+
	"terminal will enter interactive mode automatically\n"+
	"FOR EXAMPLE: -update\n")

var github = flag.Bool("github", false, "Description of github api parameters:\n"+
	"[Username / Organization]: Owner of the repository\n"+
	"[Repository]: Name of the repository\n"+
	"[Title]: Title of the issue\n"+
	"[Body]: Contents of the issue\n"+
	"[Milestone]: Milestone of the issue, which is a number\n"+
	"[State]: State of the issue, which is 'closed' or 'open'\n"+
	"[Labels]: Labels of the issue, including: {bug, documentation, duplicate, enhancement"+
	"good first issue, help wanted, invalid, question, wontfix}\n"+
	"[Assignees]: Assignees of the issue, which is an array of user ids\n")

func main() {
	flag.Parse()

	if *tip {
		usage := flag.Lookup("tip").Usage
		fmt.Println(usage)
		return
	}

	if *github {
		usage := flag.Lookup("github").Usage
		fmt.Println(usage)
		return
	}

	if illegalFlag() {
		issue.Invalid()
		return
	}

	if *interact {
		if *get {
			issue.Interact("get")
		} else if *put {
			issue.Put()
		} else if *closed {
			issue.Interact("close")
		} else if *update {
			issue.Update()
		}
	} else {
		if *get {
			if len(os.Args) != 4 {
				issue.Invalid()
				return
			}
			issue.Get(os.Args[2], os.Args[3])
		} else if *put {
			issue.Put()
		} else if *closed {
			if len(os.Args) != 5 {
				issue.Invalid()
				return
			}
			issue.Close(os.Args[2], os.Args[3], os.Args[4])
		} else if *update {
			issue.Update()
		}
	}
}

func illegalFlag() bool {
	count := 0
	if *get {
		count++
	}
	if *put {
		count++
	}
	if *closed {
		count++
	}
	if *update {
		count++
	}
	return count != 1
}
