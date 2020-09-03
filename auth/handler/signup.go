package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
)

// signupReq is not exported
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup does what it says!
func (e *Env) Signup(c *gin.Context) {
	var req signupReq

	// Bind incoming json to struct and check for validation errors

	if ok := bindData(c, &req); !ok {
		return
	}

	u, err := e.UserService.SignUp(req.Email, req.Password)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	tokens, err := e.TokenService.NewPairFromUser(u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())
		log.Printf("Rolling back user creation for user: %v\n", u)

		// rollback - add good unit test because
		// this is tough to recreate
		e.UserService.Delete(u.UID)

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
	})
}
