package repository

import (
	"context"
	"fmt"
	"log"
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
		return rerrors.NewInternal()
	}
	return nil

}

// ValidateRefreshToken checks for userID:tokenID key in repository of valid refresh tokens
func (r *TokenRepository) ValidateRefreshToken(userID string, tokenID string) error {
	// for deleting multiple keys
	key := fmt.Sprintf("%s:%s", userID, tokenID)

	_, err := r.Redis.Get(context.Background(), key).Result()

	if err == redis.Nil {
		log.Printf("Could not find refresh token in redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return rerrors.NewInternal()
	}

	if err != nil {
		log.Printf("Key not found: %s\n", key)
		return rerrors.NewAuthorization("User is not logged in")
	}

	return nil
}

// DeleteRefreshToken used to delete old  refresh tokens
// Services my access this to revolve tokens
func (r *TokenRepository) DeleteRefreshToken(userID string, tokenID string) error {
	key := fmt.Sprintf("%s:%s", userID, tokenID)
	if err := r.Redis.Del(context.Background(), key).Err(); err != nil {
		log.Printf("Could not delete refresh token to redis for userID/tokenID: %s/%s: %v\n", userID, tokenID, err)
		return rerrors.NewInternal()
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
