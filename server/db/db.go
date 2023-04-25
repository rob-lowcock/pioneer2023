package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type Db struct {
}

func (d *Db) Connect() (*pgx.Conn, error) {
	return pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
}
