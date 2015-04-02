package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var updateScript = flag.String("update-script", "", "update shell path")

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

func init() {
	flag.Parse()
}

func main() {
	if *updateScript == "" {
		log.Fatal("update shell required")
	}

	stat, err := os.Stat(*updateScript)
	if err != nil {
		log.Fatal(err)
	}
	if stat.Mode()&0100 == 0 {
		log.Fatal("script has no exec perm")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/push", pushEventHandler)
	mux.HandleFunc("/push/tag", pushTagEventHandler)

	log.Fatal(http.ListenAndServe(":12138", mux))
}

func pushEventHandler(w http.ResponseWriter, req *http.Request) {
	event, err := parseEvent(req.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if event.Ref != "refs/head/develop" {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	cmd := exec.Command(*updateScript)
	if err = cmd.Run(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func pushTagEventHandler(w http.ResponseWriter, req *http.Request) {
	event, err := parseEvent(req.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("%v\n", event)
}

func parseEvent(r io.Reader) (*PushEvent, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	event := &PushEvent{}
	err = json.Unmarshal(data, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}
