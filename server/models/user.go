package models

type User struct {
	ID       string `jsonapi:"primary,user"`
	Username string `jsonapi:"attr,username"`
	Password string `jsonapi:"attr,password"`
}
