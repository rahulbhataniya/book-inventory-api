# Health check
healthcheck:
	@echo "Checking health..."
	@curl -s http://localhost:8080/health || echo "Health check failed"

CreaBook:
	@echo "Creating a book..."
	@curl -X POST http://localhost:8080/books \
		-H "Content-Type: application/json" \
		-d '{"title":"Clean Code", "author":"Robert C. Martin", "year":2008}'

GetBooks:
	@echo "Fetching all books..."
	@curl -s http://localhost:8080/books


# Delete a book (use ID from list)
DeleteBook:
	@echo "Deleting a book..."
	@# Replace '1' with the actual ID of the book you want to delete
	@curl -X DELETE http://localhost:8080/books/1 || echo "Delete failed, book not found"

