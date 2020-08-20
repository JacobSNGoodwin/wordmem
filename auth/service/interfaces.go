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
	Create(u *model.User) (*model.User, error)
	Delete(uid uuid.UUID) error
	FindByID(uid uuid.UUID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Update(u *model.User) (*model.User, error)
}

// ITokenRepository defines methods it expects a repository
// it interacts with to implement
type ITokenRepository interface {
	SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(userID string, prevTokenID string) error
	DeleteUserRefreshTokens(userID string) error
}

// IImageRepositroy defines methods it expects
// any image repostiroy the application interacts with to implement
type IImageRepositroy interface {
	UploadUserImage(uid string, imageFile multipart.File) error
}
