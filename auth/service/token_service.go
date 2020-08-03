package service

import "github.com/jacobsngoodwin/wordmem/auth/model"

// TokenService used for injecting an implementation of ITokenRepository
// for use in service methods
type TokenService struct {
	TokenRepository ITokenRepository
}

// Create for creating new token set for a user
func (s *TokenService) Create(user *model.User) error {
	panic("not implemented") // TODO: Implement
}
