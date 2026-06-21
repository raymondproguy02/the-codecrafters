package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	http.HandleFunc("/hello", query)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/calculate", calculate)
	log.Println("Server running on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
