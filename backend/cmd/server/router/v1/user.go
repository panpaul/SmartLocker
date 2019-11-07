package v1

import (
	"SmartLocker/e"
	"SmartLocker/model"
	"SmartLocker/service/auth"
	"SmartLocker/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		return
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

func Check(c *gin.Context) {
	token := c.PostForm("token")
	_, err := auth.CheckToken(token)
	c.JSON(http.StatusOK, wrap(err, nil))
}

func Info(c *gin.Context) {
	username := c.PostForm("username")
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		id = 0
	}

	if username == "" && id == 0 {
		c.JSON(http.StatusOK, wrap(e.InvalidParams, nil))
		return
	}
	if username != "" && id != 0 { // 两个参数二选一
		c.JSON(http.StatusOK, wrap(e.InvalidParams, nil))
		return
	}

	// 检查权限
	token := c.PostForm("token")
	claim, errInt := auth.CheckToken(token)
	if errInt != e.Success {
		c.JSON(http.StatusOK, wrap(errInt, nil))
		return
	}

	if claim.Role != model.ADMIN && !(claim.Id == id || claim.Username == username) {
		c.JSON(http.StatusOK, wrap(e.PermissionDenied, "Only admin can access it"))
		return
	}

	// 获取用户信息
	u := user.User{Username: username, Id: id}
	errInt = u.Get()

	c.JSON(http.StatusOK, wrap(errInt, u))
}
