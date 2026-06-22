package main

import (
	"net/http"
)

func legacy(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://127.0.0.1:8080/v2", http.StatusMovedPermanently)
}