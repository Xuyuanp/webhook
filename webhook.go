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

// User struct
type User struct {
	Name      string `json:"name"`
	UserName  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// IssuesObjectAttributes struct
type IssuesObjectAttributes struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	AssigneeID  int    `json:"assigneee_id"`
	AuthorID    int    `json:"author_id"`
	ProjectID   int    `json:"project_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Position    string `json:"position"`
	BranchName  string `json:"branch_name"`
	Description string `json:"description"`
	MilestoneID int    `json:"milestone_id"`
	State       string `json:"state"`
	IID         int    `json:"iid"`
	URL         string `json:url`
	Action      string `json:action`
}

// IssuesEvent struct
type IssuesEvent struct {
	ObjectKind       string                 `json:"object_kind"`
	User             User                   `json:"user"`
	ObjectAttributes IssuesObjectAttributes `json:"object_attributes"`
}

// Peer struct
type Peer struct {
	Name            string `json:"name"`
	SSHURL          string `json:"ssh_url"`
	HTTPURL         string `json:"http_url"`
	VisibilityLevel int    `json:"visibility_level"`
	Namespace       string `json:"namespace"`
}

type MergeRequestObjectAttributes struct {
	ID              int    `json:"id"`
	TargetBranch    string `json:"target_branch"`
	SourceBranch    string `json:"source_branch"`
	SourceProjectID int    `json:"source_project_id"`
	AuthorID        int    `json:"author_id"`
	AssigneeID      int    `json:"assigneee_id"`
	Title           string `json:"title"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	StCommits       string `json:"st_commits"`
	StDiffs         string `json:"st_diffs"`
	MilestoneID     int    `json:"milestone_id"`
	State           string `json:"state"`
	MergeStatus     string `json:"merge_status"`
	TargetProjectID int    `json:"target_project_id"`
	IID             int    `json:"iid"`
	Description     string `json:"description"`
	Source          Peer   `json:"source"`
	Target          Peer   `json:"target"`
	LastCommit      Commit `json:"last_commit"`
}

func init() {
	flag.Parse()
}

func main() {
	if *updateScript == "" {
		log.Fatal("update shell required")
	}

	_, err := os.Stat(*updateScript)
	if err != nil {
		log.Fatal(err)
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

	if event.Ref != "refs/heads/develop" || event.UserName != "pangxuyuan" || event.UserName != "Xuyuan Pang" {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	cmd := exec.Command("sh", *updateScript)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(string(out))
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
