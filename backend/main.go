package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
	"github.com/pasquale-sergi/expense-tracker/handlers"
)

func main() {
	db := databaseLogic.DbConnection()
	mux := http.NewServeMux()

	mux.HandleFunc("/login", handlers.LoginHandler(db))
	mux.HandleFunc("/register", handlers.RegisterHandler(db))
	mux.HandleFunc("/expenses", handlers.ExpensesHandler(db))

	// response, err := user.ShowExpensesOfTheMonth(db, 1)
	// if err != nil {
	// 	fmt.Printf("Error returning the expenses list: ", err)
	// }
	// fmt.Print(response)

	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":8090", handler))

}
