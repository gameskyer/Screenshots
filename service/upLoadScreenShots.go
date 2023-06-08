package service

import (
	"Screenshots/dao"
	"Screenshots/module"
	"Screenshots/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Response struct {
	Message string       `json:"message"`
	Status  string       `json:"status"`
	Data    ResponseData `json:"data"`
}
type ResponseData struct {
	Url     string
	Md5     string `json:"md5"`
	Path    string `json:"path"`
	Domain  string `json:"domain"`
	Scene   string `json:"scene"`
	Size    int    `json:"size"`
	Mtime   int    `json:"mtime"`
	Scenes  string `json:"scenes"`
	Retmsg  string `json:"retmsg"`
	Retcode int    `json:"retcode"`
	Src     string `json:"src"`
}

type UploadData struct {
	GameName string
	Img      []string
}

//通过正则获取指定目录下的所有图片
func UpLoadScreenShots(filePath string) {
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		//正则匹配图片文件结尾
		reg, err := regexp.MatchString(`[^\s]+(\.(?i)(jpg|png|gif|bmp))$`, path)
		if err != nil {
			fmt.Println(err)
		}
		if reg {
			upLoadFile(path)
		}
		return nil
	})

}

func upLoadFile(filepath string) {
	//设置go-fastdfs文件路径
	paths := strings.SplitAfter(filepath, "\\")
	name := paths[len(paths)-3][:len(paths[len(paths)-3])-1]
	scene := paths[len(paths)-2][:len(paths[len(paths)-2])-1]
	path := paths[len(paths)-3] + paths[len(paths)-2][:len(paths[len(paths)-2])-1]
	path = strings.ReplaceAll(path, "\\", "/")

	var obj Response
	req := httplib.Post("http://" + util.Config.FileServerConfig.FileServerAddress + "/group1/upload")
	//设置文件名称
	var fileName string
	file, _ := os.Stat(filepath)
	fileName = file.Name()

	//上传文件
	req.PostFile("file", filepath) //注意不是全路径
	req.Param("filename", fileName)
	req.Param("output", "json2")
	req.Param("scene", "screen")
	req.Param("path", path)
	req.ToJSON(&obj)
	databyte, _ := json.Marshal(obj)
	fmt.Println(string(databyte))
	//上传文件后删除文件
	os.Remove(filepath)
	fileNames := strings.Split(fileName, ".")
	var insertModle module.Screenshot
	insertModle.Name = name
	insertModle.Scene = scene
	insertModle.Path = obj.Data.Path
	insertModle.Url = strings.ReplaceAll(obj.Data.Url, "\\u0026", "&")
	insertModle.Domain = obj.Data.Domain
	insertModle.FileName = fileNames[0]
	dao.InsertSql(insertModle)

}
