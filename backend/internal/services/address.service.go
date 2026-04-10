package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type AddressService interface {
	Create(ctx context.Context, req *adapters.CreateAddressRequest) (*adapters.Address, error)
	Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error)
	Get(ctx context.Context, id string) (*adapters.Address, error)
	Update(ctx context.Context, id string, req *adapters.UpdateAddressRequest) (interface{}, error)
}

type addressServiceImpl struct {
	addressDAO dao.AddressDao
}

func NewAddressService(db *sql.DB) AddressService {
	return &addressServiceImpl{
		addressDAO: dao.NewAddressDao(db),
	}
}

func (s *addressServiceImpl) Create(ctx context.Context, req *adapters.CreateAddressRequest) (*adapters.Address, error) {

	res, err := s.addressDAO.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return &adapters.Address{
		Id:          res.ID,
		UserID:      res.UserID,
		AddressType: res.AddressesType,
		Line1:       res.Line1,
		Line2:       utils.ExtractNullString(res.Line2),
		Landmark:    utils.ExtractNullString(res.Landmark),
		City:        res.City,
		Pincode:     res.Pincode,
		District:    res.District,
		State:       res.State,
		Country:     res.Country,
	}, nil
}

func (s *addressServiceImpl) Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error) {

	res, err := s.addressDAO.Search(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func (s *addressServiceImpl) Get(ctx context.Context, id string) (*adapters.Address, error) {

	res, err := s.addressDAO.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return &adapters.Address{
		Id:          res.ID,
		UserID:      res.UserID,
		AddressType: res.AddressesType,
		Line1:       res.Line1,
		Line2:       utils.ExtractNullString(res.Line2),
		Landmark:    utils.ExtractNullString(res.Landmark),
		City:        res.City,
		Pincode:     res.Pincode,
		District:    res.District,
		State:       res.State,
		Country:     res.Country,
	}, nil
}

func (s *addressServiceImpl) Update(ctx context.Context, id string, req *adapters.UpdateAddressRequest) (interface{}, error) {

	res, err := s.addressDAO.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
