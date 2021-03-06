package util

import (
	"Screenshots/module"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/client/httplib"
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
	//var uploadDatas []UploadData
	filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		//正则匹配图片文件结尾
		reg, _ := regexp.MatchString(`[^\s]+(\.(?i)(jpg|png|gif|bmp))$`, path)
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
	fmt.Println(path)
	var obj Response
	req := httplib.Post("http://192.168.221.128:8080/group1/upload")
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
	var insertModle module.InsertSql
	insertModle.Name = name
	insertModle.Scene = scene
	insertModle.Path = obj.Data.Path
	insertModle.Url = strings.ReplaceAll(obj.Data.Url, "\\u0026", "&")
	insertModle.Domain = obj.Data.Domain
	InsertSql(insertModle)

}
