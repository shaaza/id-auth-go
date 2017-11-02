package db

import "database/sql"

type DB struct {
	Instance *sql.DB
}
