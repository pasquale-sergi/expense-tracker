package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
	"github.com/pasquale-sergi/expense-tracker/middleware"
	"github.com/pasquale-sergi/expense-tracker/user"
)

func init() {
	databaseLogic.LoadEnvVariables()
	databaseLogic.DbConnection()
}

func main() {

	r := gin.Default()

	r.POST("/signup", user.Signup)
	r.POST("/login", user.Login)

	r.GET("/validate", middleware.RequireAuth, user.Validate)

	r.Run()

	// mux.Handle("/expenses", handlers.ExpensesHandler(db))
	// mux.Handle("/expensesHistory", handlers.ExpenseHistoryHandler(db))

	// log.Fatal(http.ListenAndServe(":8090", handler))

}
