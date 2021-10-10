package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Agung setiawan",
		"bio":  "A Softaware developer",
	})
}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello",
	})
}

func BooksHandler(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func PostBookHandler(c *gin.Context) {
	var BookInput book.BookRequest

	err := c.ShouldBindJSON(&BookInput)
	if err != nil {
		// log.Fatal(err)
		erroMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			erroMessages = append(erroMessages, errorMsg)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": erroMessages,
		})

		// c.JSON(http.StatusBadRequest, erroMessages)
		// fmt.Println(erroMessages)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"title":     BookInput.Title,
		"price":     BookInput.Price,
		"sub_title": BookInput.SubTitle,
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}
