package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	AddRoutes(v1)

	return r
}

func AddRoutes(superRoute *gin.RouterGroup) {
	Routes(superRoute)
}
