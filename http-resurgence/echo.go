package main

import (
	"fmt"
	"io"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "must be post", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "body cannot be empty", 400)
		defer r.Body.Close()
		return
	}
	fmt.Fprintln(w, string(body))
}
