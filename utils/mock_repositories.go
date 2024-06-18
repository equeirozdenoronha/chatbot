package utils

import (
    "chatbot/models"
    "github.com/stretchr/testify/mock"
)

type MockRepository struct {
    mock.Mock
}

func (m *MockRepository) SaveInteraction(interaction models.Interaction) error {
    args := m.Called(interaction)
    return args.Error(0)
}

func (m *MockRepository) SaveReview(review models.Review) error {
    args := m.Called(review)
    return args.Error(0)
}

func (m *MockRepository) SaveCustomer(customer models.Customer) error {
    args := m.Called(customer)
    return args.Error(0)
}
