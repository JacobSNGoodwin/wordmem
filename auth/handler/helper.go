package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

// bindData is helper function
func bindData(c *gin.Context, req interface{}) {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data\nErr: %v\n", err)
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
}
