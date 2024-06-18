package repositories

import (
	"database/sql"
    "log"

    _ "github.com/lib/pq"
    models "chatbot/models"
)

var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("postgres", "user=username dbname=review_chatbot sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
}

func CloseDB() {
    db.Close()
}

type Repository interface {
    SaveInteraction(interaction models.Interaction) error
    SaveReview(review models.Review) error
    SaveCustomer(customer models.Customer) error
}

type InteractionRepository struct{}

func (r *InteractionRepository) SaveInteraction(interaction models.Interaction) error {
    // Example Implementantion
	// _, err := db.Exec("INSERT INTO interactions (customer_id, message) VALUES ($1, $2)",
    //     interaction.CustomerID, interaction.Message)
    return nil
}

func (r *InteractionRepository) SaveReview(review models.Review) error {
    // Example Implementation
	// _, err := db.Exec("INSERT INTO reviews (customer_id, product_id, rating, review) VALUES ($1, $2, $3, $4)",
    //     review.CustomerID, review.ProductID, review.Rating, review.Review)
    return nil
}

func (r *InteractionRepository) SaveCustomer(customer models.Customer) error {
	// Example Implementation
    // _, err := db.Exec("INSERT INTO customers (name, email) VALUES ($1, $2)", customer.Name, customer.Email)
    return nil
}
