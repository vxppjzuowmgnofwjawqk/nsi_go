package muxer

import (
	"net/http"
	"nsi_go/internal/storage"
)

type muxer struct {
	rootHandler
}

func (mu *muxer) RootHandler() http.Handler {
	return mu.rootHandler
}

func New(storage storage.Storage) *muxer {
	return &muxer{
		rootHandler: newRootHandler(storage),
	}
}
