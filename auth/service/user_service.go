package service

import (
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository IUserRepository
}

// SignUp creates a new user based on data in model.User
func (s *UserService) SignUp(u *model.User) (*model.User, error) {
	// In this case, we have a one-to-one correspondence between service method "SignUp" and repository method "Create"
	// This is not always the case, though I can understand why this looks redundant
	return s.UserRepository.Create(u)
}

// Remove user used to roll back user creation on failed token creation on signup
// (only signup)
func (s *UserService) Remove(u *model.User) error {
	return s.UserRepository.Delete(u)
}
