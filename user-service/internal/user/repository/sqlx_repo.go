package repository

import "database/sql"

type SQLXRepository struct {
	db *sql.DB
}
