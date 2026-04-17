package main

import (
	"Bookstore/config"
	"Bookstore/handlers"
	"Bookstore/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// Authors
	r.GET("/authors", handlers.GetAuthors)
	r.POST("/authors", handlers.AddAuthor)

	// Categories
	r.GET("/categories", handlers.GetCategories)
	r.POST("/categories", handlers.AddCategory)

	// Books
	r.GET("/books", handlers.GetBooks)
	r.POST("/books", handlers.CreateBook)
	r.GET("/books/:id", handlers.GetBookByID)
	r.PUT("/books/:id", handlers.UpdateBook)
	r.DELETE("/books/:id", handlers.DeleteBook)

	// Favorites
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/books/favorites", handlers.GetFavorites)
		auth.PUT("/books/:id/favorites", handlers.AddFavorite)
		auth.DELETE("/books/:id/favorites", handlers.RemoveFavorite)
	}

	r.Run(":8080")
}
