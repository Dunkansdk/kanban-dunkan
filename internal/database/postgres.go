package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PSQLDB struct {
	instance *sql.DB
}

// https://go.dev/doc/database/sql-injection

func (db *PSQLDB) GetConnection() *sql.DB {
	if db.instance != nil {
		return db.instance
	}

	connStr := "postgres://postgres:postgres@localhost:5432/kandundb?sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(100)

	if err != nil {
		log.Fatal(err)
	}

	db.instance = conn

	return db.instance
}
