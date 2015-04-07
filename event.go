/*
 * Copyright 2015 Xuyuan Pang
 * Author: Xuyuan Pang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package webhook

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

// MergeRequestObjectAttributes struct
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

// MergeRequestEvent struct
type MergeRequestEvent struct {
	ObjectKind       string                       `json:"object_kind"`
	User             User                         `json:"user"`
	ObjectAttributes MergeRequestObjectAttributes `json:"object_attributes"`
}
