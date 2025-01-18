package routers

import (
	"github.com/allensuvorov/diary/internal/remote/handlers"
	"github.com/go-chi/chi"
)

func NewRouter(note handlers.NoteHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/", note.Welcome)
	r.Get("/notes/{noteId}", note.GetNote)    // Display a specific note
	r.Get("/notes/new", note.ShowNewNoteForm) // Display a form for creating a new note
	r.Post("/notes", note.CreateNote)
	return r
}
