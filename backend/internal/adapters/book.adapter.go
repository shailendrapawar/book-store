package adapters

import "github.com/shopspring/decimal"

type Book struct {
	ID          string
	Title       string
	Description string
	Author      string

	Isbn string

	Price    decimal.Decimal
	Stock    int32
	IsActive bool
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`

	Isbn string `json:"isbn" binding:"required"`

	Price decimal.Decimal `json:"price" binding:"required gte=0" `
	Stock int32           `json:"stock" binding:"required gte=0"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Author      *string `json:"author"`

	Isbn *string `json:"isbn"`

	Price    *decimal.Decimal `json:"price" binding:"gte=0"`
	Stock    *int32           `json:"stock" binding:"gte=0"`
	Reserved *int32           `json:"reserved" binding:"gte=0"`
}
