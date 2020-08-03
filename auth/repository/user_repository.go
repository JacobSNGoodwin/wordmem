package repository

import (
	"log"
	"net/http"

	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/util"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// UserRepositoryInit intializes a UserRepository by injecting in a db ref
func UserRepositoryInit(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// UserRepository is data/repository implementation
// of service layer IUserRepository
type UserRepository struct {
	DB *sqlx.DB
}

// Create reacher out to database SQLX api
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	queryString := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING uid, name, email"

	n := &model.User{}

	pw, err := util.HashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to create password has for user: %v\n", u.Email)
		return n, errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := r.DB.Get(n, queryString, u.Name, u.Email, pw); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
			return n, errors.NewAlreadyExists("email", u.Email)
		}

		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return n, errors.NewUnknown(http.StatusInternalServerError)
	}
	return n, nil
}
