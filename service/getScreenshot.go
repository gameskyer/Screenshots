package service

import (
	"Screenshots/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetScreenshot(c *gin.Context) {
	name := c.PostForm("name")
	scene := c.PostForm("scene")

	screenshots := dao.SelectScreenshotBySceneAndName(name, scene)
	screenshot, _ := json.Marshal(screenshots)
	c.String(http.StatusOK, string(screenshot))
}
