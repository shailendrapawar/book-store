package dao

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/middlewares"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type AddressDao interface {
	Create(ctx context.Context, address *adapters.CreateAddressRequest) (interface{}, error)
	Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error)
	Get(ctx context.Context, id string) (*models.Address, error)
	Update(ctx context.Context, id string, payload *adapters.UpdateAddressRequest) (interface{}, error)
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
		ID:            omit.From(strings.ToLower(addressUUID)),
		UserID:        omit.From(strings.ToLower(user.UserID)),
		AddressesType: omit.From(strings.ToLower(address.AddressType)),
		Line1:         omit.From(strings.ToLower(address.Line1)),
		Line2:         omitnull.From(strings.ToLower(address.Line2)),
		Landmark:      omitnull.From(strings.ToLower(address.Landmark)),
		City:          omit.From(strings.ToLower(address.City)),
		Pincode:       omit.From(strings.ToLower(address.Pincode)),
		District:      omit.From(strings.ToLower(address.District)),
		State:         omit.From(strings.ToLower(address.State)),
		Country:       omit.From(strings.ToLower(address.Country)),
		CreatedAt:     omit.From(time.Now()),
		UpdatedAt:     omit.From(time.Now()),
	}
	newAddress, err := models.Addresses.Insert(setter).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return newAddress, nil
}

func (d *addressDaoImpl) Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error) {

	var mods []bob.Mod[*dialect.SelectQuery]

	//append pagination
	mods = append(mods, sm.Limit(pagination.Limit))
	mods = append(mods, sm.Offset(pagination.Offset))

	//filters
	if filters.UserID != nil {
		mods = append(mods, models.SelectWhere.Addresses.UserID.EQ(*filters.UserID))
	}
	if filters.City != nil {
		mods = append(mods, models.SelectWhere.Addresses.City.EQ(*filters.City))
	}
	if filters.State != nil {
		mods = append(mods, models.SelectWhere.Addresses.State.EQ(*filters.State))
	}
	if filters.District != nil {
		mods = append(mods, models.SelectWhere.Addresses.District.EQ(*filters.District))
	}
	if filters.IsDeleted != nil {
		mods = append(mods, models.SelectWhere.Addresses.IsDeleted.EQ(*filters.IsDeleted))
	}

	result, err := models.Addresses.Query(mods...).All(ctx, d.db)

	if err != nil {
		return nil, err
	}
	var addresses []*adapters.Address

	for _, v := range result {
		a := &adapters.Address{
			Id:          v.ID,
			UserID:      v.UserID,
			AddressType: v.AddressesType,
			Line1:       v.Line1,
			Line2:       utils.ExtractNullString(v.Line2),
			Landmark:    utils.ExtractNullString(v.Landmark),
			City:        v.City,
			Pincode:     v.Pincode,
			District:    v.District,
			State:       v.State,
			Country:     v.Country,
		}
		addresses = append(addresses, a)
	}

	return addresses, nil
}
func (d *addressDaoImpl) Get(ctx context.Context, id string) (*models.Address, error) {

	address, err := models.Addresses.Query(
		sm.Where(models.Addresses.Columns.ID.EQ(psql.Arg(id))),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}

	return address, nil
}

func (d *addressDaoImpl) Update(ctx context.Context, id string, payload *adapters.UpdateAddressRequest) (interface{}, error) {

	//get address by id
	address, err := d.Get(ctx, id)

	if err != nil {
		return nil, errors.New("Book not found")
	}

	setter := d.setModel(payload, address)

	newAddress, err := models.Addresses.Update(
		models.UpdateWhere.Addresses.ID.EQ(id),
		setter.UpdateMod(),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}

	return newAddress, nil
}

func (d *addressDaoImpl) setModel(model *adapters.UpdateAddressRequest, entity *models.Address) *models.AddressSetter {

	setter := &models.AddressSetter{}

	if model.AddressType != nil {
		setter.AddressesType = omit.From(strings.ToLower(*model.AddressType))
	}

	if model.Line1 != nil {
		setter.Line1 = omit.From(strings.ToLower(*model.Line1))
	}

	if model.Line2 != nil {
		setter.Line2 = omitnull.From(strings.ToLower(*model.Line2))
	}

	if model.Landmark != nil {
		setter.Landmark = omitnull.From(strings.ToLower(*model.Landmark))
	}

	if model.City != nil {
		setter.City = omit.From(strings.ToLower(*model.City))
	}

	if model.Pincode != nil {
		setter.Pincode = omit.From(strings.ToLower(*model.Pincode))
	}

	if model.District != nil {
		setter.District = omit.From(strings.ToLower(*model.District))
	}

	if model.State != nil {
		setter.State = omit.From(strings.ToLower(*model.State))
	}

	if model.Country != nil {
		setter.Country = omit.From(strings.ToLower(*model.Country))
	}

	return setter
}
