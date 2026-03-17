package adapters

import "github.com/shopspring/decimal"

type Book struct {
	ID     string
	Isbn   string
	Title  string
	Author string
	Price  decimal.Decimal
	Stock  int32
}

type CreateBookRequest struct {
	Title  string
	Isbn   string
	Author string
	Price  decimal.Decimal
	Stock  int32
}
