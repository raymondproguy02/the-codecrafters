package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calculate(w http.ResponseWriter, r *http.Request) {
	op := r.URL.Query().Get("op")
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	con, err := strconv.Atoi(a)
	if err != nil {
		http.Error(w, "not intiger value", 400)
	}
	conv, err := strconv.Atoi(b)
	if err != nil {
		http.Error(w, "not intiger value", 400)
	}

	if op == "add" {
		fmt.Fprintln(w, con + conv)
	} 
	if op == "subtract" {
		fmt.Fprintln(w, con - conv)
	} 
	if op == "multiply" {
		fmt.Fprintln(w, con * conv)
	}
		http.Error(w, "400 Bad Request", 400)
}