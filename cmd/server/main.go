package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/mod-ctf/internal/handlers"
	"github.com/yourusername/mod-ctf/pkg/auth"
)

func main() {
	r := mux.NewRouter()

	// Аутентификация middleware
	r.Use(auth.AuthMiddleware)

	// Этапы CTF
	r.HandleFunc("/start", handlers.StartHandler).Methods("GET")
	r.HandleFunc("/osint", handlers.OSINTHandler).Methods("POST")
	r.HandleFunc("/reverse", handlers.ReverseHandler).Methods("POST")
	r.HandleFunc("/websec", handlers.WebSecHandler).Methods("POST")
	r.HandleFunc("/moderate", handlers.ModerateHandler).Methods("POST")

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
