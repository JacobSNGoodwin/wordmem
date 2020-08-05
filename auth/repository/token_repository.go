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
func (r *TokenRepository) SetRefreshToken(userID string, tokenID string, expiresIn time.Duration) error {
	// Should we create a context and add timeout?

	// We'll store userID with token id so we can scan (non-blocking)
	// over the user's tokens and delete them in case of token leakage
	// Then can delete these tokens with a pipeline to avoid creating many connections
	// for deleting multiple keys
	key := fmt.Sprintf("%s:%s", userID, tokenID)
	if err := r.Redis.Set(context.Background(), key, 0, expiresIn).Err(); err != nil {
		return fmt.Errorf("Could not add refresh token to redis %w", err)
	}
	return nil
}
