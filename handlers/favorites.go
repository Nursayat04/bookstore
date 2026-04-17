package handlers

import (
	"Bookstore/config"
	"Bookstore/middleware"
	"Bookstore/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var favorites []models.Favorite
	var total int64

	config.DB.Where("user_id = ?", userID).Model(&models.Favorite{}).Count(&total)
	result := config.DB.Where("user_id = ?", userID).
		Preload("Book").
		Limit(limit).Offset(offset).
		Find(&favorites)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch favorites"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       favorites,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func AddFavorite(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id")) // ← исправлено
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	fav := models.Favorite{UserID: userID, BookID: uint(bookID)}
	result := config.DB.Where(fav).FirstOrCreate(&fav)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add favorite"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to favorites"})
}

func RemoveFavorite(c *gin.Context) {
	userID, ok := middleware.GetUserIDFromToken(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("id")) // ← исправлено
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	result := config.DB.
		Where("user_id = ? AND book_id = ?", userID, bookID).
		Delete(&models.Favorite{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Favorite not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Removed from favorites"})
}
