package muxer

import (
	"encoding/json"
	"net/http"
	"nsi_go/internal/scheme"
	"nsi_go/internal/storage"
)

type countryHandler struct {
	storage storage.Storage
}

func newCountryHandler(storage storage.Storage) countryHandler {
	return countryHandler{
		storage: storage,
	}
}

func (h countryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "":
		switch r.Method {
		case http.MethodGet:
			id := r.URL.Query().Get("id")
			result := h.storage.GetCountryById(id)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodPost:
			name := r.URL.Query().Get("name")
			result := h.storage.CreateCountry(name)
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		case http.MethodDelete:
			id := r.URL.Query().Get("id")
			h.storage.DeleteCountry(id)
		case http.MethodPatch:
			var element scheme.NSI
			dec := json.NewDecoder(r.Body)
			_ = dec.Decode(&element)
			h.storage.UpdateCountry(element.Id, element.Name)
		}
	case "all":
		switch r.Method {
		case http.MethodGet:
			result := h.storage.GetAllCountries()
			data, _ := json.Marshal(&result)
			_, _ = w.Write(data)
		}
	}
}
