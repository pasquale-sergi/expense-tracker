package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

func UserBalance(c *gin.Context) {

	var totalExpenses float64 = 0
	var totalIncome float64 = 0
	var balance float64

	databaseLogic.DB.Model(&Expense{}).Select("sum(amount) as total").Where("userid = ?", userID).Group("userid").First(&totalExpenses)
	fmt.Print("Total Expenses = ", totalExpenses)
	databaseLogic.DB.Model(&IncomeRecords{}).Select("sum(amount) as total").Where("userid = ?", userID).Group("userid").Order("total").Find(&totalIncome)
	fmt.Print("Total Income = ", totalIncome)

	balance = totalIncome - totalExpenses
	c.JSON(http.StatusOK, balance)

}
