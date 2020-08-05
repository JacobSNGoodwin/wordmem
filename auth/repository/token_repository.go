package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// TokenRepository is data/repository implementation
// of service layer ITokenRepository
type TokenRepository struct {
	Redis *redis.Client
}

// SetRefreshToken stores a refresh tokeen in the data source (redis)
// for the given user and returns the token
func (r *TokenRepository) SetRefreshToken(u *model.User) (string, error) {
	panic("not implemented") // TODO: Implement
}

// GetRefreshToken stores a refresh tokeen in the data source (redis)
// for the given user and returns the token
func (r *TokenRepository) GetRefreshToken(tokenID string) (string, error) {
	panic("not implemented") // TODO: Implement
}
