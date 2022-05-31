package main

import (
	"net/http"
	"nsi_go/internal/muxer"
	"nsi_go/internal/storage"
)

func main() {
	mu := muxer.New(storage.New())
	http.ListenAndServe(":8080", mu.RootHandler())
}
