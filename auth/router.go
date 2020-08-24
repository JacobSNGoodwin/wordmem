package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jacobsngoodwin/wordmem/auth/handler"
)

// Router holds a reference to a router with access to services
// container in handler.Env
type Router struct {
	r *gin.Engine
}

// Init sets up route controllers by providing them access to application services
func (router *Router) Init(ic *InjectionContainer) {
	// Injection of database to create concrete Service and Repository impls
	h := ic.handlerEnv

	r := gin.Default()

	// set MaxBodySize to 4 MB
	r.Use(h.LimitBodySize(handler.MaxBodySize))

	r.GET("/me", h.AuthUser(), h.Me)
	r.PUT("/details", h.AuthUser(), h.Details)
	r.POST("/signup", h.Signup)
	r.POST("/signin", h.Signin)
	r.POST("/tokens", h.Tokens)
	r.POST("/signout", h.AuthUser(), h.Signout)

	router.r = r
}
