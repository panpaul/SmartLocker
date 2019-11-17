package v1

import (
	"SmartLocker/e"
	"SmartLocker/service/article"
	"SmartLocker/service/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func OccupyArticle(c *gin.Context) {
	claim := c.MustGet("claim").(*auth.Claims)

	cid, err := strconv.Atoi(c.PostForm("cid"))
	if err != nil {
		c.JSON(http.StatusOK, Wrap(e.InternalError, nil))
		return
	}
	if cid == 0 {
		c.JSON(http.StatusOK, Wrap(e.InvalidParams, nil))
		return
	}

	a := article.Article{UserId: claim.Id, CabinetId: cid}
	errInt := a.RandomOccupy()
	c.JSON(http.StatusOK, Wrap(errInt, a))
}

func ReleaseArticle(c *gin.Context) {
	claim := c.MustGet("claim").(*auth.Claims)

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusOK, Wrap(e.InternalError, nil))
		return
	}
	if id == 0 {
		c.JSON(http.StatusOK, Wrap(e.InvalidParams, nil))
		return
	}

	a := article.Article{Id: id, UserId: claim.Id}
	errInt := a.Update(true)
	c.JSON(http.StatusOK, Wrap(errInt, nil))
}

func LockerInfo(c *gin.Context) {
	claim := c.MustGet("claim").(*auth.Claims)
	a, err := article.GetArticles(claim.Id)
	c.JSON(http.StatusOK, Wrap(err, a))
}
