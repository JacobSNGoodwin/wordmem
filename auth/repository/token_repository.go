package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// TokenRepositoryInit creates a concrete implemetation
// of TokenRepository by inject the required redis data source
func TokenRepositoryInit(r *redis.Client) *TokenRepository {
	return &TokenRepository{
		Redis: r,
	}
}

// TokenRepository is data/repository implementation
// of service layer ITokenRepository
type TokenRepository struct {
	Redis *redis.Client
}

// Create creates a token for users
func (r *TokenRepository) Create(u *model.User) error {
	panic("not implemented") // TODO: Implement
}
