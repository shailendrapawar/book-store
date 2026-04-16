package dao

import (
	"context"
	"database/sql"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/shailendrapawar/book-store/internal/utils"
	"github.com/stephenafamo/bob"
)

type CartItemsDAO interface {
	Create(ctx context.Context, payload adapters.CreateCartItemPayload) (*adapters.CartItem, error)
	Get(ctx context.Context, payload adapters.GetCartItemPayload) (*adapters.CartItem, error)
	GetByCartID(ctx context.Context, cartID string) ([]adapters.CartItem, error)
	Update(ctx context.Context, payload adapters.UpdateCartItemPayload) (*adapters.CartItem, error)
	Delete(ctx context.Context, id, cartId string) (int, error)
}

type cartItemsDAOImpl struct {
	db bob.DB
}

func NewCartItemsDAO(db *sql.DB) CartItemsDAO {
	return &cartItemsDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *cartItemsDAOImpl) Create(ctx context.Context, payload adapters.CreateCartItemPayload) (*adapters.CartItem, error) {

	id := utils.CreateUUID()
	setter := models.CartItemSetter{
		ID:       omit.From(id),
		CartID:   omit.From(payload.CartID),
		BookID:   omit.From(payload.BookID),
		Quantity: omit.From(int32(payload.Quantity)),
	}

	cartItem, err := models.CartItems.Insert(&setter).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return &adapters.CartItem{
		Id:       cartItem.ID,
		CartID:   cartItem.CartID,
		BookID:   cartItem.BookID,
		Quantity: int(cartItem.Quantity),
	}, nil
}

func (d *cartItemsDAOImpl) Get(ctx context.Context, payload adapters.GetCartItemPayload) (*adapters.CartItem, error) {

	cartItem, err := models.CartItems.Query(
		models.SelectWhere.CartItems.CartID.EQ(payload.CartID),
		models.SelectWhere.CartItems.BookID.EQ(payload.BookID),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}
	return &adapters.CartItem{
		Id:       cartItem.ID,
		CartID:   cartItem.CartID,
		BookID:   cartItem.BookID,
		Quantity: int(cartItem.Quantity),
	}, nil
}

func (d *cartItemsDAOImpl) GetByCartID(ctx context.Context, cartID string) ([]adapters.CartItem, error) {

	cartItems, err := models.CartItems.Query(
		models.SelectWhere.CartItems.CartID.EQ(cartID),
	).All(ctx, d.db)

	if err != nil {
		return nil, err
	}

	var items []adapters.CartItem
	for _, cartItem := range cartItems {
		items = append(items, adapters.CartItem{
			Id:       cartItem.ID,
			CartID:   cartItem.CartID,
			BookID:   cartItem.BookID,
			Quantity: int(cartItem.Quantity),
		})
	}
	return items, nil
}

func (d *cartItemsDAOImpl) Update(ctx context.Context, payload adapters.UpdateCartItemPayload) (*adapters.CartItem, error) {

	setter := &models.CartItemSetter{
		Quantity: omit.From(int32(payload.Quantity)),
	}

	cartItem, err := models.CartItems.Update(
		models.UpdateWhere.CartItems.BookID.EQ(payload.BookID),
		models.UpdateWhere.CartItems.CartID.EQ(payload.CartID),
		setter.UpdateMod(),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}

	return &adapters.CartItem{
		Id:       cartItem.ID,
		CartID:   cartItem.CartID,
		BookID:   cartItem.BookID,
		Quantity: int(cartItem.Quantity),
	}, nil
}

func (d *cartItemsDAOImpl) Delete(ctx context.Context, id, cartId string) (int, error) {

	_, err := models.CartItems.Delete(
		models.DeleteWhere.CartItems.ID.EQ(id),
		models.DeleteWhere.CartItems.CartID.EQ(cartId),
	).One(ctx, d.db)

	if err != nil {
		return 0, err
	}
	return 1, nil
}
