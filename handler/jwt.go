package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	Models "go_restful_api_playground/models"
	Utils "go_restful_api_playground/utils"
	"net/http"
	"time"
)

var jwtKey = []byte("QAQ")

type LoginUserBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

// JwtApp godoc
// @Summary Get Jwt Token
// @Schemes
// @Description Get Jwt Token
// @Tags Jwt
// @Accept json
// @Produce json
// @Param LoginUserBody body LoginUserBody true "account password"
// @Success 200 {object} Models.User
// @Router /jwt/get_token [post]
func (h *Handler) JwtRetrieve(c *gin.Context) {

	var user LoginUserBody
	if err := c.BindJSON(&user); err != nil {
		return
	}

	var existUser Models.User

	result := h.db.First(&existUser, "account = ?", user.Account)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Your account or password is error!",
		})
		return
	}

	if Utils.CheckPasswordHash(user.Password, existUser.Password) {
		token, _ := GenerateToken(existUser)
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Your account or password is error!",
		})
		return
	}

}

func GenerateToken(user Models.User) (string, error) {
	type CustomClaims struct {
		Account string `json:"account"`
		jwt.RegisteredClaims
	}

	claims := CustomClaims{
		user.Account, jwt.RegisteredClaims{
			Issuer:    "Go Restful API Demo",
			Subject:   "Go Restful API Demo",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtKey)
	fmt.Printf("%v %v", ss, err)
	return ss, err

}
