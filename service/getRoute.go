package service

import (
	"Screenshots/dao"
	"Screenshots/module"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoute(c *gin.Context) {
	names := dao.GetName()
	var screen []module.ScreenshotModule
	for _, v := range names[:] {
		scr := dao.GetScreenshotScene(v)
		screen = append(screen, scr)
	}
	data, _ := json.Marshal(screen)
	c.String(http.StatusOK, string(data))
}
