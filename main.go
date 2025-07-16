package main

import (
    "log"
    "net/http"

    "github.com/dkotegaonkar/internal-transfers/db"
    "github.com/dkotegaonkar/internal-transfers/handlers"
    "github.com/go-chi/chi/v5"
)

func main() {
    if err := db.InitDB(); err != nil {
        log.Fatal("DB connection failed: ", err)
    }

    r := chi.NewRouter()

    r.Post("/accounts", handlers.CreateAccount)
    r.Get("/accounts/{id}", handlers.GetAccount)
    r.Post("/transactions", handlers.CreateTransaction)

    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", r)
}
