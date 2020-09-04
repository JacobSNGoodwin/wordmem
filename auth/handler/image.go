package handler

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/rerrors"
	"github.com/jacobsngoodwin/wordmem/auth/util"
)

// Image handles storing a user image to cloud storage and returning the image url
func (e *Env) Image(c *gin.Context) {
	claims, exists := c.Get("user")

	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := rerrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	// I had a lot of trouble just adding this to a middleware. So I'm placing it here
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodySize)

	form, err := c.MultipartForm()

	if err != nil {
		// should be a validation error
		log.Printf("Unable parse mutlipart/form: %v", err)

		if err.Error() == "http: request body too large" {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"error": fmt.Sprintf("Max request body size is %v bytes\n", MaxBodySize),
			})
		} else {
			e := rerrors.NewBadRequest("Unable to parse multipart/form-data")
			c.JSON(e.Status(), gin.H{
				"error": e,
			})
			return
		}
	}

	files := form.File["imageFile"]

	var imageFile *multipart.FileHeader

	if len(files) > 0 {
		imageFile = files[0]
	}

	userClaims := claims.(*util.IDTokenCustomClaims)

	url, err := e.UserService.SetProfileImage(userClaims.User.UID, imageFile)

	if err != nil {
		log.Printf("Failed to update profile image: %v\n", err.Error())

		c.JSON(rerrors.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"imageUrl": url,
		"message":  "success",
	})
}
