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

	// // CREATE BOOK
	// book := book.Book{}
	// book.Title = "Naruto"
	// book.Price = 60000
	// book.Discount = 10
	// book.Description = "chapter 442"
	// book.Rating = 4

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error creating book record")
	// 	fmt.Println("==========================")
	// }

	// // GET BOOK
	// var books []book.Book

	// err = db.Find(&books).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error find book record")
	// 	fmt.Println("==========================")
	// }
	// // list book
	// for _, e := range books {
	// 	fmt.Println(e.Title)
	// 	fmt.Println(e.Price)
	// }

	// // Update
	// var book book.Book
	// err = db.Where("id = ?", 1).First(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error find book record")
	// 	fmt.Println("==========================")
	// }

	// book.Title = "One Piece New Cover"
	// err = db.Save(&book).Error
	// if err != nil {
	// 	fmt.Println("==========================")
	// 	fmt.Println("error update book record")
	// 	fmt.Println("==========================")
	// }

	// // Delete
	var book book.Book
	err = db.Where("id = ?", 1).First(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("error find book record")
		fmt.Println("==========================")
	}
	err = db.Delete(&book).Error
	if err != nil {
		fmt.Println("==========================")
		fmt.Println("error delete book record")
		fmt.Println("==========================")
	}

	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)

	v1.GET("/hello", handler.HelloHandler)

	v1.GET("/book/:id", handler.BooksHandler)

	v1.POST("/book", handler.PostBookHandler)

	v1.GET("/query", handler.QueryHandler)

	router.Run(":8888")
}
