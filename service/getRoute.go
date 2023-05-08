package service

import (
	"Screenshots/module"
	"Screenshots/util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoute(c *gin.Context) {
	names := util.GetName()
	var screen []module.ScreenshotModule
	for _, v := range names[:] {
		scr := util.GetScreenshotScene(v)
		screen = append(screen, scr)
	}
	data, _ := json.Marshal(screen)
	c.String(http.StatusOK, string(data))
}
