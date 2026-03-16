package dao

import (
	"database/sql"

	"github.com/stephenafamo/bob"
)

type UserDAO interface {
}

type userDAOImpl struct {
	db bob.DB
}

func NewUserDAO(db *sql.DB) UserDAO {
	return &userDAOImpl{
		db: bob.NewDB(db),
	}
}
