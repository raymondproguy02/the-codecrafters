package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func status(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("code")
	if text == "" || len(text) == 0 {
		http.Error(w, "code parameter is required", 400)
		return
	}
	conv, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "code must be a valid integer", 400)
		return
	}
	if conv < 100 || conv > 599 {
		http.Error(w, "code must be a valid HTTP status code (100–599)", 400)
		return
	}
	w.WriteHeader(conv)
	w.WriteHeader(conv)
	fmt.Fprintf(w, "Responding with status [%s]", conv)
}
