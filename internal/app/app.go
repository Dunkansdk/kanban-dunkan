package app

import "database/sql"

type Application struct {
	db *sql.DB
}
