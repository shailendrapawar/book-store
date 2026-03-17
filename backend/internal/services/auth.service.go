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

type AuthService interface {
	Register(ctx context.Context, req adapters.RegisterRequest) (interface{}, error)
	Login(ctx context.Context, req *adapters.LoginRequest) (interface{}, error)
}

type AuthServiceImpl struct {
	userDAO dao.UserDAO
}

func NewAuthService(db *sql.DB) AuthService {
	return &AuthServiceImpl{
		userDAO: dao.NewUserDAO(db),
	}
}

func (s *AuthServiceImpl) Register(ctx context.Context, req adapters.RegisterRequest) (interface{}, error) {
	//1 check if already exists : (TODO)(OPTIONAL)

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

func (s *AuthServiceImpl) Login(ctx context.Context, req *adapters.LoginRequest) (interface{}, error) {

	// 1: fetch user first
	user, err := s.userDAO.GetByEmail(ctx, req.Email)
	if err != nil {
		//user dosent exists
		return nil, errors.New("User dosen't exists")
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	// 2: decryp hashpass
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		//invalid credentials
		return nil, errors.New("invalid credentials")
	}

	return nil, nil
}
