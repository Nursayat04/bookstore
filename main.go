package main

import (
	"Bookstore/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users", handlers.GetAuthors)
	r.POST("/users", handlers.AddAuthor)

	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.DELETE("/books/:id", handlers.DeleteBook)

	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.CreateCategory)
	r.Run("localhost:8080")
}
