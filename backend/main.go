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
	r.POST("/addIncome", middleware.RequireAuth, user.AddIncome)
	r.GET("/incomeHistory", middleware.RequireAuth, user.IncomeHistory)
	r.GET("/balance", middleware.RequireAuth, user.UserBalance)
	r.GET("/categories", middleware.RequireAuth, user.GetCategories)
	r.POST("/addCategory", middleware.RequireAuth, user.AddCategory)
	r.GET("/username", middleware.RequireAuth, user.GetUsername)
	r.GET("/summaryByCategoryMonth", middleware.RequireAuth, user.GroupExpensesOfTheMonthByCategory)
	r.GET("/summaryByCategory", middleware.RequireAuth, user.GroupAllExpensesByCategory)
	r.POST("/fixedExpenses", middleware.RequireAuth, user.SetFixedExpenses)
	r.DELETE("/deleteExpense/:description", middleware.RequireAuth, user.DeleteExpense)

	r.Run()
}
