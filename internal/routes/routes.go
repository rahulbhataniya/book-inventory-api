package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulbhataniya/book-inventory-api/internal/handler"
)

// SetupRouter initializes the routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// Book routes
	r.GET("/books", handler.GetBooks)
	r.GET("/books/:id", handler.GetBook)
	r.POST("/books", handler.CreateBook)
	r.DELETE("/books/:id", handler.DeleteBook)

	return r
}
