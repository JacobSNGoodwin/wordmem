package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
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

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBindJSON(&req); err != nil {

		// this type check appears to be extra cautious as I could not
		// find a case where this error was anything other than InvalidValidationError
		// see https://godoc.org/github.com/go-playground/validator#InvalidValidationError
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown Error"})
		}

		// vErr is serializable because it has struct tags!
		vErr := errors.NewFromValidationErrors(err.(validator.ValidationErrors))
		c.JSON(vErr.Status, gin.H{"error": vErr})

		return
	}

	resp, err := e.UserService.SignUp(&model.User{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		//
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
