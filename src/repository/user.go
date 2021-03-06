package repository

import "context"

type User struct {
	ID       string `json:"id,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(id string) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByEmailAndPassword(email string, password string) (User, error)
}
