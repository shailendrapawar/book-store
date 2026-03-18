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
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`

	Isbn string `json:"isbn"`

	Price decimal.Decimal `json:"price"`
	Stock int32           `json:"stock"`
}
