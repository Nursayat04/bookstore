package handlers

import (
	"Bookstore/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var categories = []models.Category{
	{ID: 1, Name: "Fantasy"},
	{ID: 2, Name: "Science"},
	{ID: 3, Name: "Action"},
}

func GetCategories(c *gin.Context) {
	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if category.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category name is required"})
		return
	}

	category.ID = len(categories) + 1
	categories = append(categories, category)
	c.JSON(http.StatusCreated, category)
}
