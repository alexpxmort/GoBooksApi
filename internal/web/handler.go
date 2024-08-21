package web

import (
	"fmt"
	"gobooks/internal/models"
	"gobooks/internal/service"
	"gobooks/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandlers struct {
	service *service.BookService
}

func (h *BookHandlers) CreateBook(c *gin.Context) {
	validate := validator.New()

	var book models.BookEntity
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := validate.Struct(book); err != nil {
		customErrors := utils.GetCustomErrorMessages(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": customErrors})
		return
	}

	if err := h.service.CreateBook(&book); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandlers) GetBooks(c *gin.Context) {
	books, err := h.service.GeAll()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get books"})
		return
	}

	c.JSON(http.StatusOK, books)
}
func (h *BookHandlers) GetBookById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := h.service.GetById(id)

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Book Not Found with ID %d", id)})
		return
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get book by id"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandlers) UpdateBook(c *gin.Context) {
	validate := validator.New()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var updatedBook models.BookEntity
	if err := c.BindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := validate.Struct(updatedBook); err != nil {
		customErrors := utils.GetCustomErrorMessages(err)
		c.JSON(http.StatusBadRequest, gin.H{"errors": customErrors})
		return
	}

	book, err := h.service.UpdateBook(id, &updatedBook)

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// MarkBookAsRead marca um livro como lido
func (h *BookHandlers) MarkBookAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := h.service.MarkBookAsRead(id)

	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update book to read"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func NewBookHandlers(service *service.BookService) *BookHandlers {
	return &BookHandlers{service: service}
}
