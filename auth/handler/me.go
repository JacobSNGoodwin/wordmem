package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// Me handler fetches user from ID token
// so that user can be verified by the server and returned
func (e *Env) Me(c *gin.Context) {
	userClaims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errors.NewUnknown(http.StatusInternalServerError),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": userClaims.(*util.IDTokenCustomClaims),
	})
}
