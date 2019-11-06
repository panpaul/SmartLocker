package router

import (
	"SmartLocker/config"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	// Setup a gin instance
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// config
	r.MaxMultipartMemory = 8 << 20
	gin.SetMode(config.Conf.Mode)

	// adding pprof
	if config.Conf.Mode == "debug" {
		pprof.Register(r)
	}

	// register static file routes
	r.StaticFS("/resources", http.Dir("resources"))

	// handle api v1
	apiv1 := r.Group("/api/v1")
	initAPIv1(apiv1)

	// return the instance
	return r
}
