package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

func UserBalance(c *gin.Context) {

	totalExpenses := 0
	totalIncome := 0

	databaseLogic.DB.Model(&Expense{}).Select("sum(amount) as total").Where("userid = ?", userID).Group("userid").First(&totalExpenses)
	fmt.Print("Total Expenses = ", totalExpenses)
	databaseLogic.DB.Model(&IncomeRecords{}).Select("sum(amount) as total").Where("userid = ? AND amount >= 0", userID).Group("userid").Order("total").Find(&totalIncome)
	fmt.Print("Total Income = ", totalIncome)

	balance := totalIncome - totalExpenses
	c.JSON(http.StatusOK, balance)

}
