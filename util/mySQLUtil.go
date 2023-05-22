package util

import (
	"Screenshots/module"
	"fmt"
)

type Data struct {
	Path   string `json:"path" db:"path"`
	Domain string `json:"domain" db:"domain"`
}

func InsertSql(insertParm module.Screenshot) {
	sqlStr := "INSERT INTO screenshot(name, url,path,scene,domain ) VALUES(?,?,?,?,?)"
	result, err := DB.Exec(sqlStr, insertParm.Name, insertParm.Url, insertParm.Path, insertParm.Scene, insertParm.Domain)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert data success, id:%d\n", insertID)
}

func SelectGroupByName() []string {
	var names []string
	DB.Select(&names, "select name from screenshot GROUP BY name")
	return names
}

func SelectByName(name string) module.Router {
	var router module.Router
	sqlStr := fmt.Sprintf("select scene,path,name,url,domain from screenshot where name = '%s'", name)
	var data []module.ChildrenRoute
	DB.Select(&data, sqlStr)
	router.Name = name
	router.Path = "/" + name
	for i := 0; i < len(data); i++ {
		data[i].Url = data[i].Domain + data[i].Path
	}
	router.ChildrenRoute = data
	return router
}

func GetName() []string {
	var names []string
	sqlStr := fmt.Sprintf("select name from screenshot Group By name")
	DB.Select(&names, sqlStr)
	return names
}

func GetScreenshotScene(name string) module.ScreenshotModule {
	var router module.ScreenshotModule
	sqlStr := fmt.Sprintf("select scene from screenshot where name = '%s' Group By scene", name)
	var data []module.Scene
	DB.Select(&data, sqlStr)
	router.Name = name
	router.Path = "/" + name
	for i := 0; i < len(data); i++ {
		data[i].Path = "/scene"
		//data[i].Name = data[i].Sence
	}
	router.Scene = data
	return router
}

func SelectScreenshotBySceneAndName(name, scene string) []string {
	var screenshotData []string
	sqlStr := fmt.Sprintf("select path,domain from screenshot where name = '%s' and scene = '%s'", name, scene)
	fmt.Println(sqlStr)
	var data []Data
	DB.Select(&data, sqlStr)
	for _, v := range data {
		str := v.Domain + v.Path
		screenshotData = append(screenshotData, str)
	}
	return screenshotData
}
