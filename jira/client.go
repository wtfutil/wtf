package main

import (
	"fmt"
	"os"

	gojira "gopkg.in/andygrunwald/go-jira.v1"
)

func main() {
	data := Fetch()

	fmt.Printf("+%v\n", data)
}

func Fetch() string {
	tp := gojira.BasicAuthTransport{
		Username: os.Getenv("WTF_JIRA_USERNAME"),
		Password: os.Getenv("WTF_JIRA_PASSWORD"),
	}

	client, err := gojira.NewClient(tp.Client(), "https://lendesk.atlassian.net")
	if err != nil {
		panic(err)
	}

	issue, req, err := client.Issue.Get("CORE-1464", nil)
	if err != nil {
		fmt.Printf(">> %v\n", req)
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)

	return "ok"
}
