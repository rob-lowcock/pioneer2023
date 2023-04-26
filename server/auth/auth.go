package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v4"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/models"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Db     *pgx.Conn
	DbUser db.User
}

func (a *Auth) ValidateCredentials(username, password string) (models.User, error) {
	user, err := a.DbUser.GetUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a *Auth) GenerateToken(user models.User) (string, error) {
	key := os.Getenv("JWT_SECRET")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
	})
	return t.SignedString([]byte(key))
}

func (a *Auth) ValidateToken(token string) error {

	jwt.WithValidMethods([]string{"HS256"})
	_, err := jwt.NewParser(jwt.WithValidMethods([]string{"HS256"})).ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return err
}
