package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    models "chatbot/models"
    utils "chatbot/utils"
)

func setupTest() (*utils.MockChatGPTGateway, *utils.MockRepository) {
    mockChatGPT := new(utils.MockChatGPTGateway)
    mockRepo := new(utils.MockRepository)

    return mockChatGPT, mockRepo
}

func TestChatHandler(t *testing.T) {

    tests := []struct {
        name              string
        interaction       models.Interaction
        chatGPTResponse   string
        chatGPTErr        error
        expectedStatus    int
        expectedResponse  string
    }{
        {
            name:            "Successful interaction",
            interaction:     models.Interaction{CustomerID: 1, Message: "Hello"},
            chatGPTResponse: "Hi there!",
            chatGPTErr:      nil,
            expectedStatus:  http.StatusOK,
            expectedResponse: "Hi there!",
        },
        {
            name:            "ChatGPT error",
            interaction:     models.Interaction{CustomerID: 2, Message: "Hello"},
            chatGPTResponse: "",
            chatGPTErr:      assert.AnError,
            expectedStatus:  http.StatusOK,
            expectedResponse: "Sorry, I couldn't process your request at the moment. Please try again later.",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockChatGPT, mockRepo := setupTest()
            mockChatGPT.On("GetChatGPTResponse", "test-api-key", tt.interaction.Message).Return(tt.chatGPTResponse, tt.chatGPTErr)
            mockRepo.On("SaveInteraction", tt.interaction).Return(nil)

            jsonData, _ := json.Marshal(tt.interaction)
            request, err := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonData))

            if err != nil {
                t.Fatal(err)
            }

            responseRecorder := httptest.NewRecorder()
            handler := ChatHandler(mockChatGPT, mockRepo)

            handler.ServeHTTP(responseRecorder, request)

            assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
            assert.Equal(t, tt.expectedResponse, responseRecorder.Body.String())
        })
    }
}
