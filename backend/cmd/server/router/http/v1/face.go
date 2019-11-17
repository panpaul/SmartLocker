package v1

import (
	"SmartLocker/e"
	"SmartLocker/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RecognizeFace(c *gin.Context) {
	file, err := c.FormFile("img")
	if err != nil {
		c.JSON(http.StatusOK, Wrap(e.UploadFailed, nil))
		return
	}

	// TODO check content-type locally
	ext, errInt := util.VerifyImg(file.Header.Get("Content-Type"))
	if errInt != e.Success {
		c.JSON(http.StatusOK, Wrap(errInt, nil))
		return
	}

	filename := fmt.Sprintf("%s.%s",
		util.EncodeSha256(file.Filename+util.RandString(5)), ext)

	if err := c.SaveUploadedFile(file,
		fmt.Sprintf("./resources/uploads/%s", filename)); err != nil {
		c.JSON(http.StatusOK, Wrap(e.UploadFailed, nil))
		return
	}

	url := fmt.Sprintf("/resources/uploads/%s", filename)
	c.JSON(http.StatusOK, Wrap(e.Success, url))
	return
	//TODO send the file to verify server
}
