package handler

import (
	"github.com/gin-gonic/gin"
)

func Routes(superRoute *gin.RouterGroup, handler *Handler) {
	pingRouter := superRoute.Group("/ping")
	{
		pingRouter.GET("/", handler.Ping)
	}
	ping2Router := superRoute.Group("/ping2")
	{
		ping2Router.GET("/", handler.Ping2)
	}
	usersRouter := superRoute.Group("/users")
	{
		usersRouter.GET("/", handler.UserList)
	}
	userRouter := superRoute.Group("/user")
	{
		userRouter.GET("/:id", handler.UserRetrieve)
	}
}
