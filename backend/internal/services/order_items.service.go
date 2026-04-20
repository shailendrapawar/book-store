package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type OrderItemsService interface {
	Create(ctx context.Context,
		orderID string,
		books []*adapters.Book, priceMap map[string]adapters.OrderItemMap) ([]*adapters.OrderItem, error)

	Search(ctx context.Context, filters adapters.SearchOrderItemsFilters) ([]*adapters.OrderItem, error)
}
type orderItemsServiceImpl struct {
	orderItemDAO dao.OrderItemDAO
}

func NewOrderItemsService(db *sql.DB) OrderItemsService {
	return orderItemsServiceImpl{
		orderItemDAO: dao.NewOrderItemDAO(db),
	}
}

func (s orderItemsServiceImpl) Create(ctx context.Context,
	orderID string, books []*adapters.Book,
	priceMap map[string]adapters.OrderItemMap) ([]*adapters.OrderItem, error) {

	result := []*adapters.OrderItem{}

	//1:create payload for each order item and insert into db
	for _, v := range books {
		createOrderItemPayload := &adapters.CreateOrderItemPayload{
			OrderID:    orderID,
			BookID:     v.ID,
			Quantity:   priceMap[v.ID].Quantity,
			Title:      v.Title,
			Price:      priceMap[v.ID].Price,
			TotalPrice: priceMap[v.ID].Price * float64(priceMap[v.ID].Quantity),
		}

		//use wait group and wait for all goroutines to finish before returning response
		res, err := s.orderItemDAO.Create(ctx, createOrderItemPayload)

		if err != nil {
			return nil, err
		}

		orderItem := s.toAdapter(*res)
		result = append(result, orderItem)
	}

	//return mapped data
	return result, nil
}

func (s orderItemsServiceImpl) Search(ctx context.Context, filters adapters.SearchOrderItemsFilters) ([]*adapters.OrderItem, error) {

	items, err := s.orderItemDAO.Search(ctx, filters)
	if err != nil {
		return nil, err
	}

	var result []*adapters.OrderItem
	for _, v := range items {
		result = append(result, s.toAdapter(*v))
	}
	return result, nil
}

func (s *orderItemsServiceImpl) toAdapter(data models.OrderItem) *adapters.OrderItem {
	return &adapters.OrderItem{
		ID:         data.ID,
		OrderID:    data.OrderID,
		BookID:     data.BookID,
		Title:      data.Title,
		Price:      utils.ExtractFloat(data.Price),
		Quantity:   int(data.Quantity),
		TotalPrice: utils.ExtractFloat(data.TotalPrice),
	}

}
