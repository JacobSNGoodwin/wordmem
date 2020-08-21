package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// LimitBodySize limits the size of requests to prevent malicious
// or unwanted large requests
func (e *Env) LimitBodySize(size int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, size)
		c.Next()
	}
}
