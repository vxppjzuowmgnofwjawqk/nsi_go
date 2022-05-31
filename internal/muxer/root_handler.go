package muxer

import (
	"net/http"
	"nsi_go/internal/storage"
)

type rootHandler struct {
	countryHandler
	languageHandler
	statusHandler
}

func newRootHandler(storage storage.Storage) rootHandler {
	return rootHandler{
		countryHandler:  newCountryHandler(storage),
		languageHandler: newLanguageHandler(storage),
		statusHandler:   newStatusHandler(storage),
	}
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "countries":
		h.countryHandler.ServeHTTP(w, r)
	case "languages":
		h.languageHandler.ServeHTTP(w, r)
	case "statuses":
		h.statusHandler.ServeHTTP(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
