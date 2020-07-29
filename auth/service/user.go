package service

import (
	"log"

	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// UserService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type UserService struct {
	UserRepository IUserRepository
}

// SignUp creates a new user based on data in model.User
func (s *UserService) SignUp(u *model.User) error {
	if err := s.UserRepository.Create(u); err != nil {
		log.Printf("Failed to create user: %v\n", u.UID)
		return err
	}
	return nil
}
