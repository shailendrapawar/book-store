package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/google/uuid"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/stephenafamo/bob"
)

type BookDAO interface {
	//methods
	Create(ctx context.Context, book *adapters.CreateBookRequest) (interface{}, error)
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

	//set values
	setter := &models.BookSetter{
		ID:        omit.From(bookUuid),
		Isbn:      omit.From(book.Isbn),
		Title:     omit.From(book.Title),
		Author:    omit.From(book.Author),
		Price:     omit.From(book.Price),
		Stock:     omit.From(book.Stock),
		Reserved:  omit.From(int32(0)),
		IsActive:  omit.From(true),
		CreatedAt: omit.From(time.Now()),
		UpdatedAt: omit.From(time.Now()),
	}

	row, err := models.Books.Insert(setter).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return row, nil
}
