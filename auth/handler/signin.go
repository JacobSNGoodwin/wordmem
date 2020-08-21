package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signinReq is not exported
type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signin used to authenticate extant user
func (e *Env) Signin(c *gin.Context) {
	var req signinReq

	bindData(c, &req)

	u, err := e.UserService.SignIn(req.Email, req.Password)

	if err != nil {
		//
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	tokens, err := e.TokenService.NewPairFromUser(u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":   u,
		"tokens": tokens,
	})
}
