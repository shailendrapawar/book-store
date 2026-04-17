package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
)

type BookService interface {
	Create(ctx context.Context, req *adapters.CreateBookRequest) (*adapters.Book, error)
	Get(ctx context.Context, identifier string) (*adapters.Book, error)
	Update(ctx context.Context, id string, payload adapters.UpdateBookRequest) (*adapters.Book, error)
	Search(ctx context.Context, pagination adapters.PaginationRequest) ([]*adapters.Book, error)
}

type bookService struct {
	bookDao dao.BookDAO
}

func NewBookService(db *sql.DB) BookService {
	return &bookService{
		bookDao: dao.NewBookDAO(db),
	}
}

func (s *bookService) Create(ctx context.Context, req *adapters.CreateBookRequest) (*adapters.Book, error) {

	result, err := s.bookDao.Create(ctx, req)

	if err != nil {
		return nil, err
	}

	return &adapters.Book{
		ID:          result.ID,
		Title:       result.Title,
		Description: utils.ExtractNullString(result.Description),
		Price:       utils.ExtractFloat(result.Price),
		Isbn:        result.Isbn,
		Stock:       result.Stock,
		Author:      result.Author,
	}, nil
}

func (s *bookService) Get(ctx context.Context, identifier string) (*adapters.Book, error) {

	var book *models.Book
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

	return &adapters.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: utils.ExtractNullString(book.Description),
		Price:       utils.ExtractFloat(book.Price),
		Isbn:        book.Isbn,
		Stock:       book.Stock,
		Author:      book.Author,
	}, nil

}

func (s *bookService) Update(ctx context.Context, id string, payload adapters.UpdateBookRequest) (*adapters.Book, error) {

	res, err := s.bookDao.Update(ctx, id, payload)
	if err != nil {
		return nil, err
	}
	return &adapters.Book{
		ID:          res.ID,
		Title:       res.Title,
		Description: utils.ExtractNullString(res.Description),
		Price:       utils.ExtractFloat(res.Price),
		Isbn:        res.Isbn,
		Stock:       res.Stock,
		Author:      res.Author,
	}, nil
}

func (s *bookService) Search(ctx context.Context, pagination adapters.PaginationRequest) ([]*adapters.Book, error) {

	res, err := s.bookDao.Search(ctx, pagination)
	if err != nil {
		return nil, err
	}

	items := make([]*adapters.Book, len(res))
	for i, book := range res {
		items[i] = &adapters.Book{
			ID:          book.ID,
			Title:       book.Title,
			Description: utils.ExtractNullString(book.Description),
			Price:       utils.ExtractFloat(book.Price),
			Isbn:        book.Isbn,
			Stock:       book.Stock,
			Author:      book.Author,
		}
	}
	return items, nil
}
