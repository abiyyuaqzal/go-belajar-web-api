package main

import (
	"fmt"
	"log"

	"github.com/abiyyuaqzal/go-belajar-web-api/book"
	"github.com/abiyyuaqzal/go-belajar-web-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}
	fmt.Println("Db connection success")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/books/:id", bookHandler.GetBookHandler)
	v1.POST("/books", bookHandler.PostBooksHandler)
	v1.PUT("/books/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	router.Run()

	// main
	// handler
	// service
	// repository
	// db
	// mysql
}
