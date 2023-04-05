package models

import (
	"database/sql"
)

type DBModel struct {
	DB *sql.DB
}

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}
