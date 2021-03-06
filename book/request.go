package book

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"` // validation
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description"`
	Rating      json.Number `json:"rating"`
}
