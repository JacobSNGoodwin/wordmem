package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
)

// signinReq is not exported
type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signin used to authenticate extant user
func (e *Env) Signin(c *gin.Context) {
	var req signinReq

	if ok := bindData(c, &req); !ok {
		return
	}

	u, err := e.UserService.SignIn(req.Email, req.Password)

	if err != nil {
		log.Printf("Failed to sign in user: %v\n", err.Error())
		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := e.TokenService.NewPairFromUser(u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}
