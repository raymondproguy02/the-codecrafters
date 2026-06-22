package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter, r *http.Request) {
	custom := r.Header.Get("X-Custom-Token")
	if custom == "" {
		http.Error(w, "X-Custom-Token header is missing", 400)
		return
	}
	fmt.Fprintf(w, "Token received: %s", custom)

	content := r.Header.Get("Content-Type")
	if content == "" {
		fmt.Fprintf(w, "Content-Type %s", content)
		return
	}
	fmt.Fprintln(w, "Content-Type not provided")
}
