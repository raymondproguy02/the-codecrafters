package main

import (
	"fmt"
	"net/http"
)

func dasboard(w http.ResponseWriter, r *http.Request) {
	headers := r.Header.Get("X-API-Key")
	if headers != "secret123" {
		http.Error(w, "bad key", http.StatusUnauthorized)
	}
	fmt.Fprintln(w, "Welcome")
}