package main

import "github.com/gin-gonic/gin"

// Router holds a reference to a router with access to services
// container in handler.Env
type Router struct {
	r *gin.Engine
}

// Init sets up route controllers by providing them access to application services
func (router *Router) Init(ic *InjectionContainer) {
	// Injection of database to create concrete Service and Repository impls
	handler := ic.handlerEnv

	r := gin.Default()

	r.POST("/signup", handler.Signup)

	router.r = r
}
