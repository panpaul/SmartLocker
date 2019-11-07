package v1

import (
	"SmartLocker/e"
	"github.com/gin-gonic/gin"
)

func wrap(err int, body interface{}) gin.H {
	if err != e.Success {
		return gin.H{
			"code": err,
			"msg":  e.GetMsg(err),
		}
	}
	return gin.H{
		"code": e.Success,
		"msg":  e.GetMsg(e.Success),
		"body": body,
	}
}
