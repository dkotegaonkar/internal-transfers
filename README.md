# 🏦 Internal Transfers System

A backend system for processing internal financial transactions between accounts using **Go** and **PostgreSQL**. Exposes RESTful endpoints for creating accounts, querying balances, and securely transferring funds between accounts.

---

## 📌 Project Overview

This project implements a robust, lightweight internal transfer service using Go. It provides HTTP endpoints for:

- Creating accounts
- Querying account balances
- Processing internal transactions atomically

It ensures data integrity, consistent state transitions, and structured error handling.

---

## 🗂️ Project Structure

internal-transfers/
├── go.mod # Go module definition
├── main.go # Application entry point
├── db/
│ └── db.go # DB connection initialization
├── handlers/
│ ├── account.go # Account handlers
│ ├── transaction.go # Transaction handler
│ └── utils.go # Shared error handling function
├── models/
│ └── models.go # Structs for DB and JSON mapping
└── README.md # Project documentation

---

## ⚙️ Setup Instructions

### 🧱 1. Install Dependencies

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

Here's the updated section of your `README.md`, incorporating the use of **pgAdmin 4** for creating the database and **psql CLI** for running the schema:

---

## 🗄️ 2. Create Database Schema

### 🔧 Tools Used

- **pgAdmin 4**: Used to create the database `transfers`.
- **psql CLI (Command Line Interface)**: Used to execute SQL commands and implement the schema.

### 📘 Steps to Setup the Schema

1. Open **pgAdmin 4** and:

   - Create a new database named: `internal_transfers`
   - Confirm connection with user `postgres`

2. Open your terminal and launch the **psql command-line tool**:

```bash
psql -U postgres -d transfers
```

3. Inside the `psql` shell, run the following schema:

```sql
CREATE TABLE accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(20, 5) NOT NULL
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    source_account_id BIGINT,
    destination_account_id BIGINT,
    amount NUMERIC(20, 5) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (source_account_id) REFERENCES accounts(account_id),
    FOREIGN KEY (destination_account_id) REFERENCES accounts(account_id)
);
```

▶️ Run the Server
git clone https://github.com/dkotegaonkar/internal-transfers.git
cd internal-transfers
go mod tidy

4.  In your db/db.go file, configure the connection string using your database credentials:

## connStr := "host=localhost port=5432 user=your_user password=your_password dbname=internal_transfers sslmode=disable"

go run main.go
Server will run on:
http://localhost:8080

🧪 API Testing (via Postman)
✅ Create Account
Endpoint: POST /accounts
{
"account_id": 1001,
"balance": "500.00"
}
![Create Account](assets/account_create.png)

📤 Get Account Balance
Endpoint: GET /accounts/1001
Response:
{
"account_id": 1001,
"balance": "500.00000"
}
(assets/get_balance.png)

Endpoint: GET /accounts/1002
Response:
{
"account_id": 1001,
"balance": "500.00000"
}

🔁 Create Transaction
Endpoint: POST /transactions
{
"source_account_id": 1001,
"destination_account_id": 1002,
"amount": "100.00"
}
(assets/transaction1.png)

(assets/transaction2.png)

✅ Key Features
| Feature | Description |
| ------------------------ | ------------------------------------------------------- |
| **Clean endpoints** | Exposes `/accounts` and `/transactions` with validation |
| **Atomic operations** | Uses SQL transaction to ensure atomic debit-credit |
| **Validation** | Ensures amount > 0 and account existence |
| **JSON error responses** | Consistent and readable errors |
| **Modular code** | Separation of concerns: models, handlers, DB, utils |

| Assessment Goal           | How Addressed                                             |
| ------------------------- | --------------------------------------------------------- |
| HTTP endpoint correctness | All endpoints are tested and return appropriate responses |
| Accurate processing       | SQL transactions ensure consistency                       |
| Code quality              | Modular, validated inputs, reusable helpers               |
| Documentation             | This `README.md` + inline comments                        |
| Testing                   | Manual test cases outlined via Postman                    |

📌 Assumptions
All accounts use the same currency

No authentication or authorization implemented

balance and amount are passed as strings in JSON to maintain precision

📬 Author
Dhruv Kotegaonkar
GitHub: @dkotegaonkar
