package handler

import (
	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// Env is a struct used for injected the repository into the various route handler
type Env struct {
	UserService  IUserService
	TokenService ITokenService
}

// IUserService defines methods the handler expects to
// interact with to perform CRUD operations on users
// These methods will be called from route handlers
type IUserService interface {
	SignUp(email string, password string) (*model.User, error)
	SignIn(email string, password string) (*model.User, error)
	Get(uid uuid.UUID) (*model.User, error)
}

// ITokenService defines methods handler expects to interact
// with for storing and retrieving tokens
type ITokenService interface {
	NewPairFromUser(user *model.User) (*model.TokenPair, error)
	ValidateRefreshToken(tokenString string) (*util.RefreshTokenCustomClaims, error)
}
