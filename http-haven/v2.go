package main

import (
	"fmt"
	"net/http"
)

func v2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to version 2")
}