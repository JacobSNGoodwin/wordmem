package repository

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
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
		log.Printf("Could not SET refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return rerrors.NewUnknown(http.StatusInternalServerError)
	}
	return nil
}

// DeleteRefreshToken used to delete old  refresh tokens
// Services my access this to revolve tokens
func (r *TokenRepository) DeleteRefreshToken(userID string, tokenID string) error {
	key := fmt.Sprintf("%s:%s", userID, tokenID)
	if err := r.Redis.Del(context.Background(), key).Err(); err != nil {
		log.Printf("Could not delete refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return rerrors.NewUnknown(http.StatusInternalServerError)
	}

	return nil
}

// DeleteUserRefreshTokens scans all keys in repository matching the
// current userID and deletes them in a safe way
func (r *TokenRepository) DeleteUserRefreshTokens(uid string) error {
	log.Printf("Deleting refresh tokens for UID: %s\n", uid)
	ctx := context.Background()
	pattern := fmt.Sprintf("%s*", uid)

	iter := r.Redis.Scan(ctx, 0, pattern, 0).Iterator()

	for iter.Next(ctx) {
		if err := r.Redis.Del(ctx, iter.Val()).Err(); err != nil {
			log.Printf("Failed to delete refresh token: %s\n", iter.Val())
		}
	}

	if err := iter.Err(); err != nil {
		log.Printf("Failed to delete refresh token: %s\n", iter.Val())
	}

	return nil
}
