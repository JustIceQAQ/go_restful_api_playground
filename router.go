package main

import (
	"github.com/gin-gonic/gin"
	v1 "go_restful_api_playground/api/v1"
)

func Routes(superRoute *gin.RouterGroup) {
	JwtRouter(superRoute)
	PingRouter(superRoute)
	UsersRouter(superRoute)
	UploadsRouter(superRoute)
}

// Jwt API
func JwtRouter(superRoute *gin.RouterGroup) {
	jwtRouter := superRoute.Group("/jwt")
	{
		jwtRouter.POST("/get_token", v1.JwtRetrieve)
		jwtRouter.GET("/captcha", v1.JwtCaptcha)

	}
}

// Ping API
func PingRouter(superRoute *gin.RouterGroup) {
	pingRouter := superRoute.Group("/ping")
	{
		pingRouter.GET("/", v1.Ping)
	}
}

// User API
func UsersRouter(superRoute *gin.RouterGroup) {
	usersRouter := superRoute.Group("/users")

	usersRouter.Use(v1.VerifyToken)
	{
		usersRouter.GET("/", v1.UserList)
		usersRouter.POST("/", v1.UserCreate)
	}

	userRouter := superRoute.Group("/user")

	userRouter.Use(v1.VerifyToken)
	{
		userRouter.GET("/:id", v1.UserRetrieve)
		userRouter.PUT("/:id", v1.UserUpdate)
		userRouter.DELETE("/:id", v1.UserDelete)
	}
}

// Uploads API
func UploadsRouter(superRoute *gin.RouterGroup) {
	// uploadsRouter.Use(v1.VerifyToken)

	uploadsRouter := superRoute.Group("/upload")
	{
		uploadsRouter.POST("/", v1.UploadFile)
	}
}
