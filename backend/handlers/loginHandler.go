package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/pasquale-sergi/expense-tracker/user"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func LoginHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := user.LoginUser(db, req.Username, req.Password)
		var resp LoginResponse
		if err != nil {
			resp = LoginResponse{Success: false, Message: "Invalid username or password"}
		} else {
			resp = LoginResponse{Success: true, Message: "Login successful"}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
