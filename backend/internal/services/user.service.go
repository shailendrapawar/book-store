package services

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shailendrapawar/book-store/internal/adapters"
	"github.com/shailendrapawar/book-store/internal/dao"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, req adapters.RegisterRequest) (interface{}, error)
}

type UserServiceImpl struct {
	userDAO dao.UserDAO
}

func NewUserService(db *sql.DB) UserService {
	return &UserServiceImpl{
		userDAO: dao.NewUserDAO(db),
	}
}
func (s *UserServiceImpl) Create(ctx context.Context, req adapters.RegisterRequest) (interface{}, error) {

	//1 check if already exists : (TODO)

	//2: genrate password hash
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 3: create user
	newUser := &adapters.User{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashPassword),
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	//4 call dao
	user, err := s.userDAO.Create(ctx, newUser)

	return user, err
}
