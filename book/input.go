package book

import "encoding/json"

type BookInput struct {
	Title string      `json:"title" binding:"required"` // validation
	Price json.Number `json:"price" binding:"required,number"`
}
