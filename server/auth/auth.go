package auth

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rob-lowcock/pioneer2023/db"
	"github.com/rob-lowcock/pioneer2023/models"
	"golang.org/x/crypto/bcrypt"
)

/*
This comes with a massive health warning. JWTs are vulnerable to CSRF attacks. In time this will all need to be replaced with Auth0 or similar, but for the MVP this will have to do.
*/
type Auth struct {
	Db     *pgxpool.Pool
	DbUser db.User
}

// ValidateCredentials validates the provided username and password
// against the corresponding user in the database.
// It takes a username string and a password string as input parameters.
// It returns a models.User struct and an error if the validation fails.
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
