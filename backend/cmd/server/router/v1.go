package router

import (
	"SmartLocker/cmd/server/router/v1"
	"github.com/gin-gonic/gin"
)

func initAPIv1(apiv1 *gin.RouterGroup) {
	user := apiv1.Group("/user")
	v1.InitAPIV1User(user)
}
