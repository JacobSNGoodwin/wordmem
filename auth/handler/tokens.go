package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

type tokensReq struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// Tokens returns new id and refresh tokens for a user
// provided a valid refresh token is present in the req
func (e *Env) Tokens(c *gin.Context) {
	var req tokensReq

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown Error"})
		}

		// vErr is serializable because it has struct tags!
		vErr := errors.NewFromValidationErrors(err.(validator.ValidationErrors))
		log.Printf("%v", vErr)
		c.JSON(vErr.Status, gin.H{"error": vErr})

		return
	}

	// verify token and get userID
	uid, err := e.TokenService.UserIDFromRefreshToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	// get user from UID
	u, err := e.UserService.Get(uid)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	// create new tokens

	// get a newTokenPair
	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
