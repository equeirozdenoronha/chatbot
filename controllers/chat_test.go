package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "chatbot/models"
	utils "chatbot/utils"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*utils.MockAIGateway, *utils.MockRepository) {
	mockAI := new(utils.MockAIGateway)
	mockRepo := new(utils.MockRepository)

	return mockAI, mockRepo
}

func TestChatHandler(t *testing.T) {

	tests := []struct {
		name             string
		interaction      models.Interaction
		aIResponse       string
		aIErr            error
		expectedStatus   int
		expectedResponse string
	}{
		{
			name:             "Successful interaction",
			interaction:      models.Interaction{CustomerID: 1, Message: "Hello"},
			aIResponse:       "Hi there!",
			aIErr:            nil,
			expectedStatus:   http.StatusOK,
			expectedResponse: "Hi there!",
		},
		{
			name:             "AI error",
			interaction:      models.Interaction{CustomerID: 2, Message: "Hello"},
			aIResponse:       "",
			aIErr:            assert.AnError,
			expectedStatus:   http.StatusOK,
			expectedResponse: "Sorry, I couldn't process your request at the moment. Please try again later.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAI, mockRepo := setupTest()
			mockAI.On("GetAIResponse", "test-api-key", tt.interaction.Message).Return(tt.aIResponse, tt.aIErr)
			mockRepo.On("SaveInteraction", tt.interaction).Return(nil)

			jsonData, _ := json.Marshal(tt.interaction)
			request, err := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonData))

			if err != nil {
				t.Fatal(err)
			}

			responseRecorder := httptest.NewRecorder()
			handler := ChatHandler(mockAI, mockRepo)

			handler.ServeHTTP(responseRecorder, request)

			assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
			assert.Equal(t, tt.expectedResponse, responseRecorder.Body.String())
		})
	}
}
