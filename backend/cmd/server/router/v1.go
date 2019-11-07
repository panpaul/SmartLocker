package router

import (
	v1 "SmartLocker/cmd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func initAPIv1(apiv1 *gin.RouterGroup) {
	user := apiv1.Group("/user")
	user.POST("/register", v1.Register)
	user.POST("/login", v1.Login)
	user.POST("/check", v1.Check)
	user.POST("/info", v1.Info)

}
