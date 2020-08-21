package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

// signinReq is not exported
type signinReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signin used to authenticate extant user
func (e *Env) Signin(c *gin.Context) {
	var req signinReq

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(&req); err != nil {

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
