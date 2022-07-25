package main

import (
	"log"
	"net/http"

	"github.com/trite8q1/go-gorm/internal/repository/postgres"
	"github.com/trite8q1/go-gorm/internal/rest/routes"
)

func main() {
	router := routes.NewRouter()
	postgres.ConnectToDB()

	log.Printf("Server started on port %s\n", "8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
