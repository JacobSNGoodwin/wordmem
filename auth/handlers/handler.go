package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/repository"
)

// Handler extending gin so handlers can hold repo reference
type Handler struct {
	*gin.Engine
	repo repository.Repository
}
