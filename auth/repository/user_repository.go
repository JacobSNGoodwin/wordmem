package repository

import (
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jmoiron/sqlx"
)

// UserRepositoryInit intializes a UserRepository by injecting in a db ref
func UserRepositoryInit(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// UserRepository is data/repository implementation
// of model layer UserRepository
type UserRepository struct {
	*sqlx.DB
}

// CreateUser adds a user to table associated with UserRepository
func (r *UserRepository) CreateUser(u *model.User) error {
	panic("not implemented") // TODO: Implement
}
