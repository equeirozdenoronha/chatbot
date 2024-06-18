package gateways

import (
    "testing"
	"github.com/stretchr/testify/assert"
	utils "chatbot/utils"
)

func TestGetChatGPTResponse(t *testing.T) {
    // Mock API key and message
    apiKey := "test-api-key"
    message := "Hello, how are you?"
	expectedResponse := "Hi there!"

	mockChatGPT := new(utils.MockChatGPTGateway)
	mockChatGPT.On("GetChatGPTResponse", apiKey, message).Return(expectedResponse, nil)
    // Mock response from ChatGPT
    response, _ := mockChatGPT.GetChatGPTResponse(apiKey, message)
    
	assert.Equal(t, expectedResponse, response)
}
