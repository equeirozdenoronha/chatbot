package models

type Interaction struct {
    CustomerID int    `json:"customer_id"`
    Message    string `json:"message"`
}

type Review struct {
    CustomerID int    `json:"customer_id"`
    ProductID  string `json:"product_id"`
    Rating     int    `json:"rating"`
    Review     string `json:"review"`
}

type Customer struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
