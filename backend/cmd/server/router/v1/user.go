package v1

import (
	"SmartLocker/e"
	"SmartLocker/service/auth"
	"SmartLocker/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, wrap(e.InvalidParams, nil))
		return
	}

	u := user.User{Username: username, Password: password}
	err := u.Register()
	c.JSON(http.StatusOK, wrap(err, nil))
}

func Login(c *gin.Context) { //密码是明文
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusOK, wrap(e.InvalidParams, nil))
		return
	}

	//验证用户密码
	u := user.User{Username: username, Password: password}
	t, err := u.Verify()

	if err != e.Success {
		c.JSON(http.StatusOK, wrap(err, nil))
		return
	}
	if t == false {
		c.JSON(http.StatusOK, wrap(e.Unauthorized, nil))
	}

	//生成jwt token
	expTime := time.Now().Add(12 * time.Hour).Unix()
	claim := auth.Claims{Username: u.Username, Role: u.Role, Id: u.Id}
	claim.ExpiresAt = expTime

	token := claim.GenerateToken()

	if token == "" {
		c.JSON(http.StatusOK, wrap(e.InternalError, nil))
		return
	}

	c.JSON(http.StatusOK, wrap(e.Success, token))

}
