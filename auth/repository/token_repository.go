package repository

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
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
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v", userID, tokenID, err)
		return errors.NewUnknown(http.StatusInternalServerError)
	}
	return nil
}

// ReplaceRefreshToken generarates a new refresh token
// after validating the previous token
// Note that there could be some debate about whether taking a step to
// delete an old redis key with a TTL is useful. I will do this to
// prevent old Refresh tokens from floating around
func (r *TokenRepository) ReplaceRefreshToken(userID string, prevTokenID string, newTokenID string) (string, error) {
	key := fmt.Sprintf("%s:%s", userID, prevTokenID)
	refreshToken, err := r.Redis.Get(context.Background(), key).Result()

	if err != nil {
		log.Printf("Could not GET refresh token to redis for userID/tokenID: %s/%s: %v", userID, prevTokenID, err)
		return "", errors.NewUnknown(http.StatusInternalServerError)
	}

	return refreshToken, nil
}
