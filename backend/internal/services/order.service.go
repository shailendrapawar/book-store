package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type OrderService interface {
	Create(ctx context.Context,
		payload *adapters.CreateOrderRequest,
		userCart *adapters.Cart, user *adapters.Claims,
	) (*adapters.Order, error)
}

type orderServiceImpl struct {
	cartItemsDao dao.CartItemsDAO
	orderDao     dao.Orders

	bookService       BookService
	addressService    AddressService
	orderItemsService OrderItemsService
	cartService       CartService
}

func NewOrderService(db *sql.DB) OrderService {
	return &orderServiceImpl{
		cartItemsDao: dao.NewCartItemsDAO(db),
		orderDao:     dao.NewOrdersDAO(db),

		bookService:       NewBookService(db),
		addressService:    NewAddressService(db),
		orderItemsService: NewOrderItemsService(db),
		cartService:       NewCartService(db),
	}
}

func (s *orderServiceImpl) Create(
	ctx context.Context,
	payload *adapters.CreateOrderRequest,
	userCart *adapters.Cart,
	user *adapters.Claims,
) (*adapters.Order, error) {

	//1: get all books from cart with prices
	books, orderGrossAmount, priceMap, err := s.getCartBooksWithPrice(ctx, userCart)

	if err != nil {
		return nil, errors.New("Error while mapping cart items")
	}

	// 2: get shipping address
	address, err := s.addressService.Get(ctx, payload.AddressId)

	if err != nil {
		return nil, errors.New("Error while getting user shipping address")
	}

	//2: create order with payload first
	orderPayload := &adapters.CreateOrderPayload{
		User:          user,
		Address:       address,
		Coupon:        "",
		Books:         books,
		GrossAmount:   *orderGrossAmount,
		NetAmount:     *orderGrossAmount,
		PaymentMethod: "cod",
		Currency:      "INR",
	}
	order, err := s.orderDao.Create(ctx, orderPayload)
	if err != nil {
		return nil, errors.New("Error while creating order")
	}

	// 3: create order items secondly
	orderItems, err := s.orderItemsService.Create(ctx, order.ID, books, priceMap)

	if err != nil {
		return nil, errors.New("Error while creating order items")
	}

	// 4: TODO:clear cart of user
	_, err = s.cartService.Delete(ctx, userCart.ID)

	if err != nil {
		return nil, errors.New("Error while clearing cart")
	}

	// merge both order details and items
	return &adapters.Order{
		Id:     order.ID,
		UserId: order.UserID,
		Status: order.Status,

		DiscountValue: utils.ExtractFloat(order.DiscountValue),
		DiscountType:  order.DiscountType,

		GrossAmount: *orderGrossAmount,
		NetAmount:   utils.ExtractFloat(order.NetAmount),

		ShippingAddress: string(order.ShippingAddress.Val),
		ShippingCity:    order.ShippingCity,
		ShippingState:   order.ShippingState,
		ShippingPincode: order.ShippingPincode,

		PaymentMethod: order.PaymentMethod,
		PaymentStatus: order.PaymentStatus,

		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Items:     orderItems,
	}, nil
}

func (s *orderServiceImpl) getCartBooksWithPrice(ctx context.Context, userCart *adapters.Cart) ([]*adapters.Book, *float64, map[string]adapters.OrderItemMap, error) {

	//1: get all books from cart
	books := []*adapters.Book{}
	orderGrossAmount := 0.0

	priceMap := make(map[string]float64)
	orderItemMap := make(map[string]adapters.OrderItemMap)

	for _, item := range userCart.Items {
		book, err := s.bookService.Get(ctx, item.BookID)
		if err != nil {
			return nil, nil, nil, err
		}
		orderItemMap[book.ID] = adapters.OrderItemMap{
			Price:    book.Price,
			Quantity: item.Quantity,
		}

		priceMap[book.ID] = book.Price                            //add to price map
		orderGrossAmount += (book.Price * float64(item.Quantity)) //add to total price for each entry
		books = append(books, book)                               //append book itself
	}

	return books, &orderGrossAmount, orderItemMap, nil

}
