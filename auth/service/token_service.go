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
	RefreshSecret   string
}

// NewPairFromUser creates fresh id and refresh tokens for the current user
func (s *TokenService) NewPairFromUser(u *model.User) (*model.TokenPair, error) {
	// No need to use a repository for idToken as it is unrelated to any data source
	idToken, err := util.GenerateIDToken(u, s.PrivKey)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	refreshToken, err := util.GenerateRefreshToken(u.UID, s.RefreshSecret)

	if err != nil {
		log.Printf("Error generating refreshToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	if err := s.TokenRepository.SetRefreshToken(u.UID.String(), refreshToken.ID, refreshToken.ExpiresIn); err != nil {
		log.Printf("Error storing tokenID for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, errors.NewUnknown(http.StatusInternalServerError)
	}

	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken.SS,
	}, nil
}
