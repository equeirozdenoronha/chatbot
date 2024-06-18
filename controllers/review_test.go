package controllers

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/assert"
    "chatbot/models"
    "chatbot/utils"
)

func setupReviewTest() *utils.MockRepository {
    mockRepo := new(utils.MockRepository)

    return mockRepo
}
func TestReviewHandler(t *testing.T) {

    tests := []struct {
        name              string
        review            models.Review
        saveErr           error
        expectedStatus    int
        expectedResponse  string
    }{
        {
            name:            "Successful review",
            review:          models.Review{CustomerID: 1, ProductID: "iphone13", Rating: 5, Review: "Excellent product!"},
            saveErr:         nil,
            expectedStatus:  http.StatusOK,
            expectedResponse: "Thank you for your review!",
        },
        {
            name:            "Database save error",
            review:          models.Review{CustomerID: 2, ProductID: "iphone14", Rating: 4, Review: "Good product!"},
            saveErr:         assert.AnError,
            expectedStatus:  http.StatusInternalServerError,
            expectedResponse: "assert.AnError general error for testing\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockRepo := setupReviewTest()
            mockRepo.On("SaveReview", tt.review).Return(tt.saveErr)
            jsonData, _ := json.Marshal(tt.review)
            request, _ := http.NewRequest(http.MethodPost, "/review", bytes.NewBuffer(jsonData))

            responseRecorder := httptest.NewRecorder()
            handler := ReviewHandler(mockRepo)

            handler.ServeHTTP(responseRecorder, request)

            assert.Equal(t, tt.expectedStatus, responseRecorder.Code)
            assert.Equal(t, tt.expectedResponse, responseRecorder.Body.String())
        })
    }
}
