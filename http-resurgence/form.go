package main

import (
	"fmt"
	"net/http"
)

func form(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "not post", 405)
		return
	}
	r.ParseForm()
	name := r.FormValue("username")
	lang := r.FormValue("language")

	if name == "" {
		http.Error(w, "username is required", 400)
		return
	}
	if lang == "" {
		http.Error(w, "language is required", 400)
		return
	}
	fmt.Fprintf(w, "Hello [%s], you are coding in [%s]", name, lang)
}
