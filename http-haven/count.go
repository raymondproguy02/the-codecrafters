package main

import (
	"fmt"
	"io"
	"net/http"
)

func counter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "Send a POST request with text to count words")
	}
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintln(w, len(string(body)))
	}
}
