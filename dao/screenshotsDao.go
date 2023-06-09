package dao

import (
	"Screenshots/db"
	"Screenshots/module"
	"Screenshots/util"
	"fmt"
)

func init() {
	sql := `create table if not exists "` + db.SCRENSHOTSTABLE + `" (
		"id" integer primary key autoincrement,
		"name" text not null,
		"url" text not null,
		"path" text not null,
		"scene" text not null,
		"fileName" text not null,
		"domain" text not null
	)`
	tx := db.DB.Exec(sql)
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
}

var screenshotDB = db.DB

func SelectScreenshotBySceneAndName(name, scene string) []module.Screenshot {
	//var urls []string
	var screenshotData []module.Screenshot
	tx := screenshotDB.Model(new(module.Screenshot))
	tx.Select("path", "domain", "fileName")
	tx.Where("name = ? ", name).Where("scene = ?", scene)
	tx.Scan(&screenshotData)
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
	//for _, v := range screenshotData {
	//	str := v.Domain + v.Path
	//	urls = append(urls, str)
	//}
	return screenshotData
}
func GetScreenshotScene(name string) module.ScreenshotModule {
	var router module.ScreenshotModule
	var data []module.Scene
	tx := screenshotDB.Model(new(module.Screenshot))
	tx.Select("scene,path")
	tx.Where("name = ?", name)
	tx.Group("scene")
	tx.Scan(&data)
	router.Name = name
	router.Path = "/" + name
	for i := 0; i < len(data); i++ {
		data[i].Path = "/scene"
		//data[i].Name = data[i].Sence
	}
	router.Scene = data
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
	return router
}
func GetName() []string {
	var names []string
	tx := screenshotDB.Model(new(module.Screenshot))
	tx.Select("name")
	tx.Group("name")
	tx.Scan(&names)
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
	return names
}
func SelectByName(name string) module.Router {
	var router module.Router
	tx := screenshotDB.Model(new(module.Screenshot))
	tx.Where("name = ?", name)
	tx.Scan(&router)
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
	return router
}
func SelectGroupByName() []string {
	var names []string
	tx := screenshotDB.Model(new(module.Screenshot))
	tx.Select("name")
	tx.Group("name")
	tx.Scan(&names)
	err := tx.Error
	if err != nil {
		util.Logger.Error(err)
	}
	return names
}
func InsertSql(insertParm module.Screenshot) {
	tx := screenshotDB.Create(&insertParm)
	err := tx.Error
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	insertID := insertParm.Id
	if err != nil {
		fmt.Printf("get insert id failed, err:%v\n", insertID)
		return
	}
	fmt.Printf("insert data success, id:%d\n", insertID)
}
