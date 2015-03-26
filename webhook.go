package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/codegangsta/cli"
)

const (
	appName    = "webhook"
	appUsage   = "handle git webhook and auto update repo"
	appVersion = "v0.0.1"
	appAuthor  = "Xuyuan Pang"
	appEmail   = "pangxuyuan@gmail.com"
)

// Repository struct
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

// Commit struct
type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

// Author struct
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// PushEvent struct
type PushEvent struct {
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	ProjectID         int        `json:"project_id"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
}

func main() {
	app := cli.NewApp()

	app.Name = appName
	app.Usage = appUsage
	app.Version = appVersion
	app.Author = appAuthor
	app.Email = appEmail
	app.HideHelp = true

	app.Action = func(ctx *cli.Context) {

		mux := http.NewServeMux()

		mux.HandleFunc("/push", pushEventHandler)
		mux.HandleFunc("/push/tag", pushTagEventHandler)

		http.ListenAndServe(":12138", mux)
	}

	app.Run(os.Args)
}

func pushEventHandler(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := &PushEvent{}
	err = json.Unmarshal(data, event)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func pushTagEventHandler(w http.ResponseWriter, req *http.Request) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event := &PushEvent{}
	err = json.Unmarshal(data, event)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("%v\n", event)
}
