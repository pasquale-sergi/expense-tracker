package main

import (
	"time"

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

	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // Specify your frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Allow credentials (cookies, authorization headers, etc.)
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/signup", user.Signup)
	r.POST("/login", user.Login)
	r.POST("/logout", user.Logout)

	// r.GET("/validate", middleware.RequireAuth, user.Validate)
	r.GET("/expensesHistory", middleware.RequireAuth, user.ListExpenses)
	r.GET("/currentExpenses", middleware.RequireAuth, user.ListExpensesCurrentMonth)
	r.POST("/addExpense", middleware.RequireAuth, user.AddExpense)
	r.GET("/expenses", middleware.RequireAuth, user.GetAllExpenses)

	r.Run()

	// mux.Handle("/expenses", handlers.ExpensesHandler(db))
	// mux.Handle("/expensesHistory", handlers.ExpenseHistoryHandler(db))

	// log.Fatal(http.ListenAndServe(":8090", handler))

}
