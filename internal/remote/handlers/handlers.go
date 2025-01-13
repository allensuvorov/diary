package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type NoteHandler struct {
}

func NewNoteHandler() NoteHandler {
	return NoteHandler{}
}

func (nh NoteHandler) Welcome(w http.ResponseWriter, r *http.Request) {
	log.Print("Welcome Handler - start\n")
	now := time.Now()
	fmt.Fprintf(w, "Welcome to your daily diary!\nCurrent date and time: %s\n", now.Format(time.Layout))
	log.Print("Welcome Handler - end\n")
}

func (nh NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
}

func (nh NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
}
