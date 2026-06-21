package main

import (
	"fmt"
	"net/http"
)

func query(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not allowed", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprintln(w, "Hello, Guest!")
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
