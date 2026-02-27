package controllers

import (
	"go-crud-backend/models"
	"go-crud-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetBooks godoc
// @Summary      Get all books
// @Description  Responds with the list of all books as JSON.
// @Tags         books
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, services.GetAllBooks())
}

// GetBookByID godoc
// @Summary      Get a book by ID
// @Description  Returns a single book based on the ID parameter.
// @Tags         books
// @Produce      json
// @Param        id   path      string  true  "Book ID"
// @Success      200  {object}  models.Book
// @Failure      404  {object}  nil
// @Router       /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := services.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook godoc
// @Summary      Add a new book
// @Description  Takes a JSON input and adds it to the book collection.
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      models.Book  true  "Add book"
// @Success      201   {object}  models.Book
// @Router       /books [post]
func CreateBook(c *gin.Context) {
	var newBook models.Book

	// 1. Bind JSON to the struct
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// 2. Call the service (which now returns an error for validation/uniqueness)
	result, err := services.CreateBook(newBook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// UpdateBook godoc
// @Summary      Update an existing book
// @Description  Update the details of a book by its ID
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "Book ID"
// @Param        book  body      models.UpdateBookInput  true  "Updated book object"
// @Success      200   {object}  models.Book
// @Failure      400   {object}  nil "Invalid input"
// @Failure      404   {object}  nil "Book not found"
// @Router       /books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook models.UpdateBookInput

	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	book, err := services.UpdateBook(id, updatedBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary      Delete a book
// @Description  Remove a book from the store by its ID
// @Tags         books
// @Param        id   path      string  true  "Book ID"
// @Success      204  {object}  nil     "No Content"
// @Failure      404  {object}  nil "Book not found"
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
