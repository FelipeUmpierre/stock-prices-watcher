package repository

import (
	"github.com/jmoiron/sqlx"
)

type (
	Stocks struct {
		db *sqlx.DB
	}
)

func NewStocks(db *sqlx.DB) *Stocks {
	return &Stocks{db: db}
}
