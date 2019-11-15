package v1

import (
	"SmartLocker/service/cabinet"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCabinetLocations(c *gin.Context) {
	l, err := cabinet.GetLocations()
	c.JSON(http.StatusOK, Wrap(err, l))
	return
}

func GetCabinetsByLocation(c *gin.Context) {
	l := c.PostForm("location")
	cab, err := cabinet.GetCabinets(l)
	c.JSON(http.StatusOK, Wrap(err, cab))
}
