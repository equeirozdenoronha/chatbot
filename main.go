package main

import (
	"log"
	"net/http"

	"chatbot/controllers"
	"chatbot/gateways"
	"chatbot/repositories"
)

func main() {
	// Initialize dependencies
	aIGateway := &gateways.RealAIGateway{}
	repo := &repositories.InteractionRepository{}

	// Define HTTP handlers
	http.HandleFunc("/chat", controllers.ChatHandler(aIGateway, repo))
	http.HandleFunc("/review", controllers.ReviewHandler(repo))

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
