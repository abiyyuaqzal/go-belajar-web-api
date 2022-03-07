package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abiyyuaqzal/go-belajar-web-api/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, getBook := range books {
		bookResponse := convertToBookResponse(getBook)
		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(200, gin.H{
		"Data": booksResponse,
	})
}

func (h *bookHandler) GetBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	getBook, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(getBook)

	c.JSON(200, gin.H{
		"Data": bookResponse,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"Data": convertToBookResponse(book)})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {
	var bookR book.BookRequest

	err := c.ShouldBindJSON(&bookR)
	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	book, err := h.bookService.Update(id, bookR)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{"Data": convertToBookResponse(book)})
}

func (h *bookHandler) DeleteBookHandler(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	getBook, err := h.bookService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := convertToBookResponse(getBook)

	c.JSON(200, gin.H{
		"Data": bookResponse,
	})
}

func convertToBookResponse(theBook book.Book) book.BookResponse {
	return book.BookResponse{
		ID:          theBook.ID,
		Title:       theBook.Title,
		Description: theBook.Description,
		Price:       theBook.Price,
		Rating:      theBook.Rating,
	}
}
