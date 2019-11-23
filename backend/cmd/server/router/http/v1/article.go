package v1

import (
	"SmartLocker/e"
	"SmartLocker/service/article"
	"SmartLocker/service/auth"
	"SmartLocker/service/task"
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
	if errInt != e.Success {
		c.JSON(http.StatusOK, Wrap(errInt, nil))
		return
	}

	errInt = a.Fill()
	if errInt != e.Success {
		c.JSON(http.StatusOK, Wrap(errInt, nil))
		return
	}

	task.AddBackendTask(strconv.Itoa(cid), 0, a.Position, a.UserId)

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
	// if do it later->uid=null
	errInt := a.Fill()
	if errInt != e.Success {
		c.JSON(http.StatusOK, Wrap(errInt, nil))
		return
	}

	errInt = a.Update(true)
	if errInt != e.Success {
		c.JSON(http.StatusOK, Wrap(errInt, nil))
		return
	}

	task.AddBackendTask(strconv.Itoa(a.CabinetId), 0, a.Position, a.UserId)

	c.JSON(http.StatusOK, Wrap(errInt, nil))
}

func LockerInfo(c *gin.Context) {
	claim := c.MustGet("claim").(*auth.Claims)
	a, err := article.GetArticles(claim.Id)
	c.JSON(http.StatusOK, Wrap(err, a))
}
