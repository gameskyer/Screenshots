package util

import (
	"Screenshots/been"
	"fmt"
)

func InsertSql(insertParm been.InsertSql){
	sqlStr := "INSERT INTO screenshot(name, url,path,scene,domain ) VALUES(?,?,?,?,?)"
	result, err := DB.Exec(sqlStr,insertParm.Name,insertParm.Url,insertParm.Path,insertParm.Scene,insertParm.Domain)
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

func SelectGroupByName() []string  {
	var names []string
	DB.Select(&names,"select name from screenshot GROUP BY name")
	return names
}

func SelectByName(name string)  been.Router{
	var router been.Router
	sqlStr := fmt.Sprintf("select scene,path,name,url,domain from screenshot where name = '%s'",name)
	var data []been.ChildrenRoute
	DB.Select(&data,sqlStr)
	router.Name = name
	router.Path = "/"+name
	for i := 0; i < len(data); i++ {
		data[i].Url = data[i].Domain + data[i].Path
	}
	router.ChildrenRoute = data
	return router
}
