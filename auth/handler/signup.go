package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// signupReq is not exported
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup does what it says!
func (e *Env) Signup(c *gin.Context) {
	var req signupReq

	// Bind incoming json to struct - Need to create custom validation error
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := e.UserService.SignUp(&model.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(409, gin.H{
			"message": "Failed to sign up a new user",
		})
		return
	}

	c.JSON(200, resp)
}
