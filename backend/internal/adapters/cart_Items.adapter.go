package adapters

type CartItem struct {
	Id       string `json:"id"`
	CartID   string `json:"cart_id" binding:"required,min=1" `
	BookID   string `json:"product_id" binding:"required,min=1" `
	Quantity int    `json:"quantity" binding:"required,min=1" `
}

type CreateCartItemPayload struct {
	CartID   string `json:"cart_id" binding:"required,min=1" `
	BookID   string `json:"product_id" binding:"required,min=1" `
	Quantity int    `json:"quantity" binding:"required,min=1" `
}

type AddItemToCartRequest struct {
	BookID   string `json:"product_id" binding:"required,min=1" `
	Quantity int    `json:"quantity" binding:"required,min=1" `
}

type GetCartItemPayload struct {
	CartID string `json:"cart_id" binding:"required,min=1" `
	BookID string `json:"product_id" binding:"required,min=1" `
}

type UpdateCartItemRequest struct {
	BookID   string `json:"book_id" binding:"required,min=1" `
	Quantity int    `json:"quantity" binding:"required,min=1" `
}
type UpdateCartItemPayload struct {
	CartID   string `json:"cart_id" binding:"required,min=1" `
	BookID   string `json:"product_id" binding:"required,min=1" `
	Quantity int    `json:"quantity" binding:"required,min=1" `
}
