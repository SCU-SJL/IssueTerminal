package util

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Token = "token " + GetAccessToken()

func GetAccessToken() string {
	token, err := ioutil.ReadFile("access_token.txt")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Open File Error: %v", err)
	}
	res := string(token)
	res = strings.ReplaceAll(res, "\r\n", "")
	res = strings.ReplaceAll(res, "\n", "")
	return res
}

func GetQueryUrl(username, repo string) string {
	return "/" + username + "/" + repo + "/issues"
}

func ReadParam(key string, reader *bufio.Reader) string {
	param, err := reader.ReadString('\n')
	for err != nil {
		fmt.Println("Invalid input")
		fmt.Print(key + ": ")
		param, err = reader.ReadString('\n')
	}
	param = strings.Replace(param, "\r\n", "", -1)
	param = strings.Replace(param, "\n", "", -1)
	return param
}

func Invalid() {
	fmt.Println("Invalid input, try -h")
}
