package muxer

import (
	"encoding/json"
	"net/http"
	"nsi_go/internal/scheme"
	"nsi_go/internal/storage"
)

type languageHandler struct {
	storage storage.Storage
}

func newLanguageHandler(storage storage.Storage) languageHandler {
	return languageHandler{
		storage: storage,
	}
}

func (h languageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "":
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			result := h.storage.GetLanguageById(id)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodPost:
			name := r.URL.Query().Get("name")
			result := h.storage.CreateLanguage(name)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodDelete:
			id := r.URL.Query().Get("id")
			h.storage.DeleteLanguage(id)
		case http.MethodPatch:
			var element scheme.NSI
			dec := json.NewDecoder(r.Body)
			_ = dec.Decode(&element)
			h.storage.UpdateLanguage(element.Id, element.Name)
		}
	case "all":
		switch r.Method {
		case http.MethodGet:
			result := h.storage.GetAllLanguages()
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		}
	}
}
