package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rahulbhataniya/book-inventory-api/internal/config"
	"github.com/rahulbhataniya/book-inventory-api/internal/model"
)

// CreateBook handles POST /api/books/
func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// GetBooks handles GET /api/books/
func GetBooks(c *gin.Context) {
	var books []model.Book
	config.DB.Find(&books)
	c.JSON(http.StatusOK, books)
}

// GetBookByID handles GET /api/books/:id
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook handles PUT /api/books/:id
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&book)
	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /api/books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	config.DB.Delete(&model.Book{}, idInt)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
