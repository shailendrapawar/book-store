package services

import (
	"context"
	"errors"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
)

type CartItemsService interface {
	CreateCartItem(ctx context.Context, payload adapters.CreateCartItemPayload) (*adapters.CartItem, error)
}

type cartItemsServiceImpl struct {
	cartItemsDAO dao.CartItemsDAO
}

func NewCartItemsService(cartItemsDAO dao.CartItemsDAO) CartItemsService {
	return &cartItemsServiceImpl{
		cartItemsDAO: cartItemsDAO,
	}
}

func (s *cartItemsServiceImpl) CreateCartItem(ctx context.Context, payload adapters.CreateCartItemPayload) (*adapters.CartItem, error) {

	// 1: first check if item already exists for that, cart and book id
	var cartItem adapters.CartItem
	res, err := s.cartItemsDAO.Get(ctx, adapters.GetCartItemPayload{CartID: payload.CartID, BookID: payload.BookID})

	//this can also be restricted
	// if err == nil { //item already exists
	// 	cartItem = *res
	// 	//increase and update qty of that product and return
	// 	updatePayload := adapters.UpdateCartItemPayload{
	// 		BookID:   payload.BookID,
	// 		CartID:   payload.CartID,
	// 		Quantity: cartItem.Quantity + payload.Quantity,
	// 	}
	// 	res, err := s.cartItemsDAO.Update(ctx, updatePayload)
	// 	if err != nil {
	// 		return nil, errors.New("failed to update cart item")
	// 	}
	// 	cartItem = *res
	// 	return &cartItem, nil
	// }

	// 2: create new cart item if not found
	res, err = s.cartItemsDAO.Create(ctx, payload)
	if err != nil {
		return nil, errors.New("failed to create cart item")
	}
	cartItem = *res

	return &cartItem, nil
}
