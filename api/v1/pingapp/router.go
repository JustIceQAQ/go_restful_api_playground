package pingapp

import "github.com/gin-gonic/gin"

func Routes(superRoute *gin.RouterGroup) {
	pingRouter := superRoute.Group("/ping")
	{
		pingRouter.GET("/", Ping)
	}
	ping2Router := superRoute.Group("/ping2")
	{
		ping2Router.GET("/", Ping2)
	}
}
