package main

import "github.com/gin-gonic/gin"

func initRouter() *gin.Engine {
	// Injection of database to create concrete Service and Repository impls

	ic := InjectionContainer{}
	handler := ic.Init()

	r := gin.Default()
	r.POST("/signup", handler.Signup)

	return r
}
