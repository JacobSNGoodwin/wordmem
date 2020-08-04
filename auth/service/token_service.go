package service

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// TokenService used for injecting an implementation of ITokenRepository
// for use in service methods
type TokenService struct {
	TokenRepository ITokenRepository
}

// NewSetFromUser creates fresh id and refresh tokens for the current user
func (s *TokenService) NewSetFromUser(u *model.User) (*model.TokenPair, error) {
	// No need to use a repository for idToken as it is unrelated to any data source
	idToken, err := generateIDToken(u)

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

// myCustomClaims holds structure of jwt claims
type myCustomClaims struct {
	UID   uuid.UUID `json:"uid"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.StandardClaims
}

func generateIDToken(u *model.User) (string, error) {
	unixTime := time.Now().UnixNano() / 1000000 // time in ms
	tokenExp := unixTime + 60*1000*15           // 15 minutes from current time
	claims := myCustomClaims{
		UID:   u.UID,
		Name:  u.Name,
		Email: u.Email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  unixTime,
			ExpiresAt: tokenExp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// sign with private rs256 pem
	// for info on generating keys
	// go to https://cloud.google.com/iot/docs/how-tos/credentials/keys
	// TODO: Store in env variable?
	priv, err := ioutil.ReadFile("./rsa_private.pem")

	if err != nil {
		log.Println("Failed to read private key")
		return "", err
	}

	rsaKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

	if err != nil {
		log.Println("Failed to parse RSA key from file")
		return "", err
	}

	ss, err := token.SignedString(rsaKey)

	if err != nil {
		log.Println("Failed to sign token string")
		return "", err
	}

	return ss, nil
}
