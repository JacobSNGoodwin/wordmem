package main

import (
	"github.com/gin-gonic/gin"
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

	// configure simple CORS - can also proxy in development on front end
	// for vue, see CLI config ref
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://127.0.0.1:3000", "http://127.0.0.1:3001"}
	// r.Use(cors.New(config))

	// set MaxBodySize to 4 MB - set directly in handler. Had trouble in middleware
	// r.Use(h.LimitBodySize(handler.MaxBodySize))
	// r.Use(limits.RequestSizeLimiter(handler.MaxBodySize))

	// consider using env variable
	baseGroup := r.Group("/api/account")

	baseGroup.GET("/me", h.AuthUser(), h.Me)
	baseGroup.POST("/signup", h.Signup)
	baseGroup.POST("/signin", h.Signin)
	baseGroup.POST("/tokens", h.Tokens)
	baseGroup.POST("/signout", h.AuthUser(), h.Signout)
	baseGroup.POST("/image", h.AuthUser(), h.Image)
	baseGroup.DELETE("/image", h.AuthUser(), h.DeleteImage)
	baseGroup.PUT("/details", h.AuthUser(), h.Details)

	router.r = r
}
