package services

import (
	"context"
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/dao"
)

type CartService interface {
	Create(ctx context.Context, userID string) (interface{}, error)
}

type cartServiceImpl struct {
	cartDao dao.CartDAO
}

func NewCartService(db *sql.DB) CartService {
	return &cartServiceImpl{
		cartDao: dao.NewCartDao(db),
	}
}

func (s *cartServiceImpl) Create(ctx context.Context, userID string) (interface{}, error) {

	//check if cart already exists
	// res, err := s.cartDao.GetByUserID(ctx, userID)
	return s.cartDao.Create(ctx, userID)
}
