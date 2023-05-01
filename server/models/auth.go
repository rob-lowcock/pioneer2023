package models

type Auth struct {
	ID       string `jsonapi:"primary,auth"`
	Email    string `jsonapi:"attr,email"`
	Password string `jsonapi:"attr,password"`
}
