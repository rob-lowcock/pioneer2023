package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type Db struct {
}

func (d *Db) Connect() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}
