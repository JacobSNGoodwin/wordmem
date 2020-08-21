package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens returns new id and refresh tokens for a user
// provided a valid refresh token is present in the req
func (e *Env) Tokens(c *gin.Context) {
	var req tokensReq

	bindData(c, &req)

	// verify token - get token claims if valid
	refreshClaims, err := e.TokenService.ValidateRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	// get user from UID
	u, err := e.UserService.Get(refreshClaims.UID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	// create new tokens
	tokens, err := e.TokenService.NewPairFromUser(u, refreshClaims.Id)

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}
