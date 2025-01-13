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

type NoteHandler struct {
}

func NewNoteHandler() NoteHandler {
	return NoteHandler{}
}

func (nh NoteHandler) Welcome(w http.ResponseWriter, r *http.Request) {
}

func (nh NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
}

func (nh NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
}
