package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type BookService interface {
	Create(ctx context.Context, req *adapters.CreateBookRequest) (interface{}, error)
	Get(ctx context.Context, identifier string) (interface{}, error)
}

type bookService struct {
	bookDao dao.BookDAO
}

func NewBookService(db *sql.DB) BookService {
	return &bookService{
		bookDao: dao.NewBookDAO(db),
	}
}

func (s *bookService) Create(ctx context.Context, req *adapters.CreateBookRequest) (interface{}, error) {

	result, err := s.bookDao.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *bookService) Get(ctx context.Context, identifier string) (interface{}, error) {

	var book interface{}
	//find whether its id or isbn
	if utils.IsUUID(identifier) {
		fmt.Print("UUID=====>")
		//its uuid -> call getById
		res, err := s.bookDao.GetById(ctx, identifier)
		if err != nil {
			return nil, err
		}
		book = res
	} else if utils.IsISBN(identifier) {
		fmt.Print("ISBN=====>")
		//its ISBN number -> cal getByISBN
		res, err := s.bookDao.GetByISBN(ctx, identifier)
		if err != nil {
			return nil, err
		}
		book = res
	} else {
		//invalid characters
		return nil, errors.New("Invalid identifier/id")
	}
	return book, nil
}
