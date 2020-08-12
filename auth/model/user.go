package model

import (
	"github.com/google/uuid"
)

// User defines domain model and it's json and db representations
type User struct {
	UID      uuid.UUID  `db:"uid" json:"uid,omitempty"`
	Email    string     `db:"email" json:"email,omitempty"`
	Password string     `db:"password,omitempty" json:"password,omitempty"`
	Name     NullString `db:"name,omitempty" json:"name,omitempty"`
	ImageURL NullString `db:"image_url,omitempty" json:"imageUrl,omitempty"`
	Website  NullString `db:"website,omitempty" json:"website,omitempty"`
}
