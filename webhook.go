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

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// PushEventHandler function
type PushEventHandler func(event *PushEvent)

// IssuesEventHandler function
type IssuesEventHandler func(event *IssuesEvent)

// MergeRequestEventHandler function
type MergeRequestEventHandler func(event *MergeRequestEvent)

// WebHook struct
type WebHook struct {
	PushEventHandler         PushEventHandler
	IssuesEventHandler       IssuesEventHandler
	MergeRequestEventHandler MergeRequestEventHandler
}

// New return new WebHook
func New() *WebHook {
	return &WebHook{
		PushEventHandler: func(event *PushEvent) {
			log.Printf("%T\n", event)
		},
		IssuesEventHandler: func(event *IssuesEvent) {
			log.Printf("%T\n", event)
		},
		MergeRequestEventHandler: func(event *MergeRequestEvent) {
			log.Printf("%T\n", event)
		},
	}
}

func (wh *WebHook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	v, err := parseEvent(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch event := v.(type) {
	case *PushEvent:
		wh.PushEventHandler(event)
	case *IssuesEvent:
		wh.IssuesEventHandler(event)
	case *MergeRequestEvent:
		wh.MergeRequestEventHandler(event)
	default:
	}
	w.WriteHeader(http.StatusOK)
}

func parseEvent(r io.Reader) (interface{}, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	pushEvent := &PushEvent{}
	err = json.Unmarshal(data, pushEvent)
	if err != nil {
		return pushEvent, nil
	}
	issuesEvent := &IssuesEvent{}
	err = json.Unmarshal(data, issuesEvent)
	if err != nil {
		return issuesEvent, nil
	}
	mergeRequestEvent := &MergeRequestEvent{}
	err = json.Unmarshal(data, mergeRequestEvent)
	if err != nil {
		return mergeRequestEvent, nil
	}
	return nil, errors.New("unknown event")
}
