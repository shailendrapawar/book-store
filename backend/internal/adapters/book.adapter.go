package adapters

type Book struct {
	ID          string
	Title       string
	Description string
	Author      string

	Isbn string

	Price    float64
	Stock    int32
	IsActive bool
}

type CreateBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`

	Isbn string `json:"isbn" binding:"required"`

	Price float64 `json:"price" binding:"required gte=0" `
	Stock int32   `json:"stock" binding:"required gte=0"`
}

type UpdateBookRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Author      *string `json:"author"`

	Isbn *string `json:"isbn"`

	Price    *float64 `json:"price" binding:"omitempty,gte=0"`
	Stock    *int32   `json:"stock" binding:"omitempty,gte=0"`
	Reserved *int32   `json:"reserved" binding:"omitempty,gte=0"`

	IsActive *bool `json:"isActive" binding:"omitempty"`
}
type SearchBookResponse struct {
	data       []Book
	pagination Pagination
}
