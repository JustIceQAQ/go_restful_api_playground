package handler

import (
	"github.com/gin-gonic/gin"
	Models "go_restful_api_playground/models"
	"net/http"
)

// UserApp godoc
// @Summary List User
// @Schemes
// @Description GET All Users
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {array} Models.User
// @Router /users [get]
func (h *Handler) UserList(c *gin.Context) {

	var users []Models.User
	if result := h.db.Find(&users); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, &users)
}

// UserApp godoc
// @Summary Retrieve User
// @Schemes
// @Description GET One User
// @Tags User
// @Accept json
// @Produce json
// @param id path int true "id"
// @Success 200 {object} Models.User
// @Router /user/{id} [get]
func (h *Handler) UserRetrieve(c *gin.Context) {
	id := c.Param("id")

	var user Models.User
	if result := h.db.First(&user, id); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, &user)
}
