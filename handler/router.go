package handler

import (
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup, handler *Handler) {
	PingRouter(superRoute, handler)
	UsersRouter(superRoute, handler)
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
	{
		usersRouter.GET("/", handler.UserList)
		usersRouter.POST("/", handler.UserCreate)
	}
	userRouter := superRoute.Group("/user")
	{
		userRouter.GET("/:id", handler.UserRetrieve)
		userRouter.PUT("/:id", handler.UserUpdate)
		userRouter.DELETE("/:id", handler.UserDelete)
	}
}
