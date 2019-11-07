package router

import (
	v1 "SmartLocker/cmd/server/router/v1"
	"SmartLocker/cmd/server/router/v1/middleware"
	"github.com/gin-gonic/gin"
)

func initAPIv1(apiV1 *gin.RouterGroup) {
	user := apiV1.Group("/user")
	user.POST("/register", v1.UserRegister)
	user.POST("/login", v1.UserLogin)
	user.POST("/check", v1.UserCheck)
	user.POST("/info", v1.UserInfo)

	article := apiV1.Group("/article")
	article.Use(middleware.Jwt())
	article.POST("/occupy", v1.OccupyArticle)
	article.POST("/release", v1.ReleaseArticle)
	article.POST("/info", v1.LockerInfo)

}
