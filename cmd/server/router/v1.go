package router

import (
	"SmartLocker/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initAPIv1(apiv1 *gin.RouterGroup) {
	user := apiv1.Group("/user")
	user.GET("/hello", func(c *gin.Context) {
		err := database.AddUser("test", "test", 0, "")
		if err == nil {
			c.String(http.StatusOK, "123")
			return
		}
		c.String(http.StatusOK, err.Error())
	})

}
