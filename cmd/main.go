package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Workout Tracker!")
	})
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Hello, Workout Tracker!")
}