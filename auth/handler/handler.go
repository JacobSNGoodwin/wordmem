package handler

import "github.com/jacobsngoodwin/wordmem/auth/model"

// Env is a struct used for injected the repository into the various route handler
type Env struct {
	UserService  IUserService
	TokenService ITokenService
}

// IUserService defines methods the handler expects to
// interact with to perform CRUD operations on users
// These methods will be called from route handlers
type IUserService interface {
	SignUp(user *model.User) (*model.User, error)
}

// ITokenService defines methods handler expects to interact
// with for storing and retrieving tokens
type ITokenService interface {
	NewSetFromUser(user *model.User) (*model.TokenPair, error)
}
