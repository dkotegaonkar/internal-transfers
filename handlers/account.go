package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/dkotegaonkar/internal-transfers/db"
	"github.com/dkotegaonkar/internal-transfers/models"
	"github.com/go-chi/chi/v5"
)



// POST /accounts
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc models.Account
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Validate account ID and balance
	if acc.AccountID <= 0 {
		respondWithError(w, http.StatusBadRequest, "Account ID must be positive")
		return
	}
	if acc.Balance == "" {
		respondWithError(w, http.StatusBadRequest, "Balance is required")
		return
	}
	if _, err := strconv.ParseFloat(acc.Balance, 64); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid balance format")
		return
	}

	_, err := db.DB.Exec(`INSERT INTO accounts (account_id, balance) VALUES ($1, $2)`, acc.AccountID, acc.Balance)
	if err != nil {
		log.Println("DB Insert Error:", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create account")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GET /accounts/{id}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid account ID")
		return
	}

	var acc models.Account
	err = db.DB.Get(&acc, "SELECT * FROM accounts WHERE account_id = $1", id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Account not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(acc)
}














// package handlers

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"

// 	"github.com/dkotegaonkar/internal-transfers/db"
// 	"github.com/dkotegaonkar/internal-transfers/models"
// 	"github.com/go-chi/chi/v5"
// )

// func CreateAccount(w http.ResponseWriter, r *http.Request) {
//     var acc models.Account
//     if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
//         http.Error(w, "Invalid input", http.StatusBadRequest)
//         return
//     }

//     _, err := db.DB.Exec(`INSERT INTO accounts (account_id, balance) VALUES ($1, $2)`, acc.AccountID, acc.Balance)
//     if err != nil {
//     log.Println("DB Insert Error:", err)
//     http.Error(w, "Failed to create account", http.StatusInternalServerError)
//     return
// }
//     w.WriteHeader(http.StatusCreated)
// }

// func GetAccount(w http.ResponseWriter, r *http.Request) {
//     id := chi.URLParam(r, "id")
//     var acc models.Account
//     err := db.DB.Get(&acc, "SELECT * FROM accounts WHERE account_id = $1", id)
//     if err != nil {
//         http.Error(w, "Account not found", http.StatusNotFound)
//         return
//     }
//     json.NewEncoder(w).Encode(acc)
// }
