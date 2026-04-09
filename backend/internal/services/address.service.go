package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
)

type AddressService interface {
	Create(ctx context.Context, req *adapters.CreateAddressRequest) (interface{}, error)
	Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error)
	Get(ctx context.Context, id string) (interface{}, error)
}

type addressServiceImpl struct {
	addressDAO dao.AddressDao
}

func NewAddressService(db *sql.DB) AddressService {
	return &addressServiceImpl{
		addressDAO: dao.NewAddressDao(db),
	}
}

func (s *addressServiceImpl) Create(ctx context.Context, req *adapters.CreateAddressRequest) (interface{}, error) {

	res, err := s.addressDAO.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *addressServiceImpl) Search(ctx context.Context, filters adapters.SearchAddressFilters, pagination adapters.PaginationRequest) (interface{}, error) {

	res, err := s.addressDAO.Search(ctx, filters, pagination)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (s *addressServiceImpl) Get(ctx context.Context, id string) (interface{}, error) {

	res, err := s.addressDAO.Get(ctx, id)

	if err != nil {
		return nil, err
	}

	return res, nil
}
