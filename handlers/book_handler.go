package handlers

import (
	"Bookstore/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var boks = []models.Book{
	{ID: 1, Title: "Interstellar", AuthorID: 1, CategoryID: 2, Price: 2500},
	{ID: 2, Title: "Django unchained", AuthorID: 2, CategoryID: 3, Price: 5700},
	{ID: 3, Title: "Spider-man", AuthorID: 3, CategoryID: 1, Price: 8000}}

func GetBooks(c *gin.Context) {
	categoryIDStr := c.Query("category_id")
	firstPg := c.DefaultQuery("L", "1")
	lastPg := c.DefaultQuery("R", "2")

	L, _ := strconv.Atoi(firstPg)
	R, _ := strconv.Atoi(lastPg)

	var filteredBooks []models.Book
	for _, b := range boks {
		if categoryIDStr == "" || fmt.Sprint(b.CategoryID) == categoryIDStr {
			filteredBooks = append(filteredBooks, b)
		}
	}

	startIndex := (L - 1) * R
	endIndex := startIndex + R

	if startIndex >= len(filteredBooks) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}

	if endIndex > len(filteredBooks) {
		endIndex = len(filteredBooks)
	}

	result := filteredBooks[startIndex:endIndex]
	c.JSON(http.StatusOK, result)
}
func GetBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, b := range boks {
		if fmt.Sprint(b.ID) == id {
			c.JSON(http.StatusOK, b)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title empty"})
		return
	}

	if book.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than 0"})
		return
	}

	book.ID = len(boks) + 1
	boks = append(boks, book)
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, b := range boks {
		if fmt.Sprint(b.ID) == id {
			updatedBook.ID = b.ID
			boks[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	for i, b := range boks {
		if fmt.Sprint(b.ID) == id {
			boks = append(boks[:i], boks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
}
