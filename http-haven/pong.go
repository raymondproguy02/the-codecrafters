package main

import (
	"fmt"
	"net/http"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong")
}
