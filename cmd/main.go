package main

import (
	"log"
	"net/http"

	domain "github.com/figarocms/fi-chatgpt-api/internal/domain/api"
)

func main() {
	http.HandleFunc("/ask", domain.Question)
	log.Fatal(http.ListenAndServe(":4200", nil))
}
