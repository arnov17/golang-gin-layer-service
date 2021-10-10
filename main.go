package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	fmt.Println("Database connection succeed")
	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	// GET All BOOK
	books, err := bookRepository.FindAll()
	for _, e := range books {
		fmt.Println("ttile :", e.Title)
	}

	book, err := bookRepository.FindById(2)
	fmt.Println(book)

	// // CREATE BOOK
	// book := book.Book{
	// 	Title:       "One Punch Man",
	// 	Description: "Chapter 54",
	// 	Price:       65000,
	// 	Rating:      5,
	// 	Discount:    1,
	// }
	// bookRepository.Create(book)

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	v1.GET("/hello", handler.HelloHandler)

	v1.GET("/book/:id", handler.BooksHandler)

	v1.POST("/book", handler.PostBookHandler)

	v1.GET("/query", handler.QueryHandler)

	router.Run(":8888")
}

// main
// service --> feature business, logic
// repository --> berhubungan dgn databse
// db
// mysql
