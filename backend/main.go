package main

import (
	"github.com/gin-contrib/cors"
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

	r.Use(cors.Default())

	r.POST("/signup", user.Signup)
	r.POST("/login", user.Login)

	r.GET("/validate", middleware.RequireAuth, user.Validate)
	r.GET("/expensesHistory", middleware.RequireAuth, user.ListExpenses)
	r.GET("/currentExpenses", middleware.RequireAuth, user.ListExpensesCurrentMonth)
	r.POST("/addExpense", user.AddExpense)

	r.Run()

	// mux.Handle("/expenses", handlers.ExpensesHandler(db))
	// mux.Handle("/expensesHistory", handlers.ExpenseHistoryHandler(db))

	// log.Fatal(http.ListenAndServe(":8090", handler))

}
