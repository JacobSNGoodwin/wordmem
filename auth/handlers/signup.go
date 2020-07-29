package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jacobsngoodwin/wordmem/auth/models"
)

// Signup does what it says!
func Signup(c *gin.Context) {
	uid, err := uuid.NewRandom()

	if err != nil {
		log.Println("Failed to generate a random uuid")
		c.JSON(409, gin.H{
			"message": "Failed to sign up a new user",
		})
	}

	user := &models.User{
		UID:   uid,
		Email: "bob@bob.com",
		// Name:     "Jacob Goodwin III",
		Password: "",
	}

	c.JSON(200, user)
}
