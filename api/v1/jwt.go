package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	orm "go_restful_api_playground/database"
	Models "go_restful_api_playground/models"
	Utils "go_restful_api_playground/utils"
	"net/http"
	"strings"
	"time"
)

var jwtKey = []byte("QAQ")

type LoginUserBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type UserInfoBody struct {
	ID       uint   `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
}

// JwtApp godoc
// @Summary Get Jwt Token
// @Schemes
// @Description Get Jwt Token
// @Tags Jwt
// @Accept json
// @Produce json
// @Param LoginUserBody body LoginUserBody true "account password"
// @Success 200 {string} json "{"token": "token"}"
// @Router /jwt/get_token [post]
func JwtRetrieve(c *gin.Context) {

	var user LoginUserBody
	if err := c.BindJSON(&user); err != nil {
		return
	}

	var existUser Models.User

	result := orm.Db.First(&existUser, "account = ?", user.Account)

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

// JwtApp godoc
// @Summary Captcha Jwt Token
// @Schemes
// @Description Captcha Jwt Token
// @Tags Jwt
// @Accept json
// @Produce json
// @Success 200 {object} CustomClaims
// @Router /jwt/captcha [get]
// @Security BearerAuth
func JwtCaptcha(c *gin.Context) {
	tokenString, _ := getToken(c)
	isOk, err, tokenDetail := ValidateToken(tokenString, true)
	fmt.Println(isOk, err, tokenDetail)
	if !isOk {
		c.JSON(http.StatusUnauthorized, gin.H{
			"value": http.StatusText(http.StatusUnauthorized),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"value": tokenDetail,
	})

}

type CustomClaims struct {
	UserInfoBody
	jwt.RegisteredClaims
}

func GenerateToken(user Models.User) (string, error) {
	userInfo := UserInfoBody{user.ID, user.Account, user.Username}

	claims := CustomClaims{
		userInfo, jwt.RegisteredClaims{
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
	return ss, err

}

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	result, err, _ := ValidateToken(token, false)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	if result {
		c.Next()
	}
}

func ValidateToken(tokenString string, isTokenDetail bool) (bool, error, map[string]interface{}) {

	var claims CustomClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return false, err, nil
	}
	if !token.Valid {
		return false, errors.New("invalid token"), nil
	}
	if isTokenDetail {
		dict := map[string]interface{}{
			"token":     token,
			"user-info": claims.UserInfoBody,
		}
		return true, nil, dict
	} else {
		return true, nil, nil
	}

}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
