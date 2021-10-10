package book

import "encoding/json"

type BookInput struct {
	Title string `json:"title" binding:"required"`
	// Price    int    `json:"price" binding:"required,number"`
	Price    json.Number `json:"price" binding:"required,number"`
	SubTitle string      `json:"sub_title"`
}
