package controllers

import (
	"fmt"
    "encoding/json"
    "net/http"

    "chatbot/gateways"
    "chatbot/models"
    "chatbot/repositories"
)

const chatGPTAPIKey = "YOUR_CHAT_GPT_API_KEY"

func ChatHandler(chatGPTGateway gateways.ChatGPTGateway, repo repositories.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var interaction models.Interaction
        err := json.NewDecoder(r.Body).Decode(&interaction)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        response, err := chatGPTGateway.GetChatGPTResponse(chatGPTAPIKey, interaction.Message)
        if err != nil {
			fmt.Println("ERROR FROM CHAT GPT: ", err)
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
