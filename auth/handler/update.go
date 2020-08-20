package handler

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jacobsngoodwin/wordmem/auth/errors"
)

// omitempty must be required in binding (not just omitting required)
// this is because even though the field isn't required, the other validation
// will be run
// omitempty must also be listed first (tags evaluated sequentially, I guess)
type updateReq struct {
	Name      string                `form:"name" binding:"omitempty,gte=1"`
	Email     string                `form:"email" binding:"omitempty,email"`
	Password  string                `form:"password" binding:"omitempty,gte=6,lte=30"`
	ImageFile *multipart.FileHeader `form:"imageFile" binding:"omitempty,file"`
	Website   string                `form:"website" binding:"omitempty,url"`
}

// Update handler updates account information for a user
func (e *Env) Update(c *gin.Context) {
	// _, exists := c.Get("user")

	// if !exists {
	// 	log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": errors.NewUnknown(http.StatusInternalServerError),
	// 	})

	// 	return
	// }

	var req updateReq

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

	c.JSON(http.StatusOK, gin.H{
		"name":     req.Name,
		"fileName": req.ImageFile.Filename,
	})
}
