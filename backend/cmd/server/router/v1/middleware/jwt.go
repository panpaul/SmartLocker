package middleware

import (
	v1 "SmartLocker/cmd/server/router/v1"
	"SmartLocker/e"
	"SmartLocker/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		claim, err := auth.CheckToken(token)
		if err != e.Success {
			c.JSON(http.StatusOK, v1.Wrap(err, nil))
			c.Abort()
			return
		}
		c.Set("claim", claim)
		c.Next()
	}
}
