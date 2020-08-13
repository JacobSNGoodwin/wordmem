package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// Me handler fetches user from ID token
// so that user can be verified by the server and returned
func (e *Env) Me(c *gin.Context) {
	h := authHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		// this type check appears to be extra cautious as I could not
		// find a case where this error was anything other than InvalidValidationError
		// see https://godoc.org/github.com/go-playground/validator#InvalidValidationError
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown Error"})
		}

		// vErr is serializable because it has struct tags!
		vErr := errors.NewFromValidationErrors(err.(validator.ValidationErrors))
		c.JSON(vErr.Status, gin.H{"error": vErr})

		return
	}

	idToken := strings.Split(h.IDToken, "Bearer ")[1]

	claims, err := e.TokenService.ValidateIDToken(idToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errors.NewUnauthorized("Must provide Authorization header"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": claims,
	})
}
