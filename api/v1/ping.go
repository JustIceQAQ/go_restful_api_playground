package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"go_restful_api_playground/utils"
	"net/http"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, _ := utils.Cld.Admin.Ping(ctx)

	c.JSON(http.StatusOK, gin.H{
		"API":        "pong",
		"Cloudinary": resp,
	})
}
