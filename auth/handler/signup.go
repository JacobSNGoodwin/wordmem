package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/model"
)

// Signup does what it says!
func (e *Env) Signup(c *gin.Context) {
	// TODO - get email and password from context (request body)
	// Perform validation here before creating a user model

	uid, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate a random uuid")
		c.JSON(409, gin.H{
			"message": "Failed to sign up a new user",
		})
	}

	user := &model.User{
		UID:   uid,
		Email: "bob@bob.com",
		// Name:     "Jacob Goodwin III",
		Password: "",
	}

	if err = e.UserService.SignUp(user); err != nil {
		c.JSON(409, gin.H{
			"message": "Failed to sign up a new user",
		})
	}

	c.JSON(200, user)
}
