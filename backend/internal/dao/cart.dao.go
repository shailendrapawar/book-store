package dao

import (
	"context"
	"database/sql"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type CartDAO interface {
	Create(ctx context.Context, userID string) (*adapters.Cart, error)
	Search(ctx context.Context, filters adapters.CartSearchFilters, pagination adapters.PaginationRequest) ([]*adapters.Cart, error)
	GetByID(ctx context.Context, id string) (*adapters.Cart, error)
	GetByUserID(ctx context.Context, userID string) (*adapters.Cart, error)
}

type CartDAOImpl struct {
	db bob.DB
}

func NewCartDao(db *sql.DB) CartDAO {
	return &CartDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *CartDAOImpl) Create(ctx context.Context, userID string) (*adapters.Cart, error) {

	id := utils.CreateUUID()
	newCart := models.CartSetter{
		ID:     omit.From(id),
		UserID: omit.From(userID),
	}

	cart, err := models.Carts.Insert(&newCart).One(ctx, d.db)

	if err != nil {
		return nil, err
	}

	return &adapters.Cart{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}
func (d *CartDAOImpl) GetByID(ctx context.Context, id string) (*adapters.Cart, error) {

	cart, err := models.Carts.Query(
		models.SelectWhere.Carts.ID.EQ(id),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return &adapters.Cart{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}
func (d *CartDAOImpl) GetByUserID(ctx context.Context, userID string) (*adapters.Cart, error) {

	cart, err := models.Carts.Query(
		models.SelectWhere.Carts.UserID.EQ(userID),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return &adapters.Cart{
		ID:        cart.ID,
		UserID:    cart.UserID,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}, nil
}
func (d *CartDAOImpl) Search(ctx context.Context, filters adapters.CartSearchFilters, pagination adapters.PaginationRequest) ([]*adapters.Cart, error) {

	var mods []bob.Mod[*dialect.SelectQuery]

	//append pagination
	mods = append(mods, sm.Limit(pagination.Limit))
	mods = append(mods, sm.Offset(pagination.Offset))

	res, err := models.Carts.Query(
		mods...,
	).All(ctx, d.db)

	if err != nil {
		return nil, err
	}

	var carts []*adapters.Cart
	// return &[]adapters.Cart{}, nil
	for _, v := range res {

		c := &adapters.Cart{
			ID:        v.ID,
			UserID:    v.UserID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		carts = append(carts, c)
	}
	return carts, nil
}

func (d *CartDAOImpl) Delete(ctx context.Context, cartID string) (*models.Cart, error) {

	res, err := models.Carts.Delete(
		models.DeleteWhere.Carts.ID.EQ(cartID),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return res, nil
}
