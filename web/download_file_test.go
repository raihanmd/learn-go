package main

import (
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//? Just simply add HTTP Header "Content-Disposition: attachment"
	w.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, "./uploads/test.jpg")
}

func TestDownloadFile(t *testing.T) {
	// TODO: Implement test
}
