package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

type Expense struct {
	expenseid    uint `gorm: "primaryKey"`
	Userid       uint
	Categoryname string
	Amount       float64
	Description  string
	Date         time.Time
}

func ListExpenses(c *gin.Context) {

	var expenses []Expense

	result := databaseLogic.DB.Where("userid = ?", userID).Find(&expenses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses) // Return all expenses data
}

func ListExpensesCurrentMonth(c *gin.Context) {
	// Get the current time
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Determine the start of the next month and then subtract 1 nanosecond to get the end of the current month
	startOfNextMonth := startOfMonth.AddDate(0, 1, 0)
	endOfMonth := startOfNextMonth.Add(-time.Nanosecond)

	// Create a slice to hold the existing expenses
	var expenses []Expense
	var tempExpenses []Expense

	// Query the database
	fmt.Print("USERID INTO THE EXPENSE FUNC : ", userID)
	result := databaseLogic.DB.Where("userid = ? AND date BETWEEN ? AND ?", userID, startOfMonth, endOfMonth).Find(&tempExpenses)

	// Check for errors
	if result.Error != nil {
		fmt.Println("Failed to retrieve records:", result.Error)
		return
	}

	expenses = append(expenses, tempExpenses...)
	fmt.Print(expenses)
	c.JSON(http.StatusOK, expenses)

}

func AddExpense(c *gin.Context) {
	var body struct {
		Amount       float64
		Categoryname string
		Description  string
		Date         string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	layout := "01/02/2006"

	parsedDate, err := time.Parse(layout, body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	newExpense := Expense{
		Userid:       uint(userID),
		Categoryname: body.Categoryname,
		Amount:       body.Amount,
		Description:  body.Description,
		Date:         parsedDate,
	}

	result := databaseLogic.DB.Create(&newExpense)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add expense",
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{})
}

func GetAllExpenses(c *gin.Context) {
	var expenses []Expense

	result := databaseLogic.DB.Find(&expenses)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, expenses) // Return all expenses data
}
