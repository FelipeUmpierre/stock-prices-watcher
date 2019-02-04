package repository

import "github.com/jmoiron/sqlx"

type EventStore struct {
	db *sqlx.DB
}

func NewEventstore(db *sqlx.DB) *EventStore {
	return &EventStore{db: db}
}

func (e *EventStore) Emit() {

}
