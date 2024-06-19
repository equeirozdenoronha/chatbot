package gateways

import (
	utils "chatbot/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAIResponse(t *testing.T) {
	// Mock API key and message
	apiKey := "test-api-key"
	message := "Hello, how are you?"
	expectedResponse := "Hi there!"

	mockAI := new(utils.MockAIGateway)
	mockAI.On("GetAIResponse", apiKey, message).Return(expectedResponse, nil)
	// Mock response from AI
	response, _ := mockAI.GetAIResponse(apiKey, message)

	assert.Equal(t, expectedResponse, response)
}
