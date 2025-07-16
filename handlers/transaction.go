package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dkotegaonkar/internal-transfers/db"
	"github.com/dkotegaonkar/internal-transfers/models"
)



func CreateTransaction(w http.ResponseWriter, r *http.Request) {
    var t models.Transaction
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid input")
        return
    }

    if t.SourceAccountID == t.DestinationAccountID {
        respondWithError(w, http.StatusBadRequest, "Source and destination accounts must be different")
        return
    }

    amountFloat, err := strconv.ParseFloat(t.Amount, 64)
    if err != nil || amountFloat <= 0 {
        respondWithError(w, http.StatusBadRequest, "Invalid amount")
        return
    }

    tx, err := db.DB.Beginx()
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Transaction error")
        return
    }
    defer tx.Rollback()


    var sourceBalance float64
    err = tx.Get(&sourceBalance, "SELECT balance FROM accounts WHERE account_id = $1", t.SourceAccountID)
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Source account not found")
        return
    }
    if sourceBalance < amountFloat {
        respondWithError(w, http.StatusBadRequest, "Insufficient funds")
        return
    }


    var destExists bool
    err = tx.Get(&destExists, "SELECT EXISTS (SELECT 1 FROM accounts WHERE account_id = $1)", t.DestinationAccountID)
    if err != nil || !destExists {
        respondWithError(w, http.StatusBadRequest, "Destination account not found")
        return
    }


    _, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE account_id = $2", amountFloat, t.SourceAccountID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Failed to debit source account")
        return
    }

    _, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE account_id = $2", amountFloat, t.DestinationAccountID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Failed to credit destination account")
        return
    }

    _, err = tx.Exec(`
        INSERT INTO transactions (source_account_id, destination_account_id, amount)
        VALUES ($1, $2, $3)`,
        t.SourceAccountID, t.DestinationAccountID, amountFloat,
    )
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, "Failed to record transaction")
        return
    }

    if err = tx.Commit(); err != nil {
        respondWithError(w, http.StatusInternalServerError, "Commit failed")
        return
    }

    w.WriteHeader(http.StatusCreated)
}









