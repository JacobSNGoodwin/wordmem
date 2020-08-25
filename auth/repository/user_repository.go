package repository

import (
	"log"
	"net/http"

	"github.com/google/uuid"
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
	query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING uid, email"

	newU := &model.User{}

	pw, err := util.HashPassword(u.Password)

	if err != nil {
		log.Printf("Unable to create password has for user: %v\n", u.Email)
		return newU, errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := r.DB.Get(newU, query, u.Email, pw); err != nil {
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
func (r *UserRepository) Delete(uid uuid.UUID) error {
	query := "DELETE FROM users WHERE uid=$1"
	_, err := r.DB.Exec(query, uid)
	if err != nil {
		return err
	}
	return nil
}

// FindByID fetches user by id
func (r *UserRepository) FindByID(uid uuid.UUID) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE uid=$1"

	// we need to actually check errors as it could be something other than not found
	if err := r.DB.Get(user, query, uid); err != nil {
		return user, errors.NewNotFound("uid", uid.String())
	}

	return user, nil
}

// FindByEmail fetches user by email address
func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	user := &model.User{}

	query := "SELECT * FROM users WHERE email=$1"

	if err := r.DB.Get(user, query, email); err != nil {
		log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
		return user, errors.NewNotFound("email", email)
	}

	return user, nil
}

// Update updates a user's properties
func (r *UserRepository) Update(u *model.User) error {
	query := `
		UPDATE users 
		SET name=:name, email=:email, website=:website
		WHERE uid=:uid
		RETURNING *;
	`

	nstmt, err := r.DB.PrepareNamed(query)

	if err != nil {
		log.Printf("Unable to prepare user update query: %v\n", err)
		return errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := nstmt.Get(u, u); err != nil {
		log.Printf("Failed to update details for user: %v\n", u)
		return errors.NewUnknown(http.StatusNotFound)
	}

	return nil
}

// UpdateImage is used to separately update a user's image separate from
// other account details
func (r *UserRepository) UpdateImage(uid uuid.UUID, imageURL string) error {
	query := `
		UPDATE users 
		SET image_url=$1
		WHERE uid=$2
	`
	_, err := r.DB.Exec(query, imageURL, uid)

	if err != nil {
		log.Printf("Error update image_url in database: %v\n", err)
		return errors.NewUnknown(http.StatusInternalServerError)
	}

	return nil
}
