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
	return s.UserRepository.Create(u)
}
