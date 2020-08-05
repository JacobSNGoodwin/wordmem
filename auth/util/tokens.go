package util

import (
	"crypto/rsa"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// myCustomClaims holds structure of jwt claims
type myCustomClaims struct {
	UID   uuid.UUID `json:"uid"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.StandardClaims
}

// GenerateIDToken generates an IDToken which is a jwt with myCustomClaims
func GenerateIDToken(u *model.User, key *rsa.PrivateKey) (string, error) {
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
	ss, err := token.SignedString(key)

	if err != nil {
		log.Println("Failed to sign token string")
		return "", err
	}

	return ss, nil
}
