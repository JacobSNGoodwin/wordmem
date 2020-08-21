package handler

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
	"github.com/jacobsngoodwin/wordmem/auth/service"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// omitempty must be required in binding (not just omitting required)
// this is because even though the field isn't required, the other validation
// will be run
// omitempty must also be listed first (tags evaluated sequentially, I guess)
type updateReq struct {
	Name      string                `form:"name" binding:"omitempty,gte=1"`
	Email     string                `form:"email" binding:"omitempty,email"`
	Password  string                `form:"password" binding:"omitempty,gte=6,lte=30"`
	ImageFile *multipart.FileHeader `form:"imageFile" binding:"omitempty,file"`
	Website   string                `form:"website" binding:"omitempty,url"`
}

// Update handler updates account information for a user
func (e *Env) Update(c *gin.Context) {
	claims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": errors.NewUnknown(http.StatusInternalServerError),
		})

		return
	}

	var req updateReq

	if ok := bindData(c, &req); !ok {
		return
	}

	userClaims := claims.(*util.IDTokenCustomClaims)

	updateOptions := service.UpdateOptions(req)

	u, err := e.UserService.Update(userClaims.UID, &updateOptions)

	if err != nil {
		log.Printf("Failed to update user: %v\n", err.Error())

		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
