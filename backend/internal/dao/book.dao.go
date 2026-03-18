package dao

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/aarondl/opt/omitnull"
	"github.com/google/uuid"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type BookDAO interface {
	//methods
	Create(ctx context.Context, book *adapters.CreateBookRequest) (interface{}, error)
	GetById(ctx context.Context, id string) (interface{}, error)
	GetByISBN(ctx context.Context, id string) (interface{}, error)
}

type bookDAOImpl struct {
	db bob.DB
}

func NewBookDAO(db *sql.DB) BookDAO {
	return &bookDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *bookDAOImpl) Create(ctx context.Context, book *adapters.CreateBookRequest) (interface{}, error) {

	// genrate uuid
	bookUuid := uuid.New().String()
	isbn := utils.NormalizeISBN(book.Isbn)

	//set values
	setter := &models.BookSetter{
		ID:          omit.From(bookUuid),
		Isbn:        omit.From(isbn),
		Title:       omit.From(book.Title),
		Description: omitnull.From(book.Description),
		Author:      omit.From(book.Author),
		Price:       omit.From(book.Price),
		Stock:       omit.From(book.Stock),
		Reserved:    omit.From(int32(0)),
		IsActive:    omit.From(true),
		CreatedAt:   omit.From(time.Now()),
		UpdatedAt:   omit.From(time.Now()),
	}

	row, err := models.Books.Insert(setter).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (d *bookDAOImpl) GetById(ctx context.Context, id string) (interface{}, error) {

	setter := models.Books.Columns.ID.EQ(psql.Arg(id))

	book, err := models.Books.Query(sm.Where(setter)).One(ctx, d.db)

	if err != nil {
		return nil, errors.New("Error  while getting book")
	}
	return book, nil
}

func (d *bookDAOImpl) GetByISBN(ctx context.Context, id string) (interface{}, error) {

	setter := models.Books.Columns.Isbn.EQ(psql.Arg(id))

	book, err := models.Books.Query(sm.Where(setter)).One(ctx, d.db)

	if err != nil {
		return nil, errors.New("Error  while getting book")
	}
	return book, nil
}
