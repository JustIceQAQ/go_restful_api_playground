package main

import "github.com/gin-gonic/gin"
import pingApp "playground_api/api/v1/pingapp"

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	AddRoutes(v1)

	return r
}

func AddRoutes(superRoute *gin.RouterGroup) {
	pingApp.Routes(superRoute)
}
