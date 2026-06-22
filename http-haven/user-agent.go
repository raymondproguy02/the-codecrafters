package main

import (
	"fmt"
	"net/http"
)

func agent(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("User-Agent")
	if user == "" {
		http.Error(w, "can't be empty", 400)
		return
	}
	fmt.Fprintln(w, user)
}