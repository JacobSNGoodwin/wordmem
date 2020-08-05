package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// TokenRepository is data/repository implementation
// of service layer ITokenRepository
type TokenRepository struct {
	Redis *redis.Client
}

// SetRefreshToken stores a refresh token with an expiry time
func (r *TokenRepository) SetRefreshToken(tokenID string, expiresIn time.Duration) error {
	// Should we create a context and add timeout?
	if err := r.Redis.Set(context.Background(), tokenID, 0, expiresIn).Err(); err != nil {
		return fmt.Errorf("Could not add refresh token to redis %w", err)
	}
	return nil
}
