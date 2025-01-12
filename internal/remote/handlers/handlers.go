package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	log.Print("Handler - hello\n")
	now := time.Now()
	fmt.Fprintf(w, "Welcome to your daily diary!\nCurrent date and time: %s\n", now.Format(time.Layout))
	log.Print("Handler - buy\n")
}
