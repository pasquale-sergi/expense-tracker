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
		userID := 2
		expenses, err := user.ShowExpensesOfTheMonth(db, userID)

		if err != nil {
			fmt.Print("error returing the expenses")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&expenses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}

	}
}

func ExpenseHistoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		userID := session.Values["userID"].(int)

		expenses, err := user.ShowExpensesHistory(db, userID)
		if err != nil {
			fmt.Print("Error returning expense history")
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&expenses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}

func ExpensesHistoryHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := Store.Get(r, "session")
		if err != nil {
			fmt.Println("Error retrieving session:", err)
			http.Error(w, "Error retrieving session", http.StatusInternalServerError)
			return
		}

		fmt.Println("Session values during expenses retrieval:", session.Values)

		userID, ok := session.Values["userID"].(int)
		if !ok {
			http.Error(w, "User ID not found in session", http.StatusUnauthorized)
			return
		}

		// Use userID in database queries to retrieve user-specific data
		expenses, err := user.GetExpensesByUserID(db, userID)
		if err != nil {
			http.Error(w, "Failed to fetch expenses", http.StatusInternalServerError)
			return
		}

		// Respond with fetched expenses
		json.NewEncoder(w).Encode(expenses)
	}
}
