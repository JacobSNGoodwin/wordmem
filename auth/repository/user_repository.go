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

// UserRepository is data/repository implementation
// of service layer IUserRepository
type UserRepository struct {
	DB *sqlx.DB
}

// Create reaches out to database SQLX api
func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING uid, name, email"

	newU := &model.User{}

	pw, err := util.HashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to create password has for user: %v\n", u.Email)
		return newU, errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := r.DB.Get(newU, query, u.Name, u.Email, pw); err != nil {
		// check unique constraint
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
			return newU, errors.NewAlreadyExists("email", u.Email)
		}

		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return newU, errors.NewUnknown(http.StatusInternalServerError)
	}
	return newU, nil
}

// Delete removes a user based on their uid
func (r *UserRepository) Delete(u *model.User) error {
	query := "DELETE FROM users WHERE uid=$1"
	_, err := r.DB.Exec(query, u.UID)
	if err != nil {
		return err
	}
	return nil
}
