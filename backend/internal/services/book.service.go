package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
)

type BookService interface {
	Create(ctx context.Context, req *adapters.CreateBookRequest) (interface{}, error)
}

type bookDAOImpl struct {
	bookDao dao.BookDAO
}

func NewBookService(db *sql.DB) BookService {
	return &bookDAOImpl{
		bookDao: dao.NewBookDAO(db),
	}
}

func (s *bookDAOImpl) Create(ctx context.Context, req *adapters.CreateBookRequest) (interface{}, error) {

	result, err := s.bookDao.Create(ctx, req)

	if err != nil {
		return nil, errors.New("Error while adding book")
	}

	return result, nil
}
