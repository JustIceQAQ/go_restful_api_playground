package main

import (
	"github.com/gin-gonic/gin"
	handler "go_restful_api_playground/handler"
)

func setupRouter(handlers *handler.Handler) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	AddRoutes(v1, handlers)

	return r
}

func AddRoutes(superRoute *gin.RouterGroup, handlers *handler.Handler) {
	handler.Routes(superRoute, handlers)
}
