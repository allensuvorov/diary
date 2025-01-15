package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/allensuvorov/diary/internal/remote/handlers"
	"github.com/allensuvorov/diary/internal/remote/routers"
)

func main() {
	noteHandler := handlers.NewNoteHandler()
	r := routers.NewRouter(noteHandler)
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
