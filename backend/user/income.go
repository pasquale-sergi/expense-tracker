package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

type IncomeRecords struct {
	id          uint `gorm: "primaryKey"`
	Amount      float64
	Description string
	Date        time.Time
	Userid      uint
}

func AddIncome(c *gin.Context) {
	// databaseLogic.DB.AutoMigrate(&Income{})
	var body struct {
		Amount      float64
		Description string
		Date        string
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read the addIncome body",
		})
	}

	layout := "01/02/2006"

	parsedDate, err := time.Parse(layout, body.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Print("BODY: ", &body)
	newIncome := IncomeRecords{
		Amount:      body.Amount,
		Description: body.Description,
		Date:        parsedDate,
		Userid:      uint(userID),
	}
	fmt.Print(&newIncome)

	result := databaseLogic.DB.Create(&newIncome)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add expense",
		})
		return
	}

	//Response
	c.JSON(http.StatusOK, gin.H{"message": "Income record added"})
}

func IncomeHistory(c *gin.Context) {
	var incomes []IncomeRecords
	res := databaseLogic.DB.Where("userID = ?", userID).Find(&incomes)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
	}

	c.JSON(http.StatusOK, incomes)

}
