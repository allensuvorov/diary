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
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
