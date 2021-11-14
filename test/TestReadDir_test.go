package test

import (
	"Screenshots/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/client/httplib"
	"testing"
)

func TestReadDir(t *testing.T)  {
	var obj util.Response
	//var obj interface{}
	req:=httplib.Post("http://192.168.18.128:8080/group1/upload")
	req.PostFile("file","D:\\截图\\APEX\\场景1\\中.jpg")//注意不是全路径
	req.Param("output","json2")
	req.Param("scene","screenshot")
	req.Param("path","test")
	req.ToJSON(&obj)
	databyte,_ := json.Marshal(obj)
	fmt.Println(string(databyte))
}

