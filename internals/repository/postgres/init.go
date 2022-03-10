package postgres

import "github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/infrastructure/database"

type PostgresRepository struct {
	db *database.DB
}

func NewRepository(db *database.DB) Repository {
	return &PostgresRepository{db: db}
}
