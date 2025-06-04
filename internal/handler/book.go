package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahulbhataniya/book-inventory-api/internal/config"
	"github.com/rahulbhataniya/book-inventory-api/internal/model"
)

// GetBooks returns all books from the database
func GetBooks(c *gin.Context) {
	var books []model.Book
	if err := config.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
		return
	}
	c.JSON(http.StatusOK, books)
}

// CreateBook creates a new book entry
func CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// DeleteBook deletes a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&model.Book{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}
