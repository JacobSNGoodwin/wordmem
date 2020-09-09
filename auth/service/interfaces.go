package service

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// IUserRepository defines methods the service expects
// any repository it interacts with to implement
// I considered using a more idiomatic naming convention for this repository,
// but I could not come up with a good verb name (best I could do was)
// UserRepositoryInteractor
type IUserRepository interface {
	Create(u *model.User) error
	Delete(uid uuid.UUID) error
	FindByID(uid uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(u *model.User) error
	UpdateImage(uid uuid.UUID, url string, storageID string) error
}

// ITokenRepository defines methods it expects a repository
// it interacts with to implement
type ITokenRepository interface {
	SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error
	ValidateRefreshToken(userID string, tokenID string) error
	DeleteRefreshToken(userID string, prevTokenID string) error
	DeleteUserRefreshTokens(userID string) error
}

// IImageRepositroy defines methods it expects
// any image repostiroy the application interacts with to implement
type IImageRepositroy interface {
	UploadUserImage(objName string, imageFile multipart.File) error
}
