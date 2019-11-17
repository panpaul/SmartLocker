package router

import (
	CabinetV1 "SmartLocker/cmd/server/router/cabinet/v1"
	HttpV1 "SmartLocker/cmd/server/router/http/v1"
	"SmartLocker/cmd/server/router/http/v1/middleware"
	"github.com/gin-gonic/gin"
)

func initAPIv1(apiV1 *gin.RouterGroup) {
	HttpApiV1 := apiV1.Group("/http")
	WSApiV1 := apiV1.Group("/cabinet")
	initHttpAPIv1(HttpApiV1)
	initCabinetAPIv1(WSApiV1)
}

func initHttpAPIv1(apiV1 *gin.RouterGroup) {
	user := apiV1.Group("/user")
	user.POST("/register", HttpV1.UserRegister)
	user.POST("/login", HttpV1.UserLogin)
	user.POST("/check", HttpV1.UserCheck)
	user.POST("/info", HttpV1.UserInfo)

	article := apiV1.Group("/article")
	article.Use(middleware.Jwt())
	article.POST("/occupy", HttpV1.OccupyArticle)
	article.POST("/release", HttpV1.ReleaseArticle)
	article.POST("/info", HttpV1.LockerInfo)

	face := apiV1.Group("/face")
	face.POST("/recognize", HttpV1.RecognizeFace)

	cabinet := apiV1.Group("/cabinet")
	cabinet.GET("/location", HttpV1.GetCabinetLocations)
	cabinet.POST("/cabinet", HttpV1.GetCabinetsByLocation)
}

func initCabinetAPIv1(apiV1 *gin.RouterGroup) {
	apiV1.GET("/ping", CabinetV1.PingPong)
	apiV1.GET("/register", CabinetV1.Register)
	apiV1.GET("/generateToken", CabinetV1.GenerateToken)
}
