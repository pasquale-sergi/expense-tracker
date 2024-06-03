package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pasquale-sergi/expense-tracker/user"
)

func ExpensesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := 1
		expenses, err := user.ShowExpensesOfTheMonth(db, userID)

		if err != nil {
			fmt.Errorf("error returing the expenses: ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&expenses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}
