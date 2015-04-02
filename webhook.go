package webhook

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// PushEventHandler function
type PushEventHandler func(event *PushEvent)

// PushTagEventHandler function
type PushTagEventHandler func(evetn *PushEvent)

// IssuesEventHandler function
type IssuesEventHandler func(event *IssuesEvent)

// MergeRequestEventHandler function
type MergeRequestEventHandler func(event *MergeRequestEvent)

// WebHook struct
type WebHook struct {
	PushEventHandler         PushEventHandler
	PushTagEventHandler      PushTagEventHandler
	IssuesEventHandler       IssuesEventHandler
	MergeRequestEventHandler MergeRequestEventHandler

	PushEventPath         string
	PushTagEventPath      string
	IssuesEventPath       string
	MergeRequestEventPath string

	ListenAddr string
}

// Run startup webhook service
func (wh *WebHook) Run() error {
	return http.ListenAndServe(wh.ListenAddr, wh)
}

// New return new WebHook
func New(addr string) *WebHook {
	return &WebHook{
		PushEventPath:         "/push",
		PushTagEventPath:      "/push/tag",
		IssuesEventPath:       "/issues",
		MergeRequestEventPath: "/mergerequest",
		ListenAddr:            addr,
	}
}

func (wh *WebHook) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch path {
	case wh.PushEventPath:
		if wh.PushEventHandler != nil {
			event := &PushEvent{}
			if err := parseEvent(req.Body, event); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			wh.PushEventHandler(event)
			w.WriteHeader(http.StatusOK)
		}
	case wh.PushTagEventPath:
		if wh.PushTagEventHandler != nil {
			event := &PushEvent{}
			if err := parseEvent(req.Body, event); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			wh.PushTagEventHandler(event)
			w.WriteHeader(http.StatusOK)
		}
	case wh.IssuesEventPath:
		if wh.IssuesEventHandler != nil {
			event := &IssuesEvent{}
			if err := parseEvent(req.Body, event); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			wh.IssuesEventHandler(event)
			w.WriteHeader(http.StatusOK)
		}
	case wh.MergeRequestEventPath:
		if wh.MergeRequestEventHandler != nil {
			event := &MergeRequestEvent{}
			if err := parseEvent(req.Body, event); err != nil {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			wh.MergeRequestEventHandler(event)
			w.WriteHeader(http.StatusOK)
		}
	default:
		http.NotFound(w, req)
	}
}

func parseEvent(r io.Reader, event interface{}) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, event)
	return err
}
