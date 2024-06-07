package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

//i want to show the current month expenses per category

func GroupExpensesOfTheMonthByCategory(c *gin.Context) {
	var expenses []struct {
		CategoryName string  `json:"category_name"`
		TotalAmount  float64 `json:"total_amount"`
	}

	month := c.Query("month")
	if month == "" {
		month = time.Now().Month().String()
	}

	parsedMonth, err := time.Parse("January", month)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
		return
	}

	now := time.Now()
	startOfMonth := time.Date(now.Year(), parsedMonth.Month(), 1, 0, 0, 0, 0, now.Location())
	startOfNextMonth := startOfMonth.AddDate(0, 1, 0)
	endOfMonth := startOfNextMonth.Add(-time.Nanosecond)

	res := databaseLogic.DB.
		Table("expenses").
		Select("categoryname as category_name, sum(amount) as total_amount").
		Where("userid = ? AND date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).
		Group("categoryname").
		Scan(&expenses)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
		return
	}

	c.JSON(http.StatusOK, expenses)
}

func GroupAllExpensesByCategory(c *gin.Context) {
	var expenses []struct {
		CategoryName string  `json:"category_name"`
		TotalAmount  float64 `json:"total_amount"`
	}

	res := databaseLogic.DB.Table("expenses").Select("categoryname as category_name, sum(amount) as total_amount").Where("userid = ?", userID).Group("categoryname").Find(&expenses)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error})
	}

	c.JSON(http.StatusOK, expenses)
}
