package controllers

import (
	"gin-gonic/models"
	"gin-gonic/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type BookController struct {
	bookService *services.BookService
}

func NewBookController(bookService *services.BookService) *BookController {
	return &BookController{bookService: bookService}
}

func (bc *BookController) CreateBook(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	newBookReq := models.CreateBookRequest{UserID: userID}

	if err := c.ShouldBindJSON(&newBookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook, err := bc.bookService.CreateBook(&newBookReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	updateBookReq := models.UpdateBookRequest{}
	if err := c.ShouldBindJSON(&updateBookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := bc.bookService.UpdateBook(&updateBookReq, userID)
	if err != nil {
		if strings.HasPrefix(err.Error(), "unauthorized") {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to update this book"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := bc.bookService.DeleteBook(uint(bookID), userID); err != nil {
		if strings.HasPrefix(err.Error(), "unauthorized") {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to delete this book"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := bc.bookService.GetBookByID(uint(bookID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bc *BookController) GetBooksByAuthor(c *gin.Context) {
	author := c.Query("author")
	if author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author query parameter is required"})
		return
	}

	books, err := bc.bookService.GetBooksByAuthor(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBooksByTitle(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title query parameter is required"})
		return
	}

	books, err := bc.bookService.GetBooksByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}
