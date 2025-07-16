package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() error {
    var err error
    connStr := "host=localhost port=5432 user=postgres password=golu1804 dbname=internal_transfers sslmode=disable"
    DB, err = sqlx.Connect("postgres", connStr)
    return err
}
