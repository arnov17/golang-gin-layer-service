package book

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Discount    int
	Rating      int
	CreateAt    time.Time
	UpdateAt    time.Time
}
