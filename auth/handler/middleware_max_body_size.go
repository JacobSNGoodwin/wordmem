package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaxBodySize limits the size of requests to prevent malicious
// or unwanted large requests
func (e *Env) MaxBodySize(sizeMB int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, sizeMB*1024*1024)
		c.Next()
	}
}
