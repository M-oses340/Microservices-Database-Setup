package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type db struct {
	db *sql.DB
}
