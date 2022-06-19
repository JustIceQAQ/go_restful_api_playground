package handler

import (
	"github.com/gin-gonic/gin"
	Models "go_restful_api_playground/models"
	"net/http"
)

// UserApp godoc
// @Summary List User
// @Schemes
// @Description GET Users
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
// @Description GET User
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

type UserBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// UserApp godoc
// @Summary Create User
// @Schemes
// @Description POST User
// @Tags User
// @Accept json
// @Produce json
// @Param UserBody body UserBody true "account password username"
// @Success 200 {object} Models.User
// @Router /users [post]
func (h *Handler) UserCreate(c *gin.Context) {
	var user Models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	if result := h.db.Create(&user); result.Error != nil {
		return
	}

	c.JSON(http.StatusCreated, &user)

}

// UserApp godoc
// @Summary Update User
// @Schemes
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @param id path int true "id"
// @Param UserBody body UserBody true "account password username"
// @Success 200 {object} Models.User
// @Router /user/{id} [put]
func (h *Handler) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	var existUser Models.User
	if result := h.db.First(&existUser, id); result.Error != nil {
		return
	}

	var user UserBody
	if err := c.BindJSON(&user); err != nil {
		return
	}

	existUser.Account = user.Account
	existUser.Password = user.Password
	existUser.Username = user.Username

	h.db.Save(&existUser)

	c.JSON(http.StatusOK, &existUser)

}

// UserApp godoc
// @Summary Delete User
// @Schemes
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @param id path int true "id"
// @Success 200 {object} Models.User
// @Router /user/{id} [delete]
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")

	if result := h.db.Delete(&Models.User{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)

}
