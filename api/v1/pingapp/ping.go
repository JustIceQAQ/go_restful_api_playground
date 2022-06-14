package pingapp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingApp godoc
// @Summary Ping example
// @Schemes
// @Description Doing Ping
// @Tags Ping
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "pong"}"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// PingApp godoc
// @Summary Ping2 example
// @Schemes
// @Description Doing Ping2
// @Tags Ping2
// @Accept json
// @Produce json
// @Success 200 {string} json "{"message": "pong2"}"
// @Router /ping2 [get]
func Ping2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong2",
	})
}
