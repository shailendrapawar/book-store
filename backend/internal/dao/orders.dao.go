package dao

import (
	"database/sql"

	"github.com/stephenafamo/bob"
)

type Orders interface {
}

type ordersDAOImpl struct {
	db bob.DB
}

func NewOrdersDAO(db *sql.DB) Orders {
	return &ordersDAOImpl{
		db: bob.NewDB(db),
	}
}

// func (d *ordersDAOImpl) CreateOrder(ctx context.Context,payload ) (interface{}, error) {

// }
