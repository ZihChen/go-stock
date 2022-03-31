package jwt

import (
	"github.com/gin-gonic/gin"
	"go-stock/app/http/response"
	"go-stock/app/service/jwtservice"
	"strings"
)

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		response.WrapContext(c).Error(401, "Get token error")
		c.Abort()
		return
	}

	jwtService := jwtservice.NewService()

	id, username, err := jwtService.ValidateToken(token)
	if err != nil {
		response.WrapContext(c).Error(401, err.Error())
		c.Abort()
		return
	}

	c.Set("id", id)
	c.Set("username", username)

	c.Next()
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
