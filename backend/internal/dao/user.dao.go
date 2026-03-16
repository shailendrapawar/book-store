package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/stephenafamo/bob"
)

type UserDAO interface {
	Create(ctx context.Context, user *adapters.User) (*adapters.RegisterResponse, error)
}

type userDAOImpl struct {
	db bob.DB
}

func NewUserDAO(db *sql.DB) UserDAO {
	return &userDAOImpl{
		db: bob.NewDB(db),
	}
}

func (d *userDAOImpl) Create(ctx context.Context, user *adapters.User) (*adapters.RegisterResponse, error) {

	setter := &models.UserSetter{
		ID:        omit.From(user.ID),
		Name:      omit.From(user.Name),
		Email:     omit.From(user.Email),
		Password:  omit.From(user.Password),
		Role:      omit.From(user.Role),
		CreatedAt: omit.From(time.Now()),
		UpdatedAt: omit.From(time.Now()),
	}

	row, err := models.Users.Insert(setter).One(ctx, d.db)

	return toAdapter(row), err
}

func (d *userDAOImpl) Get(ctx context.Context, keyword string) (*adapters.RegisterResponse, error) {

	//  _,err:= uuid.Parse()

}

func toAdapter(row *models.User) *adapters.RegisterResponse {

	return &adapters.RegisterResponse{
		Name:  row.ID,
		Email: row.Email,
		Role:  row.Role,
	}

}
