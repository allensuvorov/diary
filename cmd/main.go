package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Handler - hello\n")
	now := time.Now()
	fmt.Fprintf(w, "Welcome to your daily diary!\nCurrent date and time: %s\n", now.Format(time.Layout))
	log.Print("Handler - buy\n")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
