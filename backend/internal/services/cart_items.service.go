package services

import (
	"github.com/shailendrapawar/book-store/internal/dao"
)

type CartItemsService interface {
	// CreateCartItem(ctx context.Context, payload adapters.CreateCartItemPayload) (*adapters.CartItem, error)
}

type cartItemsServiceImpl struct {
	cartItemsDAO dao.CartItemsDAO
}

func NewCartItemsService(cartItemsDAO dao.CartItemsDAO) CartItemsService {
	return &cartItemsServiceImpl{
		cartItemsDAO: cartItemsDAO,
	}
}
