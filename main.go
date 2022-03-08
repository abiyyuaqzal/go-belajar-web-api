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
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	log.Fatal("Failed to Load")
	// }

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")

	dsn := "kodingworks:4]Ow8tb%?sRt,qZxFtr>n+6cg^3&ar|4@tcp(ls-7f1435e2ded91dacd203c6b9ccd08d46577baba1.cadoav7mk1ht.ap-southeast-1.rds.amazonaws.com:3306)/aqzal_belajar?charset=utf8mb4&parseTime=True&loc=Local"
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
