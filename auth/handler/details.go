package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// omitempty must be required in binding (not just omitting required)
// this is because even though the field isn't required, the other validation
// will be run
// omitempty must also be listed first (tags evaluated sequentially, I guess)
type detailsReq struct {
	Name    string `json:"name" binding:"omitempty,alphanum"`
	Email   string `json:"email" binding:"omitempty,email"`
	Website string `json:"website" binding:"omitempty,url"`
}

// Details handler updates account information for a user
func (e *Env) Details(c *gin.Context) {
	claims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": rerrors.NewUnknown(http.StatusInternalServerError),
		})

		return
	}

	var req detailsReq

	if ok := bindData(c, &req); !ok {
		return
	}

	userClaims := claims.(*util.IDTokenCustomClaims)

	u := model.User{
		UID:     userClaims.UID,
		Name:    req.Name,
		Email:   req.Email,
		Website: req.Website,
	}

	err := e.UserService.UpdateDetails(&u)

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
