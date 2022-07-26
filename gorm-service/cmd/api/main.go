package main

import (
	"log"
	"net/http"

	"github.com/trite8q1/go-gorm/internal/config"
	"github.com/trite8q1/go-gorm/internal/rest/routes"
)

func main() {
	config.NewConfig()
	log.Println("Starting server...")
	router := routes.NewRouter()

	log.Printf("Server started on port %s\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
