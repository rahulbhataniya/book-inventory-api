package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rahulbhataniya/book-inventory-api/internal/handler"
)

func RegisterRoutes(r *gin.Engine) {
	book := r.Group("/api/books")
	{
		book.POST("/", handler.CreateBook)
		book.GET("/", handler.GetBooks)
		book.GET("/:id", handler.GetBookByID)
		book.PUT("/:id", handler.UpdateBook)
		book.DELETE("/:id", handler.DeleteBook)
	}
}
