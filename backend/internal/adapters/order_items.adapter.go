package adapters

import "time"

type OrderItem struct {
	ID      string `json:"id"`
	OrderID string `json:"order_id"`
	BookID  string `json:"book_id"`
	Title   string `json:"title"`

	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type OrderItemMap struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}
type CreateOrderItemRequest struct {
}
type CreateOrderItemPayload struct {
	OrderID    string  `json:"order_id"`
	BookID     string  `json:"book_id"`
	Title      string  `json:"title"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
