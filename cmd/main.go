package main

import (
	"fmt"
	"net/http"

	"github.com/allensuvorov/diary/internal/remote/handlers"
	"github.com/allensuvorov/diary/internal/remote/routers"
)

func main() {
	noteHandler := handlers.NewNoteHandler()
	r := routers.NewRouter(noteHandler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// http.HandleFunc("/", handlers.Hello)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
