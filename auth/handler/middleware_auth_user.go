package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

// AuthUser extracts any user from the Authorization header
// which is of the form "Bearer token"
// It sets any user to the context if the user exsists
func (e *Env) AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		h := authHeader{}

		if err := c.ShouldBindHeader(&h); err != nil {
			// this type check appears to be extra cautious as I could not
			// find a case where this error was anything other than InvalidValidationError
			// see https://godoc.org/github.com/go-playground/validator#InvalidValidationError
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown Error"})
			}

			// vErr is serializable because it has struct tags!
			vErr := rerrors.NewFromValidationErrors(err.(validator.ValidationErrors))
			c.JSON(vErr.Status, gin.H{"error": vErr})
			c.Abort()
			return
		}

		idTokenHeader := strings.Split(h.IDToken, "Bearer ")

		if len(idTokenHeader) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": rerrors.NewUnauthorized("Must provide Authorization header with format Bearer token"),
			})
			c.Abort()
			return
		}

		claims, err := e.TokenService.ValidateIDToken(idTokenHeader[1])

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": rerrors.NewUnauthorized("Provided token is invalid"),
			})
			c.Abort()
			return
		}

		c.Set("user", claims)

		c.Next()
	}
}
