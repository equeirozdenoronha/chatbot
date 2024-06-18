package controllers

import (
    "encoding/json"
    "net/http"

    "chatbot/models"
    "chatbot/repositories"
)

func ReviewHandler(repo repositories.Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var review models.Review
        err := json.NewDecoder(r.Body).Decode(&review)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        err = repo.SaveReview(review)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write([]byte("Thank you for your review!"))
    }
}
