package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rob-lowcock/pioneer2023/models"
)

type User struct {
	Db *pgxpool.Pool
}

func (u *User) GetUserByUsername(username string) (models.User, error) {
	row := u.Db.QueryRow(context.Background(), `SELECT id, username, password FROM users WHERE username = $1`, username)

	user := models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
