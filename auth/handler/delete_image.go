package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// DeleteImage clears the image of a user resource
func (e *Env) DeleteImage(c *gin.Context) {
	claims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := rerrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	userClaims := claims.(*util.IDTokenCustomClaims)

	err := e.UserService.ClearProfileImage(userClaims.User.UID)

	if err != nil {
		log.Printf("Failed to delete profile image: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
