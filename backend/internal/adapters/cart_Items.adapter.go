package adapters

type AddItemToCartRequest struct {
	ProductID string `json:"product_id" binding:"required,min=1" `
	Quantity  int    `json:"quantity" binding:"required,min=1" `
}
