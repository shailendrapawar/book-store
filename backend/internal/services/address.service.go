package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
)

type AddressService interface {
	Create(ctx context.Context, req *adapters.CreateAddressRequest) (interface{}, error)
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
