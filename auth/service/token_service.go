package service

import (
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// TokenService used for injecting an implementation of ITokenRepository
// for use in service methods
type TokenService struct {
	TokenRepository ITokenRepository
	PrivKey         *rsa.PrivateKey
	PubKey          *rsa.PublicKey
}

// NewSetFromUser creates fresh id and refresh tokens for the current user
func (s *TokenService) NewSetFromUser(u *model.User) (*model.TokenPair, error) {
	// No need to use a repository for idToken as it is unrelated to any data source
	idToken, err := util.GenerateIDToken(u, s.PrivKey)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	// TODO: Generate refresh token

	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: "Dummy",
	}, nil
}

// NewSetFromRefreshToken returns a new id token and replaces the existing refresh
// token so that it is rotated
func (s *TokenService) NewSetFromRefreshToken(refreshToken string) (*model.TokenPair, error) {
	panic("not implemented") // TODO: Implement
}
