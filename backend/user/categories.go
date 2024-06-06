package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pasquale-sergi/expense-tracker/databaseLogic"
)

type Category struct {
	categoryid uint `gorm: "primaryKey"`
	Name       string
}

func GetCategories(c *gin.Context) {
	var categories []Category

	result := databaseLogic.DB.Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, categories) // Return all expenses data
}

func AddCategory(c *gin.Context) {
	var body struct {
		Name string
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	newCategory := Category{
		Name: body.Name,
	}

	res := databaseLogic.DB.Create(&newCategory)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed add new category, check if already exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "category added correctly"})
}
