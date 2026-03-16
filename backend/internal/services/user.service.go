package services

import (
	"database/sql"

	"github.com/shailendrapawar/book-store/internal/dao"
)

type UserService interface {
}

type UserServiceImpl struct {
	userDAO dao.UserDAO
}

func NewUserService(db *sql.DB) UserService {
	return &UserServiceImpl{
		userDAO: dao.NewUserDAO(db),
	}
}
func (*UserServiceImpl) Create() {

}
