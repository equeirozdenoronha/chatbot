package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chatbot/gateways"
	"chatbot/models"
	"chatbot/repositories"
)

const aIAPIKey = "test-api-key"

func ChatHandler(aIGateway gateways.AIGateway, repo repositories.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var interaction models.Interaction
		err := json.NewDecoder(r.Body).Decode(&interaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := aIGateway.GetAIResponse(aIAPIKey, interaction.Message)
		if err != nil {
			fmt.Println("ERROR FROM AI: ", err)
			response = "Sorry, I couldn't process your request at the moment. Please try again later."
		}

		err = repo.SaveInteraction(interaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(response))
	}
}
