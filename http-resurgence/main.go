package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/method-inspector", method)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/headers", header)
	log.Println("Server running on http:127.0.0.1:8080/")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
