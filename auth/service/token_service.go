package service

import (
	"crypto/rsa"
	"log"

	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
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
func (s *TokenService) NewPairFromUser(u *model.User, prevTokenID string) (*model.TokenPair, error) {
	// No need to use a repository for idToken as it is unrelated to any data source
	idToken, err := util.GenerateIDToken(u, s.PrivKey)

	if err != nil {
		log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, rerrors.NewInternal()
	}

	refreshToken, err := util.GenerateRefreshToken(u.UID, s.RefreshSecret)

	if err != nil {
		log.Printf("Error generating refreshToken for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, rerrors.NewInternal()
	}

	// set freshly minted refresh token to valid list
	if err := s.TokenRepository.SetRefreshToken(u.UID.String(), refreshToken.ID, refreshToken.ExpiresIn); err != nil {
		log.Printf("Error storing tokenID for uid: %v. Error: %v\n", u.UID, err.Error())
		return nil, rerrors.NewInternal()
	}

	// delete old refresh token
	if prevTokenID != "" {
		if err := s.TokenRepository.DeleteRefreshToken(u.UID.String(), prevTokenID); err != nil {
			log.Printf("Could not delete previous refreshToken for uid: %v, tokenID: %v\n", u.UID.String(), prevTokenID)
		}
	}

	return &model.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken.SS,
	}, nil
}

// SignOut reaches out to the token repository to
// remove all refresh tokens for a given user
func (s *TokenService) SignOut(uid string) error {
	return s.TokenRepository.DeleteUserRefreshTokens(uid)
}

// ValidateRefreshToken validates the refresh token jwt string
// It returns the claims on the token if is valid
func (s *TokenService) ValidateRefreshToken(refreshTokenString string) (*util.RefreshTokenCustomClaims, error) {
	// Validate the refresh toksn
	claims, err := util.ValidateRefreshToken(refreshTokenString, s.RefreshSecret)

	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		log.Printf("Unable to validate or parse refreshToken for token string: %s\n%v\n", refreshTokenString, err)
		return nil, rerrors.NewAuthorization("Unable to verify user from refresh token")
	}

	return claims, nil
}

// ValidateIDToken validates the id token jwt string
// It returns the claims on the token if is valid
func (s *TokenService) ValidateIDToken(tokenString string) (*util.IDTokenCustomClaims, error) {
	claims, err := util.ValidateIDToken(tokenString, s.PubKey) // uses public RSA key

	// We'll just return unauthorized error in all instances of failing to verify user
	if err != nil {
		log.Printf("Unable to validate or parse idToken - Error: %v\n", err)
		return nil, rerrors.NewAuthorization("Unable to verify user from idToken")
	}

	return claims, nil
}
