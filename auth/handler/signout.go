package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// Signout signs user out be invalidating all
// of a users refresh tokens
func (e *Env) Signout(c *gin.Context) {
	claims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": rerrors.NewUnknown(http.StatusInternalServerError),
		})

		return
	}

	userClaims := claims.(*util.IDTokenCustomClaims)
	uid := userClaims.User.UID.String()

	if err := e.TokenService.SignOut(uid); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
