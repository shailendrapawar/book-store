package dao

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/types"
)

type Orders interface {
	Create(ctx context.Context, payload *adapters.CreateOrderPayload) (*models.Order, error)
}

type ordersDAOImpl struct {
	db bob.DB
}

func NewOrdersDAO(db *sql.DB) Orders {
	return &ordersDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *ordersDAOImpl) Create(ctx context.Context, payload *adapters.CreateOrderPayload) (*models.Order, error) {

	orderUUID := utils.CreateUUID()

	//convert into json
	addressJSON, err := json.Marshal(payload.Address) //this converts into plain json bytes
	if err != nil {
		return nil, err
	}
	var addr types.JSON[json.RawMessage] //this converts for DB compatible json
	if err := addr.Scan(addressJSON); err != nil {
		return nil, err
	}

	orderSetter := &models.OrderSetter{
		ID:            omit.From(orderUUID),
		UserID:        omit.From(payload.User.UserID),
		Status:        omit.From("pending"),
		DiscountValue: omit.From(utils.ToDecimal(0)),
		DiscountType:  omit.From("fixed"),

		GrossAmount:     omit.From(utils.ToDecimal(payload.GrossAmount)),
		NetAmount:       omit.From(utils.ToDecimal(payload.NetAmount)),
		ShippingAddress: omit.From(addr),

		ShippingCity:    omit.From(payload.Address.City),
		ShippingState:   omit.From(payload.Address.State),
		ShippingPincode: omit.From(payload.Address.Pincode),

		PaymentMethod: omit.From(payload.PaymentMethod),
		PaymentStatus: omit.From("pending"),
	}

	order, err := models.Orders.Insert(orderSetter).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return order, nil

}
