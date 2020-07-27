package models

import "github.com/google/uuid"

// User defines domain model and it's json and db representations
type User struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password,omitempty"`
}

// UserRepository defines the interface that must be met
// by any database implementation of a UserRepository
type UserRepository interface {
	// User(id uuid.UUID) (User, error)
	// Users() ([]User, error)
	CreateUser(u *User) error
	// UpdateUser(u *User) error
}
