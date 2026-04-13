package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

type CartService interface {
	create(ctx context.Context) (*adapters.Cart, error) //don't export this
	Search(ctx context.Context, filters adapters.CartSearchFilters, pagination adapters.PaginationRequest) (interface{}, error)
	Get(ctx context.Context, id string) (any, error)
	GetByUserID(ctx context.Context, userID string) (*adapters.Cart, error)

	AddItemToCart(ctx context.Context, payload adapters.AddItemToCartRequest, user adapters.Claims) (any, error)
}

type cartServiceImpl struct {
	cartDao      dao.CartDAO
	cartItemsDao dao.CartItemsDAO
	bookDao      dao.BookDAO
}

func NewCartService(db *sql.DB) CartService {
	return &cartServiceImpl{
		cartDao:      dao.NewCartDao(db),
		cartItemsDao: dao.NewCartItemsDAO(db),
		bookDao:      dao.NewBookDAO(db),
	}
}

func (s *cartServiceImpl) create(ctx context.Context) (*adapters.Cart, error) {

	user := middlewares.GetUserFromCTX(ctx)
	//check if cart already exists
	cart, err := s.cartDao.GetByUserID(ctx, user.UserID)
	if err == nil {
		fmt.Print(cart)
		return nil, errors.New("Cart already exists for this user")
	}

	return s.cartDao.Create(ctx, user.UserID)
}

func (s *cartServiceImpl) Search(ctx context.Context, filters adapters.CartSearchFilters, pagination adapters.PaginationRequest) (interface{}, error) {

	res, err := s.cartDao.Search(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *cartServiceImpl) Get(ctx context.Context, id string) (interface{}, error) {

	res, err := s.cartDao.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *cartServiceImpl) GetByUserID(ctx context.Context, userID string) (*adapters.Cart, error) {

	res, err := s.cartDao.GetByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *cartServiceImpl) AddItemToCart(ctx context.Context, payload adapters.AddItemToCartRequest, user adapters.Claims) (interface{}, error) {

	var cart adapters.Cart
	// 1: get cart first
	existingCart, err := s.GetByUserID(ctx, user.UserID)

	if err == nil && existingCart == nil {
		//create new cart (coz cart doesn't exists)
		newCart, err := s.create(ctx)
		if err != nil {
			return nil, errors.New("Error while cart creation")
		}
		cart = *newCart //init new cart
	} else {
		cart = *existingCart //init existing cart
	}

	// 2: cart found so use this cart id and handle cart items

	//validate product /book if exists or not
	_, err = s.bookDao.GetById(ctx, payload.BookID)
	if err != nil {
		return nil, errors.New("book not found")
	}

	//check if item already exists
	searchPayload := adapters.GetCartItemPayload{CartID: cart.ID, BookID: payload.BookID}
	existingItem, err := s.cartItemsDao.Get(ctx, searchPayload)

	if err == nil && existingItem != nil {
		//item already exists so update quantity
		updatePayload := adapters.UpdateCartItemPayload{BookID: payload.BookID, CartID: cart.ID, Quantity: existingItem.Quantity + payload.Quantity}
		res, err := s.cartItemsDao.Update(ctx, updatePayload)
		if err != nil {
			return nil, errors.New("failed to update cart item")
		}
		return res, nil
	}

	//create new cart-item if not found
	createPayload := adapters.CreateCartItemPayload{CartID: cart.ID, BookID: payload.BookID, Quantity: payload.Quantity}
	res, err := s.cartItemsDao.Create(ctx, createPayload)
	if err != nil {
		return nil, errors.New("failed to create cart item")
	}

	return res, nil
}
