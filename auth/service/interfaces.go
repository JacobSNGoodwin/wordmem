package service

import (
	"time"

	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// IUserRepository defines methods the service expects
// any repository it interacts with to implement
// I considered using a more idiomatic naming convention for this repository,
// but I could not come up with a good verb name (best I could do was)
// UserRepositoryInteractor
type IUserRepository interface {
	Create(u *model.User) (*model.User, error)
}

// ITokenRepository defines methods it expects a repository
// it interacts with to implement
type ITokenRepository interface {
	SetRefreshToken(tokenID string, expiresIn time.Duration) error
}
