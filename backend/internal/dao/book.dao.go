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
	"github.com/shopspring/decimal"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type BookDAO interface {
	//methods
	Create(ctx context.Context, book *adapters.CreateBookRequest) (interface{}, error)

	GetById(ctx context.Context, id string) (*models.Book, error)
	GetByISBN(ctx context.Context, id string) (interface{}, error)

	Update(ctx context.Context, id string, payload adapters.UpdateBookRequest) (*models.Book, error)
	// setModel(model adapters.UpdateBookRequest, entity *models.Book) *models.BookSetter
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

	// generate uuid
	bookUuid := uuid.New().String()
	isbn := utils.NormalizeISBN(book.Isbn)

	//set values
	setter := &models.BookSetter{
		ID:          omit.From(bookUuid),
		Isbn:        omit.From(isbn),
		Title:       omit.From(book.Title),
		Description: omitnull.From(book.Description),
		Author:      omit.From(book.Author),
		Price:       omit.From(decimal.NewFromFloat(book.Price)),
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

func (d *bookDAOImpl) GetById(ctx context.Context, id string) (*models.Book, error) {

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

func (d *bookDAOImpl) Update(ctx context.Context, id string, payload adapters.UpdateBookRequest) (*models.Book, error) {

	book, err := d.GetById(ctx, id)

	if err != nil {
		return nil, errors.New("Book not found")
	}

	// fill setter
	setter := setModel(payload, book)

	updatedBook, err := models.Books.Update(
		models.UpdateWhere.Books.ID.EQ(id),
		models.UpdateWhere.Books.Reserved.LT(book.Stock), // condition in DB
		setter.UpdateMod(),
	).One(ctx, d.db)

	return updatedBook, nil
}

func setModel(model adapters.UpdateBookRequest, entity *models.Book) *models.BookSetter {

	setter := &models.BookSetter{}

	if model.Title != nil {
		setter.Title = omit.From(*model.Title)
	}

	if model.Description != nil {
		setter.Description = omitnull.From(*model.Description)
	}

	if model.Author != nil {
		setter.Author = omit.From(*model.Author)
	}

	if model.Price != nil {
		setter.Price = omit.From(decimal.NewFromFloat(*model.Price))

		// ======IMP==============
		if model.Stock != nil && entity.Reserved < *model.Stock {
			setter.Stock = omit.From(*model.Stock)
		}
		if model.Reserved != nil && entity.Stock > *model.Reserved {
			setter.Reserved = omit.From(*model.Reserved)
		}
	}
	if model.IsActive != nil {
		setter.IsActive = omit.From(*model.IsActive)
	}
	return setter
}
