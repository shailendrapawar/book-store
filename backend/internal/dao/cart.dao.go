package dao

import (
	"context"
	"database/sql"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
)

type CartDAO interface {
	Create(ctx context.Context, userID string) (*models.Cart, error)
}

type CartDAOImpl struct {
	db bob.DB
}

func NewCartDao(db *sql.DB) CartDAO {
	return &CartDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *CartDAOImpl) Create(ctx context.Context, userID string) (*models.Cart, error) {

	id := utils.CreateUUID()
	newCart := models.CartSetter{
		ID:     omit.From(id),
		UserID: omit.From(userID),
		Status: omit.From("active"),
	}
	cart, err := models.Carts.Insert(&newCart).One(ctx, d.db)

	if err != nil {
		return nil, nil
	}

	return &models.Cart{
		ID:        cart.ID,
		UserID:    cart.UserID,
		Status:    cart.Status,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}
