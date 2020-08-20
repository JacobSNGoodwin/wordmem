package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

type updateReq struct {
	Name     string `json:"name" binding:"gte=1"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"gte=6,lte=30"`
	ImageURL string `json:"imageUrl"`
	Website  string `json:"website" binding:"url"`
}

// Update handler updates account information for a user
func (e *Env) Update(c *gin.Context) {
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
