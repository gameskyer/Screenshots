package test

import (
	"Screenshots/module"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

func TestMySQL(t *testing.T) {
	var DB *sqlx.DB
	var err error
	mysqlConnet := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost:3306", "screenshot")
	DB, err = sqlx.Open("mysql", mysqlConnet)
	// 最大连接数
	DB.SetMaxOpenConns(10)
	// 闲置连接数
	DB.SetMaxIdleConns(20)
	// 最大连接周期
	DB.SetConnMaxLifetime(100 * time.Second)
	if err = DB.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}

	var router module.ScreenshotModule
	sqlStr := fmt.Sprintf("select scene,name from screenshot where name = '%s' Group By scene", "wolcen")
	var data []module.Scene
	DB.Select(&data, sqlStr)
	router.Name = "name"
	router.Path = "/" + "name"
	for i := 0; i < len(data); i++ {
		data[i].Name = "/scene"
	}
	router.Scene = data
	fmt.Println(router)
}
