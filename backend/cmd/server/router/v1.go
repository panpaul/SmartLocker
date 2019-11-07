package router

import (
	v1 "SmartLocker/cmd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func initAPIv1(apiv1 *gin.RouterGroup) {
	user := apiv1.Group("/user")
	user.GET("/register", v1.Register)
	user.GET("/login", v1.Login)
}
