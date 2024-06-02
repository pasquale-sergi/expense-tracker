package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/pasquale-sergi/expense-tracker/database"
	"github.com/pasquale-sergi/expense-tracker/handlers"
	// "github.com/pasquale-sergi/expense-tracker/user"
)

func main() {
	db := database.DbConnection()
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler(db))
	mux.HandleFunc("/register", handlers.RegisterHandler(db))

	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":8090", handler))

}
