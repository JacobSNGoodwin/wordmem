package service

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository IUserRepository
}

// SignUp creates a new user based on data in model.User
func (s *UserService) SignUp(email string, password string) (*model.User, error) {
	// In this case, we have a one-to-one correspondence between service method "SignUp" and repository method "Create"
	// This is not always the case, though I can understand why this looks redundant
	return s.UserRepository.Create(&model.User{
		Email:    email,
		Password: password,
	})
}

// SignIn returns a user after comparing supplied email/password with
// stored email/password
func (s *UserService) SignIn(email string, password string) (*model.User, error) {
	u, err := s.UserRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	// verify password
	match, err := util.ComparePasswords(u.Password, password)

	if err != nil {
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	if !match {
		return nil, errors.NewUnauthorized("Invalid email and password combination")
	}

	return u, nil
}

// Get retreives a user based on their uuid
func (s *UserService) Get(uid uuid.UUID) (*model.User, error) {
	return s.UserRepository.FindByID(uid)
}
