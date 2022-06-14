package pingapp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary Ping example
// @Schemes
// @Description Doing pingapp
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /pingapp [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingExample godoc
// @Summary Ping2 example
// @Schemes
// @Description Doing pingapp
// @Tags Ping2
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping2 [get]
func Ping2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong2",
	})
}
