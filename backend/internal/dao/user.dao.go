package dao

import (
	"context"
	"database/sql"
	"time"

	"github.com/aarondl/opt/omit"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/db/models"
	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/sm"
)

type UserDAO interface {
	Create(ctx context.Context, user *adapters.User) (*adapters.RegisterResponse, error)
	GetByID(ctx context.Context, userID string) (*models.User, error)
	GetByEmail(ctx context.Context, userEmail string) (*models.User, error)
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
	if err != nil {
		return nil, err
	}
	return toAdapter(row), err
}

func (d *userDAOImpl) GetByID(ctx context.Context, userID string) (*models.User, error) {
	user, err := models.Users.Query(
		sm.Where(models.Users.Columns.ID.EQ(psql.Arg(userID))),
	).One(ctx, d.db)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (d *userDAOImpl) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := models.Users.Query(
		sm.Where(models.Users.Columns.Email.EQ(psql.Arg(email))),
	).One(ctx, d.db)
	if err != nil {
		return nil, err
	}

	return user, err
}

func toAdapter(row *models.User) *adapters.RegisterResponse {

	return &adapters.RegisterResponse{
		Name:  row.ID,
		Email: row.Email,
		Role:  row.Role,
	}

}
