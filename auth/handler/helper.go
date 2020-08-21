package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %v\n", err)
		// this type check appears to be extra cautious as I could not
		// find a case where this error was anything other than InvalidValidationError
		// see https://godoc.org/github.com/go-playground/validator#InvalidValidationError
		if _, ok := err.(*validator.InvalidValidationError); ok {
			err := errors.NewUnknown(http.StatusBadRequest)
			c.JSON(err.Status, gin.H{
				"error": err,
			})
			return false
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			// vErr is serializable because it has struct tags!
			vErr := errors.NewFromValidationErrors(err)
			c.JSON(vErr.Status, gin.H{"error": vErr})
			return false
		}

		if err.Error() == "http: request body too large" {
			err := errors.NewExceedsMaxSize(MaxBodySize, c.Request.ContentLength)
			c.JSON(err.Status, gin.H{"error": err})
			return false
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.NewUnknown(http.StatusInternalServerError)})
		return false
	}

	return true
}
