package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jmoiron/sqlx"
)

// UserRepositoryInit intializes a UserRepository by injecting in a db ref
func UserRepositoryInit(db *sqlx.DB, inmem *redis.Client) *UserRepository {
	return &UserRepository{
		DB:    db,
		InMem: inmem,
	}
}

// UserRepository is data/repository implementation
// of service layer IUserRepository
type UserRepository struct {
	DB    *sqlx.DB
	InMem *redis.Client
}

// Create reacher out to database SQLX api
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	queryString := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING uid, name, email"

	n := &model.User{}

	if err := r.DB.Get(n, queryString, u.Name, u.Email, u.Password); err != nil {
		return n, err
	}
	return n, nil
}
