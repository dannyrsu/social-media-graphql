package main

import (
	"log"
	"net/http"

	"github.com/dannyrsu/social-media-graphql/server"
)

func main() {
	handler := server.InitializeServer()
	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
