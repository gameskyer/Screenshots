package main

import (
	"Screenshots/service"
	"Screenshots/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type picture struct {
	Name  string
	Dir   string
	ToDir []dirtwo
}
type dirtwo struct {
	DirTwo string
	File   []string
}

const fileTmp = "C:\\Users\\pky\\go\\src\\Screenshots\\screenshot"

func main() {
	router := gin.Default()
	router.Use(Cors())
	//router.LoadHTMLGlob("template/**/**/**/*")
	//router.LoadHTMLGlob("template/*")
	//router.StaticFS("/template",http.Dir("./template"))
	//router.POST("/getFilePath", func(c *gin.Context) {
	//	var pictures []picture
	//	var picture picture
	//	var dirs []string
	//	var todirs []string
	//	var file []string
	//	var paths []string
	//	var DirTwo dirtwo
	//
	//	filepath.Walk("template\\Pictures", func(path string, info os.FileInfo, err error) error {
	//		paths = append(paths, path)
	//		return nil
	//	})
	//	fmt.Println("文件目录：",paths)
	//	for _, v := range paths {
	//		s := strings.Split(v, "\\")
	//		if len(s) == 3{
	//			dirs = append(dirs,v)
	//		}else if len(s) == 4 {
	//			todirs = append(todirs,v)
	//		}else if len(s) == 5 {
	//			file = append(file,v)
	//		}
	//	}
	//	fmt.Printf("i=%v,j=%v,k=%v \n", len(dirs), len(todirs), len(file))
	//	for i:= 0 ;i<= len(dirs)-1;i++{
	//		var DirTwos []dirtwo
	//		for j:= 0 ;j<= len(todirs)-1;j++{
	//			var file1 []string
	//			todir := todirs[j]
	//			dir := dirs[i]
	//			if strings.Index(todir,dir) != -1{
	//				DirTwo.DirTwo = todirs[j]
	//			}else {
	//				continue
	//			}
	//			for k:= 0 ;k<= len(file)-1;k++{
	//				if strings.Index(file[k],todirs[j]) != -1{
	//					file1 = append(file1,file[k])
	//				}else {
	//					continue
	//				}
	//			}
	//			DirTwo.File = file1
	//			DirTwos = append(DirTwos,DirTwo)
	//		}
	//		picture.Name = strings.Split(dirs[i],"\\")[2]
	//		picture.Dir = dirs[i]
	//		picture.ToDir = DirTwos
	//		pictures = append(pictures,picture)
	//	}
	//	pictureByte,_ := json.Marshal(pictures)
	//	c.String(http.StatusOK, string(pictureByte))
	//})
	//router.GET("/file", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "template/TR.html", gin.H{})
	//})
	router.POST("/getScreenshot", func(c *gin.Context) {
		service.GetScreenshot(c)
	})
	router.POST("/uploadScreenshots", func(c *gin.Context) {
		var zipFile *os.File
		zip, _ := c.FormFile("screenshots")
		var FilePath = "C:\\Users\\pky\\go\\src\\Screenshots\\zip\\" + zip.Filename
		zipFile, _ = os.Create(FilePath)
		var fileByte = make([]byte, zip.Size)
		file, _ := zip.Open()
		file.Read(fileByte)
		zipFile.Write(fileByte)
		util.UnZip(FilePath, "screenshot")
		util.UpLoadScreenShots(fileTmp)
		//dir, _ := ioutil.ReadDir(fileTmp)

		defer os.Remove(FilePath)
		defer file.Close()
		defer zipFile.Close()

	})
	router.GET("/routerUrl", func(c *gin.Context) {
		service.GetRoute(c)
	})
	router.Run(":8000")
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//下面的都是乱添加的-_-~
			//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			// c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			// c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Cache-Control", "no-cache")
			c.Set("content-type", "application/json")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
