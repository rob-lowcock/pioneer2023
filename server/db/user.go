package db

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/rob-lowcock/pioneer2023/models"
)

type User struct {
	Db *pgx.Conn
}

func (u *User) GetUserByUsername(username string) (models.User, error) {
	row := u.Db.QueryRow(context.Background(), `SELECT id, username, password FROM users WHERE username = $1`, username)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Password)

	return user, err
}
