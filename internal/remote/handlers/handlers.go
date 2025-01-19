package handlers

import (
	"fmt"
	"io"
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

func (nh NoteHandler) ShowNewNoteForm(w http.ResponseWriter, r *http.Request) {
	log.Print("ShowNewNoteForm - start\n")
	// 1. Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	io.WriteString(w, "new note form is coming here")
	// 2. Render the HTML form for creating a new note
	//   - You can use a template engine like `html/template`
	//     to render the form with any necessary data.
	// err := templates.ExecuteTemplate(w, "new_note.html", nil)
	// if err != nil {
	// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }
	log.Print("ShowNewNoteForm - end\n")
}

func (nh NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	log.Print("CreateNote Handler - start\n")

	// plain text
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	log.Printf("CreateNote Handler - r.Body %s", b)

	// data - JSON with field: text
	// validation: check note is not empty
	fmt.Fprintf(w, "Thank you for your note!\n")
	log.Print("CreateNote Handler - end\n")
}

func (nh NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	log.Print("GetNote Handler - start\n")
	fmt.Fprintf(w, "Here is your note!\n")
	log.Print("GetNote Handler - end\n")
}
