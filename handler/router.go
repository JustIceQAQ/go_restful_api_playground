package handler

import (
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup, handler *Handler) {
	JwtRouter(superRoute, handler)
	PingRouter(superRoute, handler)

	UsersRouter(superRoute, handler)
}

// Jwt API
func JwtRouter(superRoute *gin.RouterGroup, handler *Handler) {
	jwtRouter := superRoute.Group("/jwt")
	{
		jwtRouter.POST("/get_token", handler.JwtRetrieve)
		jwtRouter.GET("/captcha", handler.JwtCaptcha)

	}
}

// Ping API
func PingRouter(superRoute *gin.RouterGroup, handler *Handler) {
	pingRouter := superRoute.Group("/ping")
	{
		pingRouter.GET("/", handler.Ping)
	}
}

// User API
func UsersRouter(superRoute *gin.RouterGroup, handler *Handler) {
	usersRouter := superRoute.Group("/users")

	usersRouter.Use(VerifyToken)
	{
		usersRouter.GET("/", handler.UserList)
		usersRouter.POST("/", handler.UserCreate)
	}

	userRouter := superRoute.Group("/user")

	userRouter.Use(VerifyToken)
	{
		userRouter.GET("/:id", handler.UserRetrieve)
		userRouter.PUT("/:id", handler.UserUpdate)
		userRouter.DELETE("/:id", handler.UserDelete)
	}
}
