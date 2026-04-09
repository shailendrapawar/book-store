package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/middlewares"
)

type CartService interface {
	Create(ctx context.Context) (any, error)
	Search(ctx context.Context, filters adapters.CartSearchFilters, pagination adapters.PaginationRequest) (interface{}, error)
	Get(ctx context.Context, id string) (any, error)
	GetByUserID(ctx context.Context, userID string) (any, error)
}

type cartServiceImpl struct {
	cartDao dao.CartDAO
}

func NewCartService(db *sql.DB) CartService {
	return &cartServiceImpl{
		cartDao: dao.NewCartDao(db),
	}
}

func (s *cartServiceImpl) Create(ctx context.Context) (interface{}, error) {

	//get user
	user := middlewares.GetUserFromCTX(ctx)
	//check if cart already exists

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

func (s *cartServiceImpl) GetByUserID(ctx context.Context, userID string) (interface{}, error) {

	res, err := s.cartDao.GetByUserID(ctx, userID)

	if err != nil {
		return nil, err
	}
	return res, nil
}
