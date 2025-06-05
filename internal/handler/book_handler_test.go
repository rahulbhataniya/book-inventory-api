package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"

	"strings"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"

	"github.com/rahulbhataniya/book-inventory-api/internal/config"
	"github.com/rahulbhataniya/book-inventory-api/internal/model"
)

func TestHealthCheck(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200 but got %d", w.Code)
	}
}

func setupTestRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/api/books", CreateBook)
	return r
}

func TestCreateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Setup in-memory SQLite DB
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}

	// Migrate schema
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		t.Fatalf("Failed to migrate DB schema: %v", err)
	}

	// Assign global config DB
	config.DB = db

	// Setup router
	router := setupTestRouter()

	// Test JSON (matching model.Book struct)
	bookJSON := `{
		"title": "Test Book",
		"author": "Test Author",
		"category": "Fiction",
		"published": "2023"
	}`

	req, _ := http.NewRequest("POST", "/api/books", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check response status and print body if failure
	if w.Code != http.StatusCreated {
		t.Fatalf("Expected status 201 but got %d. Response body: %s", w.Code, w.Body.String())
	}

	// Optional: Check response body contains title
	expected := `"title":"Test Book"`
	if !strings.Contains(w.Body.String(), expected) {
		t.Errorf("Expected response body to contain %s, got %s", expected, w.Body.String())
	}

	// Verify book inserted in DB
	var book model.Book
	result := db.First(&book, "title = ?", "Test Book")
	if result.Error != nil {
		t.Errorf("Book not found in DB: %v", result.Error)
	}
}
