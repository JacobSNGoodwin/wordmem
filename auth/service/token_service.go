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
// If a previous token is included, the previous token is removed from
// the tokens repository
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

// ValidateRefreshToken validates the refresh token jwt string
// It returns the claims on the token if is valid
func (s *TokenService) ValidateRefreshToken(refreshTokenString string) (*util.RefreshTokenCustomClaims, error) {
	// Validate the refresh toksn
	token, err := util.ValidateRefreshToken(refreshTokenString, s.RefreshSecret)

	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		log.Printf("Unable to validate or parse refreshToken for token string: %s\n%v\n", refreshTokenString, err)
		return nil, errors.NewUnauthorized("Unable to verify user from refresh token")
	}

	return token, nil
}
