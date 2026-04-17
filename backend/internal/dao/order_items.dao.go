package dao

import (
	"context"
	"database/sql"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
)

type OrderItemDAO interface {
	Create(ctx context.Context, payload *adapters.CreateOrderItemPayload) (*models.OrderItem, error)
}

type orderItemDAOImpl struct {
	db bob.DB
}

func NewOrderItemDAO(db *sql.DB) OrderItemDAO {
	return &orderItemDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *orderItemDAOImpl) Create(ctx context.Context, payload *adapters.CreateOrderItemPayload) (*models.OrderItem, error) {

	orderItemUUID := utils.CreateUUID()

	setter := models.OrderItemSetter{
		ID:      omit.From(orderItemUUID),
		OrderID: omit.From(payload.OrderID),

		BookID:     omit.From(payload.BookID),
		Title:      omit.From(payload.Title),
		Price:      omit.From(utils.ToDecimal(payload.Price)),
		Quantity:   omit.From(int32(payload.Quantity)),
		TotalPrice: omit.From(utils.ToDecimal(payload.TotalPrice)),
	}

	orderItem, err := models.OrderItems.Insert(&setter).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return orderItem, nil
}
