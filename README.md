# Chatbot Challenge

This project implements a review chatbot for an e-commerce application. The chatbot's main task is to gather customer reviews after a product delivery is confirmed. It also engages users in other relevant conversations, providing a more interactive and personalized user experience.

## Code Structure

```
chatbot/
├── controllers/
│ ├── chat_controller.go
│ ├── chat_controller_test.go
│ ├── review_controller.go
│ ├── review_controller_test.go
├── gateways/
│ ├── ai_gateway.go
├── models/
│ ├── interaction.go
│ ├── review.go
│ ├── customer.go
├── repositories/
│ ├── interaction_repository.go
│ ├── interaction_repository_test.go
│ ├── review_repository.go
│ ├── review_repository_test.go
│ └── customer_repository.go
│ └── customer_repository_test.go
├── utils/
│ ├── mock_ai_gateway.go
│ ├── mock_repository.go
├── main.go
├── main_test.go
└── go.mod
```

## Endpoints

### `/chat`

- **Method**: `POST`
- **Description**: Handles chat interactions. It receives a message from the customer and responds using the ai API.
- **Request Body**:
  ```json
  {
    "customer_id": 1,
    "message": "Hello"
  }
  ```

- **Response Body**:

    ```
        AI Message
    ```
### `/review`
- **Method**: `POST`
- **Description**: Handles review submissions. It saves the customer's review for a product.
- **Request Body**:
```json
    {
    "customer_id": 1,
    "product_id": "iphone13",
    "rating": 5,
    "review": "Excellent product!"
    }
```
- **Response Body**:

    ```
        "Thank you for your review!"
    ```

## Setup and Running

### Prerequisites
- Go 1.16+
- PostgreSQL

### Environment Variables

```
ai_API_KEY: Your ai API key.
```

### Database Setup

1. Initialize PostgreSQL database

```sql
CREATE DATABASE review_chatbot;
```

2. Create tables

```sql
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE interactions (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    message TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    product_id VARCHAR(50),
    rating INT,
    review TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Running the Application

1. Install dependencies:

```sh
go mod tidy
```

2. Run the application:
```
sh
go run main.go
```

### Running Tests

To run all tests, use the following command:

```sh
go test ./...
```



