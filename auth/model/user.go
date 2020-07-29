package model

import "github.com/google/uuid"

// User defines domain model and it's json and db representations
type User struct {
	UID      uuid.UUID `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password,omitempty"`
}
