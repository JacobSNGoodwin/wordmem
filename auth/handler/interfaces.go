package handler

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// IUserService defines methods the handler expects to
// interact with to perform CRUD operations on users
// These methods will be called from route handlers
type IUserService interface {
	Get(uid uuid.UUID) (*model.User, error)
	SignUp(email string, password string) (*model.User, error)
	SignIn(email string, password string) (*model.User, error)
	UpdateDetails(u *model.User) error
	Delete(uid uuid.UUID) error
	SetProfileImage(uid uuid.UUID, imageFile *multipart.FileHeader) (string, error)
	ClearProfileImage(uid uuid.UUID) error
}

// ITokenService defines methods handler expects to interact
// with for storing and retrieving tokens
type ITokenService interface {
	NewPairFromUser(user *model.User, prevTokenID string) (*model.TokenPair, error)
	SignOut(uid string) error
	ValidateIDToken(tokenString string) (*util.IDTokenCustomClaims, error)
	ValidateRefreshToken(tokenString string) (*util.RefreshTokenCustomClaims, error)
}
