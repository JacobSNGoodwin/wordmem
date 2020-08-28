package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// Me handler fetches user from database with most up-to-date info
// this allows user to have up-to-date resource for updating
// their profile data. The token could be outdated
func (e *Env) Me(c *gin.Context) {
	userClaims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": rerrors.NewUnknown(http.StatusInternalServerError),
		})

		return
	}

	uid := userClaims.(*util.IDTokenCustomClaims).User.UID
	u, err := e.UserService.Get(uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n", uid)
		c.JSON(http.StatusNotFound, gin.H{
			"error": rerrors.NewNotFound("user", "idToken"),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
