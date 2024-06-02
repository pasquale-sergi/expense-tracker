package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pasquale-sergi/expense-tracker/user"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
		}

		err := user.RegisterUser(db, req.Username, req.Email, req.Password)
		var res RegisterResponse
		if err != nil {
			res = RegisterResponse{Success: false, Message: "Failed to register user"}
		} else {
			res = RegisterResponse{Success: true, Message: "Registration Completed"}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
