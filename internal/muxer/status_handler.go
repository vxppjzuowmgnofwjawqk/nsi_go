package muxer

import (
	"encoding/json"
	"net/http"
	"nsi_go/internal/scheme"
	"nsi_go/internal/storage"
)

type statusHandler struct {
	storage storage.Storage
}

func newStatusHandler(storage storage.Storage) statusHandler {
	return statusHandler{
		storage: storage,
	}
}

func (h statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "":
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			result := h.storage.GetStatusById(id)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodPost:
			name := r.URL.Query().Get("name")
			result := h.storage.CreateStatus(name)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodDelete:
			id := r.URL.Query().Get("id")
			h.storage.DeleteStatus(id)
		case http.MethodPatch:
			var element scheme.NSI
			dec := json.NewDecoder(r.Body)
			_ = dec.Decode(&element)
			h.storage.UpdateStatus(element.Id, element.Name)
		}
	case "all":
		switch r.Method {
		case http.MethodGet:
			result := h.storage.GetAllStatuses()
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		}
	}
}
