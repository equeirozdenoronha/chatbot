package utils

import (
    "github.com/stretchr/testify/mock"
)

type MockChatGPTGateway struct {
    mock.Mock
}

func (m *MockChatGPTGateway) GetChatGPTResponse(apiKey string, message string) (string, error) {
    args := m.Called(apiKey, message)
    return args.String(0), args.Error(1)
}
