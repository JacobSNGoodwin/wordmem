package service

import "github.com/jacobsngoodwin/wordmem/auth/model"

// TokenService used for injecting an implementation of ITokenRepository
// for use in service methods
type TokenService struct {
	TokenRepository ITokenRepository
}

// NewSetFromUser creates fresh id and refresh tokens for the current user
func (s *TokenService) NewSetFromUser(user *model.User) error {
	panic("not implemented") // TODO: Implement
}

// NewSetFromToken returns a new id token and replaces the existing refresh
// token so that it is rotated
func (s *TokenService) NewSetFromToken(refreshToken string) error {
	panic("not implemented") // TODO: Implement
}
