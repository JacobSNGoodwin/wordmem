package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Tokens returns new id and refresh tokens for a user
// provided a valid refresh token is present in the req
func (e *Env) Tokens(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "This bidniz is working!",
	})
}
