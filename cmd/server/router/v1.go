package router

import (
	v1 "SmartLocker/cmd/server/router/v1"
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initAPIv1(apiv1 *gin.RouterGroup) {
	user := apiv1.Group("/user")
	user.GET("/hello", func(c *gin.Context) {
		err := model.AddUser("test", "test", 0, "")
		if err == nil {
			c.String(http.StatusOK, "123")
			return
		}
		c.String(http.StatusOK, err.Error())
	})
	user.GET("/test", func(c *gin.Context) {
		a, b := model.GetLockersByUid(1)
		if b != nil {
			c.JSON(http.StatusOK, v1.Wrap(e.InternalError, b))
			return
		}
		c.JSON(http.StatusOK, v1.Wrap(e.Success, a))
		return
	})

}
