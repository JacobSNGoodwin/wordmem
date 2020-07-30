package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// Signup does what it says!
func (e *Env) Signup(c *gin.Context) {
	// TODO - get email, password, and name from context (request body)
	// Perform validation here before creating a user model (uuid handled by postgres)

	user := &model.User{
		Email:    "bob@bob.com",
		Name:     "Jacob Goodwin III",
		Password: "blablabla",
	}

	u, err := e.UserService.SignUp(user)

	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(409, gin.H{
			"message": "Failed to sign up a new user",
		})
		return
	}

	c.JSON(200, u)
}
