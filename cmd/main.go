package main

import (
	"fmt"
	"net/http"

	"github.com/allensuvorov/diary/internal/remote/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
