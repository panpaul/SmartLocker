package v1

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitAPIV1User(user *gin.RouterGroup) {
	user.GET("/test", func(c *gin.Context) {
		a, b := model.GetCabinetsByLocation("L1")
		if b != nil {
			c.JSON(http.StatusOK, wrap(e.InternalError, b))
			return
		}
		c.JSON(http.StatusOK, wrap(e.Success, a))
		return
	})
}
