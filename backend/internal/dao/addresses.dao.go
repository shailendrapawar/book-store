package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/middlewares"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
)

type AddressDao interface {
	Create(ctx context.Context, address *adapters.CreateAddressRequest) (interface{}, error)
}

type addressDaoImpl struct {
	db bob.DB
}

func NewAddressDao(db *sql.DB) AddressDao {
	return &addressDaoImpl{
		db: bob.NewDB(db),
	}
}

func (d *addressDaoImpl) Create(ctx context.Context, address *adapters.CreateAddressRequest) (interface{}, error) {

	user := middlewares.GetUserFromCTX(ctx)
	addressUUID := utils.CreateUUID()

	setter := &models.AddressSetter{
		ID:            omit.From(addressUUID),
		UserID:        omit.From(user.UserID),
		AddressesType: omit.From(address.AddressType),
		Line1:         omit.From(address.Line1),
		Line2:         omitnull.From(address.Line2),
		Landmark:      omitnull.From(address.Landmark),
		City:          omit.From(address.City),
		Pincode:       omit.From(address.Pincode),
		District:      omit.From(address.District),
		State:         omit.From(address.State),
		Country:       omit.From(address.Country),
		IsDefault:     omit.From(address.IsDefault),
		CreatedAt:     omit.From(time.Now()),
		UpdatedAt:     omit.From(time.Now()),
	}
	newAddress, err := models.Addresses.Insert(setter).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return newAddress, nil
}
