package model

import "github.com/google/uuid"

// User defines domain model and it's json and db representations
type User struct {
	UID      uuid.UUID `db:"uid" json:"uid,omitempty"`
	Name     string    `db:"name" json:"name"`
	Email    string    `db:"email" json:"email,omitempty"`
	Password string    `db:"password,omitempty" json:"password,omitempty"`
}
