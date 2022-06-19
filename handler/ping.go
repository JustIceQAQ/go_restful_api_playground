package handler

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
func (h *Handler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
