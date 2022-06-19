package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func Routes(superRoute *gin.RouterGroup, handler *Handler) {
	JwtRouter(superRoute, handler)
	PingRouter(superRoute, handler)

	UsersRouter(superRoute, handler)
}

func verifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	result, err := validateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	if result {
		c.Next()
	}
}

func validateToken(tokenString string) (bool, error) {
	type authClaims struct {
		jwt.RegisteredClaims
	}
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return false, err
	}
	if !token.Valid {
		return false, errors.New("invalid token")
	}
	return false, nil
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

// Jwt API
func JwtRouter(superRoute *gin.RouterGroup, handler *Handler) {
	jwtRouter := superRoute.Group("/jwt")
	{
		jwtRouter.POST("/get_token", handler.JwtRetrieve)
	}
}

// Ping API
func PingRouter(superRoute *gin.RouterGroup, handler *Handler) {
	pingRouter := superRoute.Group("/ping")
	{
		pingRouter.GET("/", handler.Ping)
	}
}

// User API
func UsersRouter(superRoute *gin.RouterGroup, handler *Handler) {
	usersRouter := superRoute.Group("/users")

	usersRouter.Use(verifyToken)
	{
		usersRouter.GET("/", handler.UserList)
		usersRouter.POST("/", handler.UserCreate)
	}

	userRouter := superRoute.Group("/user")

	userRouter.Use(verifyToken)
	{
		userRouter.GET("/:id", handler.UserRetrieve)
		userRouter.PUT("/:id", handler.UserUpdate)
		userRouter.DELETE("/:id", handler.UserDelete)
	}
}
