package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

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
	mux := http.NewServeMux()

	mux.HandleFunc("/push", pushEventHandler)
	mux.HandleFunc("/push/tag", pushTagEventHandler)

	http.ListenAndServe(":12138", mux)
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

	fmt.Printf("%v\n", event)
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
